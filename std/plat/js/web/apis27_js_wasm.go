// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

type XRHitTestSource struct {
	ref js.Ref
}

func (this XRHitTestSource) Once() XRHitTestSource {
	this.ref.Once()
	return this
}

func (this XRHitTestSource) Ref() js.Ref {
	return this.ref
}

func (this XRHitTestSource) FromRef(ref js.Ref) XRHitTestSource {
	this.ref = ref
	return this
}

func (this XRHitTestSource) Free() {
	this.ref.Free()
}

// HasFuncCancel returns true if the method "XRHitTestSource.cancel" exists.
func (this XRHitTestSource) HasFuncCancel() bool {
	return js.True == bindings.HasFuncXRHitTestSourceCancel(
		this.ref,
	)
}

// FuncCancel returns the method "XRHitTestSource.cancel".
func (this XRHitTestSource) FuncCancel() (fn js.Func[func()]) {
	bindings.FuncXRHitTestSourceCancel(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Cancel calls the method "XRHitTestSource.cancel".
func (this XRHitTestSource) Cancel() (ret js.Void) {
	bindings.CallXRHitTestSourceCancel(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCancel calls the method "XRHitTestSource.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRHitTestSource) TryCancel() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRHitTestSourceCancel(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type XRHandedness uint32

const (
	_ XRHandedness = iota

	XRHandedness_NONE
	XRHandedness_LEFT
	XRHandedness_RIGHT
)

func (XRHandedness) FromRef(str js.Ref) XRHandedness {
	return XRHandedness(bindings.ConstOfXRHandedness(str))
}

func (x XRHandedness) String() (string, bool) {
	switch x {
	case XRHandedness_NONE:
		return "none", true
	case XRHandedness_LEFT:
		return "left", true
	case XRHandedness_RIGHT:
		return "right", true
	default:
		return "", false
	}
}

type XRTargetRayMode uint32

const (
	_ XRTargetRayMode = iota

	XRTargetRayMode_GAZE
	XRTargetRayMode_TRACKED_POINTER
	XRTargetRayMode_SCREEN
)

func (XRTargetRayMode) FromRef(str js.Ref) XRTargetRayMode {
	return XRTargetRayMode(bindings.ConstOfXRTargetRayMode(str))
}

func (x XRTargetRayMode) String() (string, bool) {
	switch x {
	case XRTargetRayMode_GAZE:
		return "gaze", true
	case XRTargetRayMode_TRACKED_POINTER:
		return "tracked-pointer", true
	case XRTargetRayMode_SCREEN:
		return "screen", true
	default:
		return "", false
	}
}

type XRHand struct {
	ref js.Ref
}

func (this XRHand) Once() XRHand {
	this.ref.Once()
	return this
}

func (this XRHand) Ref() js.Ref {
	return this.ref
}

func (this XRHand) FromRef(ref js.Ref) XRHand {
	this.ref = ref
	return this
}

func (this XRHand) Free() {
	this.ref.Free()
}

// Size returns the value of property "XRHand.size".
//
// It returns ok=false if there is no such property.
func (this XRHand) Size() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRHandSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGet returns true if the method "XRHand.get" exists.
func (this XRHand) HasFuncGet() bool {
	return js.True == bindings.HasFuncXRHandGet(
		this.ref,
	)
}

// FuncGet returns the method "XRHand.get".
func (this XRHand) FuncGet() (fn js.Func[func(key XRHandJoint) XRJointSpace]) {
	bindings.FuncXRHandGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "XRHand.get".
func (this XRHand) Get(key XRHandJoint) (ret XRJointSpace) {
	bindings.CallXRHandGet(
		this.ref, js.Pointer(&ret),
		uint32(key),
	)

	return
}

// TryGet calls the method "XRHand.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRHand) TryGet(key XRHandJoint) (ret XRJointSpace, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRHandGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(key),
	)

	return
}

type XRInputSource struct {
	ref js.Ref
}

func (this XRInputSource) Once() XRInputSource {
	this.ref.Once()
	return this
}

func (this XRInputSource) Ref() js.Ref {
	return this.ref
}

func (this XRInputSource) FromRef(ref js.Ref) XRInputSource {
	this.ref = ref
	return this
}

func (this XRInputSource) Free() {
	this.ref.Free()
}

// Handedness returns the value of property "XRInputSource.handedness".
//
// It returns ok=false if there is no such property.
func (this XRInputSource) Handedness() (ret XRHandedness, ok bool) {
	ok = js.True == bindings.GetXRInputSourceHandedness(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TargetRayMode returns the value of property "XRInputSource.targetRayMode".
//
// It returns ok=false if there is no such property.
func (this XRInputSource) TargetRayMode() (ret XRTargetRayMode, ok bool) {
	ok = js.True == bindings.GetXRInputSourceTargetRayMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TargetRaySpace returns the value of property "XRInputSource.targetRaySpace".
//
// It returns ok=false if there is no such property.
func (this XRInputSource) TargetRaySpace() (ret XRSpace, ok bool) {
	ok = js.True == bindings.GetXRInputSourceTargetRaySpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// GripSpace returns the value of property "XRInputSource.gripSpace".
//
// It returns ok=false if there is no such property.
func (this XRInputSource) GripSpace() (ret XRSpace, ok bool) {
	ok = js.True == bindings.GetXRInputSourceGripSpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Profiles returns the value of property "XRInputSource.profiles".
//
// It returns ok=false if there is no such property.
func (this XRInputSource) Profiles() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetXRInputSourceProfiles(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Gamepad returns the value of property "XRInputSource.gamepad".
//
// It returns ok=false if there is no such property.
func (this XRInputSource) Gamepad() (ret Gamepad, ok bool) {
	ok = js.True == bindings.GetXRInputSourceGamepad(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Hand returns the value of property "XRInputSource.hand".
//
// It returns ok=false if there is no such property.
func (this XRInputSource) Hand() (ret XRHand, ok bool) {
	ok = js.True == bindings.GetXRInputSourceHand(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRTransientInputHitTestResult struct {
	ref js.Ref
}

func (this XRTransientInputHitTestResult) Once() XRTransientInputHitTestResult {
	this.ref.Once()
	return this
}

func (this XRTransientInputHitTestResult) Ref() js.Ref {
	return this.ref
}

func (this XRTransientInputHitTestResult) FromRef(ref js.Ref) XRTransientInputHitTestResult {
	this.ref = ref
	return this
}

func (this XRTransientInputHitTestResult) Free() {
	this.ref.Free()
}

// InputSource returns the value of property "XRTransientInputHitTestResult.inputSource".
//
// It returns ok=false if there is no such property.
func (this XRTransientInputHitTestResult) InputSource() (ret XRInputSource, ok bool) {
	ok = js.True == bindings.GetXRTransientInputHitTestResultInputSource(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Results returns the value of property "XRTransientInputHitTestResult.results".
//
// It returns ok=false if there is no such property.
func (this XRTransientInputHitTestResult) Results() (ret js.FrozenArray[XRHitTestResult], ok bool) {
	ok = js.True == bindings.GetXRTransientInputHitTestResultResults(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRTransientInputHitTestSource struct {
	ref js.Ref
}

func (this XRTransientInputHitTestSource) Once() XRTransientInputHitTestSource {
	this.ref.Once()
	return this
}

func (this XRTransientInputHitTestSource) Ref() js.Ref {
	return this.ref
}

func (this XRTransientInputHitTestSource) FromRef(ref js.Ref) XRTransientInputHitTestSource {
	this.ref = ref
	return this
}

func (this XRTransientInputHitTestSource) Free() {
	this.ref.Free()
}

// HasFuncCancel returns true if the method "XRTransientInputHitTestSource.cancel" exists.
func (this XRTransientInputHitTestSource) HasFuncCancel() bool {
	return js.True == bindings.HasFuncXRTransientInputHitTestSourceCancel(
		this.ref,
	)
}

// FuncCancel returns the method "XRTransientInputHitTestSource.cancel".
func (this XRTransientInputHitTestSource) FuncCancel() (fn js.Func[func()]) {
	bindings.FuncXRTransientInputHitTestSourceCancel(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Cancel calls the method "XRTransientInputHitTestSource.cancel".
func (this XRTransientInputHitTestSource) Cancel() (ret js.Void) {
	bindings.CallXRTransientInputHitTestSourceCancel(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCancel calls the method "XRTransientInputHitTestSource.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRTransientInputHitTestSource) TryCancel() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRTransientInputHitTestSourceCancel(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type XRMeshSet struct {
	ref js.Ref
}

func (this XRMeshSet) Once() XRMeshSet {
	this.ref.Once()
	return this
}

func (this XRMeshSet) Ref() js.Ref {
	return this.ref
}

func (this XRMeshSet) FromRef(ref js.Ref) XRMeshSet {
	this.ref = ref
	return this
}

func (this XRMeshSet) Free() {
	this.ref.Free()
}

type XRAnchorSet struct {
	ref js.Ref
}

func (this XRAnchorSet) Once() XRAnchorSet {
	this.ref.Once()
	return this
}

func (this XRAnchorSet) Ref() js.Ref {
	return this.ref
}

func (this XRAnchorSet) FromRef(ref js.Ref) XRAnchorSet {
	this.ref = ref
	return this
}

func (this XRAnchorSet) Free() {
	this.ref.Free()
}

type XRFrame struct {
	ref js.Ref
}

func (this XRFrame) Once() XRFrame {
	this.ref.Once()
	return this
}

func (this XRFrame) Ref() js.Ref {
	return this.ref
}

func (this XRFrame) FromRef(ref js.Ref) XRFrame {
	this.ref = ref
	return this
}

func (this XRFrame) Free() {
	this.ref.Free()
}

// Session returns the value of property "XRFrame.session".
//
// It returns ok=false if there is no such property.
func (this XRFrame) Session() (ret XRSession, ok bool) {
	ok = js.True == bindings.GetXRFrameSession(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PredictedDisplayTime returns the value of property "XRFrame.predictedDisplayTime".
//
// It returns ok=false if there is no such property.
func (this XRFrame) PredictedDisplayTime() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetXRFramePredictedDisplayTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DetectedMeshes returns the value of property "XRFrame.detectedMeshes".
//
// It returns ok=false if there is no such property.
func (this XRFrame) DetectedMeshes() (ret XRMeshSet, ok bool) {
	ok = js.True == bindings.GetXRFrameDetectedMeshes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TrackedAnchors returns the value of property "XRFrame.trackedAnchors".
//
// It returns ok=false if there is no such property.
func (this XRFrame) TrackedAnchors() (ret XRAnchorSet, ok bool) {
	ok = js.True == bindings.GetXRFrameTrackedAnchors(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetViewerPose returns true if the method "XRFrame.getViewerPose" exists.
func (this XRFrame) HasFuncGetViewerPose() bool {
	return js.True == bindings.HasFuncXRFrameGetViewerPose(
		this.ref,
	)
}

// FuncGetViewerPose returns the method "XRFrame.getViewerPose".
func (this XRFrame) FuncGetViewerPose() (fn js.Func[func(referenceSpace XRReferenceSpace) XRViewerPose]) {
	bindings.FuncXRFrameGetViewerPose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetViewerPose calls the method "XRFrame.getViewerPose".
func (this XRFrame) GetViewerPose(referenceSpace XRReferenceSpace) (ret XRViewerPose) {
	bindings.CallXRFrameGetViewerPose(
		this.ref, js.Pointer(&ret),
		referenceSpace.Ref(),
	)

	return
}

// TryGetViewerPose calls the method "XRFrame.getViewerPose"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRFrame) TryGetViewerPose(referenceSpace XRReferenceSpace) (ret XRViewerPose, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRFrameGetViewerPose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		referenceSpace.Ref(),
	)

	return
}

// HasFuncGetPose returns true if the method "XRFrame.getPose" exists.
func (this XRFrame) HasFuncGetPose() bool {
	return js.True == bindings.HasFuncXRFrameGetPose(
		this.ref,
	)
}

// FuncGetPose returns the method "XRFrame.getPose".
func (this XRFrame) FuncGetPose() (fn js.Func[func(space XRSpace, baseSpace XRSpace) XRPose]) {
	bindings.FuncXRFrameGetPose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetPose calls the method "XRFrame.getPose".
func (this XRFrame) GetPose(space XRSpace, baseSpace XRSpace) (ret XRPose) {
	bindings.CallXRFrameGetPose(
		this.ref, js.Pointer(&ret),
		space.Ref(),
		baseSpace.Ref(),
	)

	return
}

// TryGetPose calls the method "XRFrame.getPose"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRFrame) TryGetPose(space XRSpace, baseSpace XRSpace) (ret XRPose, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRFrameGetPose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		space.Ref(),
		baseSpace.Ref(),
	)

	return
}

// HasFuncCreateAnchor returns true if the method "XRFrame.createAnchor" exists.
func (this XRFrame) HasFuncCreateAnchor() bool {
	return js.True == bindings.HasFuncXRFrameCreateAnchor(
		this.ref,
	)
}

// FuncCreateAnchor returns the method "XRFrame.createAnchor".
func (this XRFrame) FuncCreateAnchor() (fn js.Func[func(pose XRRigidTransform, space XRSpace) js.Promise[XRAnchor]]) {
	bindings.FuncXRFrameCreateAnchor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateAnchor calls the method "XRFrame.createAnchor".
func (this XRFrame) CreateAnchor(pose XRRigidTransform, space XRSpace) (ret js.Promise[XRAnchor]) {
	bindings.CallXRFrameCreateAnchor(
		this.ref, js.Pointer(&ret),
		pose.Ref(),
		space.Ref(),
	)

	return
}

// TryCreateAnchor calls the method "XRFrame.createAnchor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRFrame) TryCreateAnchor(pose XRRigidTransform, space XRSpace) (ret js.Promise[XRAnchor], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRFrameCreateAnchor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		pose.Ref(),
		space.Ref(),
	)

	return
}

// HasFuncGetLightEstimate returns true if the method "XRFrame.getLightEstimate" exists.
func (this XRFrame) HasFuncGetLightEstimate() bool {
	return js.True == bindings.HasFuncXRFrameGetLightEstimate(
		this.ref,
	)
}

// FuncGetLightEstimate returns the method "XRFrame.getLightEstimate".
func (this XRFrame) FuncGetLightEstimate() (fn js.Func[func(lightProbe XRLightProbe) XRLightEstimate]) {
	bindings.FuncXRFrameGetLightEstimate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetLightEstimate calls the method "XRFrame.getLightEstimate".
func (this XRFrame) GetLightEstimate(lightProbe XRLightProbe) (ret XRLightEstimate) {
	bindings.CallXRFrameGetLightEstimate(
		this.ref, js.Pointer(&ret),
		lightProbe.Ref(),
	)

	return
}

// TryGetLightEstimate calls the method "XRFrame.getLightEstimate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRFrame) TryGetLightEstimate(lightProbe XRLightProbe) (ret XRLightEstimate, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRFrameGetLightEstimate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		lightProbe.Ref(),
	)

	return
}

// HasFuncGetDepthInformation returns true if the method "XRFrame.getDepthInformation" exists.
func (this XRFrame) HasFuncGetDepthInformation() bool {
	return js.True == bindings.HasFuncXRFrameGetDepthInformation(
		this.ref,
	)
}

// FuncGetDepthInformation returns the method "XRFrame.getDepthInformation".
func (this XRFrame) FuncGetDepthInformation() (fn js.Func[func(view XRView) XRCPUDepthInformation]) {
	bindings.FuncXRFrameGetDepthInformation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDepthInformation calls the method "XRFrame.getDepthInformation".
func (this XRFrame) GetDepthInformation(view XRView) (ret XRCPUDepthInformation) {
	bindings.CallXRFrameGetDepthInformation(
		this.ref, js.Pointer(&ret),
		view.Ref(),
	)

	return
}

// TryGetDepthInformation calls the method "XRFrame.getDepthInformation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRFrame) TryGetDepthInformation(view XRView) (ret XRCPUDepthInformation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRFrameGetDepthInformation(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		view.Ref(),
	)

	return
}

// HasFuncGetJointPose returns true if the method "XRFrame.getJointPose" exists.
func (this XRFrame) HasFuncGetJointPose() bool {
	return js.True == bindings.HasFuncXRFrameGetJointPose(
		this.ref,
	)
}

// FuncGetJointPose returns the method "XRFrame.getJointPose".
func (this XRFrame) FuncGetJointPose() (fn js.Func[func(joint XRJointSpace, baseSpace XRSpace) XRJointPose]) {
	bindings.FuncXRFrameGetJointPose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetJointPose calls the method "XRFrame.getJointPose".
func (this XRFrame) GetJointPose(joint XRJointSpace, baseSpace XRSpace) (ret XRJointPose) {
	bindings.CallXRFrameGetJointPose(
		this.ref, js.Pointer(&ret),
		joint.Ref(),
		baseSpace.Ref(),
	)

	return
}

// TryGetJointPose calls the method "XRFrame.getJointPose"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRFrame) TryGetJointPose(joint XRJointSpace, baseSpace XRSpace) (ret XRJointPose, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRFrameGetJointPose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		joint.Ref(),
		baseSpace.Ref(),
	)

	return
}

// HasFuncFillJointRadii returns true if the method "XRFrame.fillJointRadii" exists.
func (this XRFrame) HasFuncFillJointRadii() bool {
	return js.True == bindings.HasFuncXRFrameFillJointRadii(
		this.ref,
	)
}

// FuncFillJointRadii returns the method "XRFrame.fillJointRadii".
func (this XRFrame) FuncFillJointRadii() (fn js.Func[func(jointSpaces js.Array[XRJointSpace], radii js.TypedArray[float32]) bool]) {
	bindings.FuncXRFrameFillJointRadii(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FillJointRadii calls the method "XRFrame.fillJointRadii".
func (this XRFrame) FillJointRadii(jointSpaces js.Array[XRJointSpace], radii js.TypedArray[float32]) (ret bool) {
	bindings.CallXRFrameFillJointRadii(
		this.ref, js.Pointer(&ret),
		jointSpaces.Ref(),
		radii.Ref(),
	)

	return
}

// TryFillJointRadii calls the method "XRFrame.fillJointRadii"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRFrame) TryFillJointRadii(jointSpaces js.Array[XRJointSpace], radii js.TypedArray[float32]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRFrameFillJointRadii(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		jointSpaces.Ref(),
		radii.Ref(),
	)

	return
}

// HasFuncFillPoses returns true if the method "XRFrame.fillPoses" exists.
func (this XRFrame) HasFuncFillPoses() bool {
	return js.True == bindings.HasFuncXRFrameFillPoses(
		this.ref,
	)
}

// FuncFillPoses returns the method "XRFrame.fillPoses".
func (this XRFrame) FuncFillPoses() (fn js.Func[func(spaces js.Array[XRSpace], baseSpace XRSpace, transforms js.TypedArray[float32]) bool]) {
	bindings.FuncXRFrameFillPoses(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FillPoses calls the method "XRFrame.fillPoses".
func (this XRFrame) FillPoses(spaces js.Array[XRSpace], baseSpace XRSpace, transforms js.TypedArray[float32]) (ret bool) {
	bindings.CallXRFrameFillPoses(
		this.ref, js.Pointer(&ret),
		spaces.Ref(),
		baseSpace.Ref(),
		transforms.Ref(),
	)

	return
}

// TryFillPoses calls the method "XRFrame.fillPoses"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRFrame) TryFillPoses(spaces js.Array[XRSpace], baseSpace XRSpace, transforms js.TypedArray[float32]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRFrameFillPoses(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		spaces.Ref(),
		baseSpace.Ref(),
		transforms.Ref(),
	)

	return
}

// HasFuncGetHitTestResults returns true if the method "XRFrame.getHitTestResults" exists.
func (this XRFrame) HasFuncGetHitTestResults() bool {
	return js.True == bindings.HasFuncXRFrameGetHitTestResults(
		this.ref,
	)
}

// FuncGetHitTestResults returns the method "XRFrame.getHitTestResults".
func (this XRFrame) FuncGetHitTestResults() (fn js.Func[func(hitTestSource XRHitTestSource) js.FrozenArray[XRHitTestResult]]) {
	bindings.FuncXRFrameGetHitTestResults(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetHitTestResults calls the method "XRFrame.getHitTestResults".
func (this XRFrame) GetHitTestResults(hitTestSource XRHitTestSource) (ret js.FrozenArray[XRHitTestResult]) {
	bindings.CallXRFrameGetHitTestResults(
		this.ref, js.Pointer(&ret),
		hitTestSource.Ref(),
	)

	return
}

// TryGetHitTestResults calls the method "XRFrame.getHitTestResults"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRFrame) TryGetHitTestResults(hitTestSource XRHitTestSource) (ret js.FrozenArray[XRHitTestResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRFrameGetHitTestResults(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		hitTestSource.Ref(),
	)

	return
}

// HasFuncGetHitTestResultsForTransientInput returns true if the method "XRFrame.getHitTestResultsForTransientInput" exists.
func (this XRFrame) HasFuncGetHitTestResultsForTransientInput() bool {
	return js.True == bindings.HasFuncXRFrameGetHitTestResultsForTransientInput(
		this.ref,
	)
}

// FuncGetHitTestResultsForTransientInput returns the method "XRFrame.getHitTestResultsForTransientInput".
func (this XRFrame) FuncGetHitTestResultsForTransientInput() (fn js.Func[func(hitTestSource XRTransientInputHitTestSource) js.FrozenArray[XRTransientInputHitTestResult]]) {
	bindings.FuncXRFrameGetHitTestResultsForTransientInput(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetHitTestResultsForTransientInput calls the method "XRFrame.getHitTestResultsForTransientInput".
func (this XRFrame) GetHitTestResultsForTransientInput(hitTestSource XRTransientInputHitTestSource) (ret js.FrozenArray[XRTransientInputHitTestResult]) {
	bindings.CallXRFrameGetHitTestResultsForTransientInput(
		this.ref, js.Pointer(&ret),
		hitTestSource.Ref(),
	)

	return
}

// TryGetHitTestResultsForTransientInput calls the method "XRFrame.getHitTestResultsForTransientInput"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRFrame) TryGetHitTestResultsForTransientInput(hitTestSource XRTransientInputHitTestSource) (ret js.FrozenArray[XRTransientInputHitTestResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRFrameGetHitTestResultsForTransientInput(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		hitTestSource.Ref(),
	)

	return
}

type XRReflectionFormat uint32

const (
	_ XRReflectionFormat = iota

	XRReflectionFormat_SRGBA8
	XRReflectionFormat_RGBA_16F
)

func (XRReflectionFormat) FromRef(str js.Ref) XRReflectionFormat {
	return XRReflectionFormat(bindings.ConstOfXRReflectionFormat(str))
}

func (x XRReflectionFormat) String() (string, bool) {
	switch x {
	case XRReflectionFormat_SRGBA8:
		return "srgba8", true
	case XRReflectionFormat_RGBA_16F:
		return "rgba16f", true
	default:
		return "", false
	}
}

type XRLightProbeInit struct {
	// ReflectionFormat is "XRLightProbeInit.reflectionFormat"
	//
	// Optional, defaults to "srgba8".
	ReflectionFormat XRReflectionFormat

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRLightProbeInit with all fields set.
func (p XRLightProbeInit) FromRef(ref js.Ref) XRLightProbeInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRLightProbeInit in the application heap.
func (p XRLightProbeInit) New() js.Ref {
	return bindings.XRLightProbeInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRLightProbeInit) UpdateFrom(ref js.Ref) {
	bindings.XRLightProbeInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRLightProbeInit) Update(ref js.Ref) {
	bindings.XRLightProbeInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRLightProbeInit) FreeMembers(recursive bool) {
}

type XRHitTestTrackableType uint32

const (
	_ XRHitTestTrackableType = iota

	XRHitTestTrackableType_POINT
	XRHitTestTrackableType_PLANE
	XRHitTestTrackableType_MESH
)

func (XRHitTestTrackableType) FromRef(str js.Ref) XRHitTestTrackableType {
	return XRHitTestTrackableType(bindings.ConstOfXRHitTestTrackableType(str))
}

func (x XRHitTestTrackableType) String() (string, bool) {
	switch x {
	case XRHitTestTrackableType_POINT:
		return "point", true
	case XRHitTestTrackableType_PLANE:
		return "plane", true
	case XRHitTestTrackableType_MESH:
		return "mesh", true
	default:
		return "", false
	}
}

type XRRayDirectionInit struct {
	// X is "XRRayDirectionInit.x"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_X MUST be set to true to make this field effective.
	X float64
	// Y is "XRRayDirectionInit.y"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Y MUST be set to true to make this field effective.
	Y float64
	// Z is "XRRayDirectionInit.z"
	//
	// Optional, defaults to -1.
	//
	// NOTE: FFI_USE_Z MUST be set to true to make this field effective.
	Z float64
	// W is "XRRayDirectionInit.w"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_W MUST be set to true to make this field effective.
	W float64

	FFI_USE_X bool // for X.
	FFI_USE_Y bool // for Y.
	FFI_USE_Z bool // for Z.
	FFI_USE_W bool // for W.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRRayDirectionInit with all fields set.
func (p XRRayDirectionInit) FromRef(ref js.Ref) XRRayDirectionInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRRayDirectionInit in the application heap.
func (p XRRayDirectionInit) New() js.Ref {
	return bindings.XRRayDirectionInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRRayDirectionInit) UpdateFrom(ref js.Ref) {
	bindings.XRRayDirectionInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRRayDirectionInit) Update(ref js.Ref) {
	bindings.XRRayDirectionInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRRayDirectionInit) FreeMembers(recursive bool) {
}

func NewXRRay(origin DOMPointInit, direction XRRayDirectionInit) (ret XRRay) {
	ret.ref = bindings.NewXRRayByXRRay(
		js.Pointer(&origin),
		js.Pointer(&direction))
	return
}

func NewXRRayByXRRay1(origin DOMPointInit) (ret XRRay) {
	ret.ref = bindings.NewXRRayByXRRay1(
		js.Pointer(&origin))
	return
}

func NewXRRayByXRRay2() (ret XRRay) {
	ret.ref = bindings.NewXRRayByXRRay2()
	return
}

func NewXRRayByXRRay3(transform XRRigidTransform) (ret XRRay) {
	ret.ref = bindings.NewXRRayByXRRay3(
		transform.Ref())
	return
}

type XRRay struct {
	ref js.Ref
}

func (this XRRay) Once() XRRay {
	this.ref.Once()
	return this
}

func (this XRRay) Ref() js.Ref {
	return this.ref
}

func (this XRRay) FromRef(ref js.Ref) XRRay {
	this.ref = ref
	return this
}

func (this XRRay) Free() {
	this.ref.Free()
}

// Origin returns the value of property "XRRay.origin".
//
// It returns ok=false if there is no such property.
func (this XRRay) Origin() (ret DOMPointReadOnly, ok bool) {
	ok = js.True == bindings.GetXRRayOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Direction returns the value of property "XRRay.direction".
//
// It returns ok=false if there is no such property.
func (this XRRay) Direction() (ret DOMPointReadOnly, ok bool) {
	ok = js.True == bindings.GetXRRayDirection(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Matrix returns the value of property "XRRay.matrix".
//
// It returns ok=false if there is no such property.
func (this XRRay) Matrix() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetXRRayMatrix(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRHitTestOptionsInit struct {
	// Space is "XRHitTestOptionsInit.space"
	//
	// Required
	Space XRSpace
	// EntityTypes is "XRHitTestOptionsInit.entityTypes"
	//
	// Optional
	EntityTypes js.FrozenArray[XRHitTestTrackableType]
	// OffsetRay is "XRHitTestOptionsInit.offsetRay"
	//
	// Optional
	OffsetRay XRRay

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRHitTestOptionsInit with all fields set.
func (p XRHitTestOptionsInit) FromRef(ref js.Ref) XRHitTestOptionsInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRHitTestOptionsInit in the application heap.
func (p XRHitTestOptionsInit) New() js.Ref {
	return bindings.XRHitTestOptionsInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRHitTestOptionsInit) UpdateFrom(ref js.Ref) {
	bindings.XRHitTestOptionsInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRHitTestOptionsInit) Update(ref js.Ref) {
	bindings.XRHitTestOptionsInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRHitTestOptionsInit) FreeMembers(recursive bool) {
	js.Free(
		p.Space.Ref(),
		p.EntityTypes.Ref(),
		p.OffsetRay.Ref(),
	)
	p.Space = p.Space.FromRef(js.Undefined)
	p.EntityTypes = p.EntityTypes.FromRef(js.Undefined)
	p.OffsetRay = p.OffsetRay.FromRef(js.Undefined)
}

type XRTransientInputHitTestOptionsInit struct {
	// Profile is "XRTransientInputHitTestOptionsInit.profile"
	//
	// Required
	Profile js.String
	// EntityTypes is "XRTransientInputHitTestOptionsInit.entityTypes"
	//
	// Optional
	EntityTypes js.FrozenArray[XRHitTestTrackableType]
	// OffsetRay is "XRTransientInputHitTestOptionsInit.offsetRay"
	//
	// Optional
	OffsetRay XRRay

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRTransientInputHitTestOptionsInit with all fields set.
func (p XRTransientInputHitTestOptionsInit) FromRef(ref js.Ref) XRTransientInputHitTestOptionsInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRTransientInputHitTestOptionsInit in the application heap.
func (p XRTransientInputHitTestOptionsInit) New() js.Ref {
	return bindings.XRTransientInputHitTestOptionsInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRTransientInputHitTestOptionsInit) UpdateFrom(ref js.Ref) {
	bindings.XRTransientInputHitTestOptionsInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRTransientInputHitTestOptionsInit) Update(ref js.Ref) {
	bindings.XRTransientInputHitTestOptionsInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRTransientInputHitTestOptionsInit) FreeMembers(recursive bool) {
	js.Free(
		p.Profile.Ref(),
		p.EntityTypes.Ref(),
		p.OffsetRay.Ref(),
	)
	p.Profile = p.Profile.FromRef(js.Undefined)
	p.EntityTypes = p.EntityTypes.FromRef(js.Undefined)
	p.OffsetRay = p.OffsetRay.FromRef(js.Undefined)
}

type XRVisibilityState uint32

const (
	_ XRVisibilityState = iota

	XRVisibilityState_VISIBLE
	XRVisibilityState_VISIBLE_BLURRED
	XRVisibilityState_HIDDEN
)

func (XRVisibilityState) FromRef(str js.Ref) XRVisibilityState {
	return XRVisibilityState(bindings.ConstOfXRVisibilityState(str))
}

func (x XRVisibilityState) String() (string, bool) {
	switch x {
	case XRVisibilityState_VISIBLE:
		return "visible", true
	case XRVisibilityState_VISIBLE_BLURRED:
		return "visible-blurred", true
	case XRVisibilityState_HIDDEN:
		return "hidden", true
	default:
		return "", false
	}
}

type XRRenderState struct {
	ref js.Ref
}

func (this XRRenderState) Once() XRRenderState {
	this.ref.Once()
	return this
}

func (this XRRenderState) Ref() js.Ref {
	return this.ref
}

func (this XRRenderState) FromRef(ref js.Ref) XRRenderState {
	this.ref = ref
	return this
}

func (this XRRenderState) Free() {
	this.ref.Free()
}

// DepthNear returns the value of property "XRRenderState.depthNear".
//
// It returns ok=false if there is no such property.
func (this XRRenderState) DepthNear() (ret float64, ok bool) {
	ok = js.True == bindings.GetXRRenderStateDepthNear(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DepthFar returns the value of property "XRRenderState.depthFar".
//
// It returns ok=false if there is no such property.
func (this XRRenderState) DepthFar() (ret float64, ok bool) {
	ok = js.True == bindings.GetXRRenderStateDepthFar(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InlineVerticalFieldOfView returns the value of property "XRRenderState.inlineVerticalFieldOfView".
//
// It returns ok=false if there is no such property.
func (this XRRenderState) InlineVerticalFieldOfView() (ret float64, ok bool) {
	ok = js.True == bindings.GetXRRenderStateInlineVerticalFieldOfView(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BaseLayer returns the value of property "XRRenderState.baseLayer".
//
// It returns ok=false if there is no such property.
func (this XRRenderState) BaseLayer() (ret XRWebGLLayer, ok bool) {
	ok = js.True == bindings.GetXRRenderStateBaseLayer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Layers returns the value of property "XRRenderState.layers".
//
// It returns ok=false if there is no such property.
func (this XRRenderState) Layers() (ret js.FrozenArray[XRLayer], ok bool) {
	ok = js.True == bindings.GetXRRenderStateLayers(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRInputSourceArray struct {
	ref js.Ref
}

func (this XRInputSourceArray) Once() XRInputSourceArray {
	this.ref.Once()
	return this
}

func (this XRInputSourceArray) Ref() js.Ref {
	return this.ref
}

func (this XRInputSourceArray) FromRef(ref js.Ref) XRInputSourceArray {
	this.ref = ref
	return this
}

func (this XRInputSourceArray) Free() {
	this.ref.Free()
}

// Length returns the value of property "XRInputSourceArray.length".
//
// It returns ok=false if there is no such property.
func (this XRInputSourceArray) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRInputSourceArrayLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGet returns true if the method "XRInputSourceArray." exists.
func (this XRInputSourceArray) HasFuncGet() bool {
	return js.True == bindings.HasFuncXRInputSourceArrayGet(
		this.ref,
	)
}

// FuncGet returns the method "XRInputSourceArray.".
func (this XRInputSourceArray) FuncGet() (fn js.Func[func(index uint32) XRInputSource]) {
	bindings.FuncXRInputSourceArrayGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "XRInputSourceArray.".
func (this XRInputSourceArray) Get(index uint32) (ret XRInputSource) {
	bindings.CallXRInputSourceArrayGet(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "XRInputSourceArray."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRInputSourceArray) TryGet(index uint32) (ret XRInputSource, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRInputSourceArrayGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

type XREnvironmentBlendMode uint32

const (
	_ XREnvironmentBlendMode = iota

	XREnvironmentBlendMode_OPAQUE
	XREnvironmentBlendMode_ALPHA_BLEND
	XREnvironmentBlendMode_ADDITIVE
)

func (XREnvironmentBlendMode) FromRef(str js.Ref) XREnvironmentBlendMode {
	return XREnvironmentBlendMode(bindings.ConstOfXREnvironmentBlendMode(str))
}

func (x XREnvironmentBlendMode) String() (string, bool) {
	switch x {
	case XREnvironmentBlendMode_OPAQUE:
		return "opaque", true
	case XREnvironmentBlendMode_ALPHA_BLEND:
		return "alpha-blend", true
	case XREnvironmentBlendMode_ADDITIVE:
		return "additive", true
	default:
		return "", false
	}
}

type XRInteractionMode uint32

const (
	_ XRInteractionMode = iota

	XRInteractionMode_SCREEN_SPACE
	XRInteractionMode_WORLD_SPACE
)

func (XRInteractionMode) FromRef(str js.Ref) XRInteractionMode {
	return XRInteractionMode(bindings.ConstOfXRInteractionMode(str))
}

func (x XRInteractionMode) String() (string, bool) {
	switch x {
	case XRInteractionMode_SCREEN_SPACE:
		return "screen-space", true
	case XRInteractionMode_WORLD_SPACE:
		return "world-space", true
	default:
		return "", false
	}
}

type XRDOMOverlayType uint32

const (
	_ XRDOMOverlayType = iota

	XRDOMOverlayType_SCREEN
	XRDOMOverlayType_FLOATING
	XRDOMOverlayType_HEAD_LOCKED
)

func (XRDOMOverlayType) FromRef(str js.Ref) XRDOMOverlayType {
	return XRDOMOverlayType(bindings.ConstOfXRDOMOverlayType(str))
}

func (x XRDOMOverlayType) String() (string, bool) {
	switch x {
	case XRDOMOverlayType_SCREEN:
		return "screen", true
	case XRDOMOverlayType_FLOATING:
		return "floating", true
	case XRDOMOverlayType_HEAD_LOCKED:
		return "head-locked", true
	default:
		return "", false
	}
}

type XRDOMOverlayState struct {
	// Type is "XRDOMOverlayState.type"
	//
	// Optional
	Type XRDOMOverlayType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRDOMOverlayState with all fields set.
func (p XRDOMOverlayState) FromRef(ref js.Ref) XRDOMOverlayState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRDOMOverlayState in the application heap.
func (p XRDOMOverlayState) New() js.Ref {
	return bindings.XRDOMOverlayStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRDOMOverlayState) UpdateFrom(ref js.Ref) {
	bindings.XRDOMOverlayStateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRDOMOverlayState) Update(ref js.Ref) {
	bindings.XRDOMOverlayStateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRDOMOverlayState) FreeMembers(recursive bool) {
}

type XRDepthUsage uint32

const (
	_ XRDepthUsage = iota

	XRDepthUsage_CPU_OPTIMIZED
	XRDepthUsage_GPU_OPTIMIZED
)

func (XRDepthUsage) FromRef(str js.Ref) XRDepthUsage {
	return XRDepthUsage(bindings.ConstOfXRDepthUsage(str))
}

func (x XRDepthUsage) String() (string, bool) {
	switch x {
	case XRDepthUsage_CPU_OPTIMIZED:
		return "cpu-optimized", true
	case XRDepthUsage_GPU_OPTIMIZED:
		return "gpu-optimized", true
	default:
		return "", false
	}
}

type XRDepthDataFormat uint32

const (
	_ XRDepthDataFormat = iota

	XRDepthDataFormat_LUMINANCE_ALPHA
	XRDepthDataFormat_FLOAT32
)

func (XRDepthDataFormat) FromRef(str js.Ref) XRDepthDataFormat {
	return XRDepthDataFormat(bindings.ConstOfXRDepthDataFormat(str))
}

func (x XRDepthDataFormat) String() (string, bool) {
	switch x {
	case XRDepthDataFormat_LUMINANCE_ALPHA:
		return "luminance-alpha", true
	case XRDepthDataFormat_FLOAT32:
		return "float32", true
	default:
		return "", false
	}
}

type XRSession struct {
	EventTarget
}

func (this XRSession) Once() XRSession {
	this.ref.Once()
	return this
}

func (this XRSession) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this XRSession) FromRef(ref js.Ref) XRSession {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this XRSession) Free() {
	this.ref.Free()
}

// VisibilityState returns the value of property "XRSession.visibilityState".
//
// It returns ok=false if there is no such property.
func (this XRSession) VisibilityState() (ret XRVisibilityState, ok bool) {
	ok = js.True == bindings.GetXRSessionVisibilityState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FrameRate returns the value of property "XRSession.frameRate".
//
// It returns ok=false if there is no such property.
func (this XRSession) FrameRate() (ret float32, ok bool) {
	ok = js.True == bindings.GetXRSessionFrameRate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SupportedFrameRates returns the value of property "XRSession.supportedFrameRates".
//
// It returns ok=false if there is no such property.
func (this XRSession) SupportedFrameRates() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetXRSessionSupportedFrameRates(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RenderState returns the value of property "XRSession.renderState".
//
// It returns ok=false if there is no such property.
func (this XRSession) RenderState() (ret XRRenderState, ok bool) {
	ok = js.True == bindings.GetXRSessionRenderState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InputSources returns the value of property "XRSession.inputSources".
//
// It returns ok=false if there is no such property.
func (this XRSession) InputSources() (ret XRInputSourceArray, ok bool) {
	ok = js.True == bindings.GetXRSessionInputSources(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EnabledFeatures returns the value of property "XRSession.enabledFeatures".
//
// It returns ok=false if there is no such property.
func (this XRSession) EnabledFeatures() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetXRSessionEnabledFeatures(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsSystemKeyboardSupported returns the value of property "XRSession.isSystemKeyboardSupported".
//
// It returns ok=false if there is no such property.
func (this XRSession) IsSystemKeyboardSupported() (ret bool, ok bool) {
	ok = js.True == bindings.GetXRSessionIsSystemKeyboardSupported(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EnvironmentBlendMode returns the value of property "XRSession.environmentBlendMode".
//
// It returns ok=false if there is no such property.
func (this XRSession) EnvironmentBlendMode() (ret XREnvironmentBlendMode, ok bool) {
	ok = js.True == bindings.GetXRSessionEnvironmentBlendMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InteractionMode returns the value of property "XRSession.interactionMode".
//
// It returns ok=false if there is no such property.
func (this XRSession) InteractionMode() (ret XRInteractionMode, ok bool) {
	ok = js.True == bindings.GetXRSessionInteractionMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DomOverlayState returns the value of property "XRSession.domOverlayState".
//
// It returns ok=false if there is no such property.
func (this XRSession) DomOverlayState() (ret XRDOMOverlayState, ok bool) {
	ok = js.True == bindings.GetXRSessionDomOverlayState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DepthUsage returns the value of property "XRSession.depthUsage".
//
// It returns ok=false if there is no such property.
func (this XRSession) DepthUsage() (ret XRDepthUsage, ok bool) {
	ok = js.True == bindings.GetXRSessionDepthUsage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DepthDataFormat returns the value of property "XRSession.depthDataFormat".
//
// It returns ok=false if there is no such property.
func (this XRSession) DepthDataFormat() (ret XRDepthDataFormat, ok bool) {
	ok = js.True == bindings.GetXRSessionDepthDataFormat(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PersistentAnchors returns the value of property "XRSession.persistentAnchors".
//
// It returns ok=false if there is no such property.
func (this XRSession) PersistentAnchors() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetXRSessionPersistentAnchors(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PreferredReflectionFormat returns the value of property "XRSession.preferredReflectionFormat".
//
// It returns ok=false if there is no such property.
func (this XRSession) PreferredReflectionFormat() (ret XRReflectionFormat, ok bool) {
	ok = js.True == bindings.GetXRSessionPreferredReflectionFormat(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncUpdateRenderState returns true if the method "XRSession.updateRenderState" exists.
func (this XRSession) HasFuncUpdateRenderState() bool {
	return js.True == bindings.HasFuncXRSessionUpdateRenderState(
		this.ref,
	)
}

// FuncUpdateRenderState returns the method "XRSession.updateRenderState".
func (this XRSession) FuncUpdateRenderState() (fn js.Func[func(state XRRenderStateInit)]) {
	bindings.FuncXRSessionUpdateRenderState(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpdateRenderState calls the method "XRSession.updateRenderState".
func (this XRSession) UpdateRenderState(state XRRenderStateInit) (ret js.Void) {
	bindings.CallXRSessionUpdateRenderState(
		this.ref, js.Pointer(&ret),
		js.Pointer(&state),
	)

	return
}

// TryUpdateRenderState calls the method "XRSession.updateRenderState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRSession) TryUpdateRenderState(state XRRenderStateInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRSessionUpdateRenderState(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&state),
	)

	return
}

// HasFuncUpdateRenderState1 returns true if the method "XRSession.updateRenderState" exists.
func (this XRSession) HasFuncUpdateRenderState1() bool {
	return js.True == bindings.HasFuncXRSessionUpdateRenderState1(
		this.ref,
	)
}

// FuncUpdateRenderState1 returns the method "XRSession.updateRenderState".
func (this XRSession) FuncUpdateRenderState1() (fn js.Func[func()]) {
	bindings.FuncXRSessionUpdateRenderState1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpdateRenderState1 calls the method "XRSession.updateRenderState".
func (this XRSession) UpdateRenderState1() (ret js.Void) {
	bindings.CallXRSessionUpdateRenderState1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryUpdateRenderState1 calls the method "XRSession.updateRenderState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRSession) TryUpdateRenderState1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRSessionUpdateRenderState1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncUpdateTargetFrameRate returns true if the method "XRSession.updateTargetFrameRate" exists.
func (this XRSession) HasFuncUpdateTargetFrameRate() bool {
	return js.True == bindings.HasFuncXRSessionUpdateTargetFrameRate(
		this.ref,
	)
}

// FuncUpdateTargetFrameRate returns the method "XRSession.updateTargetFrameRate".
func (this XRSession) FuncUpdateTargetFrameRate() (fn js.Func[func(rate float32) js.Promise[js.Void]]) {
	bindings.FuncXRSessionUpdateTargetFrameRate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpdateTargetFrameRate calls the method "XRSession.updateTargetFrameRate".
func (this XRSession) UpdateTargetFrameRate(rate float32) (ret js.Promise[js.Void]) {
	bindings.CallXRSessionUpdateTargetFrameRate(
		this.ref, js.Pointer(&ret),
		float32(rate),
	)

	return
}

// TryUpdateTargetFrameRate calls the method "XRSession.updateTargetFrameRate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRSession) TryUpdateTargetFrameRate(rate float32) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRSessionUpdateTargetFrameRate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(rate),
	)

	return
}

// HasFuncRequestReferenceSpace returns true if the method "XRSession.requestReferenceSpace" exists.
func (this XRSession) HasFuncRequestReferenceSpace() bool {
	return js.True == bindings.HasFuncXRSessionRequestReferenceSpace(
		this.ref,
	)
}

// FuncRequestReferenceSpace returns the method "XRSession.requestReferenceSpace".
func (this XRSession) FuncRequestReferenceSpace() (fn js.Func[func(typ XRReferenceSpaceType) js.Promise[XRReferenceSpace]]) {
	bindings.FuncXRSessionRequestReferenceSpace(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestReferenceSpace calls the method "XRSession.requestReferenceSpace".
func (this XRSession) RequestReferenceSpace(typ XRReferenceSpaceType) (ret js.Promise[XRReferenceSpace]) {
	bindings.CallXRSessionRequestReferenceSpace(
		this.ref, js.Pointer(&ret),
		uint32(typ),
	)

	return
}

// TryRequestReferenceSpace calls the method "XRSession.requestReferenceSpace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRSession) TryRequestReferenceSpace(typ XRReferenceSpaceType) (ret js.Promise[XRReferenceSpace], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRSessionRequestReferenceSpace(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
	)

	return
}

// HasFuncRequestAnimationFrame returns true if the method "XRSession.requestAnimationFrame" exists.
func (this XRSession) HasFuncRequestAnimationFrame() bool {
	return js.True == bindings.HasFuncXRSessionRequestAnimationFrame(
		this.ref,
	)
}

// FuncRequestAnimationFrame returns the method "XRSession.requestAnimationFrame".
func (this XRSession) FuncRequestAnimationFrame() (fn js.Func[func(callback js.Func[func(time DOMHighResTimeStamp, frame XRFrame)]) uint32]) {
	bindings.FuncXRSessionRequestAnimationFrame(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestAnimationFrame calls the method "XRSession.requestAnimationFrame".
func (this XRSession) RequestAnimationFrame(callback js.Func[func(time DOMHighResTimeStamp, frame XRFrame)]) (ret uint32) {
	bindings.CallXRSessionRequestAnimationFrame(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryRequestAnimationFrame calls the method "XRSession.requestAnimationFrame"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRSession) TryRequestAnimationFrame(callback js.Func[func(time DOMHighResTimeStamp, frame XRFrame)]) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRSessionRequestAnimationFrame(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncCancelAnimationFrame returns true if the method "XRSession.cancelAnimationFrame" exists.
func (this XRSession) HasFuncCancelAnimationFrame() bool {
	return js.True == bindings.HasFuncXRSessionCancelAnimationFrame(
		this.ref,
	)
}

// FuncCancelAnimationFrame returns the method "XRSession.cancelAnimationFrame".
func (this XRSession) FuncCancelAnimationFrame() (fn js.Func[func(handle uint32)]) {
	bindings.FuncXRSessionCancelAnimationFrame(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CancelAnimationFrame calls the method "XRSession.cancelAnimationFrame".
func (this XRSession) CancelAnimationFrame(handle uint32) (ret js.Void) {
	bindings.CallXRSessionCancelAnimationFrame(
		this.ref, js.Pointer(&ret),
		uint32(handle),
	)

	return
}

// TryCancelAnimationFrame calls the method "XRSession.cancelAnimationFrame"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRSession) TryCancelAnimationFrame(handle uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRSessionCancelAnimationFrame(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(handle),
	)

	return
}

// HasFuncEnd returns true if the method "XRSession.end" exists.
func (this XRSession) HasFuncEnd() bool {
	return js.True == bindings.HasFuncXRSessionEnd(
		this.ref,
	)
}

// FuncEnd returns the method "XRSession.end".
func (this XRSession) FuncEnd() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncXRSessionEnd(
		this.ref, js.Pointer(&fn),
	)
	return
}

// End calls the method "XRSession.end".
func (this XRSession) End() (ret js.Promise[js.Void]) {
	bindings.CallXRSessionEnd(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryEnd calls the method "XRSession.end"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRSession) TryEnd() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRSessionEnd(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRestorePersistentAnchor returns true if the method "XRSession.restorePersistentAnchor" exists.
func (this XRSession) HasFuncRestorePersistentAnchor() bool {
	return js.True == bindings.HasFuncXRSessionRestorePersistentAnchor(
		this.ref,
	)
}

// FuncRestorePersistentAnchor returns the method "XRSession.restorePersistentAnchor".
func (this XRSession) FuncRestorePersistentAnchor() (fn js.Func[func(uuid js.String) js.Promise[XRAnchor]]) {
	bindings.FuncXRSessionRestorePersistentAnchor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RestorePersistentAnchor calls the method "XRSession.restorePersistentAnchor".
func (this XRSession) RestorePersistentAnchor(uuid js.String) (ret js.Promise[XRAnchor]) {
	bindings.CallXRSessionRestorePersistentAnchor(
		this.ref, js.Pointer(&ret),
		uuid.Ref(),
	)

	return
}

// TryRestorePersistentAnchor calls the method "XRSession.restorePersistentAnchor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRSession) TryRestorePersistentAnchor(uuid js.String) (ret js.Promise[XRAnchor], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRSessionRestorePersistentAnchor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uuid.Ref(),
	)

	return
}

// HasFuncDeletePersistentAnchor returns true if the method "XRSession.deletePersistentAnchor" exists.
func (this XRSession) HasFuncDeletePersistentAnchor() bool {
	return js.True == bindings.HasFuncXRSessionDeletePersistentAnchor(
		this.ref,
	)
}

// FuncDeletePersistentAnchor returns the method "XRSession.deletePersistentAnchor".
func (this XRSession) FuncDeletePersistentAnchor() (fn js.Func[func(uuid js.String) js.Promise[js.Void]]) {
	bindings.FuncXRSessionDeletePersistentAnchor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeletePersistentAnchor calls the method "XRSession.deletePersistentAnchor".
func (this XRSession) DeletePersistentAnchor(uuid js.String) (ret js.Promise[js.Void]) {
	bindings.CallXRSessionDeletePersistentAnchor(
		this.ref, js.Pointer(&ret),
		uuid.Ref(),
	)

	return
}

// TryDeletePersistentAnchor calls the method "XRSession.deletePersistentAnchor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRSession) TryDeletePersistentAnchor(uuid js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRSessionDeletePersistentAnchor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uuid.Ref(),
	)

	return
}

// HasFuncRequestLightProbe returns true if the method "XRSession.requestLightProbe" exists.
func (this XRSession) HasFuncRequestLightProbe() bool {
	return js.True == bindings.HasFuncXRSessionRequestLightProbe(
		this.ref,
	)
}

// FuncRequestLightProbe returns the method "XRSession.requestLightProbe".
func (this XRSession) FuncRequestLightProbe() (fn js.Func[func(options XRLightProbeInit) js.Promise[XRLightProbe]]) {
	bindings.FuncXRSessionRequestLightProbe(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestLightProbe calls the method "XRSession.requestLightProbe".
func (this XRSession) RequestLightProbe(options XRLightProbeInit) (ret js.Promise[XRLightProbe]) {
	bindings.CallXRSessionRequestLightProbe(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRequestLightProbe calls the method "XRSession.requestLightProbe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRSession) TryRequestLightProbe(options XRLightProbeInit) (ret js.Promise[XRLightProbe], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRSessionRequestLightProbe(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRequestLightProbe1 returns true if the method "XRSession.requestLightProbe" exists.
func (this XRSession) HasFuncRequestLightProbe1() bool {
	return js.True == bindings.HasFuncXRSessionRequestLightProbe1(
		this.ref,
	)
}

// FuncRequestLightProbe1 returns the method "XRSession.requestLightProbe".
func (this XRSession) FuncRequestLightProbe1() (fn js.Func[func() js.Promise[XRLightProbe]]) {
	bindings.FuncXRSessionRequestLightProbe1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestLightProbe1 calls the method "XRSession.requestLightProbe".
func (this XRSession) RequestLightProbe1() (ret js.Promise[XRLightProbe]) {
	bindings.CallXRSessionRequestLightProbe1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestLightProbe1 calls the method "XRSession.requestLightProbe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRSession) TryRequestLightProbe1() (ret js.Promise[XRLightProbe], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRSessionRequestLightProbe1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestHitTestSource returns true if the method "XRSession.requestHitTestSource" exists.
func (this XRSession) HasFuncRequestHitTestSource() bool {
	return js.True == bindings.HasFuncXRSessionRequestHitTestSource(
		this.ref,
	)
}

// FuncRequestHitTestSource returns the method "XRSession.requestHitTestSource".
func (this XRSession) FuncRequestHitTestSource() (fn js.Func[func(options XRHitTestOptionsInit) js.Promise[XRHitTestSource]]) {
	bindings.FuncXRSessionRequestHitTestSource(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestHitTestSource calls the method "XRSession.requestHitTestSource".
func (this XRSession) RequestHitTestSource(options XRHitTestOptionsInit) (ret js.Promise[XRHitTestSource]) {
	bindings.CallXRSessionRequestHitTestSource(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRequestHitTestSource calls the method "XRSession.requestHitTestSource"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRSession) TryRequestHitTestSource(options XRHitTestOptionsInit) (ret js.Promise[XRHitTestSource], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRSessionRequestHitTestSource(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRequestHitTestSourceForTransientInput returns true if the method "XRSession.requestHitTestSourceForTransientInput" exists.
func (this XRSession) HasFuncRequestHitTestSourceForTransientInput() bool {
	return js.True == bindings.HasFuncXRSessionRequestHitTestSourceForTransientInput(
		this.ref,
	)
}

// FuncRequestHitTestSourceForTransientInput returns the method "XRSession.requestHitTestSourceForTransientInput".
func (this XRSession) FuncRequestHitTestSourceForTransientInput() (fn js.Func[func(options XRTransientInputHitTestOptionsInit) js.Promise[XRTransientInputHitTestSource]]) {
	bindings.FuncXRSessionRequestHitTestSourceForTransientInput(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestHitTestSourceForTransientInput calls the method "XRSession.requestHitTestSourceForTransientInput".
func (this XRSession) RequestHitTestSourceForTransientInput(options XRTransientInputHitTestOptionsInit) (ret js.Promise[XRTransientInputHitTestSource]) {
	bindings.CallXRSessionRequestHitTestSourceForTransientInput(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRequestHitTestSourceForTransientInput calls the method "XRSession.requestHitTestSourceForTransientInput"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRSession) TryRequestHitTestSourceForTransientInput(options XRTransientInputHitTestOptionsInit) (ret js.Promise[XRTransientInputHitTestSource], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRSessionRequestHitTestSourceForTransientInput(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

type XRDOMOverlayInit struct {
	// Root is "XRDOMOverlayInit.root"
	//
	// Required
	Root Element

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRDOMOverlayInit with all fields set.
func (p XRDOMOverlayInit) FromRef(ref js.Ref) XRDOMOverlayInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRDOMOverlayInit in the application heap.
func (p XRDOMOverlayInit) New() js.Ref {
	return bindings.XRDOMOverlayInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRDOMOverlayInit) UpdateFrom(ref js.Ref) {
	bindings.XRDOMOverlayInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRDOMOverlayInit) Update(ref js.Ref) {
	bindings.XRDOMOverlayInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRDOMOverlayInit) FreeMembers(recursive bool) {
	js.Free(
		p.Root.Ref(),
	)
	p.Root = p.Root.FromRef(js.Undefined)
}

type XRDepthStateInit struct {
	// UsagePreference is "XRDepthStateInit.usagePreference"
	//
	// Required
	UsagePreference js.Array[XRDepthUsage]
	// DataFormatPreference is "XRDepthStateInit.dataFormatPreference"
	//
	// Required
	DataFormatPreference js.Array[XRDepthDataFormat]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRDepthStateInit with all fields set.
func (p XRDepthStateInit) FromRef(ref js.Ref) XRDepthStateInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRDepthStateInit in the application heap.
func (p XRDepthStateInit) New() js.Ref {
	return bindings.XRDepthStateInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRDepthStateInit) UpdateFrom(ref js.Ref) {
	bindings.XRDepthStateInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRDepthStateInit) Update(ref js.Ref) {
	bindings.XRDepthStateInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRDepthStateInit) FreeMembers(recursive bool) {
	js.Free(
		p.UsagePreference.Ref(),
		p.DataFormatPreference.Ref(),
	)
	p.UsagePreference = p.UsagePreference.FromRef(js.Undefined)
	p.DataFormatPreference = p.DataFormatPreference.FromRef(js.Undefined)
}

type XRSessionInit struct {
	// RequiredFeatures is "XRSessionInit.requiredFeatures"
	//
	// Optional
	RequiredFeatures js.Array[js.String]
	// OptionalFeatures is "XRSessionInit.optionalFeatures"
	//
	// Optional
	OptionalFeatures js.Array[js.String]
	// DomOverlay is "XRSessionInit.domOverlay"
	//
	// Optional
	//
	// NOTE: DomOverlay.FFI_USE MUST be set to true to get DomOverlay used.
	DomOverlay XRDOMOverlayInit
	// DepthSensing is "XRSessionInit.depthSensing"
	//
	// Optional
	//
	// NOTE: DepthSensing.FFI_USE MUST be set to true to get DepthSensing used.
	DepthSensing XRDepthStateInit

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRSessionInit with all fields set.
func (p XRSessionInit) FromRef(ref js.Ref) XRSessionInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRSessionInit in the application heap.
func (p XRSessionInit) New() js.Ref {
	return bindings.XRSessionInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRSessionInit) UpdateFrom(ref js.Ref) {
	bindings.XRSessionInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRSessionInit) Update(ref js.Ref) {
	bindings.XRSessionInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRSessionInit) FreeMembers(recursive bool) {
	js.Free(
		p.RequiredFeatures.Ref(),
		p.OptionalFeatures.Ref(),
	)
	p.RequiredFeatures = p.RequiredFeatures.FromRef(js.Undefined)
	p.OptionalFeatures = p.OptionalFeatures.FromRef(js.Undefined)
	if recursive {
		p.DomOverlay.FreeMembers(true)
		p.DepthSensing.FreeMembers(true)
	}
}

type XRSystem struct {
	EventTarget
}

func (this XRSystem) Once() XRSystem {
	this.ref.Once()
	return this
}

func (this XRSystem) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this XRSystem) FromRef(ref js.Ref) XRSystem {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this XRSystem) Free() {
	this.ref.Free()
}

// HasFuncIsSessionSupported returns true if the method "XRSystem.isSessionSupported" exists.
func (this XRSystem) HasFuncIsSessionSupported() bool {
	return js.True == bindings.HasFuncXRSystemIsSessionSupported(
		this.ref,
	)
}

// FuncIsSessionSupported returns the method "XRSystem.isSessionSupported".
func (this XRSystem) FuncIsSessionSupported() (fn js.Func[func(mode XRSessionMode) js.Promise[js.Boolean]]) {
	bindings.FuncXRSystemIsSessionSupported(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsSessionSupported calls the method "XRSystem.isSessionSupported".
func (this XRSystem) IsSessionSupported(mode XRSessionMode) (ret js.Promise[js.Boolean]) {
	bindings.CallXRSystemIsSessionSupported(
		this.ref, js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryIsSessionSupported calls the method "XRSystem.isSessionSupported"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRSystem) TryIsSessionSupported(mode XRSessionMode) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRSystemIsSessionSupported(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

// HasFuncRequestSession returns true if the method "XRSystem.requestSession" exists.
func (this XRSystem) HasFuncRequestSession() bool {
	return js.True == bindings.HasFuncXRSystemRequestSession(
		this.ref,
	)
}

// FuncRequestSession returns the method "XRSystem.requestSession".
func (this XRSystem) FuncRequestSession() (fn js.Func[func(mode XRSessionMode, options XRSessionInit) js.Promise[XRSession]]) {
	bindings.FuncXRSystemRequestSession(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestSession calls the method "XRSystem.requestSession".
func (this XRSystem) RequestSession(mode XRSessionMode, options XRSessionInit) (ret js.Promise[XRSession]) {
	bindings.CallXRSystemRequestSession(
		this.ref, js.Pointer(&ret),
		uint32(mode),
		js.Pointer(&options),
	)

	return
}

// TryRequestSession calls the method "XRSystem.requestSession"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRSystem) TryRequestSession(mode XRSessionMode, options XRSessionInit) (ret js.Promise[XRSession], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRSystemRequestSession(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		js.Pointer(&options),
	)

	return
}

// HasFuncRequestSession1 returns true if the method "XRSystem.requestSession" exists.
func (this XRSystem) HasFuncRequestSession1() bool {
	return js.True == bindings.HasFuncXRSystemRequestSession1(
		this.ref,
	)
}

// FuncRequestSession1 returns the method "XRSystem.requestSession".
func (this XRSystem) FuncRequestSession1() (fn js.Func[func(mode XRSessionMode) js.Promise[XRSession]]) {
	bindings.FuncXRSystemRequestSession1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestSession1 calls the method "XRSystem.requestSession".
func (this XRSystem) RequestSession1(mode XRSessionMode) (ret js.Promise[XRSession]) {
	bindings.CallXRSystemRequestSession1(
		this.ref, js.Pointer(&ret),
		uint32(mode),
	)

	return
}

// TryRequestSession1 calls the method "XRSystem.requestSession"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRSystem) TryRequestSession1(mode XRSessionMode) (ret js.Promise[XRSession], exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRSystemRequestSession1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
	)

	return
}

type NotificationDirection uint32

const (
	_ NotificationDirection = iota

	NotificationDirection_AUTO
	NotificationDirection_LTR
	NotificationDirection_RTL
)

func (NotificationDirection) FromRef(str js.Ref) NotificationDirection {
	return NotificationDirection(bindings.ConstOfNotificationDirection(str))
}

func (x NotificationDirection) String() (string, bool) {
	switch x {
	case NotificationDirection_AUTO:
		return "auto", true
	case NotificationDirection_LTR:
		return "ltr", true
	case NotificationDirection_RTL:
		return "rtl", true
	default:
		return "", false
	}
}

type NotificationAction struct {
	// Action is "NotificationAction.action"
	//
	// Required
	Action js.String
	// Title is "NotificationAction.title"
	//
	// Required
	Title js.String
	// Icon is "NotificationAction.icon"
	//
	// Optional
	Icon js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NotificationAction with all fields set.
func (p NotificationAction) FromRef(ref js.Ref) NotificationAction {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NotificationAction in the application heap.
func (p NotificationAction) New() js.Ref {
	return bindings.NotificationActionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NotificationAction) UpdateFrom(ref js.Ref) {
	bindings.NotificationActionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NotificationAction) Update(ref js.Ref) {
	bindings.NotificationActionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NotificationAction) FreeMembers(recursive bool) {
	js.Free(
		p.Action.Ref(),
		p.Title.Ref(),
		p.Icon.Ref(),
	)
	p.Action = p.Action.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Icon = p.Icon.FromRef(js.Undefined)
}

type NotificationOptions struct {
	// Dir is "NotificationOptions.dir"
	//
	// Optional, defaults to "auto".
	Dir NotificationDirection
	// Lang is "NotificationOptions.lang"
	//
	// Optional, defaults to "".
	Lang js.String
	// Body is "NotificationOptions.body"
	//
	// Optional, defaults to "".
	Body js.String
	// Tag is "NotificationOptions.tag"
	//
	// Optional, defaults to "".
	Tag js.String
	// Image is "NotificationOptions.image"
	//
	// Optional
	Image js.String
	// Icon is "NotificationOptions.icon"
	//
	// Optional
	Icon js.String
	// Badge is "NotificationOptions.badge"
	//
	// Optional
	Badge js.String
	// Vibrate is "NotificationOptions.vibrate"
	//
	// Optional
	Vibrate VibratePattern
	// Timestamp is "NotificationOptions.timestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timestamp MUST be set to true to make this field effective.
	Timestamp EpochTimeStamp
	// Renotify is "NotificationOptions.renotify"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Renotify MUST be set to true to make this field effective.
	Renotify bool
	// Silent is "NotificationOptions.silent"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_Silent MUST be set to true to make this field effective.
	Silent bool
	// RequireInteraction is "NotificationOptions.requireInteraction"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_RequireInteraction MUST be set to true to make this field effective.
	RequireInteraction bool
	// Data is "NotificationOptions.data"
	//
	// Optional, defaults to null.
	Data js.Any
	// Actions is "NotificationOptions.actions"
	//
	// Optional, defaults to [].
	Actions js.Array[NotificationAction]

	FFI_USE_Timestamp          bool // for Timestamp.
	FFI_USE_Renotify           bool // for Renotify.
	FFI_USE_Silent             bool // for Silent.
	FFI_USE_RequireInteraction bool // for RequireInteraction.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NotificationOptions with all fields set.
func (p NotificationOptions) FromRef(ref js.Ref) NotificationOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NotificationOptions in the application heap.
func (p NotificationOptions) New() js.Ref {
	return bindings.NotificationOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NotificationOptions) UpdateFrom(ref js.Ref) {
	bindings.NotificationOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NotificationOptions) Update(ref js.Ref) {
	bindings.NotificationOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NotificationOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Lang.Ref(),
		p.Body.Ref(),
		p.Tag.Ref(),
		p.Image.Ref(),
		p.Icon.Ref(),
		p.Badge.Ref(),
		p.Vibrate.Ref(),
		p.Data.Ref(),
		p.Actions.Ref(),
	)
	p.Lang = p.Lang.FromRef(js.Undefined)
	p.Body = p.Body.FromRef(js.Undefined)
	p.Tag = p.Tag.FromRef(js.Undefined)
	p.Image = p.Image.FromRef(js.Undefined)
	p.Icon = p.Icon.FromRef(js.Undefined)
	p.Badge = p.Badge.FromRef(js.Undefined)
	p.Vibrate = p.Vibrate.FromRef(js.Undefined)
	p.Data = p.Data.FromRef(js.Undefined)
	p.Actions = p.Actions.FromRef(js.Undefined)
}

type NotificationPermission uint32

const (
	_ NotificationPermission = iota

	NotificationPermission_DEFAULT
	NotificationPermission_DENIED
	NotificationPermission_GRANTED
)

func (NotificationPermission) FromRef(str js.Ref) NotificationPermission {
	return NotificationPermission(bindings.ConstOfNotificationPermission(str))
}

func (x NotificationPermission) String() (string, bool) {
	switch x {
	case NotificationPermission_DEFAULT:
		return "default", true
	case NotificationPermission_DENIED:
		return "denied", true
	case NotificationPermission_GRANTED:
		return "granted", true
	default:
		return "", false
	}
}

type NotificationPermissionCallbackFunc func(this js.Ref, permission NotificationPermission) js.Ref

func (fn NotificationPermissionCallbackFunc) Register() js.Func[func(permission NotificationPermission)] {
	return js.RegisterCallback[func(permission NotificationPermission)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn NotificationPermissionCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		NotificationPermission(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type NotificationPermissionCallback[T any] struct {
	Fn  func(arg T, this js.Ref, permission NotificationPermission) js.Ref
	Arg T
}

func (cb *NotificationPermissionCallback[T]) Register() js.Func[func(permission NotificationPermission)] {
	return js.RegisterCallback[func(permission NotificationPermission)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *NotificationPermissionCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		NotificationPermission(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

func NewNotification(title js.String, options NotificationOptions) (ret Notification) {
	ret.ref = bindings.NewNotificationByNotification(
		title.Ref(),
		js.Pointer(&options))
	return
}

func NewNotificationByNotification1(title js.String) (ret Notification) {
	ret.ref = bindings.NewNotificationByNotification1(
		title.Ref())
	return
}

type Notification struct {
	EventTarget
}

func (this Notification) Once() Notification {
	this.ref.Once()
	return this
}

func (this Notification) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this Notification) FromRef(ref js.Ref) Notification {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this Notification) Free() {
	this.ref.Free()
}

// Permission returns the value of property "Notification.permission".
//
// It returns ok=false if there is no such property.
func (this Notification) Permission() (ret NotificationPermission, ok bool) {
	ok = js.True == bindings.GetNotificationPermission(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxActions returns the value of property "Notification.maxActions".
//
// It returns ok=false if there is no such property.
func (this Notification) MaxActions() (ret uint32, ok bool) {
	ok = js.True == bindings.GetNotificationMaxActions(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Title returns the value of property "Notification.title".
//
// It returns ok=false if there is no such property.
func (this Notification) Title() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNotificationTitle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Dir returns the value of property "Notification.dir".
//
// It returns ok=false if there is no such property.
func (this Notification) Dir() (ret NotificationDirection, ok bool) {
	ok = js.True == bindings.GetNotificationDir(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Lang returns the value of property "Notification.lang".
//
// It returns ok=false if there is no such property.
func (this Notification) Lang() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNotificationLang(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Body returns the value of property "Notification.body".
//
// It returns ok=false if there is no such property.
func (this Notification) Body() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNotificationBody(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Tag returns the value of property "Notification.tag".
//
// It returns ok=false if there is no such property.
func (this Notification) Tag() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNotificationTag(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Image returns the value of property "Notification.image".
//
// It returns ok=false if there is no such property.
func (this Notification) Image() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNotificationImage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Icon returns the value of property "Notification.icon".
//
// It returns ok=false if there is no such property.
func (this Notification) Icon() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNotificationIcon(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Badge returns the value of property "Notification.badge".
//
// It returns ok=false if there is no such property.
func (this Notification) Badge() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNotificationBadge(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Vibrate returns the value of property "Notification.vibrate".
//
// It returns ok=false if there is no such property.
func (this Notification) Vibrate() (ret js.FrozenArray[uint32], ok bool) {
	ok = js.True == bindings.GetNotificationVibrate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Timestamp returns the value of property "Notification.timestamp".
//
// It returns ok=false if there is no such property.
func (this Notification) Timestamp() (ret EpochTimeStamp, ok bool) {
	ok = js.True == bindings.GetNotificationTimestamp(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Renotify returns the value of property "Notification.renotify".
//
// It returns ok=false if there is no such property.
func (this Notification) Renotify() (ret bool, ok bool) {
	ok = js.True == bindings.GetNotificationRenotify(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Silent returns the value of property "Notification.silent".
//
// It returns ok=false if there is no such property.
func (this Notification) Silent() (ret bool, ok bool) {
	ok = js.True == bindings.GetNotificationSilent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RequireInteraction returns the value of property "Notification.requireInteraction".
//
// It returns ok=false if there is no such property.
func (this Notification) RequireInteraction() (ret bool, ok bool) {
	ok = js.True == bindings.GetNotificationRequireInteraction(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Data returns the value of property "Notification.data".
//
// It returns ok=false if there is no such property.
func (this Notification) Data() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetNotificationData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Actions returns the value of property "Notification.actions".
//
// It returns ok=false if there is no such property.
func (this Notification) Actions() (ret js.FrozenArray[NotificationAction], ok bool) {
	ok = js.True == bindings.GetNotificationActions(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncRequestPermission returns true if the static method "Notification.requestPermission" exists.
func (this Notification) HasFuncRequestPermission() bool {
	return js.True == bindings.HasFuncNotificationRequestPermission(
		this.ref,
	)
}

// FuncRequestPermission returns the static method "Notification.requestPermission".
func (this Notification) FuncRequestPermission() (fn js.Func[func(deprecatedCallback js.Func[func(permission NotificationPermission)]) js.Promise[NotificationPermission]]) {
	bindings.FuncNotificationRequestPermission(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestPermission calls the static method "Notification.requestPermission".
func (this Notification) RequestPermission(deprecatedCallback js.Func[func(permission NotificationPermission)]) (ret js.Promise[NotificationPermission]) {
	bindings.CallNotificationRequestPermission(
		this.ref, js.Pointer(&ret),
		deprecatedCallback.Ref(),
	)

	return
}

// TryRequestPermission calls the static method "Notification.requestPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Notification) TryRequestPermission(deprecatedCallback js.Func[func(permission NotificationPermission)]) (ret js.Promise[NotificationPermission], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNotificationRequestPermission(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		deprecatedCallback.Ref(),
	)

	return
}

// HasFuncRequestPermission1 returns true if the static method "Notification.requestPermission" exists.
func (this Notification) HasFuncRequestPermission1() bool {
	return js.True == bindings.HasFuncNotificationRequestPermission1(
		this.ref,
	)
}

// FuncRequestPermission1 returns the static method "Notification.requestPermission".
func (this Notification) FuncRequestPermission1() (fn js.Func[func() js.Promise[NotificationPermission]]) {
	bindings.FuncNotificationRequestPermission1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestPermission1 calls the static method "Notification.requestPermission".
func (this Notification) RequestPermission1() (ret js.Promise[NotificationPermission]) {
	bindings.CallNotificationRequestPermission1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestPermission1 calls the static method "Notification.requestPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Notification) TryRequestPermission1() (ret js.Promise[NotificationPermission], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNotificationRequestPermission1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClose returns true if the method "Notification.close" exists.
func (this Notification) HasFuncClose() bool {
	return js.True == bindings.HasFuncNotificationClose(
		this.ref,
	)
}

// FuncClose returns the method "Notification.close".
func (this Notification) FuncClose() (fn js.Func[func()]) {
	bindings.FuncNotificationClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "Notification.close".
func (this Notification) Close() (ret js.Void) {
	bindings.CallNotificationClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "Notification.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Notification) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNotificationClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type GetNotificationOptions struct {
	// Tag is "GetNotificationOptions.tag"
	//
	// Optional, defaults to "".
	Tag js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetNotificationOptions with all fields set.
func (p GetNotificationOptions) FromRef(ref js.Ref) GetNotificationOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetNotificationOptions in the application heap.
func (p GetNotificationOptions) New() js.Ref {
	return bindings.GetNotificationOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetNotificationOptions) UpdateFrom(ref js.Ref) {
	bindings.GetNotificationOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetNotificationOptions) Update(ref js.Ref) {
	bindings.GetNotificationOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetNotificationOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Tag.Ref(),
	)
	p.Tag = p.Tag.FromRef(js.Undefined)
}

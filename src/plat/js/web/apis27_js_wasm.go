// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

type XRHitTestSource struct {
	ref js.Ref
}

func (this XRHitTestSource) Once() XRHitTestSource {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Cancel calls the method "XRHitTestSource.cancel".
//
// The returned bool will be false if there is no such method.
func (this XRHitTestSource) Cancel() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXRHitTestSourceCancel(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CancelFunc returns the method "XRHitTestSource.cancel".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRHitTestSource) CancelFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.XRHitTestSourceCancelFunc(
			this.Ref(),
		),
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Size returns the value of property "XRHand.size".
//
// The returned bool will be false if there is no such property.
func (this XRHand) Size() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRHandSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Get calls the method "XRHand.get".
//
// The returned bool will be false if there is no such method.
func (this XRHand) Get(key XRHandJoint) (XRJointSpace, bool) {
	var _ok bool
	_ret := bindings.CallXRHandGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(key),
	)

	return XRJointSpace{}.FromRef(_ret), _ok
}

// GetFunc returns the method "XRHand.get".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRHand) GetFunc() (fn js.Func[func(key XRHandJoint) XRJointSpace]) {
	return fn.FromRef(
		bindings.XRHandGetFunc(
			this.Ref(),
		),
	)
}

type XRInputSource struct {
	ref js.Ref
}

func (this XRInputSource) Once() XRInputSource {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Handedness returns the value of property "XRInputSource.handedness".
//
// The returned bool will be false if there is no such property.
func (this XRInputSource) Handedness() (XRHandedness, bool) {
	var _ok bool
	_ret := bindings.GetXRInputSourceHandedness(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRHandedness(_ret), _ok
}

// TargetRayMode returns the value of property "XRInputSource.targetRayMode".
//
// The returned bool will be false if there is no such property.
func (this XRInputSource) TargetRayMode() (XRTargetRayMode, bool) {
	var _ok bool
	_ret := bindings.GetXRInputSourceTargetRayMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRTargetRayMode(_ret), _ok
}

// TargetRaySpace returns the value of property "XRInputSource.targetRaySpace".
//
// The returned bool will be false if there is no such property.
func (this XRInputSource) TargetRaySpace() (XRSpace, bool) {
	var _ok bool
	_ret := bindings.GetXRInputSourceTargetRaySpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRSpace{}.FromRef(_ret), _ok
}

// GripSpace returns the value of property "XRInputSource.gripSpace".
//
// The returned bool will be false if there is no such property.
func (this XRInputSource) GripSpace() (XRSpace, bool) {
	var _ok bool
	_ret := bindings.GetXRInputSourceGripSpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRSpace{}.FromRef(_ret), _ok
}

// Profiles returns the value of property "XRInputSource.profiles".
//
// The returned bool will be false if there is no such property.
func (this XRInputSource) Profiles() (js.FrozenArray[js.String], bool) {
	var _ok bool
	_ret := bindings.GetXRInputSourceProfiles(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[js.String]{}.FromRef(_ret), _ok
}

// Gamepad returns the value of property "XRInputSource.gamepad".
//
// The returned bool will be false if there is no such property.
func (this XRInputSource) Gamepad() (Gamepad, bool) {
	var _ok bool
	_ret := bindings.GetXRInputSourceGamepad(
		this.Ref(), js.Pointer(&_ok),
	)
	return Gamepad{}.FromRef(_ret), _ok
}

// Hand returns the value of property "XRInputSource.hand".
//
// The returned bool will be false if there is no such property.
func (this XRInputSource) Hand() (XRHand, bool) {
	var _ok bool
	_ret := bindings.GetXRInputSourceHand(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRHand{}.FromRef(_ret), _ok
}

type XRTransientInputHitTestResult struct {
	ref js.Ref
}

func (this XRTransientInputHitTestResult) Once() XRTransientInputHitTestResult {
	this.Ref().Once()
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
	this.Ref().Free()
}

// InputSource returns the value of property "XRTransientInputHitTestResult.inputSource".
//
// The returned bool will be false if there is no such property.
func (this XRTransientInputHitTestResult) InputSource() (XRInputSource, bool) {
	var _ok bool
	_ret := bindings.GetXRTransientInputHitTestResultInputSource(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRInputSource{}.FromRef(_ret), _ok
}

// Results returns the value of property "XRTransientInputHitTestResult.results".
//
// The returned bool will be false if there is no such property.
func (this XRTransientInputHitTestResult) Results() (js.FrozenArray[XRHitTestResult], bool) {
	var _ok bool
	_ret := bindings.GetXRTransientInputHitTestResultResults(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[XRHitTestResult]{}.FromRef(_ret), _ok
}

type XRTransientInputHitTestSource struct {
	ref js.Ref
}

func (this XRTransientInputHitTestSource) Once() XRTransientInputHitTestSource {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Cancel calls the method "XRTransientInputHitTestSource.cancel".
//
// The returned bool will be false if there is no such method.
func (this XRTransientInputHitTestSource) Cancel() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXRTransientInputHitTestSourceCancel(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CancelFunc returns the method "XRTransientInputHitTestSource.cancel".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRTransientInputHitTestSource) CancelFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.XRTransientInputHitTestSourceCancelFunc(
			this.Ref(),
		),
	)
}

type XRMeshSet struct {
	ref js.Ref
}

func (this XRMeshSet) Once() XRMeshSet {
	this.Ref().Once()
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
	this.Ref().Free()
}

type XRAnchorSet struct {
	ref js.Ref
}

func (this XRAnchorSet) Once() XRAnchorSet {
	this.Ref().Once()
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
	this.Ref().Free()
}

type XRFrame struct {
	ref js.Ref
}

func (this XRFrame) Once() XRFrame {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Session returns the value of property "XRFrame.session".
//
// The returned bool will be false if there is no such property.
func (this XRFrame) Session() (XRSession, bool) {
	var _ok bool
	_ret := bindings.GetXRFrameSession(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRSession{}.FromRef(_ret), _ok
}

// PredictedDisplayTime returns the value of property "XRFrame.predictedDisplayTime".
//
// The returned bool will be false if there is no such property.
func (this XRFrame) PredictedDisplayTime() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetXRFramePredictedDisplayTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// DetectedMeshes returns the value of property "XRFrame.detectedMeshes".
//
// The returned bool will be false if there is no such property.
func (this XRFrame) DetectedMeshes() (XRMeshSet, bool) {
	var _ok bool
	_ret := bindings.GetXRFrameDetectedMeshes(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRMeshSet{}.FromRef(_ret), _ok
}

// TrackedAnchors returns the value of property "XRFrame.trackedAnchors".
//
// The returned bool will be false if there is no such property.
func (this XRFrame) TrackedAnchors() (XRAnchorSet, bool) {
	var _ok bool
	_ret := bindings.GetXRFrameTrackedAnchors(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRAnchorSet{}.FromRef(_ret), _ok
}

// GetViewerPose calls the method "XRFrame.getViewerPose".
//
// The returned bool will be false if there is no such method.
func (this XRFrame) GetViewerPose(referenceSpace XRReferenceSpace) (XRViewerPose, bool) {
	var _ok bool
	_ret := bindings.CallXRFrameGetViewerPose(
		this.Ref(), js.Pointer(&_ok),
		referenceSpace.Ref(),
	)

	return XRViewerPose{}.FromRef(_ret), _ok
}

// GetViewerPoseFunc returns the method "XRFrame.getViewerPose".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRFrame) GetViewerPoseFunc() (fn js.Func[func(referenceSpace XRReferenceSpace) XRViewerPose]) {
	return fn.FromRef(
		bindings.XRFrameGetViewerPoseFunc(
			this.Ref(),
		),
	)
}

// GetPose calls the method "XRFrame.getPose".
//
// The returned bool will be false if there is no such method.
func (this XRFrame) GetPose(space XRSpace, baseSpace XRSpace) (XRPose, bool) {
	var _ok bool
	_ret := bindings.CallXRFrameGetPose(
		this.Ref(), js.Pointer(&_ok),
		space.Ref(),
		baseSpace.Ref(),
	)

	return XRPose{}.FromRef(_ret), _ok
}

// GetPoseFunc returns the method "XRFrame.getPose".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRFrame) GetPoseFunc() (fn js.Func[func(space XRSpace, baseSpace XRSpace) XRPose]) {
	return fn.FromRef(
		bindings.XRFrameGetPoseFunc(
			this.Ref(),
		),
	)
}

// CreateAnchor calls the method "XRFrame.createAnchor".
//
// The returned bool will be false if there is no such method.
func (this XRFrame) CreateAnchor(pose XRRigidTransform, space XRSpace) (js.Promise[XRAnchor], bool) {
	var _ok bool
	_ret := bindings.CallXRFrameCreateAnchor(
		this.Ref(), js.Pointer(&_ok),
		pose.Ref(),
		space.Ref(),
	)

	return js.Promise[XRAnchor]{}.FromRef(_ret), _ok
}

// CreateAnchorFunc returns the method "XRFrame.createAnchor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRFrame) CreateAnchorFunc() (fn js.Func[func(pose XRRigidTransform, space XRSpace) js.Promise[XRAnchor]]) {
	return fn.FromRef(
		bindings.XRFrameCreateAnchorFunc(
			this.Ref(),
		),
	)
}

// GetLightEstimate calls the method "XRFrame.getLightEstimate".
//
// The returned bool will be false if there is no such method.
func (this XRFrame) GetLightEstimate(lightProbe XRLightProbe) (XRLightEstimate, bool) {
	var _ok bool
	_ret := bindings.CallXRFrameGetLightEstimate(
		this.Ref(), js.Pointer(&_ok),
		lightProbe.Ref(),
	)

	return XRLightEstimate{}.FromRef(_ret), _ok
}

// GetLightEstimateFunc returns the method "XRFrame.getLightEstimate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRFrame) GetLightEstimateFunc() (fn js.Func[func(lightProbe XRLightProbe) XRLightEstimate]) {
	return fn.FromRef(
		bindings.XRFrameGetLightEstimateFunc(
			this.Ref(),
		),
	)
}

// GetDepthInformation calls the method "XRFrame.getDepthInformation".
//
// The returned bool will be false if there is no such method.
func (this XRFrame) GetDepthInformation(view XRView) (XRCPUDepthInformation, bool) {
	var _ok bool
	_ret := bindings.CallXRFrameGetDepthInformation(
		this.Ref(), js.Pointer(&_ok),
		view.Ref(),
	)

	return XRCPUDepthInformation{}.FromRef(_ret), _ok
}

// GetDepthInformationFunc returns the method "XRFrame.getDepthInformation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRFrame) GetDepthInformationFunc() (fn js.Func[func(view XRView) XRCPUDepthInformation]) {
	return fn.FromRef(
		bindings.XRFrameGetDepthInformationFunc(
			this.Ref(),
		),
	)
}

// GetJointPose calls the method "XRFrame.getJointPose".
//
// The returned bool will be false if there is no such method.
func (this XRFrame) GetJointPose(joint XRJointSpace, baseSpace XRSpace) (XRJointPose, bool) {
	var _ok bool
	_ret := bindings.CallXRFrameGetJointPose(
		this.Ref(), js.Pointer(&_ok),
		joint.Ref(),
		baseSpace.Ref(),
	)

	return XRJointPose{}.FromRef(_ret), _ok
}

// GetJointPoseFunc returns the method "XRFrame.getJointPose".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRFrame) GetJointPoseFunc() (fn js.Func[func(joint XRJointSpace, baseSpace XRSpace) XRJointPose]) {
	return fn.FromRef(
		bindings.XRFrameGetJointPoseFunc(
			this.Ref(),
		),
	)
}

// FillJointRadii calls the method "XRFrame.fillJointRadii".
//
// The returned bool will be false if there is no such method.
func (this XRFrame) FillJointRadii(jointSpaces js.Array[XRJointSpace], radii js.TypedArray[float32]) (bool, bool) {
	var _ok bool
	_ret := bindings.CallXRFrameFillJointRadii(
		this.Ref(), js.Pointer(&_ok),
		jointSpaces.Ref(),
		radii.Ref(),
	)

	return _ret == js.True, _ok
}

// FillJointRadiiFunc returns the method "XRFrame.fillJointRadii".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRFrame) FillJointRadiiFunc() (fn js.Func[func(jointSpaces js.Array[XRJointSpace], radii js.TypedArray[float32]) bool]) {
	return fn.FromRef(
		bindings.XRFrameFillJointRadiiFunc(
			this.Ref(),
		),
	)
}

// FillPoses calls the method "XRFrame.fillPoses".
//
// The returned bool will be false if there is no such method.
func (this XRFrame) FillPoses(spaces js.Array[XRSpace], baseSpace XRSpace, transforms js.TypedArray[float32]) (bool, bool) {
	var _ok bool
	_ret := bindings.CallXRFrameFillPoses(
		this.Ref(), js.Pointer(&_ok),
		spaces.Ref(),
		baseSpace.Ref(),
		transforms.Ref(),
	)

	return _ret == js.True, _ok
}

// FillPosesFunc returns the method "XRFrame.fillPoses".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRFrame) FillPosesFunc() (fn js.Func[func(spaces js.Array[XRSpace], baseSpace XRSpace, transforms js.TypedArray[float32]) bool]) {
	return fn.FromRef(
		bindings.XRFrameFillPosesFunc(
			this.Ref(),
		),
	)
}

// GetHitTestResults calls the method "XRFrame.getHitTestResults".
//
// The returned bool will be false if there is no such method.
func (this XRFrame) GetHitTestResults(hitTestSource XRHitTestSource) (js.FrozenArray[XRHitTestResult], bool) {
	var _ok bool
	_ret := bindings.CallXRFrameGetHitTestResults(
		this.Ref(), js.Pointer(&_ok),
		hitTestSource.Ref(),
	)

	return js.FrozenArray[XRHitTestResult]{}.FromRef(_ret), _ok
}

// GetHitTestResultsFunc returns the method "XRFrame.getHitTestResults".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRFrame) GetHitTestResultsFunc() (fn js.Func[func(hitTestSource XRHitTestSource) js.FrozenArray[XRHitTestResult]]) {
	return fn.FromRef(
		bindings.XRFrameGetHitTestResultsFunc(
			this.Ref(),
		),
	)
}

// GetHitTestResultsForTransientInput calls the method "XRFrame.getHitTestResultsForTransientInput".
//
// The returned bool will be false if there is no such method.
func (this XRFrame) GetHitTestResultsForTransientInput(hitTestSource XRTransientInputHitTestSource) (js.FrozenArray[XRTransientInputHitTestResult], bool) {
	var _ok bool
	_ret := bindings.CallXRFrameGetHitTestResultsForTransientInput(
		this.Ref(), js.Pointer(&_ok),
		hitTestSource.Ref(),
	)

	return js.FrozenArray[XRTransientInputHitTestResult]{}.FromRef(_ret), _ok
}

// GetHitTestResultsForTransientInputFunc returns the method "XRFrame.getHitTestResultsForTransientInput".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRFrame) GetHitTestResultsForTransientInputFunc() (fn js.Func[func(hitTestSource XRTransientInputHitTestSource) js.FrozenArray[XRTransientInputHitTestResult]]) {
	return fn.FromRef(
		bindings.XRFrameGetHitTestResultsForTransientInputFunc(
			this.Ref(),
		),
	)
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
func (p XRLightProbeInit) UpdateFrom(ref js.Ref) {
	bindings.XRLightProbeInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRLightProbeInit) Update(ref js.Ref) {
	bindings.XRLightProbeInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p XRRayDirectionInit) UpdateFrom(ref js.Ref) {
	bindings.XRRayDirectionInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRRayDirectionInit) Update(ref js.Ref) {
	bindings.XRRayDirectionInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewXRRay(origin DOMPointInit, direction XRRayDirectionInit) XRRay {
	return XRRay{}.FromRef(
		bindings.NewXRRayByXRRay(
			js.Pointer(&origin),
			js.Pointer(&direction)),
	)
}

func NewXRRayByXRRay1(origin DOMPointInit) XRRay {
	return XRRay{}.FromRef(
		bindings.NewXRRayByXRRay1(
			js.Pointer(&origin)),
	)
}

func NewXRRayByXRRay2() XRRay {
	return XRRay{}.FromRef(
		bindings.NewXRRayByXRRay2(),
	)
}

func NewXRRayByXRRay3(transform XRRigidTransform) XRRay {
	return XRRay{}.FromRef(
		bindings.NewXRRayByXRRay3(
			transform.Ref()),
	)
}

type XRRay struct {
	ref js.Ref
}

func (this XRRay) Once() XRRay {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Origin returns the value of property "XRRay.origin".
//
// The returned bool will be false if there is no such property.
func (this XRRay) Origin() (DOMPointReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetXRRayOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMPointReadOnly{}.FromRef(_ret), _ok
}

// Direction returns the value of property "XRRay.direction".
//
// The returned bool will be false if there is no such property.
func (this XRRay) Direction() (DOMPointReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetXRRayDirection(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMPointReadOnly{}.FromRef(_ret), _ok
}

// Matrix returns the value of property "XRRay.matrix".
//
// The returned bool will be false if there is no such property.
func (this XRRay) Matrix() (js.TypedArray[float32], bool) {
	var _ok bool
	_ret := bindings.GetXRRayMatrix(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[float32]{}.FromRef(_ret), _ok
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
func (p XRHitTestOptionsInit) UpdateFrom(ref js.Ref) {
	bindings.XRHitTestOptionsInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRHitTestOptionsInit) Update(ref js.Ref) {
	bindings.XRHitTestOptionsInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p XRTransientInputHitTestOptionsInit) UpdateFrom(ref js.Ref) {
	bindings.XRTransientInputHitTestOptionsInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRTransientInputHitTestOptionsInit) Update(ref js.Ref) {
	bindings.XRTransientInputHitTestOptionsInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// DepthNear returns the value of property "XRRenderState.depthNear".
//
// The returned bool will be false if there is no such property.
func (this XRRenderState) DepthNear() (float64, bool) {
	var _ok bool
	_ret := bindings.GetXRRenderStateDepthNear(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// DepthFar returns the value of property "XRRenderState.depthFar".
//
// The returned bool will be false if there is no such property.
func (this XRRenderState) DepthFar() (float64, bool) {
	var _ok bool
	_ret := bindings.GetXRRenderStateDepthFar(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// InlineVerticalFieldOfView returns the value of property "XRRenderState.inlineVerticalFieldOfView".
//
// The returned bool will be false if there is no such property.
func (this XRRenderState) InlineVerticalFieldOfView() (float64, bool) {
	var _ok bool
	_ret := bindings.GetXRRenderStateInlineVerticalFieldOfView(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// BaseLayer returns the value of property "XRRenderState.baseLayer".
//
// The returned bool will be false if there is no such property.
func (this XRRenderState) BaseLayer() (XRWebGLLayer, bool) {
	var _ok bool
	_ret := bindings.GetXRRenderStateBaseLayer(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRWebGLLayer{}.FromRef(_ret), _ok
}

// Layers returns the value of property "XRRenderState.layers".
//
// The returned bool will be false if there is no such property.
func (this XRRenderState) Layers() (js.FrozenArray[XRLayer], bool) {
	var _ok bool
	_ret := bindings.GetXRRenderStateLayers(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[XRLayer]{}.FromRef(_ret), _ok
}

type XRInputSourceArray struct {
	ref js.Ref
}

func (this XRInputSourceArray) Once() XRInputSourceArray {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Length returns the value of property "XRInputSourceArray.length".
//
// The returned bool will be false if there is no such property.
func (this XRInputSourceArray) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRInputSourceArrayLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Get calls the method "XRInputSourceArray.".
//
// The returned bool will be false if there is no such method.
func (this XRInputSourceArray) Get(index uint32) (XRInputSource, bool) {
	var _ok bool
	_ret := bindings.CallXRInputSourceArrayGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return XRInputSource{}.FromRef(_ret), _ok
}

// GetFunc returns the method "XRInputSourceArray.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRInputSourceArray) GetFunc() (fn js.Func[func(index uint32) XRInputSource]) {
	return fn.FromRef(
		bindings.XRInputSourceArrayGetFunc(
			this.Ref(),
		),
	)
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
func (p XRDOMOverlayState) UpdateFrom(ref js.Ref) {
	bindings.XRDOMOverlayStateJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRDOMOverlayState) Update(ref js.Ref) {
	bindings.XRDOMOverlayStateJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// VisibilityState returns the value of property "XRSession.visibilityState".
//
// The returned bool will be false if there is no such property.
func (this XRSession) VisibilityState() (XRVisibilityState, bool) {
	var _ok bool
	_ret := bindings.GetXRSessionVisibilityState(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRVisibilityState(_ret), _ok
}

// FrameRate returns the value of property "XRSession.frameRate".
//
// The returned bool will be false if there is no such property.
func (this XRSession) FrameRate() (float32, bool) {
	var _ok bool
	_ret := bindings.GetXRSessionFrameRate(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// SupportedFrameRates returns the value of property "XRSession.supportedFrameRates".
//
// The returned bool will be false if there is no such property.
func (this XRSession) SupportedFrameRates() (js.TypedArray[float32], bool) {
	var _ok bool
	_ret := bindings.GetXRSessionSupportedFrameRates(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[float32]{}.FromRef(_ret), _ok
}

// RenderState returns the value of property "XRSession.renderState".
//
// The returned bool will be false if there is no such property.
func (this XRSession) RenderState() (XRRenderState, bool) {
	var _ok bool
	_ret := bindings.GetXRSessionRenderState(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRRenderState{}.FromRef(_ret), _ok
}

// InputSources returns the value of property "XRSession.inputSources".
//
// The returned bool will be false if there is no such property.
func (this XRSession) InputSources() (XRInputSourceArray, bool) {
	var _ok bool
	_ret := bindings.GetXRSessionInputSources(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRInputSourceArray{}.FromRef(_ret), _ok
}

// EnabledFeatures returns the value of property "XRSession.enabledFeatures".
//
// The returned bool will be false if there is no such property.
func (this XRSession) EnabledFeatures() (js.FrozenArray[js.String], bool) {
	var _ok bool
	_ret := bindings.GetXRSessionEnabledFeatures(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[js.String]{}.FromRef(_ret), _ok
}

// IsSystemKeyboardSupported returns the value of property "XRSession.isSystemKeyboardSupported".
//
// The returned bool will be false if there is no such property.
func (this XRSession) IsSystemKeyboardSupported() (bool, bool) {
	var _ok bool
	_ret := bindings.GetXRSessionIsSystemKeyboardSupported(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// EnvironmentBlendMode returns the value of property "XRSession.environmentBlendMode".
//
// The returned bool will be false if there is no such property.
func (this XRSession) EnvironmentBlendMode() (XREnvironmentBlendMode, bool) {
	var _ok bool
	_ret := bindings.GetXRSessionEnvironmentBlendMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return XREnvironmentBlendMode(_ret), _ok
}

// InteractionMode returns the value of property "XRSession.interactionMode".
//
// The returned bool will be false if there is no such property.
func (this XRSession) InteractionMode() (XRInteractionMode, bool) {
	var _ok bool
	_ret := bindings.GetXRSessionInteractionMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRInteractionMode(_ret), _ok
}

// DomOverlayState returns the value of property "XRSession.domOverlayState".
//
// The returned bool will be false if there is no such property.
func (this XRSession) DomOverlayState() (XRDOMOverlayState, bool) {
	var _ret XRDOMOverlayState
	_ok := js.True == bindings.GetXRSessionDomOverlayState(
		this.Ref(),
		js.Pointer(&_ret),
	)

	return _ret, _ok
}

// DepthUsage returns the value of property "XRSession.depthUsage".
//
// The returned bool will be false if there is no such property.
func (this XRSession) DepthUsage() (XRDepthUsage, bool) {
	var _ok bool
	_ret := bindings.GetXRSessionDepthUsage(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRDepthUsage(_ret), _ok
}

// DepthDataFormat returns the value of property "XRSession.depthDataFormat".
//
// The returned bool will be false if there is no such property.
func (this XRSession) DepthDataFormat() (XRDepthDataFormat, bool) {
	var _ok bool
	_ret := bindings.GetXRSessionDepthDataFormat(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRDepthDataFormat(_ret), _ok
}

// PersistentAnchors returns the value of property "XRSession.persistentAnchors".
//
// The returned bool will be false if there is no such property.
func (this XRSession) PersistentAnchors() (js.FrozenArray[js.String], bool) {
	var _ok bool
	_ret := bindings.GetXRSessionPersistentAnchors(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[js.String]{}.FromRef(_ret), _ok
}

// PreferredReflectionFormat returns the value of property "XRSession.preferredReflectionFormat".
//
// The returned bool will be false if there is no such property.
func (this XRSession) PreferredReflectionFormat() (XRReflectionFormat, bool) {
	var _ok bool
	_ret := bindings.GetXRSessionPreferredReflectionFormat(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRReflectionFormat(_ret), _ok
}

// UpdateRenderState calls the method "XRSession.updateRenderState".
//
// The returned bool will be false if there is no such method.
func (this XRSession) UpdateRenderState(state XRRenderStateInit) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXRSessionUpdateRenderState(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&state),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UpdateRenderStateFunc returns the method "XRSession.updateRenderState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRSession) UpdateRenderStateFunc() (fn js.Func[func(state XRRenderStateInit)]) {
	return fn.FromRef(
		bindings.XRSessionUpdateRenderStateFunc(
			this.Ref(),
		),
	)
}

// UpdateRenderState1 calls the method "XRSession.updateRenderState".
//
// The returned bool will be false if there is no such method.
func (this XRSession) UpdateRenderState1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXRSessionUpdateRenderState1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UpdateRenderState1Func returns the method "XRSession.updateRenderState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRSession) UpdateRenderState1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.XRSessionUpdateRenderState1Func(
			this.Ref(),
		),
	)
}

// UpdateTargetFrameRate calls the method "XRSession.updateTargetFrameRate".
//
// The returned bool will be false if there is no such method.
func (this XRSession) UpdateTargetFrameRate(rate float32) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallXRSessionUpdateTargetFrameRate(
		this.Ref(), js.Pointer(&_ok),
		float32(rate),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// UpdateTargetFrameRateFunc returns the method "XRSession.updateTargetFrameRate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRSession) UpdateTargetFrameRateFunc() (fn js.Func[func(rate float32) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.XRSessionUpdateTargetFrameRateFunc(
			this.Ref(),
		),
	)
}

// RequestReferenceSpace calls the method "XRSession.requestReferenceSpace".
//
// The returned bool will be false if there is no such method.
func (this XRSession) RequestReferenceSpace(typ XRReferenceSpaceType) (js.Promise[XRReferenceSpace], bool) {
	var _ok bool
	_ret := bindings.CallXRSessionRequestReferenceSpace(
		this.Ref(), js.Pointer(&_ok),
		uint32(typ),
	)

	return js.Promise[XRReferenceSpace]{}.FromRef(_ret), _ok
}

// RequestReferenceSpaceFunc returns the method "XRSession.requestReferenceSpace".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRSession) RequestReferenceSpaceFunc() (fn js.Func[func(typ XRReferenceSpaceType) js.Promise[XRReferenceSpace]]) {
	return fn.FromRef(
		bindings.XRSessionRequestReferenceSpaceFunc(
			this.Ref(),
		),
	)
}

// RequestAnimationFrame calls the method "XRSession.requestAnimationFrame".
//
// The returned bool will be false if there is no such method.
func (this XRSession) RequestAnimationFrame(callback js.Func[func(time DOMHighResTimeStamp, frame XRFrame)]) (uint32, bool) {
	var _ok bool
	_ret := bindings.CallXRSessionRequestAnimationFrame(
		this.Ref(), js.Pointer(&_ok),
		callback.Ref(),
	)

	return uint32(_ret), _ok
}

// RequestAnimationFrameFunc returns the method "XRSession.requestAnimationFrame".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRSession) RequestAnimationFrameFunc() (fn js.Func[func(callback js.Func[func(time DOMHighResTimeStamp, frame XRFrame)]) uint32]) {
	return fn.FromRef(
		bindings.XRSessionRequestAnimationFrameFunc(
			this.Ref(),
		),
	)
}

// CancelAnimationFrame calls the method "XRSession.cancelAnimationFrame".
//
// The returned bool will be false if there is no such method.
func (this XRSession) CancelAnimationFrame(handle uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXRSessionCancelAnimationFrame(
		this.Ref(), js.Pointer(&_ok),
		uint32(handle),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CancelAnimationFrameFunc returns the method "XRSession.cancelAnimationFrame".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRSession) CancelAnimationFrameFunc() (fn js.Func[func(handle uint32)]) {
	return fn.FromRef(
		bindings.XRSessionCancelAnimationFrameFunc(
			this.Ref(),
		),
	)
}

// End calls the method "XRSession.end".
//
// The returned bool will be false if there is no such method.
func (this XRSession) End() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallXRSessionEnd(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// EndFunc returns the method "XRSession.end".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRSession) EndFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.XRSessionEndFunc(
			this.Ref(),
		),
	)
}

// RestorePersistentAnchor calls the method "XRSession.restorePersistentAnchor".
//
// The returned bool will be false if there is no such method.
func (this XRSession) RestorePersistentAnchor(uuid js.String) (js.Promise[XRAnchor], bool) {
	var _ok bool
	_ret := bindings.CallXRSessionRestorePersistentAnchor(
		this.Ref(), js.Pointer(&_ok),
		uuid.Ref(),
	)

	return js.Promise[XRAnchor]{}.FromRef(_ret), _ok
}

// RestorePersistentAnchorFunc returns the method "XRSession.restorePersistentAnchor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRSession) RestorePersistentAnchorFunc() (fn js.Func[func(uuid js.String) js.Promise[XRAnchor]]) {
	return fn.FromRef(
		bindings.XRSessionRestorePersistentAnchorFunc(
			this.Ref(),
		),
	)
}

// DeletePersistentAnchor calls the method "XRSession.deletePersistentAnchor".
//
// The returned bool will be false if there is no such method.
func (this XRSession) DeletePersistentAnchor(uuid js.String) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallXRSessionDeletePersistentAnchor(
		this.Ref(), js.Pointer(&_ok),
		uuid.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// DeletePersistentAnchorFunc returns the method "XRSession.deletePersistentAnchor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRSession) DeletePersistentAnchorFunc() (fn js.Func[func(uuid js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.XRSessionDeletePersistentAnchorFunc(
			this.Ref(),
		),
	)
}

// RequestLightProbe calls the method "XRSession.requestLightProbe".
//
// The returned bool will be false if there is no such method.
func (this XRSession) RequestLightProbe(options XRLightProbeInit) (js.Promise[XRLightProbe], bool) {
	var _ok bool
	_ret := bindings.CallXRSessionRequestLightProbe(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[XRLightProbe]{}.FromRef(_ret), _ok
}

// RequestLightProbeFunc returns the method "XRSession.requestLightProbe".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRSession) RequestLightProbeFunc() (fn js.Func[func(options XRLightProbeInit) js.Promise[XRLightProbe]]) {
	return fn.FromRef(
		bindings.XRSessionRequestLightProbeFunc(
			this.Ref(),
		),
	)
}

// RequestLightProbe1 calls the method "XRSession.requestLightProbe".
//
// The returned bool will be false if there is no such method.
func (this XRSession) RequestLightProbe1() (js.Promise[XRLightProbe], bool) {
	var _ok bool
	_ret := bindings.CallXRSessionRequestLightProbe1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[XRLightProbe]{}.FromRef(_ret), _ok
}

// RequestLightProbe1Func returns the method "XRSession.requestLightProbe".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRSession) RequestLightProbe1Func() (fn js.Func[func() js.Promise[XRLightProbe]]) {
	return fn.FromRef(
		bindings.XRSessionRequestLightProbe1Func(
			this.Ref(),
		),
	)
}

// RequestHitTestSource calls the method "XRSession.requestHitTestSource".
//
// The returned bool will be false if there is no such method.
func (this XRSession) RequestHitTestSource(options XRHitTestOptionsInit) (js.Promise[XRHitTestSource], bool) {
	var _ok bool
	_ret := bindings.CallXRSessionRequestHitTestSource(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[XRHitTestSource]{}.FromRef(_ret), _ok
}

// RequestHitTestSourceFunc returns the method "XRSession.requestHitTestSource".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRSession) RequestHitTestSourceFunc() (fn js.Func[func(options XRHitTestOptionsInit) js.Promise[XRHitTestSource]]) {
	return fn.FromRef(
		bindings.XRSessionRequestHitTestSourceFunc(
			this.Ref(),
		),
	)
}

// RequestHitTestSourceForTransientInput calls the method "XRSession.requestHitTestSourceForTransientInput".
//
// The returned bool will be false if there is no such method.
func (this XRSession) RequestHitTestSourceForTransientInput(options XRTransientInputHitTestOptionsInit) (js.Promise[XRTransientInputHitTestSource], bool) {
	var _ok bool
	_ret := bindings.CallXRSessionRequestHitTestSourceForTransientInput(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[XRTransientInputHitTestSource]{}.FromRef(_ret), _ok
}

// RequestHitTestSourceForTransientInputFunc returns the method "XRSession.requestHitTestSourceForTransientInput".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRSession) RequestHitTestSourceForTransientInputFunc() (fn js.Func[func(options XRTransientInputHitTestOptionsInit) js.Promise[XRTransientInputHitTestSource]]) {
	return fn.FromRef(
		bindings.XRSessionRequestHitTestSourceForTransientInputFunc(
			this.Ref(),
		),
	)
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
func (p XRDOMOverlayInit) UpdateFrom(ref js.Ref) {
	bindings.XRDOMOverlayInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRDOMOverlayInit) Update(ref js.Ref) {
	bindings.XRDOMOverlayInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p XRDepthStateInit) UpdateFrom(ref js.Ref) {
	bindings.XRDepthStateInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRDepthStateInit) Update(ref js.Ref) {
	bindings.XRDepthStateInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	DomOverlay XRDOMOverlayInit
	// DepthSensing is "XRSessionInit.depthSensing"
	//
	// Optional
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
func (p XRSessionInit) UpdateFrom(ref js.Ref) {
	bindings.XRSessionInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRSessionInit) Update(ref js.Ref) {
	bindings.XRSessionInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type XRSystem struct {
	EventTarget
}

func (this XRSystem) Once() XRSystem {
	this.Ref().Once()
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
	this.Ref().Free()
}

// IsSessionSupported calls the method "XRSystem.isSessionSupported".
//
// The returned bool will be false if there is no such method.
func (this XRSystem) IsSessionSupported(mode XRSessionMode) (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallXRSystemIsSessionSupported(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// IsSessionSupportedFunc returns the method "XRSystem.isSessionSupported".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRSystem) IsSessionSupportedFunc() (fn js.Func[func(mode XRSessionMode) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.XRSystemIsSessionSupportedFunc(
			this.Ref(),
		),
	)
}

// RequestSession calls the method "XRSystem.requestSession".
//
// The returned bool will be false if there is no such method.
func (this XRSystem) RequestSession(mode XRSessionMode, options XRSessionInit) (js.Promise[XRSession], bool) {
	var _ok bool
	_ret := bindings.CallXRSystemRequestSession(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
		js.Pointer(&options),
	)

	return js.Promise[XRSession]{}.FromRef(_ret), _ok
}

// RequestSessionFunc returns the method "XRSystem.requestSession".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRSystem) RequestSessionFunc() (fn js.Func[func(mode XRSessionMode, options XRSessionInit) js.Promise[XRSession]]) {
	return fn.FromRef(
		bindings.XRSystemRequestSessionFunc(
			this.Ref(),
		),
	)
}

// RequestSession1 calls the method "XRSystem.requestSession".
//
// The returned bool will be false if there is no such method.
func (this XRSystem) RequestSession1(mode XRSessionMode) (js.Promise[XRSession], bool) {
	var _ok bool
	_ret := bindings.CallXRSystemRequestSession1(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
	)

	return js.Promise[XRSession]{}.FromRef(_ret), _ok
}

// RequestSession1Func returns the method "XRSystem.requestSession".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRSystem) RequestSession1Func() (fn js.Func[func(mode XRSessionMode) js.Promise[XRSession]]) {
	return fn.FromRef(
		bindings.XRSystemRequestSession1Func(
			this.Ref(),
		),
	)
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
func (p NotificationAction) UpdateFrom(ref js.Ref) {
	bindings.NotificationActionJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NotificationAction) Update(ref js.Ref) {
	bindings.NotificationActionJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p NotificationOptions) UpdateFrom(ref js.Ref) {
	bindings.NotificationOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NotificationOptions) Update(ref js.Ref) {
	bindings.NotificationOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
		assert.Throw("invalid", "callback", "invocation")
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

func NewNotification(title js.String, options NotificationOptions) Notification {
	return Notification{}.FromRef(
		bindings.NewNotificationByNotification(
			title.Ref(),
			js.Pointer(&options)),
	)
}

func NewNotificationByNotification1(title js.String) Notification {
	return Notification{}.FromRef(
		bindings.NewNotificationByNotification1(
			title.Ref()),
	)
}

type Notification struct {
	EventTarget
}

func (this Notification) Once() Notification {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Permission returns the value of property "Notification.permission".
//
// The returned bool will be false if there is no such property.
func (this Notification) Permission() (NotificationPermission, bool) {
	var _ok bool
	_ret := bindings.GetNotificationPermission(
		this.Ref(), js.Pointer(&_ok),
	)
	return NotificationPermission(_ret), _ok
}

// MaxActions returns the value of property "Notification.maxActions".
//
// The returned bool will be false if there is no such property.
func (this Notification) MaxActions() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetNotificationMaxActions(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Title returns the value of property "Notification.title".
//
// The returned bool will be false if there is no such property.
func (this Notification) Title() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNotificationTitle(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Dir returns the value of property "Notification.dir".
//
// The returned bool will be false if there is no such property.
func (this Notification) Dir() (NotificationDirection, bool) {
	var _ok bool
	_ret := bindings.GetNotificationDir(
		this.Ref(), js.Pointer(&_ok),
	)
	return NotificationDirection(_ret), _ok
}

// Lang returns the value of property "Notification.lang".
//
// The returned bool will be false if there is no such property.
func (this Notification) Lang() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNotificationLang(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Body returns the value of property "Notification.body".
//
// The returned bool will be false if there is no such property.
func (this Notification) Body() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNotificationBody(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Tag returns the value of property "Notification.tag".
//
// The returned bool will be false if there is no such property.
func (this Notification) Tag() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNotificationTag(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Image returns the value of property "Notification.image".
//
// The returned bool will be false if there is no such property.
func (this Notification) Image() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNotificationImage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Icon returns the value of property "Notification.icon".
//
// The returned bool will be false if there is no such property.
func (this Notification) Icon() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNotificationIcon(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Badge returns the value of property "Notification.badge".
//
// The returned bool will be false if there is no such property.
func (this Notification) Badge() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNotificationBadge(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Vibrate returns the value of property "Notification.vibrate".
//
// The returned bool will be false if there is no such property.
func (this Notification) Vibrate() (js.FrozenArray[uint32], bool) {
	var _ok bool
	_ret := bindings.GetNotificationVibrate(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[uint32]{}.FromRef(_ret), _ok
}

// Timestamp returns the value of property "Notification.timestamp".
//
// The returned bool will be false if there is no such property.
func (this Notification) Timestamp() (EpochTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetNotificationTimestamp(
		this.Ref(), js.Pointer(&_ok),
	)
	return EpochTimeStamp(_ret), _ok
}

// Renotify returns the value of property "Notification.renotify".
//
// The returned bool will be false if there is no such property.
func (this Notification) Renotify() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNotificationRenotify(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Silent returns the value of property "Notification.silent".
//
// The returned bool will be false if there is no such property.
func (this Notification) Silent() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNotificationSilent(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// RequireInteraction returns the value of property "Notification.requireInteraction".
//
// The returned bool will be false if there is no such property.
func (this Notification) RequireInteraction() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNotificationRequireInteraction(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Data returns the value of property "Notification.data".
//
// The returned bool will be false if there is no such property.
func (this Notification) Data() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetNotificationData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// Actions returns the value of property "Notification.actions".
//
// The returned bool will be false if there is no such property.
func (this Notification) Actions() (js.FrozenArray[NotificationAction], bool) {
	var _ok bool
	_ret := bindings.GetNotificationActions(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[NotificationAction]{}.FromRef(_ret), _ok
}

// RequestPermission calls the staticmethod "Notification.requestPermission".
//
// The returned bool will be false if there is no such method.
func (this Notification) RequestPermission(deprecatedCallback js.Func[func(permission NotificationPermission)]) (js.Promise[NotificationPermission], bool) {
	var _ok bool
	_ret := bindings.CallNotificationRequestPermission(
		this.Ref(), js.Pointer(&_ok),
		deprecatedCallback.Ref(),
	)

	return js.Promise[NotificationPermission]{}.FromRef(_ret), _ok
}

// RequestPermissionFunc returns the staticmethod "Notification.requestPermission".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Notification) RequestPermissionFunc() (fn js.Func[func(deprecatedCallback js.Func[func(permission NotificationPermission)]) js.Promise[NotificationPermission]]) {
	return fn.FromRef(
		bindings.NotificationRequestPermissionFunc(
			this.Ref(),
		),
	)
}

// RequestPermission1 calls the staticmethod "Notification.requestPermission".
//
// The returned bool will be false if there is no such method.
func (this Notification) RequestPermission1() (js.Promise[NotificationPermission], bool) {
	var _ok bool
	_ret := bindings.CallNotificationRequestPermission1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[NotificationPermission]{}.FromRef(_ret), _ok
}

// RequestPermission1Func returns the staticmethod "Notification.requestPermission".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Notification) RequestPermission1Func() (fn js.Func[func() js.Promise[NotificationPermission]]) {
	return fn.FromRef(
		bindings.NotificationRequestPermission1Func(
			this.Ref(),
		),
	)
}

// Close calls the method "Notification.close".
//
// The returned bool will be false if there is no such method.
func (this Notification) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallNotificationClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "Notification.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Notification) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.NotificationCloseFunc(
			this.Ref(),
		),
	)
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
func (p GetNotificationOptions) UpdateFrom(ref js.Ref) {
	bindings.GetNotificationOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GetNotificationOptions) Update(ref js.Ref) {
	bindings.GetNotificationOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

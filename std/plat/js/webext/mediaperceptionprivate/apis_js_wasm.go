// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package mediaperceptionprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/mediaperceptionprivate/bindings"
)

type AudioSpectrogram struct {
	// Values is "AudioSpectrogram.values"
	//
	// Optional
	Values js.Array[float64]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioSpectrogram with all fields set.
func (p AudioSpectrogram) FromRef(ref js.Ref) AudioSpectrogram {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioSpectrogram in the application heap.
func (p AudioSpectrogram) New() js.Ref {
	return bindings.AudioSpectrogramJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioSpectrogram) UpdateFrom(ref js.Ref) {
	bindings.AudioSpectrogramJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioSpectrogram) Update(ref js.Ref) {
	bindings.AudioSpectrogramJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioSpectrogram) FreeMembers(recursive bool) {
	js.Free(
		p.Values.Ref(),
	)
	p.Values = p.Values.FromRef(js.Undefined)
}

type AudioHumanPresenceDetection struct {
	// HumanPresenceLikelihood is "AudioHumanPresenceDetection.humanPresenceLikelihood"
	//
	// Optional
	//
	// NOTE: FFI_USE_HumanPresenceLikelihood MUST be set to true to make this field effective.
	HumanPresenceLikelihood float64
	// NoiseSpectrogram is "AudioHumanPresenceDetection.noiseSpectrogram"
	//
	// Optional
	//
	// NOTE: NoiseSpectrogram.FFI_USE MUST be set to true to get NoiseSpectrogram used.
	NoiseSpectrogram AudioSpectrogram
	// FrameSpectrogram is "AudioHumanPresenceDetection.frameSpectrogram"
	//
	// Optional
	//
	// NOTE: FrameSpectrogram.FFI_USE MUST be set to true to get FrameSpectrogram used.
	FrameSpectrogram AudioSpectrogram

	FFI_USE_HumanPresenceLikelihood bool // for HumanPresenceLikelihood.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioHumanPresenceDetection with all fields set.
func (p AudioHumanPresenceDetection) FromRef(ref js.Ref) AudioHumanPresenceDetection {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioHumanPresenceDetection in the application heap.
func (p AudioHumanPresenceDetection) New() js.Ref {
	return bindings.AudioHumanPresenceDetectionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioHumanPresenceDetection) UpdateFrom(ref js.Ref) {
	bindings.AudioHumanPresenceDetectionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioHumanPresenceDetection) Update(ref js.Ref) {
	bindings.AudioHumanPresenceDetectionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioHumanPresenceDetection) FreeMembers(recursive bool) {
	if recursive {
		p.NoiseSpectrogram.FreeMembers(true)
		p.FrameSpectrogram.FreeMembers(true)
	}
}

type AudioLocalization struct {
	// AzimuthRadians is "AudioLocalization.azimuthRadians"
	//
	// Optional
	//
	// NOTE: FFI_USE_AzimuthRadians MUST be set to true to make this field effective.
	AzimuthRadians float64
	// AzimuthScores is "AudioLocalization.azimuthScores"
	//
	// Optional
	AzimuthScores js.Array[float64]

	FFI_USE_AzimuthRadians bool // for AzimuthRadians.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioLocalization with all fields set.
func (p AudioLocalization) FromRef(ref js.Ref) AudioLocalization {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioLocalization in the application heap.
func (p AudioLocalization) New() js.Ref {
	return bindings.AudioLocalizationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioLocalization) UpdateFrom(ref js.Ref) {
	bindings.AudioLocalizationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioLocalization) Update(ref js.Ref) {
	bindings.AudioLocalizationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioLocalization) FreeMembers(recursive bool) {
	js.Free(
		p.AzimuthScores.Ref(),
	)
	p.AzimuthScores = p.AzimuthScores.FromRef(js.Undefined)
}

type HotwordType uint32

const (
	_ HotwordType = iota

	HotwordType_UNKNOWN_TYPE
	HotwordType_OK_GOOGLE
)

func (HotwordType) FromRef(str js.Ref) HotwordType {
	return HotwordType(bindings.ConstOfHotwordType(str))
}

func (x HotwordType) String() (string, bool) {
	switch x {
	case HotwordType_UNKNOWN_TYPE:
		return "UNKNOWN_TYPE", true
	case HotwordType_OK_GOOGLE:
		return "OK_GOOGLE", true
	default:
		return "", false
	}
}

type Hotword struct {
	// Id is "Hotword.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// Type is "Hotword.type"
	//
	// Optional
	Type HotwordType
	// FrameId is "Hotword.frameId"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameId MUST be set to true to make this field effective.
	FrameId int32
	// StartTimestampMs is "Hotword.startTimestampMs"
	//
	// Optional
	//
	// NOTE: FFI_USE_StartTimestampMs MUST be set to true to make this field effective.
	StartTimestampMs int32
	// EndTimestampMs is "Hotword.endTimestampMs"
	//
	// Optional
	//
	// NOTE: FFI_USE_EndTimestampMs MUST be set to true to make this field effective.
	EndTimestampMs int32
	// Confidence is "Hotword.confidence"
	//
	// Optional
	//
	// NOTE: FFI_USE_Confidence MUST be set to true to make this field effective.
	Confidence float64

	FFI_USE_Id               bool // for Id.
	FFI_USE_FrameId          bool // for FrameId.
	FFI_USE_StartTimestampMs bool // for StartTimestampMs.
	FFI_USE_EndTimestampMs   bool // for EndTimestampMs.
	FFI_USE_Confidence       bool // for Confidence.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Hotword with all fields set.
func (p Hotword) FromRef(ref js.Ref) Hotword {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Hotword in the application heap.
func (p Hotword) New() js.Ref {
	return bindings.HotwordJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Hotword) UpdateFrom(ref js.Ref) {
	bindings.HotwordJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Hotword) Update(ref js.Ref) {
	bindings.HotwordJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Hotword) FreeMembers(recursive bool) {
}

type HotwordDetection struct {
	// Hotwords is "HotwordDetection.hotwords"
	//
	// Optional
	Hotwords js.Array[Hotword]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HotwordDetection with all fields set.
func (p HotwordDetection) FromRef(ref js.Ref) HotwordDetection {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HotwordDetection in the application heap.
func (p HotwordDetection) New() js.Ref {
	return bindings.HotwordDetectionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HotwordDetection) UpdateFrom(ref js.Ref) {
	bindings.HotwordDetectionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HotwordDetection) Update(ref js.Ref) {
	bindings.HotwordDetectionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HotwordDetection) FreeMembers(recursive bool) {
	js.Free(
		p.Hotwords.Ref(),
	)
	p.Hotwords = p.Hotwords.FromRef(js.Undefined)
}

type AudioPerception struct {
	// TimestampUs is "AudioPerception.timestampUs"
	//
	// Optional
	//
	// NOTE: FFI_USE_TimestampUs MUST be set to true to make this field effective.
	TimestampUs float64
	// AudioLocalization is "AudioPerception.audioLocalization"
	//
	// Optional
	//
	// NOTE: AudioLocalization.FFI_USE MUST be set to true to get AudioLocalization used.
	AudioLocalization AudioLocalization
	// AudioHumanPresenceDetection is "AudioPerception.audioHumanPresenceDetection"
	//
	// Optional
	//
	// NOTE: AudioHumanPresenceDetection.FFI_USE MUST be set to true to get AudioHumanPresenceDetection used.
	AudioHumanPresenceDetection AudioHumanPresenceDetection
	// HotwordDetection is "AudioPerception.hotwordDetection"
	//
	// Optional
	//
	// NOTE: HotwordDetection.FFI_USE MUST be set to true to get HotwordDetection used.
	HotwordDetection HotwordDetection

	FFI_USE_TimestampUs bool // for TimestampUs.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioPerception with all fields set.
func (p AudioPerception) FromRef(ref js.Ref) AudioPerception {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioPerception in the application heap.
func (p AudioPerception) New() js.Ref {
	return bindings.AudioPerceptionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioPerception) UpdateFrom(ref js.Ref) {
	bindings.AudioPerceptionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioPerception) Update(ref js.Ref) {
	bindings.AudioPerceptionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioPerception) FreeMembers(recursive bool) {
	if recursive {
		p.AudioLocalization.FreeMembers(true)
		p.AudioHumanPresenceDetection.FreeMembers(true)
		p.HotwordDetection.FreeMembers(true)
	}
}

type AudioVisualHumanPresenceDetection struct {
	// HumanPresenceLikelihood is "AudioVisualHumanPresenceDetection.humanPresenceLikelihood"
	//
	// Optional
	//
	// NOTE: FFI_USE_HumanPresenceLikelihood MUST be set to true to make this field effective.
	HumanPresenceLikelihood float64

	FFI_USE_HumanPresenceLikelihood bool // for HumanPresenceLikelihood.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioVisualHumanPresenceDetection with all fields set.
func (p AudioVisualHumanPresenceDetection) FromRef(ref js.Ref) AudioVisualHumanPresenceDetection {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioVisualHumanPresenceDetection in the application heap.
func (p AudioVisualHumanPresenceDetection) New() js.Ref {
	return bindings.AudioVisualHumanPresenceDetectionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioVisualHumanPresenceDetection) UpdateFrom(ref js.Ref) {
	bindings.AudioVisualHumanPresenceDetectionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioVisualHumanPresenceDetection) Update(ref js.Ref) {
	bindings.AudioVisualHumanPresenceDetectionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioVisualHumanPresenceDetection) FreeMembers(recursive bool) {
}

type AudioVisualPerception struct {
	// TimestampUs is "AudioVisualPerception.timestampUs"
	//
	// Optional
	//
	// NOTE: FFI_USE_TimestampUs MUST be set to true to make this field effective.
	TimestampUs float64
	// AudioVisualHumanPresenceDetection is "AudioVisualPerception.audioVisualHumanPresenceDetection"
	//
	// Optional
	//
	// NOTE: AudioVisualHumanPresenceDetection.FFI_USE MUST be set to true to get AudioVisualHumanPresenceDetection used.
	AudioVisualHumanPresenceDetection AudioVisualHumanPresenceDetection

	FFI_USE_TimestampUs bool // for TimestampUs.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioVisualPerception with all fields set.
func (p AudioVisualPerception) FromRef(ref js.Ref) AudioVisualPerception {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioVisualPerception in the application heap.
func (p AudioVisualPerception) New() js.Ref {
	return bindings.AudioVisualPerceptionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioVisualPerception) UpdateFrom(ref js.Ref) {
	bindings.AudioVisualPerceptionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioVisualPerception) Update(ref js.Ref) {
	bindings.AudioVisualPerceptionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioVisualPerception) FreeMembers(recursive bool) {
	if recursive {
		p.AudioVisualHumanPresenceDetection.FreeMembers(true)
	}
}

type Point struct {
	// X is "Point.x"
	//
	// Optional
	//
	// NOTE: FFI_USE_X MUST be set to true to make this field effective.
	X float64
	// Y is "Point.y"
	//
	// Optional
	//
	// NOTE: FFI_USE_Y MUST be set to true to make this field effective.
	Y float64

	FFI_USE_X bool // for X.
	FFI_USE_Y bool // for Y.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Point with all fields set.
func (p Point) FromRef(ref js.Ref) Point {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Point in the application heap.
func (p Point) New() js.Ref {
	return bindings.PointJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Point) UpdateFrom(ref js.Ref) {
	bindings.PointJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Point) Update(ref js.Ref) {
	bindings.PointJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Point) FreeMembers(recursive bool) {
}

type BoundingBox struct {
	// Normalized is "BoundingBox.normalized"
	//
	// Optional
	//
	// NOTE: FFI_USE_Normalized MUST be set to true to make this field effective.
	Normalized bool
	// TopLeft is "BoundingBox.topLeft"
	//
	// Optional
	//
	// NOTE: TopLeft.FFI_USE MUST be set to true to get TopLeft used.
	TopLeft Point
	// BottomRight is "BoundingBox.bottomRight"
	//
	// Optional
	//
	// NOTE: BottomRight.FFI_USE MUST be set to true to get BottomRight used.
	BottomRight Point

	FFI_USE_Normalized bool // for Normalized.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BoundingBox with all fields set.
func (p BoundingBox) FromRef(ref js.Ref) BoundingBox {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BoundingBox in the application heap.
func (p BoundingBox) New() js.Ref {
	return bindings.BoundingBoxJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BoundingBox) UpdateFrom(ref js.Ref) {
	bindings.BoundingBoxJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BoundingBox) Update(ref js.Ref) {
	bindings.BoundingBoxJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BoundingBox) FreeMembers(recursive bool) {
	if recursive {
		p.TopLeft.FreeMembers(true)
		p.BottomRight.FreeMembers(true)
	}
}

type ComponentType uint32

const (
	_ ComponentType = iota

	ComponentType_LIGHT
	ComponentType_FULL
)

func (ComponentType) FromRef(str js.Ref) ComponentType {
	return ComponentType(bindings.ConstOfComponentType(str))
}

func (x ComponentType) String() (string, bool) {
	switch x {
	case ComponentType_LIGHT:
		return "LIGHT", true
	case ComponentType_FULL:
		return "FULL", true
	default:
		return "", false
	}
}

type Component struct {
	// Type is "Component.type"
	//
	// Optional
	Type ComponentType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Component with all fields set.
func (p Component) FromRef(ref js.Ref) Component {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Component in the application heap.
func (p Component) New() js.Ref {
	return bindings.ComponentJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Component) UpdateFrom(ref js.Ref) {
	bindings.ComponentJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Component) Update(ref js.Ref) {
	bindings.ComponentJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Component) FreeMembers(recursive bool) {
}

type ComponentInstallationError uint32

const (
	_ ComponentInstallationError = iota

	ComponentInstallationError_UNKNOWN_COMPONENT
	ComponentInstallationError_INSTALL_FAILURE
	ComponentInstallationError_MOUNT_FAILURE
	ComponentInstallationError_COMPATIBILITY_CHECK_FAILED
	ComponentInstallationError_NOT_FOUND
)

func (ComponentInstallationError) FromRef(str js.Ref) ComponentInstallationError {
	return ComponentInstallationError(bindings.ConstOfComponentInstallationError(str))
}

func (x ComponentInstallationError) String() (string, bool) {
	switch x {
	case ComponentInstallationError_UNKNOWN_COMPONENT:
		return "UNKNOWN_COMPONENT", true
	case ComponentInstallationError_INSTALL_FAILURE:
		return "INSTALL_FAILURE", true
	case ComponentInstallationError_MOUNT_FAILURE:
		return "MOUNT_FAILURE", true
	case ComponentInstallationError_COMPATIBILITY_CHECK_FAILED:
		return "COMPATIBILITY_CHECK_FAILED", true
	case ComponentInstallationError_NOT_FOUND:
		return "NOT_FOUND", true
	default:
		return "", false
	}
}

type ComponentStatus uint32

const (
	_ ComponentStatus = iota

	ComponentStatus_UNKNOWN
	ComponentStatus_INSTALLED
	ComponentStatus_FAILED_TO_INSTALL
)

func (ComponentStatus) FromRef(str js.Ref) ComponentStatus {
	return ComponentStatus(bindings.ConstOfComponentStatus(str))
}

func (x ComponentStatus) String() (string, bool) {
	switch x {
	case ComponentStatus_UNKNOWN:
		return "UNKNOWN", true
	case ComponentStatus_INSTALLED:
		return "INSTALLED", true
	case ComponentStatus_FAILED_TO_INSTALL:
		return "FAILED_TO_INSTALL", true
	default:
		return "", false
	}
}

type ComponentState struct {
	// Status is "ComponentState.status"
	//
	// Optional
	Status ComponentStatus
	// Version is "ComponentState.version"
	//
	// Optional
	Version js.String
	// InstallationErrorCode is "ComponentState.installationErrorCode"
	//
	// Optional
	InstallationErrorCode ComponentInstallationError

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ComponentState with all fields set.
func (p ComponentState) FromRef(ref js.Ref) ComponentState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ComponentState in the application heap.
func (p ComponentState) New() js.Ref {
	return bindings.ComponentStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ComponentState) UpdateFrom(ref js.Ref) {
	bindings.ComponentStateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ComponentState) Update(ref js.Ref) {
	bindings.ComponentStateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ComponentState) FreeMembers(recursive bool) {
	js.Free(
		p.Version.Ref(),
	)
	p.Version = p.Version.FromRef(js.Undefined)
}

type ComponentStateCallbackFunc func(this js.Ref, componentState *ComponentState) js.Ref

func (fn ComponentStateCallbackFunc) Register() js.Func[func(componentState *ComponentState)] {
	return js.RegisterCallback[func(componentState *ComponentState)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ComponentStateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ComponentState
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ComponentStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, componentState *ComponentState) js.Ref
	Arg T
}

func (cb *ComponentStateCallback[T]) Register() js.Func[func(componentState *ComponentState)] {
	return js.RegisterCallback[func(componentState *ComponentState)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ComponentStateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ComponentState
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ServiceError uint32

const (
	_ ServiceError = iota

	ServiceError_SERVICE_UNREACHABLE
	ServiceError_SERVICE_NOT_RUNNING
	ServiceError_SERVICE_BUSY_LAUNCHING
	ServiceError_SERVICE_NOT_INSTALLED
	ServiceError_MOJO_CONNECTION_FAILURE
)

func (ServiceError) FromRef(str js.Ref) ServiceError {
	return ServiceError(bindings.ConstOfServiceError(str))
}

func (x ServiceError) String() (string, bool) {
	switch x {
	case ServiceError_SERVICE_UNREACHABLE:
		return "SERVICE_UNREACHABLE", true
	case ServiceError_SERVICE_NOT_RUNNING:
		return "SERVICE_NOT_RUNNING", true
	case ServiceError_SERVICE_BUSY_LAUNCHING:
		return "SERVICE_BUSY_LAUNCHING", true
	case ServiceError_SERVICE_NOT_INSTALLED:
		return "SERVICE_NOT_INSTALLED", true
	case ServiceError_MOJO_CONNECTION_FAILURE:
		return "MOJO_CONNECTION_FAILURE", true
	default:
		return "", false
	}
}

type EntityType uint32

const (
	_ EntityType = iota

	EntityType_UNSPECIFIED
	EntityType_FACE
	EntityType_PERSON
	EntityType_MOTION_REGION
	EntityType_LABELED_REGION
)

func (EntityType) FromRef(str js.Ref) EntityType {
	return EntityType(bindings.ConstOfEntityType(str))
}

func (x EntityType) String() (string, bool) {
	switch x {
	case EntityType_UNSPECIFIED:
		return "UNSPECIFIED", true
	case EntityType_FACE:
		return "FACE", true
	case EntityType_PERSON:
		return "PERSON", true
	case EntityType_MOTION_REGION:
		return "MOTION_REGION", true
	case EntityType_LABELED_REGION:
		return "LABELED_REGION", true
	default:
		return "", false
	}
}

type DistanceUnits uint32

const (
	_ DistanceUnits = iota

	DistanceUnits_UNSPECIFIED
	DistanceUnits_METERS
	DistanceUnits_PIXELS
)

func (DistanceUnits) FromRef(str js.Ref) DistanceUnits {
	return DistanceUnits(bindings.ConstOfDistanceUnits(str))
}

func (x DistanceUnits) String() (string, bool) {
	switch x {
	case DistanceUnits_UNSPECIFIED:
		return "UNSPECIFIED", true
	case DistanceUnits_METERS:
		return "METERS", true
	case DistanceUnits_PIXELS:
		return "PIXELS", true
	default:
		return "", false
	}
}

type Distance struct {
	// Units is "Distance.units"
	//
	// Optional
	Units DistanceUnits
	// Magnitude is "Distance.magnitude"
	//
	// Optional
	//
	// NOTE: FFI_USE_Magnitude MUST be set to true to make this field effective.
	Magnitude float64

	FFI_USE_Magnitude bool // for Magnitude.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Distance with all fields set.
func (p Distance) FromRef(ref js.Ref) Distance {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Distance in the application heap.
func (p Distance) New() js.Ref {
	return bindings.DistanceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Distance) UpdateFrom(ref js.Ref) {
	bindings.DistanceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Distance) Update(ref js.Ref) {
	bindings.DistanceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Distance) FreeMembers(recursive bool) {
}

type Entity struct {
	// Id is "Entity.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// Type is "Entity.type"
	//
	// Optional
	Type EntityType
	// EntityLabel is "Entity.entityLabel"
	//
	// Optional
	EntityLabel js.String
	// BoundingBox is "Entity.boundingBox"
	//
	// Optional
	//
	// NOTE: BoundingBox.FFI_USE MUST be set to true to get BoundingBox used.
	BoundingBox BoundingBox
	// Confidence is "Entity.confidence"
	//
	// Optional
	//
	// NOTE: FFI_USE_Confidence MUST be set to true to make this field effective.
	Confidence float64
	// Depth is "Entity.depth"
	//
	// Optional
	//
	// NOTE: Depth.FFI_USE MUST be set to true to get Depth used.
	Depth Distance

	FFI_USE_Id         bool // for Id.
	FFI_USE_Confidence bool // for Confidence.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Entity with all fields set.
func (p Entity) FromRef(ref js.Ref) Entity {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Entity in the application heap.
func (p Entity) New() js.Ref {
	return bindings.EntityJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Entity) UpdateFrom(ref js.Ref) {
	bindings.EntityJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Entity) Update(ref js.Ref) {
	bindings.EntityJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Entity) FreeMembers(recursive bool) {
	js.Free(
		p.EntityLabel.Ref(),
	)
	p.EntityLabel = p.EntityLabel.FromRef(js.Undefined)
	if recursive {
		p.BoundingBox.FreeMembers(true)
		p.Depth.FreeMembers(true)
	}
}

type PacketLatency struct {
	// PacketLabel is "PacketLatency.packetLabel"
	//
	// Optional
	PacketLabel js.String
	// LatencyUsec is "PacketLatency.latencyUsec"
	//
	// Optional
	//
	// NOTE: FFI_USE_LatencyUsec MUST be set to true to make this field effective.
	LatencyUsec int32

	FFI_USE_LatencyUsec bool // for LatencyUsec.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PacketLatency with all fields set.
func (p PacketLatency) FromRef(ref js.Ref) PacketLatency {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PacketLatency in the application heap.
func (p PacketLatency) New() js.Ref {
	return bindings.PacketLatencyJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PacketLatency) UpdateFrom(ref js.Ref) {
	bindings.PacketLatencyJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PacketLatency) Update(ref js.Ref) {
	bindings.PacketLatencyJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PacketLatency) FreeMembers(recursive bool) {
	js.Free(
		p.PacketLabel.Ref(),
	)
	p.PacketLabel = p.PacketLabel.FromRef(js.Undefined)
}

type LightCondition uint32

const (
	_ LightCondition = iota

	LightCondition_UNSPECIFIED
	LightCondition_NO_CHANGE
	LightCondition_TURNED_ON
	LightCondition_TURNED_OFF
	LightCondition_DIMMER
	LightCondition_BRIGHTER
	LightCondition_BLACK_FRAME
)

func (LightCondition) FromRef(str js.Ref) LightCondition {
	return LightCondition(bindings.ConstOfLightCondition(str))
}

func (x LightCondition) String() (string, bool) {
	switch x {
	case LightCondition_UNSPECIFIED:
		return "UNSPECIFIED", true
	case LightCondition_NO_CHANGE:
		return "NO_CHANGE", true
	case LightCondition_TURNED_ON:
		return "TURNED_ON", true
	case LightCondition_TURNED_OFF:
		return "TURNED_OFF", true
	case LightCondition_DIMMER:
		return "DIMMER", true
	case LightCondition_BRIGHTER:
		return "BRIGHTER", true
	case LightCondition_BLACK_FRAME:
		return "BLACK_FRAME", true
	default:
		return "", false
	}
}

type VideoHumanPresenceDetection struct {
	// HumanPresenceLikelihood is "VideoHumanPresenceDetection.humanPresenceLikelihood"
	//
	// Optional
	//
	// NOTE: FFI_USE_HumanPresenceLikelihood MUST be set to true to make this field effective.
	HumanPresenceLikelihood float64
	// MotionDetectedLikelihood is "VideoHumanPresenceDetection.motionDetectedLikelihood"
	//
	// Optional
	//
	// NOTE: FFI_USE_MotionDetectedLikelihood MUST be set to true to make this field effective.
	MotionDetectedLikelihood float64
	// LightCondition is "VideoHumanPresenceDetection.lightCondition"
	//
	// Optional
	LightCondition LightCondition
	// LightConditionLikelihood is "VideoHumanPresenceDetection.lightConditionLikelihood"
	//
	// Optional
	//
	// NOTE: FFI_USE_LightConditionLikelihood MUST be set to true to make this field effective.
	LightConditionLikelihood float64

	FFI_USE_HumanPresenceLikelihood  bool // for HumanPresenceLikelihood.
	FFI_USE_MotionDetectedLikelihood bool // for MotionDetectedLikelihood.
	FFI_USE_LightConditionLikelihood bool // for LightConditionLikelihood.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoHumanPresenceDetection with all fields set.
func (p VideoHumanPresenceDetection) FromRef(ref js.Ref) VideoHumanPresenceDetection {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VideoHumanPresenceDetection in the application heap.
func (p VideoHumanPresenceDetection) New() js.Ref {
	return bindings.VideoHumanPresenceDetectionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *VideoHumanPresenceDetection) UpdateFrom(ref js.Ref) {
	bindings.VideoHumanPresenceDetectionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VideoHumanPresenceDetection) Update(ref js.Ref) {
	bindings.VideoHumanPresenceDetectionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VideoHumanPresenceDetection) FreeMembers(recursive bool) {
}

type FramePerceptionType uint32

const (
	_ FramePerceptionType = iota

	FramePerceptionType_UNKNOWN_TYPE
	FramePerceptionType_FACE_DETECTION
	FramePerceptionType_PERSON_DETECTION
	FramePerceptionType_MOTION_DETECTION
)

func (FramePerceptionType) FromRef(str js.Ref) FramePerceptionType {
	return FramePerceptionType(bindings.ConstOfFramePerceptionType(str))
}

func (x FramePerceptionType) String() (string, bool) {
	switch x {
	case FramePerceptionType_UNKNOWN_TYPE:
		return "UNKNOWN_TYPE", true
	case FramePerceptionType_FACE_DETECTION:
		return "FACE_DETECTION", true
	case FramePerceptionType_PERSON_DETECTION:
		return "PERSON_DETECTION", true
	case FramePerceptionType_MOTION_DETECTION:
		return "MOTION_DETECTION", true
	default:
		return "", false
	}
}

type FramePerception struct {
	// FrameId is "FramePerception.frameId"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameId MUST be set to true to make this field effective.
	FrameId int32
	// FrameWidthInPx is "FramePerception.frameWidthInPx"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameWidthInPx MUST be set to true to make this field effective.
	FrameWidthInPx int32
	// FrameHeightInPx is "FramePerception.frameHeightInPx"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameHeightInPx MUST be set to true to make this field effective.
	FrameHeightInPx int32
	// Timestamp is "FramePerception.timestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timestamp MUST be set to true to make this field effective.
	Timestamp float64
	// Entities is "FramePerception.entities"
	//
	// Optional
	Entities js.Array[Entity]
	// PacketLatency is "FramePerception.packetLatency"
	//
	// Optional
	PacketLatency js.Array[PacketLatency]
	// VideoHumanPresenceDetection is "FramePerception.videoHumanPresenceDetection"
	//
	// Optional
	//
	// NOTE: VideoHumanPresenceDetection.FFI_USE MUST be set to true to get VideoHumanPresenceDetection used.
	VideoHumanPresenceDetection VideoHumanPresenceDetection
	// FramePerceptionTypes is "FramePerception.framePerceptionTypes"
	//
	// Optional
	FramePerceptionTypes js.Array[FramePerceptionType]

	FFI_USE_FrameId         bool // for FrameId.
	FFI_USE_FrameWidthInPx  bool // for FrameWidthInPx.
	FFI_USE_FrameHeightInPx bool // for FrameHeightInPx.
	FFI_USE_Timestamp       bool // for Timestamp.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FramePerception with all fields set.
func (p FramePerception) FromRef(ref js.Ref) FramePerception {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FramePerception in the application heap.
func (p FramePerception) New() js.Ref {
	return bindings.FramePerceptionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FramePerception) UpdateFrom(ref js.Ref) {
	bindings.FramePerceptionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FramePerception) Update(ref js.Ref) {
	bindings.FramePerceptionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FramePerception) FreeMembers(recursive bool) {
	js.Free(
		p.Entities.Ref(),
		p.PacketLatency.Ref(),
		p.FramePerceptionTypes.Ref(),
	)
	p.Entities = p.Entities.FromRef(js.Undefined)
	p.PacketLatency = p.PacketLatency.FromRef(js.Undefined)
	p.FramePerceptionTypes = p.FramePerceptionTypes.FromRef(js.Undefined)
	if recursive {
		p.VideoHumanPresenceDetection.FreeMembers(true)
	}
}

type ImageFormat uint32

const (
	_ ImageFormat = iota

	ImageFormat_RAW
	ImageFormat_PNG
	ImageFormat_JPEG
)

func (ImageFormat) FromRef(str js.Ref) ImageFormat {
	return ImageFormat(bindings.ConstOfImageFormat(str))
}

func (x ImageFormat) String() (string, bool) {
	switch x {
	case ImageFormat_RAW:
		return "RAW", true
	case ImageFormat_PNG:
		return "PNG", true
	case ImageFormat_JPEG:
		return "JPEG", true
	default:
		return "", false
	}
}

type ImageFrame struct {
	// Width is "ImageFrame.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int32
	// Height is "ImageFrame.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int32
	// Format is "ImageFrame.format"
	//
	// Optional
	Format ImageFormat
	// DataLength is "ImageFrame.dataLength"
	//
	// Optional
	//
	// NOTE: FFI_USE_DataLength MUST be set to true to make this field effective.
	DataLength int32
	// Frame is "ImageFrame.frame"
	//
	// Optional
	Frame js.ArrayBuffer

	FFI_USE_Width      bool // for Width.
	FFI_USE_Height     bool // for Height.
	FFI_USE_DataLength bool // for DataLength.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ImageFrame with all fields set.
func (p ImageFrame) FromRef(ref js.Ref) ImageFrame {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ImageFrame in the application heap.
func (p ImageFrame) New() js.Ref {
	return bindings.ImageFrameJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ImageFrame) UpdateFrom(ref js.Ref) {
	bindings.ImageFrameJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ImageFrame) Update(ref js.Ref) {
	bindings.ImageFrameJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ImageFrame) FreeMembers(recursive bool) {
	js.Free(
		p.Frame.Ref(),
	)
	p.Frame = p.Frame.FromRef(js.Undefined)
}

type Metadata struct {
	// VisualExperienceControllerVersion is "Metadata.visualExperienceControllerVersion"
	//
	// Optional
	VisualExperienceControllerVersion js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Metadata with all fields set.
func (p Metadata) FromRef(ref js.Ref) Metadata {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Metadata in the application heap.
func (p Metadata) New() js.Ref {
	return bindings.MetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Metadata) UpdateFrom(ref js.Ref) {
	bindings.MetadataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Metadata) Update(ref js.Ref) {
	bindings.MetadataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Metadata) FreeMembers(recursive bool) {
	js.Free(
		p.VisualExperienceControllerVersion.Ref(),
	)
	p.VisualExperienceControllerVersion = p.VisualExperienceControllerVersion.FromRef(js.Undefined)
}

type PerceptionSample struct {
	// FramePerception is "PerceptionSample.framePerception"
	//
	// Optional
	//
	// NOTE: FramePerception.FFI_USE MUST be set to true to get FramePerception used.
	FramePerception FramePerception
	// ImageFrame is "PerceptionSample.imageFrame"
	//
	// Optional
	//
	// NOTE: ImageFrame.FFI_USE MUST be set to true to get ImageFrame used.
	ImageFrame ImageFrame
	// AudioPerception is "PerceptionSample.audioPerception"
	//
	// Optional
	//
	// NOTE: AudioPerception.FFI_USE MUST be set to true to get AudioPerception used.
	AudioPerception AudioPerception
	// AudioVisualPerception is "PerceptionSample.audioVisualPerception"
	//
	// Optional
	//
	// NOTE: AudioVisualPerception.FFI_USE MUST be set to true to get AudioVisualPerception used.
	AudioVisualPerception AudioVisualPerception
	// Metadata is "PerceptionSample.metadata"
	//
	// Optional
	//
	// NOTE: Metadata.FFI_USE MUST be set to true to get Metadata used.
	Metadata Metadata

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PerceptionSample with all fields set.
func (p PerceptionSample) FromRef(ref js.Ref) PerceptionSample {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PerceptionSample in the application heap.
func (p PerceptionSample) New() js.Ref {
	return bindings.PerceptionSampleJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PerceptionSample) UpdateFrom(ref js.Ref) {
	bindings.PerceptionSampleJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PerceptionSample) Update(ref js.Ref) {
	bindings.PerceptionSampleJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PerceptionSample) FreeMembers(recursive bool) {
	if recursive {
		p.FramePerception.FreeMembers(true)
		p.ImageFrame.FreeMembers(true)
		p.AudioPerception.FreeMembers(true)
		p.AudioVisualPerception.FreeMembers(true)
		p.Metadata.FreeMembers(true)
	}
}

type Diagnostics struct {
	// ServiceError is "Diagnostics.serviceError"
	//
	// Optional
	ServiceError ServiceError
	// PerceptionSamples is "Diagnostics.perceptionSamples"
	//
	// Optional
	PerceptionSamples js.Array[PerceptionSample]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Diagnostics with all fields set.
func (p Diagnostics) FromRef(ref js.Ref) Diagnostics {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Diagnostics in the application heap.
func (p Diagnostics) New() js.Ref {
	return bindings.DiagnosticsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Diagnostics) UpdateFrom(ref js.Ref) {
	bindings.DiagnosticsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Diagnostics) Update(ref js.Ref) {
	bindings.DiagnosticsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Diagnostics) FreeMembers(recursive bool) {
	js.Free(
		p.PerceptionSamples.Ref(),
	)
	p.PerceptionSamples = p.PerceptionSamples.FromRef(js.Undefined)
}

type DiagnosticsCallbackFunc func(this js.Ref, diagnostics *Diagnostics) js.Ref

func (fn DiagnosticsCallbackFunc) Register() js.Func[func(diagnostics *Diagnostics)] {
	return js.RegisterCallback[func(diagnostics *Diagnostics)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DiagnosticsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Diagnostics
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DiagnosticsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, diagnostics *Diagnostics) js.Ref
	Arg T
}

func (cb *DiagnosticsCallback[T]) Register() js.Func[func(diagnostics *Diagnostics)] {
	return js.RegisterCallback[func(diagnostics *Diagnostics)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DiagnosticsCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Diagnostics
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type Feature uint32

const (
	_ Feature = iota

	Feature_AUTOZOOM
	Feature_HOTWORD_DETECTION
	Feature_OCCUPANCY_DETECTION
	Feature_EDGE_EMBEDDINGS
	Feature_SOFTWARE_CROPPING
)

func (Feature) FromRef(str js.Ref) Feature {
	return Feature(bindings.ConstOfFeature(str))
}

func (x Feature) String() (string, bool) {
	switch x {
	case Feature_AUTOZOOM:
		return "AUTOZOOM", true
	case Feature_HOTWORD_DETECTION:
		return "HOTWORD_DETECTION", true
	case Feature_OCCUPANCY_DETECTION:
		return "OCCUPANCY_DETECTION", true
	case Feature_EDGE_EMBEDDINGS:
		return "EDGE_EMBEDDINGS", true
	case Feature_SOFTWARE_CROPPING:
		return "SOFTWARE_CROPPING", true
	default:
		return "", false
	}
}

type MediaPerception struct {
	// Timestamp is "MediaPerception.timestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timestamp MUST be set to true to make this field effective.
	Timestamp float64
	// FramePerceptions is "MediaPerception.framePerceptions"
	//
	// Optional
	FramePerceptions js.Array[FramePerception]
	// AudioPerceptions is "MediaPerception.audioPerceptions"
	//
	// Optional
	AudioPerceptions js.Array[AudioPerception]
	// AudioVisualPerceptions is "MediaPerception.audioVisualPerceptions"
	//
	// Optional
	AudioVisualPerceptions js.Array[AudioVisualPerception]
	// Metadata is "MediaPerception.metadata"
	//
	// Optional
	//
	// NOTE: Metadata.FFI_USE MUST be set to true to get Metadata used.
	Metadata Metadata

	FFI_USE_Timestamp bool // for Timestamp.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaPerception with all fields set.
func (p MediaPerception) FromRef(ref js.Ref) MediaPerception {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaPerception in the application heap.
func (p MediaPerception) New() js.Ref {
	return bindings.MediaPerceptionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MediaPerception) UpdateFrom(ref js.Ref) {
	bindings.MediaPerceptionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaPerception) Update(ref js.Ref) {
	bindings.MediaPerceptionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaPerception) FreeMembers(recursive bool) {
	js.Free(
		p.FramePerceptions.Ref(),
		p.AudioPerceptions.Ref(),
		p.AudioVisualPerceptions.Ref(),
	)
	p.FramePerceptions = p.FramePerceptions.FromRef(js.Undefined)
	p.AudioPerceptions = p.AudioPerceptions.FromRef(js.Undefined)
	p.AudioVisualPerceptions = p.AudioVisualPerceptions.FromRef(js.Undefined)
	if recursive {
		p.Metadata.FreeMembers(true)
	}
}

type OneOf_String_Float64 struct {
	ref js.Ref
}

func (x OneOf_String_Float64) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_Float64) Free() {
	x.ref.Free()
}

func (x OneOf_String_Float64) FromRef(ref js.Ref) OneOf_String_Float64 {
	return OneOf_String_Float64{
		ref: ref,
	}
}

func (x OneOf_String_Float64) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_Float64) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

type NamedTemplateArgument struct {
	// Name is "NamedTemplateArgument.name"
	//
	// Optional
	Name js.String
	// Value is "NamedTemplateArgument.value"
	//
	// Optional
	Value OneOf_String_Float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NamedTemplateArgument with all fields set.
func (p NamedTemplateArgument) FromRef(ref js.Ref) NamedTemplateArgument {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NamedTemplateArgument in the application heap.
func (p NamedTemplateArgument) New() js.Ref {
	return bindings.NamedTemplateArgumentJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NamedTemplateArgument) UpdateFrom(ref js.Ref) {
	bindings.NamedTemplateArgumentJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NamedTemplateArgument) Update(ref js.Ref) {
	bindings.NamedTemplateArgumentJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NamedTemplateArgument) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.Value.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type ProcessStatus uint32

const (
	_ ProcessStatus = iota

	ProcessStatus_UNKNOWN
	ProcessStatus_STARTED
	ProcessStatus_STOPPED
	ProcessStatus_SERVICE_ERROR
)

func (ProcessStatus) FromRef(str js.Ref) ProcessStatus {
	return ProcessStatus(bindings.ConstOfProcessStatus(str))
}

func (x ProcessStatus) String() (string, bool) {
	switch x {
	case ProcessStatus_UNKNOWN:
		return "UNKNOWN", true
	case ProcessStatus_STARTED:
		return "STARTED", true
	case ProcessStatus_STOPPED:
		return "STOPPED", true
	case ProcessStatus_SERVICE_ERROR:
		return "SERVICE_ERROR", true
	default:
		return "", false
	}
}

type ProcessState struct {
	// Status is "ProcessState.status"
	//
	// Optional
	Status ProcessStatus
	// ServiceError is "ProcessState.serviceError"
	//
	// Optional
	ServiceError ServiceError

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProcessState with all fields set.
func (p ProcessState) FromRef(ref js.Ref) ProcessState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProcessState in the application heap.
func (p ProcessState) New() js.Ref {
	return bindings.ProcessStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProcessState) UpdateFrom(ref js.Ref) {
	bindings.ProcessStateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProcessState) Update(ref js.Ref) {
	bindings.ProcessStateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProcessState) FreeMembers(recursive bool) {
}

type ProcessStateCallbackFunc func(this js.Ref, processState *ProcessState) js.Ref

func (fn ProcessStateCallbackFunc) Register() js.Func[func(processState *ProcessState)] {
	return js.RegisterCallback[func(processState *ProcessState)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ProcessStateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ProcessState
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ProcessStateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, processState *ProcessState) js.Ref
	Arg T
}

func (cb *ProcessStateCallback[T]) Register() js.Func[func(processState *ProcessState)] {
	return js.RegisterCallback[func(processState *ProcessState)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ProcessStateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ProcessState
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type Status uint32

const (
	_ Status = iota

	Status_UNINITIALIZED
	Status_STARTED
	Status_RUNNING
	Status_SUSPENDED
	Status_RESTARTING
	Status_STOPPED
	Status_SERVICE_ERROR
)

func (Status) FromRef(str js.Ref) Status {
	return Status(bindings.ConstOfStatus(str))
}

func (x Status) String() (string, bool) {
	switch x {
	case Status_UNINITIALIZED:
		return "UNINITIALIZED", true
	case Status_STARTED:
		return "STARTED", true
	case Status_RUNNING:
		return "RUNNING", true
	case Status_SUSPENDED:
		return "SUSPENDED", true
	case Status_RESTARTING:
		return "RESTARTING", true
	case Status_STOPPED:
		return "STOPPED", true
	case Status_SERVICE_ERROR:
		return "SERVICE_ERROR", true
	default:
		return "", false
	}
}

type VideoStreamParam struct {
	// Id is "VideoStreamParam.id"
	//
	// Optional
	Id js.String
	// Width is "VideoStreamParam.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int32
	// Height is "VideoStreamParam.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int32
	// FrameRate is "VideoStreamParam.frameRate"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameRate MUST be set to true to make this field effective.
	FrameRate int32

	FFI_USE_Width     bool // for Width.
	FFI_USE_Height    bool // for Height.
	FFI_USE_FrameRate bool // for FrameRate.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoStreamParam with all fields set.
func (p VideoStreamParam) FromRef(ref js.Ref) VideoStreamParam {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VideoStreamParam in the application heap.
func (p VideoStreamParam) New() js.Ref {
	return bindings.VideoStreamParamJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *VideoStreamParam) UpdateFrom(ref js.Ref) {
	bindings.VideoStreamParamJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VideoStreamParam) Update(ref js.Ref) {
	bindings.VideoStreamParamJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VideoStreamParam) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
}

type Whiteboard struct {
	// TopLeft is "Whiteboard.topLeft"
	//
	// Optional
	//
	// NOTE: TopLeft.FFI_USE MUST be set to true to get TopLeft used.
	TopLeft Point
	// TopRight is "Whiteboard.topRight"
	//
	// Optional
	//
	// NOTE: TopRight.FFI_USE MUST be set to true to get TopRight used.
	TopRight Point
	// BottomLeft is "Whiteboard.bottomLeft"
	//
	// Optional
	//
	// NOTE: BottomLeft.FFI_USE MUST be set to true to get BottomLeft used.
	BottomLeft Point
	// BottomRight is "Whiteboard.bottomRight"
	//
	// Optional
	//
	// NOTE: BottomRight.FFI_USE MUST be set to true to get BottomRight used.
	BottomRight Point
	// AspectRatio is "Whiteboard.aspectRatio"
	//
	// Optional
	//
	// NOTE: FFI_USE_AspectRatio MUST be set to true to make this field effective.
	AspectRatio float64

	FFI_USE_AspectRatio bool // for AspectRatio.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Whiteboard with all fields set.
func (p Whiteboard) FromRef(ref js.Ref) Whiteboard {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Whiteboard in the application heap.
func (p Whiteboard) New() js.Ref {
	return bindings.WhiteboardJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Whiteboard) UpdateFrom(ref js.Ref) {
	bindings.WhiteboardJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Whiteboard) Update(ref js.Ref) {
	bindings.WhiteboardJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Whiteboard) FreeMembers(recursive bool) {
	if recursive {
		p.TopLeft.FreeMembers(true)
		p.TopRight.FreeMembers(true)
		p.BottomLeft.FreeMembers(true)
		p.BottomRight.FreeMembers(true)
	}
}

type State struct {
	// Status is "State.status"
	//
	// Optional
	Status Status
	// DeviceContext is "State.deviceContext"
	//
	// Optional
	DeviceContext js.String
	// ServiceError is "State.serviceError"
	//
	// Optional
	ServiceError ServiceError
	// VideoStreamParam is "State.videoStreamParam"
	//
	// Optional
	VideoStreamParam js.Array[VideoStreamParam]
	// Configuration is "State.configuration"
	//
	// Optional
	Configuration js.String
	// Whiteboard is "State.whiteboard"
	//
	// Optional
	//
	// NOTE: Whiteboard.FFI_USE MUST be set to true to get Whiteboard used.
	Whiteboard Whiteboard
	// Features is "State.features"
	//
	// Optional
	Features js.Array[Feature]
	// NamedTemplateArguments is "State.namedTemplateArguments"
	//
	// Optional
	NamedTemplateArguments js.Array[NamedTemplateArgument]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a State with all fields set.
func (p State) FromRef(ref js.Ref) State {
	p.UpdateFrom(ref)
	return p
}

// New creates a new State in the application heap.
func (p State) New() js.Ref {
	return bindings.StateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *State) UpdateFrom(ref js.Ref) {
	bindings.StateJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *State) Update(ref js.Ref) {
	bindings.StateJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *State) FreeMembers(recursive bool) {
	js.Free(
		p.DeviceContext.Ref(),
		p.VideoStreamParam.Ref(),
		p.Configuration.Ref(),
		p.Features.Ref(),
		p.NamedTemplateArguments.Ref(),
	)
	p.DeviceContext = p.DeviceContext.FromRef(js.Undefined)
	p.VideoStreamParam = p.VideoStreamParam.FromRef(js.Undefined)
	p.Configuration = p.Configuration.FromRef(js.Undefined)
	p.Features = p.Features.FromRef(js.Undefined)
	p.NamedTemplateArguments = p.NamedTemplateArguments.FromRef(js.Undefined)
	if recursive {
		p.Whiteboard.FreeMembers(true)
	}
}

type StateCallbackFunc func(this js.Ref, state *State) js.Ref

func (fn StateCallbackFunc) Register() js.Func[func(state *State)] {
	return js.RegisterCallback[func(state *State)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn StateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 State
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type StateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, state *State) js.Ref
	Arg T
}

func (cb *StateCallback[T]) Register() js.Func[func(state *State)] {
	return js.RegisterCallback[func(state *State)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *StateCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 State
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncGetDiagnostics returns true if the function "WEBEXT.mediaPerceptionPrivate.getDiagnostics" exists.
func HasFuncGetDiagnostics() bool {
	return js.True == bindings.HasFuncGetDiagnostics()
}

// FuncGetDiagnostics returns the function "WEBEXT.mediaPerceptionPrivate.getDiagnostics".
func FuncGetDiagnostics() (fn js.Func[func() js.Promise[Diagnostics]]) {
	bindings.FuncGetDiagnostics(
		js.Pointer(&fn),
	)
	return
}

// GetDiagnostics calls the function "WEBEXT.mediaPerceptionPrivate.getDiagnostics" directly.
func GetDiagnostics() (ret js.Promise[Diagnostics]) {
	bindings.CallGetDiagnostics(
		js.Pointer(&ret),
	)

	return
}

// TryGetDiagnostics calls the function "WEBEXT.mediaPerceptionPrivate.getDiagnostics"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDiagnostics() (ret js.Promise[Diagnostics], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDiagnostics(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetState returns true if the function "WEBEXT.mediaPerceptionPrivate.getState" exists.
func HasFuncGetState() bool {
	return js.True == bindings.HasFuncGetState()
}

// FuncGetState returns the function "WEBEXT.mediaPerceptionPrivate.getState".
func FuncGetState() (fn js.Func[func() js.Promise[State]]) {
	bindings.FuncGetState(
		js.Pointer(&fn),
	)
	return
}

// GetState calls the function "WEBEXT.mediaPerceptionPrivate.getState" directly.
func GetState() (ret js.Promise[State]) {
	bindings.CallGetState(
		js.Pointer(&ret),
	)

	return
}

// TryGetState calls the function "WEBEXT.mediaPerceptionPrivate.getState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetState() (ret js.Promise[State], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetState(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnMediaPerceptionEventCallbackFunc func(this js.Ref, mediaPerception *MediaPerception) js.Ref

func (fn OnMediaPerceptionEventCallbackFunc) Register() js.Func[func(mediaPerception *MediaPerception)] {
	return js.RegisterCallback[func(mediaPerception *MediaPerception)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMediaPerceptionEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MediaPerception
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnMediaPerceptionEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, mediaPerception *MediaPerception) js.Ref
	Arg T
}

func (cb *OnMediaPerceptionEventCallback[T]) Register() js.Func[func(mediaPerception *MediaPerception)] {
	return js.RegisterCallback[func(mediaPerception *MediaPerception)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMediaPerceptionEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MediaPerception
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnMediaPerception returns true if the function "WEBEXT.mediaPerceptionPrivate.onMediaPerception.addListener" exists.
func HasFuncOnMediaPerception() bool {
	return js.True == bindings.HasFuncOnMediaPerception()
}

// FuncOnMediaPerception returns the function "WEBEXT.mediaPerceptionPrivate.onMediaPerception.addListener".
func FuncOnMediaPerception() (fn js.Func[func(callback js.Func[func(mediaPerception *MediaPerception)])]) {
	bindings.FuncOnMediaPerception(
		js.Pointer(&fn),
	)
	return
}

// OnMediaPerception calls the function "WEBEXT.mediaPerceptionPrivate.onMediaPerception.addListener" directly.
func OnMediaPerception(callback js.Func[func(mediaPerception *MediaPerception)]) (ret js.Void) {
	bindings.CallOnMediaPerception(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMediaPerception calls the function "WEBEXT.mediaPerceptionPrivate.onMediaPerception.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMediaPerception(callback js.Func[func(mediaPerception *MediaPerception)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMediaPerception(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMediaPerception returns true if the function "WEBEXT.mediaPerceptionPrivate.onMediaPerception.removeListener" exists.
func HasFuncOffMediaPerception() bool {
	return js.True == bindings.HasFuncOffMediaPerception()
}

// FuncOffMediaPerception returns the function "WEBEXT.mediaPerceptionPrivate.onMediaPerception.removeListener".
func FuncOffMediaPerception() (fn js.Func[func(callback js.Func[func(mediaPerception *MediaPerception)])]) {
	bindings.FuncOffMediaPerception(
		js.Pointer(&fn),
	)
	return
}

// OffMediaPerception calls the function "WEBEXT.mediaPerceptionPrivate.onMediaPerception.removeListener" directly.
func OffMediaPerception(callback js.Func[func(mediaPerception *MediaPerception)]) (ret js.Void) {
	bindings.CallOffMediaPerception(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMediaPerception calls the function "WEBEXT.mediaPerceptionPrivate.onMediaPerception.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMediaPerception(callback js.Func[func(mediaPerception *MediaPerception)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMediaPerception(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMediaPerception returns true if the function "WEBEXT.mediaPerceptionPrivate.onMediaPerception.hasListener" exists.
func HasFuncHasOnMediaPerception() bool {
	return js.True == bindings.HasFuncHasOnMediaPerception()
}

// FuncHasOnMediaPerception returns the function "WEBEXT.mediaPerceptionPrivate.onMediaPerception.hasListener".
func FuncHasOnMediaPerception() (fn js.Func[func(callback js.Func[func(mediaPerception *MediaPerception)]) bool]) {
	bindings.FuncHasOnMediaPerception(
		js.Pointer(&fn),
	)
	return
}

// HasOnMediaPerception calls the function "WEBEXT.mediaPerceptionPrivate.onMediaPerception.hasListener" directly.
func HasOnMediaPerception(callback js.Func[func(mediaPerception *MediaPerception)]) (ret bool) {
	bindings.CallHasOnMediaPerception(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMediaPerception calls the function "WEBEXT.mediaPerceptionPrivate.onMediaPerception.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMediaPerception(callback js.Func[func(mediaPerception *MediaPerception)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMediaPerception(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSetAnalyticsComponent returns true if the function "WEBEXT.mediaPerceptionPrivate.setAnalyticsComponent" exists.
func HasFuncSetAnalyticsComponent() bool {
	return js.True == bindings.HasFuncSetAnalyticsComponent()
}

// FuncSetAnalyticsComponent returns the function "WEBEXT.mediaPerceptionPrivate.setAnalyticsComponent".
func FuncSetAnalyticsComponent() (fn js.Func[func(component Component) js.Promise[ComponentState]]) {
	bindings.FuncSetAnalyticsComponent(
		js.Pointer(&fn),
	)
	return
}

// SetAnalyticsComponent calls the function "WEBEXT.mediaPerceptionPrivate.setAnalyticsComponent" directly.
func SetAnalyticsComponent(component Component) (ret js.Promise[ComponentState]) {
	bindings.CallSetAnalyticsComponent(
		js.Pointer(&ret),
		js.Pointer(&component),
	)

	return
}

// TrySetAnalyticsComponent calls the function "WEBEXT.mediaPerceptionPrivate.setAnalyticsComponent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetAnalyticsComponent(component Component) (ret js.Promise[ComponentState], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetAnalyticsComponent(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&component),
	)

	return
}

// HasFuncSetComponentProcessState returns true if the function "WEBEXT.mediaPerceptionPrivate.setComponentProcessState" exists.
func HasFuncSetComponentProcessState() bool {
	return js.True == bindings.HasFuncSetComponentProcessState()
}

// FuncSetComponentProcessState returns the function "WEBEXT.mediaPerceptionPrivate.setComponentProcessState".
func FuncSetComponentProcessState() (fn js.Func[func(processState ProcessState) js.Promise[ProcessState]]) {
	bindings.FuncSetComponentProcessState(
		js.Pointer(&fn),
	)
	return
}

// SetComponentProcessState calls the function "WEBEXT.mediaPerceptionPrivate.setComponentProcessState" directly.
func SetComponentProcessState(processState ProcessState) (ret js.Promise[ProcessState]) {
	bindings.CallSetComponentProcessState(
		js.Pointer(&ret),
		js.Pointer(&processState),
	)

	return
}

// TrySetComponentProcessState calls the function "WEBEXT.mediaPerceptionPrivate.setComponentProcessState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetComponentProcessState(processState ProcessState) (ret js.Promise[ProcessState], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetComponentProcessState(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&processState),
	)

	return
}

// HasFuncSetState returns true if the function "WEBEXT.mediaPerceptionPrivate.setState" exists.
func HasFuncSetState() bool {
	return js.True == bindings.HasFuncSetState()
}

// FuncSetState returns the function "WEBEXT.mediaPerceptionPrivate.setState".
func FuncSetState() (fn js.Func[func(state State) js.Promise[State]]) {
	bindings.FuncSetState(
		js.Pointer(&fn),
	)
	return
}

// SetState calls the function "WEBEXT.mediaPerceptionPrivate.setState" directly.
func SetState(state State) (ret js.Promise[State]) {
	bindings.CallSetState(
		js.Pointer(&ret),
		js.Pointer(&state),
	)

	return
}

// TrySetState calls the function "WEBEXT.mediaPerceptionPrivate.setState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetState(state State) (ret js.Promise[State], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetState(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&state),
	)

	return
}

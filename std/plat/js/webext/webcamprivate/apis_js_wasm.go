// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webcamprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/webcamprivate/bindings"
)

type AutofocusState uint32

const (
	_ AutofocusState = iota

	AutofocusState_ON
	AutofocusState_OFF
)

func (AutofocusState) FromRef(str js.Ref) AutofocusState {
	return AutofocusState(bindings.ConstOfAutofocusState(str))
}

func (x AutofocusState) String() (string, bool) {
	switch x {
	case AutofocusState_ON:
		return "on", true
	case AutofocusState_OFF:
		return "off", true
	default:
		return "", false
	}
}

type PanDirection uint32

const (
	_ PanDirection = iota

	PanDirection_STOP
	PanDirection_RIGHT
	PanDirection_LEFT
)

func (PanDirection) FromRef(str js.Ref) PanDirection {
	return PanDirection(bindings.ConstOfPanDirection(str))
}

func (x PanDirection) String() (string, bool) {
	switch x {
	case PanDirection_STOP:
		return "stop", true
	case PanDirection_RIGHT:
		return "right", true
	case PanDirection_LEFT:
		return "left", true
	default:
		return "", false
	}
}

type Protocol uint32

const (
	_ Protocol = iota

	Protocol_VISCA
)

func (Protocol) FromRef(str js.Ref) Protocol {
	return Protocol(bindings.ConstOfProtocol(str))
}

func (x Protocol) String() (string, bool) {
	switch x {
	case Protocol_VISCA:
		return "visca", true
	default:
		return "", false
	}
}

type ProtocolConfiguration struct {
	// Protocol is "ProtocolConfiguration.protocol"
	//
	// Optional
	Protocol Protocol

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProtocolConfiguration with all fields set.
func (p ProtocolConfiguration) FromRef(ref js.Ref) ProtocolConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProtocolConfiguration in the application heap.
func (p ProtocolConfiguration) New() js.Ref {
	return bindings.ProtocolConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProtocolConfiguration) UpdateFrom(ref js.Ref) {
	bindings.ProtocolConfigurationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProtocolConfiguration) Update(ref js.Ref) {
	bindings.ProtocolConfigurationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProtocolConfiguration) FreeMembers(recursive bool) {
}

type Range struct {
	// Min is "Range.min"
	//
	// Optional
	//
	// NOTE: FFI_USE_Min MUST be set to true to make this field effective.
	Min float64
	// Max is "Range.max"
	//
	// Optional
	//
	// NOTE: FFI_USE_Max MUST be set to true to make this field effective.
	Max float64

	FFI_USE_Min bool // for Min.
	FFI_USE_Max bool // for Max.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Range with all fields set.
func (p Range) FromRef(ref js.Ref) Range {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Range in the application heap.
func (p Range) New() js.Ref {
	return bindings.RangeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Range) UpdateFrom(ref js.Ref) {
	bindings.RangeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Range) Update(ref js.Ref) {
	bindings.RangeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Range) FreeMembers(recursive bool) {
}

type TiltDirection uint32

const (
	_ TiltDirection = iota

	TiltDirection_STOP
	TiltDirection_UP
	TiltDirection_DOWN
)

func (TiltDirection) FromRef(str js.Ref) TiltDirection {
	return TiltDirection(bindings.ConstOfTiltDirection(str))
}

func (x TiltDirection) String() (string, bool) {
	switch x {
	case TiltDirection_STOP:
		return "stop", true
	case TiltDirection_UP:
		return "up", true
	case TiltDirection_DOWN:
		return "down", true
	default:
		return "", false
	}
}

type WebcamConfiguration struct {
	// Pan is "WebcamConfiguration.pan"
	//
	// Optional
	//
	// NOTE: FFI_USE_Pan MUST be set to true to make this field effective.
	Pan float64
	// PanSpeed is "WebcamConfiguration.panSpeed"
	//
	// Optional
	//
	// NOTE: FFI_USE_PanSpeed MUST be set to true to make this field effective.
	PanSpeed float64
	// PanDirection is "WebcamConfiguration.panDirection"
	//
	// Optional
	PanDirection PanDirection
	// Tilt is "WebcamConfiguration.tilt"
	//
	// Optional
	//
	// NOTE: FFI_USE_Tilt MUST be set to true to make this field effective.
	Tilt float64
	// TiltSpeed is "WebcamConfiguration.tiltSpeed"
	//
	// Optional
	//
	// NOTE: FFI_USE_TiltSpeed MUST be set to true to make this field effective.
	TiltSpeed float64
	// TiltDirection is "WebcamConfiguration.tiltDirection"
	//
	// Optional
	TiltDirection TiltDirection
	// Zoom is "WebcamConfiguration.zoom"
	//
	// Optional
	//
	// NOTE: FFI_USE_Zoom MUST be set to true to make this field effective.
	Zoom float64
	// AutofocusState is "WebcamConfiguration.autofocusState"
	//
	// Optional
	AutofocusState AutofocusState
	// Focus is "WebcamConfiguration.focus"
	//
	// Optional
	//
	// NOTE: FFI_USE_Focus MUST be set to true to make this field effective.
	Focus float64

	FFI_USE_Pan       bool // for Pan.
	FFI_USE_PanSpeed  bool // for PanSpeed.
	FFI_USE_Tilt      bool // for Tilt.
	FFI_USE_TiltSpeed bool // for TiltSpeed.
	FFI_USE_Zoom      bool // for Zoom.
	FFI_USE_Focus     bool // for Focus.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebcamConfiguration with all fields set.
func (p WebcamConfiguration) FromRef(ref js.Ref) WebcamConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebcamConfiguration in the application heap.
func (p WebcamConfiguration) New() js.Ref {
	return bindings.WebcamConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *WebcamConfiguration) UpdateFrom(ref js.Ref) {
	bindings.WebcamConfigurationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WebcamConfiguration) Update(ref js.Ref) {
	bindings.WebcamConfigurationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WebcamConfiguration) FreeMembers(recursive bool) {
}

type WebcamConfigurationCallbackFunc func(this js.Ref, configuration *WebcamCurrentConfiguration) js.Ref

func (fn WebcamConfigurationCallbackFunc) Register() js.Func[func(configuration *WebcamCurrentConfiguration)] {
	return js.RegisterCallback[func(configuration *WebcamCurrentConfiguration)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn WebcamConfigurationCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 WebcamCurrentConfiguration
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

type WebcamConfigurationCallback[T any] struct {
	Fn  func(arg T, this js.Ref, configuration *WebcamCurrentConfiguration) js.Ref
	Arg T
}

func (cb *WebcamConfigurationCallback[T]) Register() js.Func[func(configuration *WebcamCurrentConfiguration)] {
	return js.RegisterCallback[func(configuration *WebcamCurrentConfiguration)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *WebcamConfigurationCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 WebcamCurrentConfiguration
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

type WebcamCurrentConfiguration struct {
	// Pan is "WebcamCurrentConfiguration.pan"
	//
	// Optional
	//
	// NOTE: FFI_USE_Pan MUST be set to true to make this field effective.
	Pan float64
	// Tilt is "WebcamCurrentConfiguration.tilt"
	//
	// Optional
	//
	// NOTE: FFI_USE_Tilt MUST be set to true to make this field effective.
	Tilt float64
	// Zoom is "WebcamCurrentConfiguration.zoom"
	//
	// Optional
	//
	// NOTE: FFI_USE_Zoom MUST be set to true to make this field effective.
	Zoom float64
	// Focus is "WebcamCurrentConfiguration.focus"
	//
	// Optional
	//
	// NOTE: FFI_USE_Focus MUST be set to true to make this field effective.
	Focus float64
	// PanRange is "WebcamCurrentConfiguration.panRange"
	//
	// Optional
	//
	// NOTE: PanRange.FFI_USE MUST be set to true to get PanRange used.
	PanRange Range
	// TiltRange is "WebcamCurrentConfiguration.tiltRange"
	//
	// Optional
	//
	// NOTE: TiltRange.FFI_USE MUST be set to true to get TiltRange used.
	TiltRange Range
	// ZoomRange is "WebcamCurrentConfiguration.zoomRange"
	//
	// Optional
	//
	// NOTE: ZoomRange.FFI_USE MUST be set to true to get ZoomRange used.
	ZoomRange Range
	// FocusRange is "WebcamCurrentConfiguration.focusRange"
	//
	// Optional
	//
	// NOTE: FocusRange.FFI_USE MUST be set to true to get FocusRange used.
	FocusRange Range

	FFI_USE_Pan   bool // for Pan.
	FFI_USE_Tilt  bool // for Tilt.
	FFI_USE_Zoom  bool // for Zoom.
	FFI_USE_Focus bool // for Focus.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebcamCurrentConfiguration with all fields set.
func (p WebcamCurrentConfiguration) FromRef(ref js.Ref) WebcamCurrentConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebcamCurrentConfiguration in the application heap.
func (p WebcamCurrentConfiguration) New() js.Ref {
	return bindings.WebcamCurrentConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *WebcamCurrentConfiguration) UpdateFrom(ref js.Ref) {
	bindings.WebcamCurrentConfigurationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WebcamCurrentConfiguration) Update(ref js.Ref) {
	bindings.WebcamCurrentConfigurationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WebcamCurrentConfiguration) FreeMembers(recursive bool) {
	if recursive {
		p.PanRange.FreeMembers(true)
		p.TiltRange.FreeMembers(true)
		p.ZoomRange.FreeMembers(true)
		p.FocusRange.FreeMembers(true)
	}
}

type WebcamIdCallbackFunc func(this js.Ref, webcamId js.String) js.Ref

func (fn WebcamIdCallbackFunc) Register() js.Func[func(webcamId js.String)] {
	return js.RegisterCallback[func(webcamId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn WebcamIdCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type WebcamIdCallback[T any] struct {
	Fn  func(arg T, this js.Ref, webcamId js.String) js.Ref
	Arg T
}

func (cb *WebcamIdCallback[T]) Register() js.Func[func(webcamId js.String)] {
	return js.RegisterCallback[func(webcamId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *WebcamIdCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncCloseWebcam returns true if the function "WEBEXT.webcamPrivate.closeWebcam" exists.
func HasFuncCloseWebcam() bool {
	return js.True == bindings.HasFuncCloseWebcam()
}

// FuncCloseWebcam returns the function "WEBEXT.webcamPrivate.closeWebcam".
func FuncCloseWebcam() (fn js.Func[func(webcamId js.String)]) {
	bindings.FuncCloseWebcam(
		js.Pointer(&fn),
	)
	return
}

// CloseWebcam calls the function "WEBEXT.webcamPrivate.closeWebcam" directly.
func CloseWebcam(webcamId js.String) (ret js.Void) {
	bindings.CallCloseWebcam(
		js.Pointer(&ret),
		webcamId.Ref(),
	)

	return
}

// TryCloseWebcam calls the function "WEBEXT.webcamPrivate.closeWebcam"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCloseWebcam(webcamId js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCloseWebcam(
		js.Pointer(&ret), js.Pointer(&exception),
		webcamId.Ref(),
	)

	return
}

// HasFuncGet returns true if the function "WEBEXT.webcamPrivate.get" exists.
func HasFuncGet() bool {
	return js.True == bindings.HasFuncGet()
}

// FuncGet returns the function "WEBEXT.webcamPrivate.get".
func FuncGet() (fn js.Func[func(webcamId js.String) js.Promise[WebcamCurrentConfiguration]]) {
	bindings.FuncGet(
		js.Pointer(&fn),
	)
	return
}

// Get calls the function "WEBEXT.webcamPrivate.get" directly.
func Get(webcamId js.String) (ret js.Promise[WebcamCurrentConfiguration]) {
	bindings.CallGet(
		js.Pointer(&ret),
		webcamId.Ref(),
	)

	return
}

// TryGet calls the function "WEBEXT.webcamPrivate.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGet(webcamId js.String) (ret js.Promise[WebcamCurrentConfiguration], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGet(
		js.Pointer(&ret), js.Pointer(&exception),
		webcamId.Ref(),
	)

	return
}

// HasFuncOpenSerialWebcam returns true if the function "WEBEXT.webcamPrivate.openSerialWebcam" exists.
func HasFuncOpenSerialWebcam() bool {
	return js.True == bindings.HasFuncOpenSerialWebcam()
}

// FuncOpenSerialWebcam returns the function "WEBEXT.webcamPrivate.openSerialWebcam".
func FuncOpenSerialWebcam() (fn js.Func[func(path js.String, protocol ProtocolConfiguration) js.Promise[js.String]]) {
	bindings.FuncOpenSerialWebcam(
		js.Pointer(&fn),
	)
	return
}

// OpenSerialWebcam calls the function "WEBEXT.webcamPrivate.openSerialWebcam" directly.
func OpenSerialWebcam(path js.String, protocol ProtocolConfiguration) (ret js.Promise[js.String]) {
	bindings.CallOpenSerialWebcam(
		js.Pointer(&ret),
		path.Ref(),
		js.Pointer(&protocol),
	)

	return
}

// TryOpenSerialWebcam calls the function "WEBEXT.webcamPrivate.openSerialWebcam"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenSerialWebcam(path js.String, protocol ProtocolConfiguration) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenSerialWebcam(
		js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		js.Pointer(&protocol),
	)

	return
}

// HasFuncReset returns true if the function "WEBEXT.webcamPrivate.reset" exists.
func HasFuncReset() bool {
	return js.True == bindings.HasFuncReset()
}

// FuncReset returns the function "WEBEXT.webcamPrivate.reset".
func FuncReset() (fn js.Func[func(webcamId js.String, config WebcamConfiguration) js.Promise[WebcamCurrentConfiguration]]) {
	bindings.FuncReset(
		js.Pointer(&fn),
	)
	return
}

// Reset calls the function "WEBEXT.webcamPrivate.reset" directly.
func Reset(webcamId js.String, config WebcamConfiguration) (ret js.Promise[WebcamCurrentConfiguration]) {
	bindings.CallReset(
		js.Pointer(&ret),
		webcamId.Ref(),
		js.Pointer(&config),
	)

	return
}

// TryReset calls the function "WEBEXT.webcamPrivate.reset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReset(webcamId js.String, config WebcamConfiguration) (ret js.Promise[WebcamCurrentConfiguration], exception js.Any, ok bool) {
	ok = js.True == bindings.TryReset(
		js.Pointer(&ret), js.Pointer(&exception),
		webcamId.Ref(),
		js.Pointer(&config),
	)

	return
}

// HasFuncRestoreCameraPreset returns true if the function "WEBEXT.webcamPrivate.restoreCameraPreset" exists.
func HasFuncRestoreCameraPreset() bool {
	return js.True == bindings.HasFuncRestoreCameraPreset()
}

// FuncRestoreCameraPreset returns the function "WEBEXT.webcamPrivate.restoreCameraPreset".
func FuncRestoreCameraPreset() (fn js.Func[func(webcamId js.String, presetNumber float64) js.Promise[WebcamCurrentConfiguration]]) {
	bindings.FuncRestoreCameraPreset(
		js.Pointer(&fn),
	)
	return
}

// RestoreCameraPreset calls the function "WEBEXT.webcamPrivate.restoreCameraPreset" directly.
func RestoreCameraPreset(webcamId js.String, presetNumber float64) (ret js.Promise[WebcamCurrentConfiguration]) {
	bindings.CallRestoreCameraPreset(
		js.Pointer(&ret),
		webcamId.Ref(),
		float64(presetNumber),
	)

	return
}

// TryRestoreCameraPreset calls the function "WEBEXT.webcamPrivate.restoreCameraPreset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRestoreCameraPreset(webcamId js.String, presetNumber float64) (ret js.Promise[WebcamCurrentConfiguration], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRestoreCameraPreset(
		js.Pointer(&ret), js.Pointer(&exception),
		webcamId.Ref(),
		float64(presetNumber),
	)

	return
}

// HasFuncSet returns true if the function "WEBEXT.webcamPrivate.set" exists.
func HasFuncSet() bool {
	return js.True == bindings.HasFuncSet()
}

// FuncSet returns the function "WEBEXT.webcamPrivate.set".
func FuncSet() (fn js.Func[func(webcamId js.String, config WebcamConfiguration) js.Promise[WebcamCurrentConfiguration]]) {
	bindings.FuncSet(
		js.Pointer(&fn),
	)
	return
}

// Set calls the function "WEBEXT.webcamPrivate.set" directly.
func Set(webcamId js.String, config WebcamConfiguration) (ret js.Promise[WebcamCurrentConfiguration]) {
	bindings.CallSet(
		js.Pointer(&ret),
		webcamId.Ref(),
		js.Pointer(&config),
	)

	return
}

// TrySet calls the function "WEBEXT.webcamPrivate.set"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySet(webcamId js.String, config WebcamConfiguration) (ret js.Promise[WebcamCurrentConfiguration], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySet(
		js.Pointer(&ret), js.Pointer(&exception),
		webcamId.Ref(),
		js.Pointer(&config),
	)

	return
}

// HasFuncSetCameraPreset returns true if the function "WEBEXT.webcamPrivate.setCameraPreset" exists.
func HasFuncSetCameraPreset() bool {
	return js.True == bindings.HasFuncSetCameraPreset()
}

// FuncSetCameraPreset returns the function "WEBEXT.webcamPrivate.setCameraPreset".
func FuncSetCameraPreset() (fn js.Func[func(webcamId js.String, presetNumber float64) js.Promise[WebcamCurrentConfiguration]]) {
	bindings.FuncSetCameraPreset(
		js.Pointer(&fn),
	)
	return
}

// SetCameraPreset calls the function "WEBEXT.webcamPrivate.setCameraPreset" directly.
func SetCameraPreset(webcamId js.String, presetNumber float64) (ret js.Promise[WebcamCurrentConfiguration]) {
	bindings.CallSetCameraPreset(
		js.Pointer(&ret),
		webcamId.Ref(),
		float64(presetNumber),
	)

	return
}

// TrySetCameraPreset calls the function "WEBEXT.webcamPrivate.setCameraPreset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetCameraPreset(webcamId js.String, presetNumber float64) (ret js.Promise[WebcamCurrentConfiguration], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetCameraPreset(
		js.Pointer(&ret), js.Pointer(&exception),
		webcamId.Ref(),
		float64(presetNumber),
	)

	return
}

// HasFuncSetHome returns true if the function "WEBEXT.webcamPrivate.setHome" exists.
func HasFuncSetHome() bool {
	return js.True == bindings.HasFuncSetHome()
}

// FuncSetHome returns the function "WEBEXT.webcamPrivate.setHome".
func FuncSetHome() (fn js.Func[func(webcamId js.String) js.Promise[WebcamCurrentConfiguration]]) {
	bindings.FuncSetHome(
		js.Pointer(&fn),
	)
	return
}

// SetHome calls the function "WEBEXT.webcamPrivate.setHome" directly.
func SetHome(webcamId js.String) (ret js.Promise[WebcamCurrentConfiguration]) {
	bindings.CallSetHome(
		js.Pointer(&ret),
		webcamId.Ref(),
	)

	return
}

// TrySetHome calls the function "WEBEXT.webcamPrivate.setHome"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetHome(webcamId js.String) (ret js.Promise[WebcamCurrentConfiguration], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetHome(
		js.Pointer(&ret), js.Pointer(&exception),
		webcamId.Ref(),
	)

	return
}

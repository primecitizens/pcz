// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

type DeviceOrientationEventInit struct {
	// Alpha is "DeviceOrientationEventInit.alpha"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_Alpha MUST be set to true to make this field effective.
	Alpha float64
	// Beta is "DeviceOrientationEventInit.beta"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_Beta MUST be set to true to make this field effective.
	Beta float64
	// Gamma is "DeviceOrientationEventInit.gamma"
	//
	// Optional, defaults to null.
	//
	// NOTE: FFI_USE_Gamma MUST be set to true to make this field effective.
	Gamma float64
	// Absolute is "DeviceOrientationEventInit.absolute"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Absolute MUST be set to true to make this field effective.
	Absolute bool
	// Bubbles is "DeviceOrientationEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "DeviceOrientationEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "DeviceOrientationEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Alpha      bool // for Alpha.
	FFI_USE_Beta       bool // for Beta.
	FFI_USE_Gamma      bool // for Gamma.
	FFI_USE_Absolute   bool // for Absolute.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeviceOrientationEventInit with all fields set.
func (p DeviceOrientationEventInit) FromRef(ref js.Ref) DeviceOrientationEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeviceOrientationEventInit in the application heap.
func (p DeviceOrientationEventInit) New() js.Ref {
	return bindings.DeviceOrientationEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeviceOrientationEventInit) UpdateFrom(ref js.Ref) {
	bindings.DeviceOrientationEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeviceOrientationEventInit) Update(ref js.Ref) {
	bindings.DeviceOrientationEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeviceOrientationEventInit) FreeMembers(recursive bool) {
}

func NewDeviceOrientationEvent(typ js.String, eventInitDict DeviceOrientationEventInit) (ret DeviceOrientationEvent) {
	ret.ref = bindings.NewDeviceOrientationEventByDeviceOrientationEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewDeviceOrientationEventByDeviceOrientationEvent1(typ js.String) (ret DeviceOrientationEvent) {
	ret.ref = bindings.NewDeviceOrientationEventByDeviceOrientationEvent1(
		typ.Ref())
	return
}

type DeviceOrientationEvent struct {
	Event
}

func (this DeviceOrientationEvent) Once() DeviceOrientationEvent {
	this.ref.Once()
	return this
}

func (this DeviceOrientationEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this DeviceOrientationEvent) FromRef(ref js.Ref) DeviceOrientationEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this DeviceOrientationEvent) Free() {
	this.ref.Free()
}

// Alpha returns the value of property "DeviceOrientationEvent.alpha".
//
// It returns ok=false if there is no such property.
func (this DeviceOrientationEvent) Alpha() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceOrientationEventAlpha(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Beta returns the value of property "DeviceOrientationEvent.beta".
//
// It returns ok=false if there is no such property.
func (this DeviceOrientationEvent) Beta() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceOrientationEventBeta(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Gamma returns the value of property "DeviceOrientationEvent.gamma".
//
// It returns ok=false if there is no such property.
func (this DeviceOrientationEvent) Gamma() (ret float64, ok bool) {
	ok = js.True == bindings.GetDeviceOrientationEventGamma(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Absolute returns the value of property "DeviceOrientationEvent.absolute".
//
// It returns ok=false if there is no such property.
func (this DeviceOrientationEvent) Absolute() (ret bool, ok bool) {
	ok = js.True == bindings.GetDeviceOrientationEventAbsolute(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncRequestPermission returns true if the static method "DeviceOrientationEvent.requestPermission" exists.
func (this DeviceOrientationEvent) HasFuncRequestPermission() bool {
	return js.True == bindings.HasFuncDeviceOrientationEventRequestPermission(
		this.ref,
	)
}

// FuncRequestPermission returns the static method "DeviceOrientationEvent.requestPermission".
func (this DeviceOrientationEvent) FuncRequestPermission() (fn js.Func[func() js.Promise[PermissionState]]) {
	bindings.FuncDeviceOrientationEventRequestPermission(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestPermission calls the static method "DeviceOrientationEvent.requestPermission".
func (this DeviceOrientationEvent) RequestPermission() (ret js.Promise[PermissionState]) {
	bindings.CallDeviceOrientationEventRequestPermission(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestPermission calls the static method "DeviceOrientationEvent.requestPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DeviceOrientationEvent) TryRequestPermission() (ret js.Promise[PermissionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeviceOrientationEventRequestPermission(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type DirectionSetting uint32

const (
	_ DirectionSetting = iota

	DirectionSetting_
	DirectionSetting_RL
	DirectionSetting_LR
)

func (DirectionSetting) FromRef(str js.Ref) DirectionSetting {
	return DirectionSetting(bindings.ConstOfDirectionSetting(str))
}

func (x DirectionSetting) String() (string, bool) {
	switch x {
	case DirectionSetting_:
		return "", true
	case DirectionSetting_RL:
		return "rl", true
	case DirectionSetting_LR:
		return "lr", true
	default:
		return "", false
	}
}

type DisplayCaptureSurfaceType uint32

const (
	_ DisplayCaptureSurfaceType = iota

	DisplayCaptureSurfaceType_MONITOR
	DisplayCaptureSurfaceType_WINDOW
	DisplayCaptureSurfaceType_BROWSER
)

func (DisplayCaptureSurfaceType) FromRef(str js.Ref) DisplayCaptureSurfaceType {
	return DisplayCaptureSurfaceType(bindings.ConstOfDisplayCaptureSurfaceType(str))
}

func (x DisplayCaptureSurfaceType) String() (string, bool) {
	switch x {
	case DisplayCaptureSurfaceType_MONITOR:
		return "monitor", true
	case DisplayCaptureSurfaceType_WINDOW:
		return "window", true
	case DisplayCaptureSurfaceType_BROWSER:
		return "browser", true
	default:
		return "", false
	}
}

type DocumentPictureInPictureEventInit struct {
	// Window is "DocumentPictureInPictureEventInit.window"
	//
	// Required
	Window Window
	// Bubbles is "DocumentPictureInPictureEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "DocumentPictureInPictureEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "DocumentPictureInPictureEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DocumentPictureInPictureEventInit with all fields set.
func (p DocumentPictureInPictureEventInit) FromRef(ref js.Ref) DocumentPictureInPictureEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DocumentPictureInPictureEventInit in the application heap.
func (p DocumentPictureInPictureEventInit) New() js.Ref {
	return bindings.DocumentPictureInPictureEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DocumentPictureInPictureEventInit) UpdateFrom(ref js.Ref) {
	bindings.DocumentPictureInPictureEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DocumentPictureInPictureEventInit) Update(ref js.Ref) {
	bindings.DocumentPictureInPictureEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DocumentPictureInPictureEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Window.Ref(),
	)
	p.Window = p.Window.FromRef(js.Undefined)
}

func NewDocumentPictureInPictureEvent(typ js.String, eventInitDict DocumentPictureInPictureEventInit) (ret DocumentPictureInPictureEvent) {
	ret.ref = bindings.NewDocumentPictureInPictureEventByDocumentPictureInPictureEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type DocumentPictureInPictureEvent struct {
	Event
}

func (this DocumentPictureInPictureEvent) Once() DocumentPictureInPictureEvent {
	this.ref.Once()
	return this
}

func (this DocumentPictureInPictureEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this DocumentPictureInPictureEvent) FromRef(ref js.Ref) DocumentPictureInPictureEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this DocumentPictureInPictureEvent) Free() {
	this.ref.Free()
}

// Window returns the value of property "DocumentPictureInPictureEvent.window".
//
// It returns ok=false if there is no such property.
func (this DocumentPictureInPictureEvent) Window() (ret Window, ok bool) {
	ok = js.True == bindings.GetDocumentPictureInPictureEventWindow(
		this.ref, js.Pointer(&ret),
	)
	return
}

type DragEventInit struct {
	// DataTransfer is "DragEventInit.dataTransfer"
	//
	// Optional, defaults to null.
	DataTransfer DataTransfer
	// MovementX is "DragEventInit.movementX"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_MovementX MUST be set to true to make this field effective.
	MovementX float64
	// MovementY is "DragEventInit.movementY"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_MovementY MUST be set to true to make this field effective.
	MovementY float64

	FFI_USE_MovementX bool // for MovementX.
	FFI_USE_MovementY bool // for MovementY.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DragEventInit with all fields set.
func (p DragEventInit) FromRef(ref js.Ref) DragEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DragEventInit in the application heap.
func (p DragEventInit) New() js.Ref {
	return bindings.DragEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DragEventInit) UpdateFrom(ref js.Ref) {
	bindings.DragEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DragEventInit) Update(ref js.Ref) {
	bindings.DragEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DragEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.DataTransfer.Ref(),
	)
	p.DataTransfer = p.DataTransfer.FromRef(js.Undefined)
}

func NewDragEvent(typ js.String, eventInitDict DragEventInit) (ret DragEvent) {
	ret.ref = bindings.NewDragEventByDragEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewDragEventByDragEvent1(typ js.String) (ret DragEvent) {
	ret.ref = bindings.NewDragEventByDragEvent1(
		typ.Ref())
	return
}

type DragEvent struct {
	MouseEvent
}

func (this DragEvent) Once() DragEvent {
	this.ref.Once()
	return this
}

func (this DragEvent) Ref() js.Ref {
	return this.MouseEvent.Ref()
}

func (this DragEvent) FromRef(ref js.Ref) DragEvent {
	this.MouseEvent = this.MouseEvent.FromRef(ref)
	return this
}

func (this DragEvent) Free() {
	this.ref.Free()
}

// DataTransfer returns the value of property "DragEvent.dataTransfer".
//
// It returns ok=false if there is no such property.
func (this DragEvent) DataTransfer() (ret DataTransfer, ok bool) {
	ok = js.True == bindings.GetDragEventDataTransfer(
		this.ref, js.Pointer(&ret),
	)
	return
}

const (
	EXT_blend_minmax_MIN_EXT GLenum = 0x8007
	EXT_blend_minmax_MAX_EXT GLenum = 0x8008
)

type EXT_blend_minmax struct {
	ref js.Ref
}

func (this EXT_blend_minmax) Once() EXT_blend_minmax {
	this.ref.Once()
	return this
}

func (this EXT_blend_minmax) Ref() js.Ref {
	return this.ref
}

func (this EXT_blend_minmax) FromRef(ref js.Ref) EXT_blend_minmax {
	this.ref = ref
	return this
}

func (this EXT_blend_minmax) Free() {
	this.ref.Free()
}

type EXT_color_buffer_float struct {
	ref js.Ref
}

func (this EXT_color_buffer_float) Once() EXT_color_buffer_float {
	this.ref.Once()
	return this
}

func (this EXT_color_buffer_float) Ref() js.Ref {
	return this.ref
}

func (this EXT_color_buffer_float) FromRef(ref js.Ref) EXT_color_buffer_float {
	this.ref = ref
	return this
}

func (this EXT_color_buffer_float) Free() {
	this.ref.Free()
}

const (
	EXT_color_buffer_half_float_RGBA16F_EXT                               GLenum = 0x881A
	EXT_color_buffer_half_float_RGB16F_EXT                                GLenum = 0x881B
	EXT_color_buffer_half_float_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE_EXT GLenum = 0x8211
	EXT_color_buffer_half_float_UNSIGNED_NORMALIZED_EXT                   GLenum = 0x8C17
)

type EXT_color_buffer_half_float struct {
	ref js.Ref
}

func (this EXT_color_buffer_half_float) Once() EXT_color_buffer_half_float {
	this.ref.Once()
	return this
}

func (this EXT_color_buffer_half_float) Ref() js.Ref {
	return this.ref
}

func (this EXT_color_buffer_half_float) FromRef(ref js.Ref) EXT_color_buffer_half_float {
	this.ref = ref
	return this
}

func (this EXT_color_buffer_half_float) Free() {
	this.ref.Free()
}

const (
	EXT_disjoint_timer_query_QUERY_COUNTER_BITS_EXT     GLenum = 0x8864
	EXT_disjoint_timer_query_CURRENT_QUERY_EXT          GLenum = 0x8865
	EXT_disjoint_timer_query_QUERY_RESULT_EXT           GLenum = 0x8866
	EXT_disjoint_timer_query_QUERY_RESULT_AVAILABLE_EXT GLenum = 0x8867
	EXT_disjoint_timer_query_TIME_ELAPSED_EXT           GLenum = 0x88BF
	EXT_disjoint_timer_query_TIMESTAMP_EXT              GLenum = 0x8E28
	EXT_disjoint_timer_query_GPU_DISJOINT_EXT           GLenum = 0x8FBB
)

type WebGLTimerQueryEXT struct {
	WebGLObject
}

func (this WebGLTimerQueryEXT) Once() WebGLTimerQueryEXT {
	this.ref.Once()
	return this
}

func (this WebGLTimerQueryEXT) Ref() js.Ref {
	return this.WebGLObject.Ref()
}

func (this WebGLTimerQueryEXT) FromRef(ref js.Ref) WebGLTimerQueryEXT {
	this.WebGLObject = this.WebGLObject.FromRef(ref)
	return this
}

func (this WebGLTimerQueryEXT) Free() {
	this.ref.Free()
}

type EXT_disjoint_timer_query struct {
	ref js.Ref
}

func (this EXT_disjoint_timer_query) Once() EXT_disjoint_timer_query {
	this.ref.Once()
	return this
}

func (this EXT_disjoint_timer_query) Ref() js.Ref {
	return this.ref
}

func (this EXT_disjoint_timer_query) FromRef(ref js.Ref) EXT_disjoint_timer_query {
	this.ref = ref
	return this
}

func (this EXT_disjoint_timer_query) Free() {
	this.ref.Free()
}

// HasFuncCreateQueryEXT returns true if the method "EXT_disjoint_timer_query.createQueryEXT" exists.
func (this EXT_disjoint_timer_query) HasFuncCreateQueryEXT() bool {
	return js.True == bindings.HasFuncEXT_disjoint_timer_queryCreateQueryEXT(
		this.ref,
	)
}

// FuncCreateQueryEXT returns the method "EXT_disjoint_timer_query.createQueryEXT".
func (this EXT_disjoint_timer_query) FuncCreateQueryEXT() (fn js.Func[func() WebGLTimerQueryEXT]) {
	bindings.FuncEXT_disjoint_timer_queryCreateQueryEXT(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateQueryEXT calls the method "EXT_disjoint_timer_query.createQueryEXT".
func (this EXT_disjoint_timer_query) CreateQueryEXT() (ret WebGLTimerQueryEXT) {
	bindings.CallEXT_disjoint_timer_queryCreateQueryEXT(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateQueryEXT calls the method "EXT_disjoint_timer_query.createQueryEXT"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EXT_disjoint_timer_query) TryCreateQueryEXT() (ret WebGLTimerQueryEXT, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEXT_disjoint_timer_queryCreateQueryEXT(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeleteQueryEXT returns true if the method "EXT_disjoint_timer_query.deleteQueryEXT" exists.
func (this EXT_disjoint_timer_query) HasFuncDeleteQueryEXT() bool {
	return js.True == bindings.HasFuncEXT_disjoint_timer_queryDeleteQueryEXT(
		this.ref,
	)
}

// FuncDeleteQueryEXT returns the method "EXT_disjoint_timer_query.deleteQueryEXT".
func (this EXT_disjoint_timer_query) FuncDeleteQueryEXT() (fn js.Func[func(query WebGLTimerQueryEXT)]) {
	bindings.FuncEXT_disjoint_timer_queryDeleteQueryEXT(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteQueryEXT calls the method "EXT_disjoint_timer_query.deleteQueryEXT".
func (this EXT_disjoint_timer_query) DeleteQueryEXT(query WebGLTimerQueryEXT) (ret js.Void) {
	bindings.CallEXT_disjoint_timer_queryDeleteQueryEXT(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryDeleteQueryEXT calls the method "EXT_disjoint_timer_query.deleteQueryEXT"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EXT_disjoint_timer_query) TryDeleteQueryEXT(query WebGLTimerQueryEXT) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEXT_disjoint_timer_queryDeleteQueryEXT(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncIsQueryEXT returns true if the method "EXT_disjoint_timer_query.isQueryEXT" exists.
func (this EXT_disjoint_timer_query) HasFuncIsQueryEXT() bool {
	return js.True == bindings.HasFuncEXT_disjoint_timer_queryIsQueryEXT(
		this.ref,
	)
}

// FuncIsQueryEXT returns the method "EXT_disjoint_timer_query.isQueryEXT".
func (this EXT_disjoint_timer_query) FuncIsQueryEXT() (fn js.Func[func(query WebGLTimerQueryEXT) bool]) {
	bindings.FuncEXT_disjoint_timer_queryIsQueryEXT(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsQueryEXT calls the method "EXT_disjoint_timer_query.isQueryEXT".
func (this EXT_disjoint_timer_query) IsQueryEXT(query WebGLTimerQueryEXT) (ret bool) {
	bindings.CallEXT_disjoint_timer_queryIsQueryEXT(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryIsQueryEXT calls the method "EXT_disjoint_timer_query.isQueryEXT"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EXT_disjoint_timer_query) TryIsQueryEXT(query WebGLTimerQueryEXT) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEXT_disjoint_timer_queryIsQueryEXT(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncBeginQueryEXT returns true if the method "EXT_disjoint_timer_query.beginQueryEXT" exists.
func (this EXT_disjoint_timer_query) HasFuncBeginQueryEXT() bool {
	return js.True == bindings.HasFuncEXT_disjoint_timer_queryBeginQueryEXT(
		this.ref,
	)
}

// FuncBeginQueryEXT returns the method "EXT_disjoint_timer_query.beginQueryEXT".
func (this EXT_disjoint_timer_query) FuncBeginQueryEXT() (fn js.Func[func(target GLenum, query WebGLTimerQueryEXT)]) {
	bindings.FuncEXT_disjoint_timer_queryBeginQueryEXT(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BeginQueryEXT calls the method "EXT_disjoint_timer_query.beginQueryEXT".
func (this EXT_disjoint_timer_query) BeginQueryEXT(target GLenum, query WebGLTimerQueryEXT) (ret js.Void) {
	bindings.CallEXT_disjoint_timer_queryBeginQueryEXT(
		this.ref, js.Pointer(&ret),
		uint32(target),
		query.Ref(),
	)

	return
}

// TryBeginQueryEXT calls the method "EXT_disjoint_timer_query.beginQueryEXT"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EXT_disjoint_timer_query) TryBeginQueryEXT(target GLenum, query WebGLTimerQueryEXT) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEXT_disjoint_timer_queryBeginQueryEXT(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		query.Ref(),
	)

	return
}

// HasFuncEndQueryEXT returns true if the method "EXT_disjoint_timer_query.endQueryEXT" exists.
func (this EXT_disjoint_timer_query) HasFuncEndQueryEXT() bool {
	return js.True == bindings.HasFuncEXT_disjoint_timer_queryEndQueryEXT(
		this.ref,
	)
}

// FuncEndQueryEXT returns the method "EXT_disjoint_timer_query.endQueryEXT".
func (this EXT_disjoint_timer_query) FuncEndQueryEXT() (fn js.Func[func(target GLenum)]) {
	bindings.FuncEXT_disjoint_timer_queryEndQueryEXT(
		this.ref, js.Pointer(&fn),
	)
	return
}

// EndQueryEXT calls the method "EXT_disjoint_timer_query.endQueryEXT".
func (this EXT_disjoint_timer_query) EndQueryEXT(target GLenum) (ret js.Void) {
	bindings.CallEXT_disjoint_timer_queryEndQueryEXT(
		this.ref, js.Pointer(&ret),
		uint32(target),
	)

	return
}

// TryEndQueryEXT calls the method "EXT_disjoint_timer_query.endQueryEXT"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EXT_disjoint_timer_query) TryEndQueryEXT(target GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEXT_disjoint_timer_queryEndQueryEXT(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
	)

	return
}

// HasFuncQueryCounterEXT returns true if the method "EXT_disjoint_timer_query.queryCounterEXT" exists.
func (this EXT_disjoint_timer_query) HasFuncQueryCounterEXT() bool {
	return js.True == bindings.HasFuncEXT_disjoint_timer_queryQueryCounterEXT(
		this.ref,
	)
}

// FuncQueryCounterEXT returns the method "EXT_disjoint_timer_query.queryCounterEXT".
func (this EXT_disjoint_timer_query) FuncQueryCounterEXT() (fn js.Func[func(query WebGLTimerQueryEXT, target GLenum)]) {
	bindings.FuncEXT_disjoint_timer_queryQueryCounterEXT(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QueryCounterEXT calls the method "EXT_disjoint_timer_query.queryCounterEXT".
func (this EXT_disjoint_timer_query) QueryCounterEXT(query WebGLTimerQueryEXT, target GLenum) (ret js.Void) {
	bindings.CallEXT_disjoint_timer_queryQueryCounterEXT(
		this.ref, js.Pointer(&ret),
		query.Ref(),
		uint32(target),
	)

	return
}

// TryQueryCounterEXT calls the method "EXT_disjoint_timer_query.queryCounterEXT"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EXT_disjoint_timer_query) TryQueryCounterEXT(query WebGLTimerQueryEXT, target GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEXT_disjoint_timer_queryQueryCounterEXT(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(target),
	)

	return
}

// HasFuncGetQueryEXT returns true if the method "EXT_disjoint_timer_query.getQueryEXT" exists.
func (this EXT_disjoint_timer_query) HasFuncGetQueryEXT() bool {
	return js.True == bindings.HasFuncEXT_disjoint_timer_queryGetQueryEXT(
		this.ref,
	)
}

// FuncGetQueryEXT returns the method "EXT_disjoint_timer_query.getQueryEXT".
func (this EXT_disjoint_timer_query) FuncGetQueryEXT() (fn js.Func[func(target GLenum, pname GLenum) js.Any]) {
	bindings.FuncEXT_disjoint_timer_queryGetQueryEXT(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetQueryEXT calls the method "EXT_disjoint_timer_query.getQueryEXT".
func (this EXT_disjoint_timer_query) GetQueryEXT(target GLenum, pname GLenum) (ret js.Any) {
	bindings.CallEXT_disjoint_timer_queryGetQueryEXT(
		this.ref, js.Pointer(&ret),
		uint32(target),
		uint32(pname),
	)

	return
}

// TryGetQueryEXT calls the method "EXT_disjoint_timer_query.getQueryEXT"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EXT_disjoint_timer_query) TryGetQueryEXT(target GLenum, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEXT_disjoint_timer_queryGetQueryEXT(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(pname),
	)

	return
}

// HasFuncGetQueryObjectEXT returns true if the method "EXT_disjoint_timer_query.getQueryObjectEXT" exists.
func (this EXT_disjoint_timer_query) HasFuncGetQueryObjectEXT() bool {
	return js.True == bindings.HasFuncEXT_disjoint_timer_queryGetQueryObjectEXT(
		this.ref,
	)
}

// FuncGetQueryObjectEXT returns the method "EXT_disjoint_timer_query.getQueryObjectEXT".
func (this EXT_disjoint_timer_query) FuncGetQueryObjectEXT() (fn js.Func[func(query WebGLTimerQueryEXT, pname GLenum) js.Any]) {
	bindings.FuncEXT_disjoint_timer_queryGetQueryObjectEXT(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetQueryObjectEXT calls the method "EXT_disjoint_timer_query.getQueryObjectEXT".
func (this EXT_disjoint_timer_query) GetQueryObjectEXT(query WebGLTimerQueryEXT, pname GLenum) (ret js.Any) {
	bindings.CallEXT_disjoint_timer_queryGetQueryObjectEXT(
		this.ref, js.Pointer(&ret),
		query.Ref(),
		uint32(pname),
	)

	return
}

// TryGetQueryObjectEXT calls the method "EXT_disjoint_timer_query.getQueryObjectEXT"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EXT_disjoint_timer_query) TryGetQueryObjectEXT(query WebGLTimerQueryEXT, pname GLenum) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEXT_disjoint_timer_queryGetQueryObjectEXT(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(pname),
	)

	return
}

const (
	EXT_disjoint_timer_query_webgl2_QUERY_COUNTER_BITS_EXT GLenum = 0x8864
	EXT_disjoint_timer_query_webgl2_TIME_ELAPSED_EXT       GLenum = 0x88BF
	EXT_disjoint_timer_query_webgl2_TIMESTAMP_EXT          GLenum = 0x8E28
	EXT_disjoint_timer_query_webgl2_GPU_DISJOINT_EXT       GLenum = 0x8FBB
)

type EXT_disjoint_timer_query_webgl2 struct {
	ref js.Ref
}

func (this EXT_disjoint_timer_query_webgl2) Once() EXT_disjoint_timer_query_webgl2 {
	this.ref.Once()
	return this
}

func (this EXT_disjoint_timer_query_webgl2) Ref() js.Ref {
	return this.ref
}

func (this EXT_disjoint_timer_query_webgl2) FromRef(ref js.Ref) EXT_disjoint_timer_query_webgl2 {
	this.ref = ref
	return this
}

func (this EXT_disjoint_timer_query_webgl2) Free() {
	this.ref.Free()
}

// HasFuncQueryCounterEXT returns true if the method "EXT_disjoint_timer_query_webgl2.queryCounterEXT" exists.
func (this EXT_disjoint_timer_query_webgl2) HasFuncQueryCounterEXT() bool {
	return js.True == bindings.HasFuncEXT_disjoint_timer_query_webgl2QueryCounterEXT(
		this.ref,
	)
}

// FuncQueryCounterEXT returns the method "EXT_disjoint_timer_query_webgl2.queryCounterEXT".
func (this EXT_disjoint_timer_query_webgl2) FuncQueryCounterEXT() (fn js.Func[func(query WebGLQuery, target GLenum)]) {
	bindings.FuncEXT_disjoint_timer_query_webgl2QueryCounterEXT(
		this.ref, js.Pointer(&fn),
	)
	return
}

// QueryCounterEXT calls the method "EXT_disjoint_timer_query_webgl2.queryCounterEXT".
func (this EXT_disjoint_timer_query_webgl2) QueryCounterEXT(query WebGLQuery, target GLenum) (ret js.Void) {
	bindings.CallEXT_disjoint_timer_query_webgl2QueryCounterEXT(
		this.ref, js.Pointer(&ret),
		query.Ref(),
		uint32(target),
	)

	return
}

// TryQueryCounterEXT calls the method "EXT_disjoint_timer_query_webgl2.queryCounterEXT"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EXT_disjoint_timer_query_webgl2) TryQueryCounterEXT(query WebGLQuery, target GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEXT_disjoint_timer_query_webgl2QueryCounterEXT(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(target),
	)

	return
}

type EXT_float_blend struct {
	ref js.Ref
}

func (this EXT_float_blend) Once() EXT_float_blend {
	this.ref.Once()
	return this
}

func (this EXT_float_blend) Ref() js.Ref {
	return this.ref
}

func (this EXT_float_blend) FromRef(ref js.Ref) EXT_float_blend {
	this.ref = ref
	return this
}

func (this EXT_float_blend) Free() {
	this.ref.Free()
}

type EXT_frag_depth struct {
	ref js.Ref
}

func (this EXT_frag_depth) Once() EXT_frag_depth {
	this.ref.Once()
	return this
}

func (this EXT_frag_depth) Ref() js.Ref {
	return this.ref
}

func (this EXT_frag_depth) FromRef(ref js.Ref) EXT_frag_depth {
	this.ref = ref
	return this
}

func (this EXT_frag_depth) Free() {
	this.ref.Free()
}

const (
	EXT_sRGB_SRGB_EXT                                  GLenum = 0x8C40
	EXT_sRGB_SRGB_ALPHA_EXT                            GLenum = 0x8C42
	EXT_sRGB_SRGB8_ALPHA8_EXT                          GLenum = 0x8C43
	EXT_sRGB_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING_EXT GLenum = 0x8210
)

type EXT_sRGB struct {
	ref js.Ref
}

func (this EXT_sRGB) Once() EXT_sRGB {
	this.ref.Once()
	return this
}

func (this EXT_sRGB) Ref() js.Ref {
	return this.ref
}

func (this EXT_sRGB) FromRef(ref js.Ref) EXT_sRGB {
	this.ref = ref
	return this
}

func (this EXT_sRGB) Free() {
	this.ref.Free()
}

type EXT_shader_texture_lod struct {
	ref js.Ref
}

func (this EXT_shader_texture_lod) Once() EXT_shader_texture_lod {
	this.ref.Once()
	return this
}

func (this EXT_shader_texture_lod) Ref() js.Ref {
	return this.ref
}

func (this EXT_shader_texture_lod) FromRef(ref js.Ref) EXT_shader_texture_lod {
	this.ref = ref
	return this
}

func (this EXT_shader_texture_lod) Free() {
	this.ref.Free()
}

const (
	EXT_texture_compression_bptc_COMPRESSED_RGBA_BPTC_UNORM_EXT         GLenum = 0x8E8C
	EXT_texture_compression_bptc_COMPRESSED_SRGB_ALPHA_BPTC_UNORM_EXT   GLenum = 0x8E8D
	EXT_texture_compression_bptc_COMPRESSED_RGB_BPTC_SIGNED_FLOAT_EXT   GLenum = 0x8E8E
	EXT_texture_compression_bptc_COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_EXT GLenum = 0x8E8F
)

type EXT_texture_compression_bptc struct {
	ref js.Ref
}

func (this EXT_texture_compression_bptc) Once() EXT_texture_compression_bptc {
	this.ref.Once()
	return this
}

func (this EXT_texture_compression_bptc) Ref() js.Ref {
	return this.ref
}

func (this EXT_texture_compression_bptc) FromRef(ref js.Ref) EXT_texture_compression_bptc {
	this.ref = ref
	return this
}

func (this EXT_texture_compression_bptc) Free() {
	this.ref.Free()
}

const (
	EXT_texture_compression_rgtc_COMPRESSED_RED_RGTC1_EXT              GLenum = 0x8DBB
	EXT_texture_compression_rgtc_COMPRESSED_SIGNED_RED_RGTC1_EXT       GLenum = 0x8DBC
	EXT_texture_compression_rgtc_COMPRESSED_RED_GREEN_RGTC2_EXT        GLenum = 0x8DBD
	EXT_texture_compression_rgtc_COMPRESSED_SIGNED_RED_GREEN_RGTC2_EXT GLenum = 0x8DBE
)

type EXT_texture_compression_rgtc struct {
	ref js.Ref
}

func (this EXT_texture_compression_rgtc) Once() EXT_texture_compression_rgtc {
	this.ref.Once()
	return this
}

func (this EXT_texture_compression_rgtc) Ref() js.Ref {
	return this.ref
}

func (this EXT_texture_compression_rgtc) FromRef(ref js.Ref) EXT_texture_compression_rgtc {
	this.ref = ref
	return this
}

func (this EXT_texture_compression_rgtc) Free() {
	this.ref.Free()
}

const (
	EXT_texture_filter_anisotropic_TEXTURE_MAX_ANISOTROPY_EXT     GLenum = 0x84FE
	EXT_texture_filter_anisotropic_MAX_TEXTURE_MAX_ANISOTROPY_EXT GLenum = 0x84FF
)

type EXT_texture_filter_anisotropic struct {
	ref js.Ref
}

func (this EXT_texture_filter_anisotropic) Once() EXT_texture_filter_anisotropic {
	this.ref.Once()
	return this
}

func (this EXT_texture_filter_anisotropic) Ref() js.Ref {
	return this.ref
}

func (this EXT_texture_filter_anisotropic) FromRef(ref js.Ref) EXT_texture_filter_anisotropic {
	this.ref = ref
	return this
}

func (this EXT_texture_filter_anisotropic) Free() {
	this.ref.Free()
}

const (
	EXT_texture_norm16_R16_EXT          GLenum = 0x822A
	EXT_texture_norm16_RG16_EXT         GLenum = 0x822C
	EXT_texture_norm16_RGB16_EXT        GLenum = 0x8054
	EXT_texture_norm16_RGBA16_EXT       GLenum = 0x805B
	EXT_texture_norm16_R16_SNORM_EXT    GLenum = 0x8F98
	EXT_texture_norm16_RG16_SNORM_EXT   GLenum = 0x8F99
	EXT_texture_norm16_RGB16_SNORM_EXT  GLenum = 0x8F9A
	EXT_texture_norm16_RGBA16_SNORM_EXT GLenum = 0x8F9B
)

type EXT_texture_norm16 struct {
	ref js.Ref
}

func (this EXT_texture_norm16) Once() EXT_texture_norm16 {
	this.ref.Once()
	return this
}

func (this EXT_texture_norm16) Ref() js.Ref {
	return this.ref
}

func (this EXT_texture_norm16) FromRef(ref js.Ref) EXT_texture_norm16 {
	this.ref = ref
	return this
}

func (this EXT_texture_norm16) Free() {
	this.ref.Free()
}

type NamedCurve = js.String

type EcKeyAlgorithm struct {
	// NamedCurve is "EcKeyAlgorithm.namedCurve"
	//
	// Required
	NamedCurve NamedCurve
	// Name is "EcKeyAlgorithm.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EcKeyAlgorithm with all fields set.
func (p EcKeyAlgorithm) FromRef(ref js.Ref) EcKeyAlgorithm {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EcKeyAlgorithm in the application heap.
func (p EcKeyAlgorithm) New() js.Ref {
	return bindings.EcKeyAlgorithmJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EcKeyAlgorithm) UpdateFrom(ref js.Ref) {
	bindings.EcKeyAlgorithmJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EcKeyAlgorithm) Update(ref js.Ref) {
	bindings.EcKeyAlgorithmJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EcKeyAlgorithm) FreeMembers(recursive bool) {
	js.Free(
		p.NamedCurve.Ref(),
		p.Name.Ref(),
	)
	p.NamedCurve = p.NamedCurve.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

type EcKeyGenParams struct {
	// NamedCurve is "EcKeyGenParams.namedCurve"
	//
	// Required
	NamedCurve NamedCurve
	// Name is "EcKeyGenParams.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EcKeyGenParams with all fields set.
func (p EcKeyGenParams) FromRef(ref js.Ref) EcKeyGenParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EcKeyGenParams in the application heap.
func (p EcKeyGenParams) New() js.Ref {
	return bindings.EcKeyGenParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EcKeyGenParams) UpdateFrom(ref js.Ref) {
	bindings.EcKeyGenParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EcKeyGenParams) Update(ref js.Ref) {
	bindings.EcKeyGenParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EcKeyGenParams) FreeMembers(recursive bool) {
	js.Free(
		p.NamedCurve.Ref(),
		p.Name.Ref(),
	)
	p.NamedCurve = p.NamedCurve.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

type EcKeyImportParams struct {
	// NamedCurve is "EcKeyImportParams.namedCurve"
	//
	// Required
	NamedCurve NamedCurve
	// Name is "EcKeyImportParams.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EcKeyImportParams with all fields set.
func (p EcKeyImportParams) FromRef(ref js.Ref) EcKeyImportParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EcKeyImportParams in the application heap.
func (p EcKeyImportParams) New() js.Ref {
	return bindings.EcKeyImportParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EcKeyImportParams) UpdateFrom(ref js.Ref) {
	bindings.EcKeyImportParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EcKeyImportParams) Update(ref js.Ref) {
	bindings.EcKeyImportParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EcKeyImportParams) FreeMembers(recursive bool) {
	js.Free(
		p.NamedCurve.Ref(),
		p.Name.Ref(),
	)
	p.NamedCurve = p.NamedCurve.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

type EcdhKeyDeriveParams struct {
	// Public is "EcdhKeyDeriveParams.public"
	//
	// Required
	Public CryptoKey
	// Name is "EcdhKeyDeriveParams.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EcdhKeyDeriveParams with all fields set.
func (p EcdhKeyDeriveParams) FromRef(ref js.Ref) EcdhKeyDeriveParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EcdhKeyDeriveParams in the application heap.
func (p EcdhKeyDeriveParams) New() js.Ref {
	return bindings.EcdhKeyDeriveParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EcdhKeyDeriveParams) UpdateFrom(ref js.Ref) {
	bindings.EcdhKeyDeriveParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EcdhKeyDeriveParams) Update(ref js.Ref) {
	bindings.EcdhKeyDeriveParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EcdhKeyDeriveParams) FreeMembers(recursive bool) {
	js.Free(
		p.Public.Ref(),
		p.Name.Ref(),
	)
	p.Public = p.Public.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

type HashAlgorithmIdentifier = AlgorithmIdentifier

type EcdsaParams struct {
	// Hash is "EcdsaParams.hash"
	//
	// Required
	Hash HashAlgorithmIdentifier
	// Name is "EcdsaParams.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EcdsaParams with all fields set.
func (p EcdsaParams) FromRef(ref js.Ref) EcdsaParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EcdsaParams in the application heap.
func (p EcdsaParams) New() js.Ref {
	return bindings.EcdsaParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EcdsaParams) UpdateFrom(ref js.Ref) {
	bindings.EcdsaParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EcdsaParams) Update(ref js.Ref) {
	bindings.EcdsaParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EcdsaParams) FreeMembers(recursive bool) {
	js.Free(
		p.Hash.Ref(),
		p.Name.Ref(),
	)
	p.Hash = p.Hash.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

type Ed448Params struct {
	// Context is "Ed448Params.context"
	//
	// Optional
	Context BufferSource
	// Name is "Ed448Params.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Ed448Params with all fields set.
func (p Ed448Params) FromRef(ref js.Ref) Ed448Params {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Ed448Params in the application heap.
func (p Ed448Params) New() js.Ref {
	return bindings.Ed448ParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Ed448Params) UpdateFrom(ref js.Ref) {
	bindings.Ed448ParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Ed448Params) Update(ref js.Ref) {
	bindings.Ed448ParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Ed448Params) FreeMembers(recursive bool) {
	js.Free(
		p.Context.Ref(),
		p.Name.Ref(),
	)
	p.Context = p.Context.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

type EffectCallbackFunc func(this js.Ref, progress float64, currentTarget OneOf_Element_CSSPseudoElement, animation Animation) js.Ref

func (fn EffectCallbackFunc) Register() js.Func[func(progress float64, currentTarget OneOf_Element_CSSPseudoElement, animation Animation)] {
	return js.RegisterCallback[func(progress float64, currentTarget OneOf_Element_CSSPseudoElement, animation Animation)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EffectCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
		OneOf_Element_CSSPseudoElement{}.FromRef(args[1+1]),
		Animation{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EffectCallback[T any] struct {
	Fn  func(arg T, this js.Ref, progress float64, currentTarget OneOf_Element_CSSPseudoElement, animation Animation) js.Ref
	Arg T
}

func (cb *EffectCallback[T]) Register() js.Func[func(progress float64, currentTarget OneOf_Element_CSSPseudoElement, animation Animation)] {
	return js.RegisterCallback[func(progress float64, currentTarget OneOf_Element_CSSPseudoElement, animation Animation)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EffectCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
		OneOf_Element_CSSPseudoElement{}.FromRef(args[1+1]),
		Animation{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EncodedVideoChunkType uint32

const (
	_ EncodedVideoChunkType = iota

	EncodedVideoChunkType_KEY
	EncodedVideoChunkType_DELTA
)

func (EncodedVideoChunkType) FromRef(str js.Ref) EncodedVideoChunkType {
	return EncodedVideoChunkType(bindings.ConstOfEncodedVideoChunkType(str))
}

func (x EncodedVideoChunkType) String() (string, bool) {
	switch x {
	case EncodedVideoChunkType_KEY:
		return "key", true
	case EncodedVideoChunkType_DELTA:
		return "delta", true
	default:
		return "", false
	}
}

type EncodedVideoChunkInit struct {
	// Type is "EncodedVideoChunkInit.type"
	//
	// Required
	Type EncodedVideoChunkType
	// Timestamp is "EncodedVideoChunkInit.timestamp"
	//
	// Required
	Timestamp int64
	// Duration is "EncodedVideoChunkInit.duration"
	//
	// Optional
	//
	// NOTE: FFI_USE_Duration MUST be set to true to make this field effective.
	Duration uint64
	// Data is "EncodedVideoChunkInit.data"
	//
	// Required
	Data AllowSharedBufferSource

	FFI_USE_Duration bool // for Duration.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EncodedVideoChunkInit with all fields set.
func (p EncodedVideoChunkInit) FromRef(ref js.Ref) EncodedVideoChunkInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EncodedVideoChunkInit in the application heap.
func (p EncodedVideoChunkInit) New() js.Ref {
	return bindings.EncodedVideoChunkInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EncodedVideoChunkInit) UpdateFrom(ref js.Ref) {
	bindings.EncodedVideoChunkInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EncodedVideoChunkInit) Update(ref js.Ref) {
	bindings.EncodedVideoChunkInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EncodedVideoChunkInit) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
}

func NewEncodedVideoChunk(init EncodedVideoChunkInit) (ret EncodedVideoChunk) {
	ret.ref = bindings.NewEncodedVideoChunkByEncodedVideoChunk(
		js.Pointer(&init))
	return
}

type EncodedVideoChunk struct {
	ref js.Ref
}

func (this EncodedVideoChunk) Once() EncodedVideoChunk {
	this.ref.Once()
	return this
}

func (this EncodedVideoChunk) Ref() js.Ref {
	return this.ref
}

func (this EncodedVideoChunk) FromRef(ref js.Ref) EncodedVideoChunk {
	this.ref = ref
	return this
}

func (this EncodedVideoChunk) Free() {
	this.ref.Free()
}

// Type returns the value of property "EncodedVideoChunk.type".
//
// It returns ok=false if there is no such property.
func (this EncodedVideoChunk) Type() (ret EncodedVideoChunkType, ok bool) {
	ok = js.True == bindings.GetEncodedVideoChunkType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Timestamp returns the value of property "EncodedVideoChunk.timestamp".
//
// It returns ok=false if there is no such property.
func (this EncodedVideoChunk) Timestamp() (ret int64, ok bool) {
	ok = js.True == bindings.GetEncodedVideoChunkTimestamp(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Duration returns the value of property "EncodedVideoChunk.duration".
//
// It returns ok=false if there is no such property.
func (this EncodedVideoChunk) Duration() (ret uint64, ok bool) {
	ok = js.True == bindings.GetEncodedVideoChunkDuration(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ByteLength returns the value of property "EncodedVideoChunk.byteLength".
//
// It returns ok=false if there is no such property.
func (this EncodedVideoChunk) ByteLength() (ret uint32, ok bool) {
	ok = js.True == bindings.GetEncodedVideoChunkByteLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncCopyTo returns true if the method "EncodedVideoChunk.copyTo" exists.
func (this EncodedVideoChunk) HasFuncCopyTo() bool {
	return js.True == bindings.HasFuncEncodedVideoChunkCopyTo(
		this.ref,
	)
}

// FuncCopyTo returns the method "EncodedVideoChunk.copyTo".
func (this EncodedVideoChunk) FuncCopyTo() (fn js.Func[func(destination AllowSharedBufferSource)]) {
	bindings.FuncEncodedVideoChunkCopyTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyTo calls the method "EncodedVideoChunk.copyTo".
func (this EncodedVideoChunk) CopyTo(destination AllowSharedBufferSource) (ret js.Void) {
	bindings.CallEncodedVideoChunkCopyTo(
		this.ref, js.Pointer(&ret),
		destination.Ref(),
	)

	return
}

// TryCopyTo calls the method "EncodedVideoChunk.copyTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EncodedVideoChunk) TryCopyTo(destination AllowSharedBufferSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEncodedVideoChunkCopyTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destination.Ref(),
	)

	return
}

type HardwareAcceleration uint32

const (
	_ HardwareAcceleration = iota

	HardwareAcceleration_NO_PREFERENCE
	HardwareAcceleration_PREFER_HARDWARE
	HardwareAcceleration_PREFER_SOFTWARE
)

func (HardwareAcceleration) FromRef(str js.Ref) HardwareAcceleration {
	return HardwareAcceleration(bindings.ConstOfHardwareAcceleration(str))
}

func (x HardwareAcceleration) String() (string, bool) {
	switch x {
	case HardwareAcceleration_NO_PREFERENCE:
		return "no-preference", true
	case HardwareAcceleration_PREFER_HARDWARE:
		return "prefer-hardware", true
	case HardwareAcceleration_PREFER_SOFTWARE:
		return "prefer-software", true
	default:
		return "", false
	}
}

type VideoDecoderConfig struct {
	// Codec is "VideoDecoderConfig.codec"
	//
	// Required
	Codec js.String
	// Description is "VideoDecoderConfig.description"
	//
	// Optional
	Description AllowSharedBufferSource
	// CodedWidth is "VideoDecoderConfig.codedWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_CodedWidth MUST be set to true to make this field effective.
	CodedWidth uint32
	// CodedHeight is "VideoDecoderConfig.codedHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_CodedHeight MUST be set to true to make this field effective.
	CodedHeight uint32
	// DisplayAspectWidth is "VideoDecoderConfig.displayAspectWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_DisplayAspectWidth MUST be set to true to make this field effective.
	DisplayAspectWidth uint32
	// DisplayAspectHeight is "VideoDecoderConfig.displayAspectHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_DisplayAspectHeight MUST be set to true to make this field effective.
	DisplayAspectHeight uint32
	// ColorSpace is "VideoDecoderConfig.colorSpace"
	//
	// Optional
	//
	// NOTE: ColorSpace.FFI_USE MUST be set to true to get ColorSpace used.
	ColorSpace VideoColorSpaceInit
	// HardwareAcceleration is "VideoDecoderConfig.hardwareAcceleration"
	//
	// Optional, defaults to "no-preference".
	HardwareAcceleration HardwareAcceleration
	// OptimizeForLatency is "VideoDecoderConfig.optimizeForLatency"
	//
	// Optional
	//
	// NOTE: FFI_USE_OptimizeForLatency MUST be set to true to make this field effective.
	OptimizeForLatency bool

	FFI_USE_CodedWidth          bool // for CodedWidth.
	FFI_USE_CodedHeight         bool // for CodedHeight.
	FFI_USE_DisplayAspectWidth  bool // for DisplayAspectWidth.
	FFI_USE_DisplayAspectHeight bool // for DisplayAspectHeight.
	FFI_USE_OptimizeForLatency  bool // for OptimizeForLatency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoDecoderConfig with all fields set.
func (p VideoDecoderConfig) FromRef(ref js.Ref) VideoDecoderConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VideoDecoderConfig in the application heap.
func (p VideoDecoderConfig) New() js.Ref {
	return bindings.VideoDecoderConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *VideoDecoderConfig) UpdateFrom(ref js.Ref) {
	bindings.VideoDecoderConfigJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VideoDecoderConfig) Update(ref js.Ref) {
	bindings.VideoDecoderConfigJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VideoDecoderConfig) FreeMembers(recursive bool) {
	js.Free(
		p.Codec.Ref(),
		p.Description.Ref(),
	)
	p.Codec = p.Codec.FromRef(js.Undefined)
	p.Description = p.Description.FromRef(js.Undefined)
	if recursive {
		p.ColorSpace.FreeMembers(true)
	}
}

type SvcOutputMetadata struct {
	// TemporalLayerId is "SvcOutputMetadata.temporalLayerId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TemporalLayerId MUST be set to true to make this field effective.
	TemporalLayerId uint32

	FFI_USE_TemporalLayerId bool // for TemporalLayerId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SvcOutputMetadata with all fields set.
func (p SvcOutputMetadata) FromRef(ref js.Ref) SvcOutputMetadata {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SvcOutputMetadata in the application heap.
func (p SvcOutputMetadata) New() js.Ref {
	return bindings.SvcOutputMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SvcOutputMetadata) UpdateFrom(ref js.Ref) {
	bindings.SvcOutputMetadataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SvcOutputMetadata) Update(ref js.Ref) {
	bindings.SvcOutputMetadataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SvcOutputMetadata) FreeMembers(recursive bool) {
}

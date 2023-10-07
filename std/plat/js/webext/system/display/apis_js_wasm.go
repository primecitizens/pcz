// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package display

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/system/display/bindings"
)

type ActiveState uint32

const (
	_ ActiveState = iota

	ActiveState_ACTIVE
	ActiveState_INACTIVE
)

func (ActiveState) FromRef(str js.Ref) ActiveState {
	return ActiveState(bindings.ConstOfActiveState(str))
}

func (x ActiveState) String() (string, bool) {
	switch x {
	case ActiveState_ACTIVE:
		return "active", true
	case ActiveState_INACTIVE:
		return "inactive", true
	default:
		return "", false
	}
}

type Bounds struct {
	// Left is "Bounds.left"
	//
	// Optional
	//
	// NOTE: FFI_USE_Left MUST be set to true to make this field effective.
	Left int32
	// Top is "Bounds.top"
	//
	// Optional
	//
	// NOTE: FFI_USE_Top MUST be set to true to make this field effective.
	Top int32
	// Width is "Bounds.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int32
	// Height is "Bounds.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int32

	FFI_USE_Left   bool // for Left.
	FFI_USE_Top    bool // for Top.
	FFI_USE_Width  bool // for Width.
	FFI_USE_Height bool // for Height.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Bounds with all fields set.
func (p Bounds) FromRef(ref js.Ref) Bounds {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Bounds in the application heap.
func (p Bounds) New() js.Ref {
	return bindings.BoundsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Bounds) UpdateFrom(ref js.Ref) {
	bindings.BoundsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Bounds) Update(ref js.Ref) {
	bindings.BoundsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Bounds) FreeMembers(recursive bool) {
}

type DisplayInfoCallbackFunc func(this js.Ref, displayInfo js.Array[DisplayUnitInfo]) js.Ref

func (fn DisplayInfoCallbackFunc) Register() js.Func[func(displayInfo js.Array[DisplayUnitInfo])] {
	return js.RegisterCallback[func(displayInfo js.Array[DisplayUnitInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DisplayInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[DisplayUnitInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DisplayInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, displayInfo js.Array[DisplayUnitInfo]) js.Ref
	Arg T
}

func (cb *DisplayInfoCallback[T]) Register() js.Func[func(displayInfo js.Array[DisplayUnitInfo])] {
	return js.RegisterCallback[func(displayInfo js.Array[DisplayUnitInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DisplayInfoCallback[T]) DispatchCallback(
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

		js.Array[DisplayUnitInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type Edid struct {
	// ManufacturerId is "Edid.manufacturerId"
	//
	// Optional
	ManufacturerId js.String
	// ProductId is "Edid.productId"
	//
	// Optional
	ProductId js.String
	// YearOfManufacture is "Edid.yearOfManufacture"
	//
	// Optional
	//
	// NOTE: FFI_USE_YearOfManufacture MUST be set to true to make this field effective.
	YearOfManufacture int32

	FFI_USE_YearOfManufacture bool // for YearOfManufacture.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Edid with all fields set.
func (p Edid) FromRef(ref js.Ref) Edid {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Edid in the application heap.
func (p Edid) New() js.Ref {
	return bindings.EdidJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Edid) UpdateFrom(ref js.Ref) {
	bindings.EdidJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Edid) Update(ref js.Ref) {
	bindings.EdidJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Edid) FreeMembers(recursive bool) {
	js.Free(
		p.ManufacturerId.Ref(),
		p.ProductId.Ref(),
	)
	p.ManufacturerId = p.ManufacturerId.FromRef(js.Undefined)
	p.ProductId = p.ProductId.FromRef(js.Undefined)
}

type Insets struct {
	// Left is "Insets.left"
	//
	// Optional
	//
	// NOTE: FFI_USE_Left MUST be set to true to make this field effective.
	Left int32
	// Top is "Insets.top"
	//
	// Optional
	//
	// NOTE: FFI_USE_Top MUST be set to true to make this field effective.
	Top int32
	// Right is "Insets.right"
	//
	// Optional
	//
	// NOTE: FFI_USE_Right MUST be set to true to make this field effective.
	Right int32
	// Bottom is "Insets.bottom"
	//
	// Optional
	//
	// NOTE: FFI_USE_Bottom MUST be set to true to make this field effective.
	Bottom int32

	FFI_USE_Left   bool // for Left.
	FFI_USE_Top    bool // for Top.
	FFI_USE_Right  bool // for Right.
	FFI_USE_Bottom bool // for Bottom.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Insets with all fields set.
func (p Insets) FromRef(ref js.Ref) Insets {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Insets in the application heap.
func (p Insets) New() js.Ref {
	return bindings.InsetsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Insets) UpdateFrom(ref js.Ref) {
	bindings.InsetsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Insets) Update(ref js.Ref) {
	bindings.InsetsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Insets) FreeMembers(recursive bool) {
}

type DisplayMode struct {
	// Width is "DisplayMode.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int32
	// Height is "DisplayMode.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int32
	// WidthInNativePixels is "DisplayMode.widthInNativePixels"
	//
	// Optional
	//
	// NOTE: FFI_USE_WidthInNativePixels MUST be set to true to make this field effective.
	WidthInNativePixels int32
	// HeightInNativePixels is "DisplayMode.heightInNativePixels"
	//
	// Optional
	//
	// NOTE: FFI_USE_HeightInNativePixels MUST be set to true to make this field effective.
	HeightInNativePixels int32
	// UiScale is "DisplayMode.uiScale"
	//
	// Optional
	//
	// NOTE: FFI_USE_UiScale MUST be set to true to make this field effective.
	UiScale float64
	// DeviceScaleFactor is "DisplayMode.deviceScaleFactor"
	//
	// Optional
	//
	// NOTE: FFI_USE_DeviceScaleFactor MUST be set to true to make this field effective.
	DeviceScaleFactor float64
	// RefreshRate is "DisplayMode.refreshRate"
	//
	// Optional
	//
	// NOTE: FFI_USE_RefreshRate MUST be set to true to make this field effective.
	RefreshRate float64
	// IsNative is "DisplayMode.isNative"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsNative MUST be set to true to make this field effective.
	IsNative bool
	// IsSelected is "DisplayMode.isSelected"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsSelected MUST be set to true to make this field effective.
	IsSelected bool
	// IsInterlaced is "DisplayMode.isInterlaced"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsInterlaced MUST be set to true to make this field effective.
	IsInterlaced bool

	FFI_USE_Width                bool // for Width.
	FFI_USE_Height               bool // for Height.
	FFI_USE_WidthInNativePixels  bool // for WidthInNativePixels.
	FFI_USE_HeightInNativePixels bool // for HeightInNativePixels.
	FFI_USE_UiScale              bool // for UiScale.
	FFI_USE_DeviceScaleFactor    bool // for DeviceScaleFactor.
	FFI_USE_RefreshRate          bool // for RefreshRate.
	FFI_USE_IsNative             bool // for IsNative.
	FFI_USE_IsSelected           bool // for IsSelected.
	FFI_USE_IsInterlaced         bool // for IsInterlaced.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DisplayMode with all fields set.
func (p DisplayMode) FromRef(ref js.Ref) DisplayMode {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DisplayMode in the application heap.
func (p DisplayMode) New() js.Ref {
	return bindings.DisplayModeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DisplayMode) UpdateFrom(ref js.Ref) {
	bindings.DisplayModeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DisplayMode) Update(ref js.Ref) {
	bindings.DisplayModeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DisplayMode) FreeMembers(recursive bool) {
}

type DisplayUnitInfo struct {
	// Id is "DisplayUnitInfo.id"
	//
	// Optional
	Id js.String
	// Name is "DisplayUnitInfo.name"
	//
	// Optional
	Name js.String
	// Edid is "DisplayUnitInfo.edid"
	//
	// Optional
	//
	// NOTE: Edid.FFI_USE MUST be set to true to get Edid used.
	Edid Edid
	// MirroringSourceId is "DisplayUnitInfo.mirroringSourceId"
	//
	// Optional
	MirroringSourceId js.String
	// MirroringDestinationIds is "DisplayUnitInfo.mirroringDestinationIds"
	//
	// Optional
	MirroringDestinationIds js.Array[js.String]
	// IsPrimary is "DisplayUnitInfo.isPrimary"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsPrimary MUST be set to true to make this field effective.
	IsPrimary bool
	// IsInternal is "DisplayUnitInfo.isInternal"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsInternal MUST be set to true to make this field effective.
	IsInternal bool
	// IsEnabled is "DisplayUnitInfo.isEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsEnabled MUST be set to true to make this field effective.
	IsEnabled bool
	// ActiveState is "DisplayUnitInfo.activeState"
	//
	// Optional
	ActiveState ActiveState
	// IsUnified is "DisplayUnitInfo.isUnified"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsUnified MUST be set to true to make this field effective.
	IsUnified bool
	// IsAutoRotationAllowed is "DisplayUnitInfo.isAutoRotationAllowed"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsAutoRotationAllowed MUST be set to true to make this field effective.
	IsAutoRotationAllowed bool
	// DpiX is "DisplayUnitInfo.dpiX"
	//
	// Optional
	//
	// NOTE: FFI_USE_DpiX MUST be set to true to make this field effective.
	DpiX float64
	// DpiY is "DisplayUnitInfo.dpiY"
	//
	// Optional
	//
	// NOTE: FFI_USE_DpiY MUST be set to true to make this field effective.
	DpiY float64
	// Rotation is "DisplayUnitInfo.rotation"
	//
	// Optional
	//
	// NOTE: FFI_USE_Rotation MUST be set to true to make this field effective.
	Rotation int32
	// Bounds is "DisplayUnitInfo.bounds"
	//
	// Optional
	//
	// NOTE: Bounds.FFI_USE MUST be set to true to get Bounds used.
	Bounds Bounds
	// Overscan is "DisplayUnitInfo.overscan"
	//
	// Optional
	//
	// NOTE: Overscan.FFI_USE MUST be set to true to get Overscan used.
	Overscan Insets
	// WorkArea is "DisplayUnitInfo.workArea"
	//
	// Optional
	//
	// NOTE: WorkArea.FFI_USE MUST be set to true to get WorkArea used.
	WorkArea Bounds
	// Modes is "DisplayUnitInfo.modes"
	//
	// Optional
	Modes js.Array[DisplayMode]
	// HasTouchSupport is "DisplayUnitInfo.hasTouchSupport"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasTouchSupport MUST be set to true to make this field effective.
	HasTouchSupport bool
	// HasAccelerometerSupport is "DisplayUnitInfo.hasAccelerometerSupport"
	//
	// Optional
	//
	// NOTE: FFI_USE_HasAccelerometerSupport MUST be set to true to make this field effective.
	HasAccelerometerSupport bool
	// AvailableDisplayZoomFactors is "DisplayUnitInfo.availableDisplayZoomFactors"
	//
	// Optional
	AvailableDisplayZoomFactors js.Array[float64]
	// DisplayZoomFactor is "DisplayUnitInfo.displayZoomFactor"
	//
	// Optional
	//
	// NOTE: FFI_USE_DisplayZoomFactor MUST be set to true to make this field effective.
	DisplayZoomFactor float64

	FFI_USE_IsPrimary               bool // for IsPrimary.
	FFI_USE_IsInternal              bool // for IsInternal.
	FFI_USE_IsEnabled               bool // for IsEnabled.
	FFI_USE_IsUnified               bool // for IsUnified.
	FFI_USE_IsAutoRotationAllowed   bool // for IsAutoRotationAllowed.
	FFI_USE_DpiX                    bool // for DpiX.
	FFI_USE_DpiY                    bool // for DpiY.
	FFI_USE_Rotation                bool // for Rotation.
	FFI_USE_HasTouchSupport         bool // for HasTouchSupport.
	FFI_USE_HasAccelerometerSupport bool // for HasAccelerometerSupport.
	FFI_USE_DisplayZoomFactor       bool // for DisplayZoomFactor.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DisplayUnitInfo with all fields set.
func (p DisplayUnitInfo) FromRef(ref js.Ref) DisplayUnitInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DisplayUnitInfo in the application heap.
func (p DisplayUnitInfo) New() js.Ref {
	return bindings.DisplayUnitInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DisplayUnitInfo) UpdateFrom(ref js.Ref) {
	bindings.DisplayUnitInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DisplayUnitInfo) Update(ref js.Ref) {
	bindings.DisplayUnitInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DisplayUnitInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Name.Ref(),
		p.MirroringSourceId.Ref(),
		p.MirroringDestinationIds.Ref(),
		p.Modes.Ref(),
		p.AvailableDisplayZoomFactors.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.MirroringSourceId = p.MirroringSourceId.FromRef(js.Undefined)
	p.MirroringDestinationIds = p.MirroringDestinationIds.FromRef(js.Undefined)
	p.Modes = p.Modes.FromRef(js.Undefined)
	p.AvailableDisplayZoomFactors = p.AvailableDisplayZoomFactors.FromRef(js.Undefined)
	if recursive {
		p.Edid.FreeMembers(true)
		p.Bounds.FreeMembers(true)
		p.Overscan.FreeMembers(true)
		p.WorkArea.FreeMembers(true)
	}
}

type LayoutPosition uint32

const (
	_ LayoutPosition = iota

	LayoutPosition_TOP
	LayoutPosition_RIGHT
	LayoutPosition_BOTTOM
	LayoutPosition_LEFT
)

func (LayoutPosition) FromRef(str js.Ref) LayoutPosition {
	return LayoutPosition(bindings.ConstOfLayoutPosition(str))
}

func (x LayoutPosition) String() (string, bool) {
	switch x {
	case LayoutPosition_TOP:
		return "top", true
	case LayoutPosition_RIGHT:
		return "right", true
	case LayoutPosition_BOTTOM:
		return "bottom", true
	case LayoutPosition_LEFT:
		return "left", true
	default:
		return "", false
	}
}

type DisplayLayout struct {
	// Id is "DisplayLayout.id"
	//
	// Optional
	Id js.String
	// ParentId is "DisplayLayout.parentId"
	//
	// Optional
	ParentId js.String
	// Position is "DisplayLayout.position"
	//
	// Optional
	Position LayoutPosition
	// Offset is "DisplayLayout.offset"
	//
	// Optional
	//
	// NOTE: FFI_USE_Offset MUST be set to true to make this field effective.
	Offset int32

	FFI_USE_Offset bool // for Offset.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DisplayLayout with all fields set.
func (p DisplayLayout) FromRef(ref js.Ref) DisplayLayout {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DisplayLayout in the application heap.
func (p DisplayLayout) New() js.Ref {
	return bindings.DisplayLayoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DisplayLayout) UpdateFrom(ref js.Ref) {
	bindings.DisplayLayoutJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DisplayLayout) Update(ref js.Ref) {
	bindings.DisplayLayoutJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DisplayLayout) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.ParentId.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.ParentId = p.ParentId.FromRef(js.Undefined)
}

type DisplayLayoutCallbackFunc func(this js.Ref, layouts js.Array[DisplayLayout]) js.Ref

func (fn DisplayLayoutCallbackFunc) Register() js.Func[func(layouts js.Array[DisplayLayout])] {
	return js.RegisterCallback[func(layouts js.Array[DisplayLayout])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DisplayLayoutCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[DisplayLayout]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DisplayLayoutCallback[T any] struct {
	Fn  func(arg T, this js.Ref, layouts js.Array[DisplayLayout]) js.Ref
	Arg T
}

func (cb *DisplayLayoutCallback[T]) Register() js.Func[func(layouts js.Array[DisplayLayout])] {
	return js.RegisterCallback[func(layouts js.Array[DisplayLayout])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DisplayLayoutCallback[T]) DispatchCallback(
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

		js.Array[DisplayLayout]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DisplayProperties struct {
	// IsUnified is "DisplayProperties.isUnified"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsUnified MUST be set to true to make this field effective.
	IsUnified bool
	// MirroringSourceId is "DisplayProperties.mirroringSourceId"
	//
	// Optional
	MirroringSourceId js.String
	// IsPrimary is "DisplayProperties.isPrimary"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsPrimary MUST be set to true to make this field effective.
	IsPrimary bool
	// Overscan is "DisplayProperties.overscan"
	//
	// Optional
	//
	// NOTE: Overscan.FFI_USE MUST be set to true to get Overscan used.
	Overscan Insets
	// Rotation is "DisplayProperties.rotation"
	//
	// Optional
	//
	// NOTE: FFI_USE_Rotation MUST be set to true to make this field effective.
	Rotation int32
	// BoundsOriginX is "DisplayProperties.boundsOriginX"
	//
	// Optional
	//
	// NOTE: FFI_USE_BoundsOriginX MUST be set to true to make this field effective.
	BoundsOriginX int32
	// BoundsOriginY is "DisplayProperties.boundsOriginY"
	//
	// Optional
	//
	// NOTE: FFI_USE_BoundsOriginY MUST be set to true to make this field effective.
	BoundsOriginY int32
	// DisplayMode is "DisplayProperties.displayMode"
	//
	// Optional
	//
	// NOTE: DisplayMode.FFI_USE MUST be set to true to get DisplayMode used.
	DisplayMode DisplayMode
	// DisplayZoomFactor is "DisplayProperties.displayZoomFactor"
	//
	// Optional
	//
	// NOTE: FFI_USE_DisplayZoomFactor MUST be set to true to make this field effective.
	DisplayZoomFactor float64

	FFI_USE_IsUnified         bool // for IsUnified.
	FFI_USE_IsPrimary         bool // for IsPrimary.
	FFI_USE_Rotation          bool // for Rotation.
	FFI_USE_BoundsOriginX     bool // for BoundsOriginX.
	FFI_USE_BoundsOriginY     bool // for BoundsOriginY.
	FFI_USE_DisplayZoomFactor bool // for DisplayZoomFactor.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DisplayProperties with all fields set.
func (p DisplayProperties) FromRef(ref js.Ref) DisplayProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DisplayProperties in the application heap.
func (p DisplayProperties) New() js.Ref {
	return bindings.DisplayPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DisplayProperties) UpdateFrom(ref js.Ref) {
	bindings.DisplayPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DisplayProperties) Update(ref js.Ref) {
	bindings.DisplayPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DisplayProperties) FreeMembers(recursive bool) {
	js.Free(
		p.MirroringSourceId.Ref(),
	)
	p.MirroringSourceId = p.MirroringSourceId.FromRef(js.Undefined)
	if recursive {
		p.Overscan.FreeMembers(true)
		p.DisplayMode.FreeMembers(true)
	}
}

type GetInfoFlags struct {
	// SingleUnified is "GetInfoFlags.singleUnified"
	//
	// Optional
	//
	// NOTE: FFI_USE_SingleUnified MUST be set to true to make this field effective.
	SingleUnified bool

	FFI_USE_SingleUnified bool // for SingleUnified.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetInfoFlags with all fields set.
func (p GetInfoFlags) FromRef(ref js.Ref) GetInfoFlags {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetInfoFlags in the application heap.
func (p GetInfoFlags) New() js.Ref {
	return bindings.GetInfoFlagsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetInfoFlags) UpdateFrom(ref js.Ref) {
	bindings.GetInfoFlagsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetInfoFlags) Update(ref js.Ref) {
	bindings.GetInfoFlagsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetInfoFlags) FreeMembers(recursive bool) {
}

type MirrorMode uint32

const (
	_ MirrorMode = iota

	MirrorMode_OFF
	MirrorMode_NORMAL
	MirrorMode_MIXED
)

func (MirrorMode) FromRef(str js.Ref) MirrorMode {
	return MirrorMode(bindings.ConstOfMirrorMode(str))
}

func (x MirrorMode) String() (string, bool) {
	switch x {
	case MirrorMode_OFF:
		return "off", true
	case MirrorMode_NORMAL:
		return "normal", true
	case MirrorMode_MIXED:
		return "mixed", true
	default:
		return "", false
	}
}

type MirrorModeInfo struct {
	// Mode is "MirrorModeInfo.mode"
	//
	// Optional
	Mode MirrorMode
	// MirroringSourceId is "MirrorModeInfo.mirroringSourceId"
	//
	// Optional
	MirroringSourceId js.String
	// MirroringDestinationIds is "MirrorModeInfo.mirroringDestinationIds"
	//
	// Optional
	MirroringDestinationIds js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MirrorModeInfo with all fields set.
func (p MirrorModeInfo) FromRef(ref js.Ref) MirrorModeInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MirrorModeInfo in the application heap.
func (p MirrorModeInfo) New() js.Ref {
	return bindings.MirrorModeInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MirrorModeInfo) UpdateFrom(ref js.Ref) {
	bindings.MirrorModeInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MirrorModeInfo) Update(ref js.Ref) {
	bindings.MirrorModeInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MirrorModeInfo) FreeMembers(recursive bool) {
	js.Free(
		p.MirroringSourceId.Ref(),
		p.MirroringDestinationIds.Ref(),
	)
	p.MirroringSourceId = p.MirroringSourceId.FromRef(js.Undefined)
	p.MirroringDestinationIds = p.MirroringDestinationIds.FromRef(js.Undefined)
}

type NativeTouchCalibrationCallbackFunc func(this js.Ref, success bool) js.Ref

func (fn NativeTouchCalibrationCallbackFunc) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn NativeTouchCalibrationCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type NativeTouchCalibrationCallback[T any] struct {
	Fn  func(arg T, this js.Ref, success bool) js.Ref
	Arg T
}

func (cb *NativeTouchCalibrationCallback[T]) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *NativeTouchCalibrationCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type Point struct {
	// X is "Point.x"
	//
	// Optional
	//
	// NOTE: FFI_USE_X MUST be set to true to make this field effective.
	X int32
	// Y is "Point.y"
	//
	// Optional
	//
	// NOTE: FFI_USE_Y MUST be set to true to make this field effective.
	Y int32

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

type SetDisplayLayoutCallbackFunc func(this js.Ref) js.Ref

func (fn SetDisplayLayoutCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetDisplayLayoutCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetDisplayLayoutCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *SetDisplayLayoutCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetDisplayLayoutCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetDisplayUnitInfoCallbackFunc func(this js.Ref) js.Ref

func (fn SetDisplayUnitInfoCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetDisplayUnitInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetDisplayUnitInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *SetDisplayUnitInfoCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetDisplayUnitInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetMirrorModeCallbackFunc func(this js.Ref) js.Ref

func (fn SetMirrorModeCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetMirrorModeCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetMirrorModeCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *SetMirrorModeCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetMirrorModeCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TouchCalibrationPair struct {
	// DisplayPoint is "TouchCalibrationPair.displayPoint"
	//
	// Optional
	//
	// NOTE: DisplayPoint.FFI_USE MUST be set to true to get DisplayPoint used.
	DisplayPoint Point
	// TouchPoint is "TouchCalibrationPair.touchPoint"
	//
	// Optional
	//
	// NOTE: TouchPoint.FFI_USE MUST be set to true to get TouchPoint used.
	TouchPoint Point

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TouchCalibrationPair with all fields set.
func (p TouchCalibrationPair) FromRef(ref js.Ref) TouchCalibrationPair {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TouchCalibrationPair in the application heap.
func (p TouchCalibrationPair) New() js.Ref {
	return bindings.TouchCalibrationPairJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TouchCalibrationPair) UpdateFrom(ref js.Ref) {
	bindings.TouchCalibrationPairJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TouchCalibrationPair) Update(ref js.Ref) {
	bindings.TouchCalibrationPairJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TouchCalibrationPair) FreeMembers(recursive bool) {
	if recursive {
		p.DisplayPoint.FreeMembers(true)
		p.TouchPoint.FreeMembers(true)
	}
}

type TouchCalibrationPairQuad struct {
	// Pair1 is "TouchCalibrationPairQuad.pair1"
	//
	// Optional
	//
	// NOTE: Pair1.FFI_USE MUST be set to true to get Pair1 used.
	Pair1 TouchCalibrationPair
	// Pair2 is "TouchCalibrationPairQuad.pair2"
	//
	// Optional
	//
	// NOTE: Pair2.FFI_USE MUST be set to true to get Pair2 used.
	Pair2 TouchCalibrationPair
	// Pair3 is "TouchCalibrationPairQuad.pair3"
	//
	// Optional
	//
	// NOTE: Pair3.FFI_USE MUST be set to true to get Pair3 used.
	Pair3 TouchCalibrationPair
	// Pair4 is "TouchCalibrationPairQuad.pair4"
	//
	// Optional
	//
	// NOTE: Pair4.FFI_USE MUST be set to true to get Pair4 used.
	Pair4 TouchCalibrationPair

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TouchCalibrationPairQuad with all fields set.
func (p TouchCalibrationPairQuad) FromRef(ref js.Ref) TouchCalibrationPairQuad {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TouchCalibrationPairQuad in the application heap.
func (p TouchCalibrationPairQuad) New() js.Ref {
	return bindings.TouchCalibrationPairQuadJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TouchCalibrationPairQuad) UpdateFrom(ref js.Ref) {
	bindings.TouchCalibrationPairQuadJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TouchCalibrationPairQuad) Update(ref js.Ref) {
	bindings.TouchCalibrationPairQuadJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TouchCalibrationPairQuad) FreeMembers(recursive bool) {
	if recursive {
		p.Pair1.FreeMembers(true)
		p.Pair2.FreeMembers(true)
		p.Pair3.FreeMembers(true)
		p.Pair4.FreeMembers(true)
	}
}

// HasFuncClearTouchCalibration returns true if the function "WEBEXT.system.display.clearTouchCalibration" exists.
func HasFuncClearTouchCalibration() bool {
	return js.True == bindings.HasFuncClearTouchCalibration()
}

// FuncClearTouchCalibration returns the function "WEBEXT.system.display.clearTouchCalibration".
func FuncClearTouchCalibration() (fn js.Func[func(id js.String)]) {
	bindings.FuncClearTouchCalibration(
		js.Pointer(&fn),
	)
	return
}

// ClearTouchCalibration calls the function "WEBEXT.system.display.clearTouchCalibration" directly.
func ClearTouchCalibration(id js.String) (ret js.Void) {
	bindings.CallClearTouchCalibration(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryClearTouchCalibration calls the function "WEBEXT.system.display.clearTouchCalibration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClearTouchCalibration(id js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryClearTouchCalibration(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncCompleteCustomTouchCalibration returns true if the function "WEBEXT.system.display.completeCustomTouchCalibration" exists.
func HasFuncCompleteCustomTouchCalibration() bool {
	return js.True == bindings.HasFuncCompleteCustomTouchCalibration()
}

// FuncCompleteCustomTouchCalibration returns the function "WEBEXT.system.display.completeCustomTouchCalibration".
func FuncCompleteCustomTouchCalibration() (fn js.Func[func(pairs TouchCalibrationPairQuad, bounds Bounds)]) {
	bindings.FuncCompleteCustomTouchCalibration(
		js.Pointer(&fn),
	)
	return
}

// CompleteCustomTouchCalibration calls the function "WEBEXT.system.display.completeCustomTouchCalibration" directly.
func CompleteCustomTouchCalibration(pairs TouchCalibrationPairQuad, bounds Bounds) (ret js.Void) {
	bindings.CallCompleteCustomTouchCalibration(
		js.Pointer(&ret),
		js.Pointer(&pairs),
		js.Pointer(&bounds),
	)

	return
}

// TryCompleteCustomTouchCalibration calls the function "WEBEXT.system.display.completeCustomTouchCalibration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCompleteCustomTouchCalibration(pairs TouchCalibrationPairQuad, bounds Bounds) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCompleteCustomTouchCalibration(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&pairs),
		js.Pointer(&bounds),
	)

	return
}

// HasFuncEnableUnifiedDesktop returns true if the function "WEBEXT.system.display.enableUnifiedDesktop" exists.
func HasFuncEnableUnifiedDesktop() bool {
	return js.True == bindings.HasFuncEnableUnifiedDesktop()
}

// FuncEnableUnifiedDesktop returns the function "WEBEXT.system.display.enableUnifiedDesktop".
func FuncEnableUnifiedDesktop() (fn js.Func[func(enabled bool)]) {
	bindings.FuncEnableUnifiedDesktop(
		js.Pointer(&fn),
	)
	return
}

// EnableUnifiedDesktop calls the function "WEBEXT.system.display.enableUnifiedDesktop" directly.
func EnableUnifiedDesktop(enabled bool) (ret js.Void) {
	bindings.CallEnableUnifiedDesktop(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TryEnableUnifiedDesktop calls the function "WEBEXT.system.display.enableUnifiedDesktop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryEnableUnifiedDesktop(enabled bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEnableUnifiedDesktop(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncGetDisplayLayout returns true if the function "WEBEXT.system.display.getDisplayLayout" exists.
func HasFuncGetDisplayLayout() bool {
	return js.True == bindings.HasFuncGetDisplayLayout()
}

// FuncGetDisplayLayout returns the function "WEBEXT.system.display.getDisplayLayout".
func FuncGetDisplayLayout() (fn js.Func[func() js.Promise[js.Array[DisplayLayout]]]) {
	bindings.FuncGetDisplayLayout(
		js.Pointer(&fn),
	)
	return
}

// GetDisplayLayout calls the function "WEBEXT.system.display.getDisplayLayout" directly.
func GetDisplayLayout() (ret js.Promise[js.Array[DisplayLayout]]) {
	bindings.CallGetDisplayLayout(
		js.Pointer(&ret),
	)

	return
}

// TryGetDisplayLayout calls the function "WEBEXT.system.display.getDisplayLayout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDisplayLayout() (ret js.Promise[js.Array[DisplayLayout]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDisplayLayout(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetInfo returns true if the function "WEBEXT.system.display.getInfo" exists.
func HasFuncGetInfo() bool {
	return js.True == bindings.HasFuncGetInfo()
}

// FuncGetInfo returns the function "WEBEXT.system.display.getInfo".
func FuncGetInfo() (fn js.Func[func(flags GetInfoFlags) js.Promise[js.Array[DisplayUnitInfo]]]) {
	bindings.FuncGetInfo(
		js.Pointer(&fn),
	)
	return
}

// GetInfo calls the function "WEBEXT.system.display.getInfo" directly.
func GetInfo(flags GetInfoFlags) (ret js.Promise[js.Array[DisplayUnitInfo]]) {
	bindings.CallGetInfo(
		js.Pointer(&ret),
		js.Pointer(&flags),
	)

	return
}

// TryGetInfo calls the function "WEBEXT.system.display.getInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetInfo(flags GetInfoFlags) (ret js.Promise[js.Array[DisplayUnitInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&flags),
	)

	return
}

type OnDisplayChangedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnDisplayChangedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDisplayChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDisplayChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnDisplayChangedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDisplayChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDisplayChanged returns true if the function "WEBEXT.system.display.onDisplayChanged.addListener" exists.
func HasFuncOnDisplayChanged() bool {
	return js.True == bindings.HasFuncOnDisplayChanged()
}

// FuncOnDisplayChanged returns the function "WEBEXT.system.display.onDisplayChanged.addListener".
func FuncOnDisplayChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnDisplayChanged(
		js.Pointer(&fn),
	)
	return
}

// OnDisplayChanged calls the function "WEBEXT.system.display.onDisplayChanged.addListener" directly.
func OnDisplayChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnDisplayChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDisplayChanged calls the function "WEBEXT.system.display.onDisplayChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDisplayChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDisplayChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDisplayChanged returns true if the function "WEBEXT.system.display.onDisplayChanged.removeListener" exists.
func HasFuncOffDisplayChanged() bool {
	return js.True == bindings.HasFuncOffDisplayChanged()
}

// FuncOffDisplayChanged returns the function "WEBEXT.system.display.onDisplayChanged.removeListener".
func FuncOffDisplayChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffDisplayChanged(
		js.Pointer(&fn),
	)
	return
}

// OffDisplayChanged calls the function "WEBEXT.system.display.onDisplayChanged.removeListener" directly.
func OffDisplayChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffDisplayChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDisplayChanged calls the function "WEBEXT.system.display.onDisplayChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDisplayChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDisplayChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDisplayChanged returns true if the function "WEBEXT.system.display.onDisplayChanged.hasListener" exists.
func HasFuncHasOnDisplayChanged() bool {
	return js.True == bindings.HasFuncHasOnDisplayChanged()
}

// FuncHasOnDisplayChanged returns the function "WEBEXT.system.display.onDisplayChanged.hasListener".
func FuncHasOnDisplayChanged() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnDisplayChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnDisplayChanged calls the function "WEBEXT.system.display.onDisplayChanged.hasListener" directly.
func HasOnDisplayChanged(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnDisplayChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDisplayChanged calls the function "WEBEXT.system.display.onDisplayChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDisplayChanged(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDisplayChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOverscanCalibrationAdjust returns true if the function "WEBEXT.system.display.overscanCalibrationAdjust" exists.
func HasFuncOverscanCalibrationAdjust() bool {
	return js.True == bindings.HasFuncOverscanCalibrationAdjust()
}

// FuncOverscanCalibrationAdjust returns the function "WEBEXT.system.display.overscanCalibrationAdjust".
func FuncOverscanCalibrationAdjust() (fn js.Func[func(id js.String, delta Insets)]) {
	bindings.FuncOverscanCalibrationAdjust(
		js.Pointer(&fn),
	)
	return
}

// OverscanCalibrationAdjust calls the function "WEBEXT.system.display.overscanCalibrationAdjust" directly.
func OverscanCalibrationAdjust(id js.String, delta Insets) (ret js.Void) {
	bindings.CallOverscanCalibrationAdjust(
		js.Pointer(&ret),
		id.Ref(),
		js.Pointer(&delta),
	)

	return
}

// TryOverscanCalibrationAdjust calls the function "WEBEXT.system.display.overscanCalibrationAdjust"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOverscanCalibrationAdjust(id js.String, delta Insets) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOverscanCalibrationAdjust(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		js.Pointer(&delta),
	)

	return
}

// HasFuncOverscanCalibrationComplete returns true if the function "WEBEXT.system.display.overscanCalibrationComplete" exists.
func HasFuncOverscanCalibrationComplete() bool {
	return js.True == bindings.HasFuncOverscanCalibrationComplete()
}

// FuncOverscanCalibrationComplete returns the function "WEBEXT.system.display.overscanCalibrationComplete".
func FuncOverscanCalibrationComplete() (fn js.Func[func(id js.String)]) {
	bindings.FuncOverscanCalibrationComplete(
		js.Pointer(&fn),
	)
	return
}

// OverscanCalibrationComplete calls the function "WEBEXT.system.display.overscanCalibrationComplete" directly.
func OverscanCalibrationComplete(id js.String) (ret js.Void) {
	bindings.CallOverscanCalibrationComplete(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryOverscanCalibrationComplete calls the function "WEBEXT.system.display.overscanCalibrationComplete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOverscanCalibrationComplete(id js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOverscanCalibrationComplete(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncOverscanCalibrationReset returns true if the function "WEBEXT.system.display.overscanCalibrationReset" exists.
func HasFuncOverscanCalibrationReset() bool {
	return js.True == bindings.HasFuncOverscanCalibrationReset()
}

// FuncOverscanCalibrationReset returns the function "WEBEXT.system.display.overscanCalibrationReset".
func FuncOverscanCalibrationReset() (fn js.Func[func(id js.String)]) {
	bindings.FuncOverscanCalibrationReset(
		js.Pointer(&fn),
	)
	return
}

// OverscanCalibrationReset calls the function "WEBEXT.system.display.overscanCalibrationReset" directly.
func OverscanCalibrationReset(id js.String) (ret js.Void) {
	bindings.CallOverscanCalibrationReset(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryOverscanCalibrationReset calls the function "WEBEXT.system.display.overscanCalibrationReset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOverscanCalibrationReset(id js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOverscanCalibrationReset(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncOverscanCalibrationStart returns true if the function "WEBEXT.system.display.overscanCalibrationStart" exists.
func HasFuncOverscanCalibrationStart() bool {
	return js.True == bindings.HasFuncOverscanCalibrationStart()
}

// FuncOverscanCalibrationStart returns the function "WEBEXT.system.display.overscanCalibrationStart".
func FuncOverscanCalibrationStart() (fn js.Func[func(id js.String)]) {
	bindings.FuncOverscanCalibrationStart(
		js.Pointer(&fn),
	)
	return
}

// OverscanCalibrationStart calls the function "WEBEXT.system.display.overscanCalibrationStart" directly.
func OverscanCalibrationStart(id js.String) (ret js.Void) {
	bindings.CallOverscanCalibrationStart(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryOverscanCalibrationStart calls the function "WEBEXT.system.display.overscanCalibrationStart"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOverscanCalibrationStart(id js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOverscanCalibrationStart(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncSetDisplayLayout returns true if the function "WEBEXT.system.display.setDisplayLayout" exists.
func HasFuncSetDisplayLayout() bool {
	return js.True == bindings.HasFuncSetDisplayLayout()
}

// FuncSetDisplayLayout returns the function "WEBEXT.system.display.setDisplayLayout".
func FuncSetDisplayLayout() (fn js.Func[func(layouts js.Array[DisplayLayout]) js.Promise[js.Void]]) {
	bindings.FuncSetDisplayLayout(
		js.Pointer(&fn),
	)
	return
}

// SetDisplayLayout calls the function "WEBEXT.system.display.setDisplayLayout" directly.
func SetDisplayLayout(layouts js.Array[DisplayLayout]) (ret js.Promise[js.Void]) {
	bindings.CallSetDisplayLayout(
		js.Pointer(&ret),
		layouts.Ref(),
	)

	return
}

// TrySetDisplayLayout calls the function "WEBEXT.system.display.setDisplayLayout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetDisplayLayout(layouts js.Array[DisplayLayout]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetDisplayLayout(
		js.Pointer(&ret), js.Pointer(&exception),
		layouts.Ref(),
	)

	return
}

// HasFuncSetDisplayProperties returns true if the function "WEBEXT.system.display.setDisplayProperties" exists.
func HasFuncSetDisplayProperties() bool {
	return js.True == bindings.HasFuncSetDisplayProperties()
}

// FuncSetDisplayProperties returns the function "WEBEXT.system.display.setDisplayProperties".
func FuncSetDisplayProperties() (fn js.Func[func(id js.String, info DisplayProperties) js.Promise[js.Void]]) {
	bindings.FuncSetDisplayProperties(
		js.Pointer(&fn),
	)
	return
}

// SetDisplayProperties calls the function "WEBEXT.system.display.setDisplayProperties" directly.
func SetDisplayProperties(id js.String, info DisplayProperties) (ret js.Promise[js.Void]) {
	bindings.CallSetDisplayProperties(
		js.Pointer(&ret),
		id.Ref(),
		js.Pointer(&info),
	)

	return
}

// TrySetDisplayProperties calls the function "WEBEXT.system.display.setDisplayProperties"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetDisplayProperties(id js.String, info DisplayProperties) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetDisplayProperties(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		js.Pointer(&info),
	)

	return
}

// HasFuncSetMirrorMode returns true if the function "WEBEXT.system.display.setMirrorMode" exists.
func HasFuncSetMirrorMode() bool {
	return js.True == bindings.HasFuncSetMirrorMode()
}

// FuncSetMirrorMode returns the function "WEBEXT.system.display.setMirrorMode".
func FuncSetMirrorMode() (fn js.Func[func(info MirrorModeInfo) js.Promise[js.Void]]) {
	bindings.FuncSetMirrorMode(
		js.Pointer(&fn),
	)
	return
}

// SetMirrorMode calls the function "WEBEXT.system.display.setMirrorMode" directly.
func SetMirrorMode(info MirrorModeInfo) (ret js.Promise[js.Void]) {
	bindings.CallSetMirrorMode(
		js.Pointer(&ret),
		js.Pointer(&info),
	)

	return
}

// TrySetMirrorMode calls the function "WEBEXT.system.display.setMirrorMode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetMirrorMode(info MirrorModeInfo) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetMirrorMode(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&info),
	)

	return
}

// HasFuncShowNativeTouchCalibration returns true if the function "WEBEXT.system.display.showNativeTouchCalibration" exists.
func HasFuncShowNativeTouchCalibration() bool {
	return js.True == bindings.HasFuncShowNativeTouchCalibration()
}

// FuncShowNativeTouchCalibration returns the function "WEBEXT.system.display.showNativeTouchCalibration".
func FuncShowNativeTouchCalibration() (fn js.Func[func(id js.String) js.Promise[js.Boolean]]) {
	bindings.FuncShowNativeTouchCalibration(
		js.Pointer(&fn),
	)
	return
}

// ShowNativeTouchCalibration calls the function "WEBEXT.system.display.showNativeTouchCalibration" directly.
func ShowNativeTouchCalibration(id js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallShowNativeTouchCalibration(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryShowNativeTouchCalibration calls the function "WEBEXT.system.display.showNativeTouchCalibration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShowNativeTouchCalibration(id js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryShowNativeTouchCalibration(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncStartCustomTouchCalibration returns true if the function "WEBEXT.system.display.startCustomTouchCalibration" exists.
func HasFuncStartCustomTouchCalibration() bool {
	return js.True == bindings.HasFuncStartCustomTouchCalibration()
}

// FuncStartCustomTouchCalibration returns the function "WEBEXT.system.display.startCustomTouchCalibration".
func FuncStartCustomTouchCalibration() (fn js.Func[func(id js.String)]) {
	bindings.FuncStartCustomTouchCalibration(
		js.Pointer(&fn),
	)
	return
}

// StartCustomTouchCalibration calls the function "WEBEXT.system.display.startCustomTouchCalibration" directly.
func StartCustomTouchCalibration(id js.String) (ret js.Void) {
	bindings.CallStartCustomTouchCalibration(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryStartCustomTouchCalibration calls the function "WEBEXT.system.display.startCustomTouchCalibration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartCustomTouchCalibration(id js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartCustomTouchCalibration(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

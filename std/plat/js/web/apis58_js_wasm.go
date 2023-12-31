// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

type XRLayerLayout uint32

const (
	_ XRLayerLayout = iota

	XRLayerLayout_DEFAULT
	XRLayerLayout_MONO
	XRLayerLayout_STEREO
	XRLayerLayout_STEREO_LEFT_RIGHT
	XRLayerLayout_STEREO_TOP_BOTTOM
)

func (XRLayerLayout) FromRef(str js.Ref) XRLayerLayout {
	return XRLayerLayout(bindings.ConstOfXRLayerLayout(str))
}

func (x XRLayerLayout) String() (string, bool) {
	switch x {
	case XRLayerLayout_DEFAULT:
		return "default", true
	case XRLayerLayout_MONO:
		return "mono", true
	case XRLayerLayout_STEREO:
		return "stereo", true
	case XRLayerLayout_STEREO_LEFT_RIGHT:
		return "stereo-left-right", true
	case XRLayerLayout_STEREO_TOP_BOTTOM:
		return "stereo-top-bottom", true
	default:
		return "", false
	}
}

type XRLayerQuality uint32

const (
	_ XRLayerQuality = iota

	XRLayerQuality_DEFAULT
	XRLayerQuality_TEXT_OPTIMIZED
	XRLayerQuality_GRAPHICS_OPTIMIZED
)

func (XRLayerQuality) FromRef(str js.Ref) XRLayerQuality {
	return XRLayerQuality(bindings.ConstOfXRLayerQuality(str))
}

func (x XRLayerQuality) String() (string, bool) {
	switch x {
	case XRLayerQuality_DEFAULT:
		return "default", true
	case XRLayerQuality_TEXT_OPTIMIZED:
		return "text-optimized", true
	case XRLayerQuality_GRAPHICS_OPTIMIZED:
		return "graphics-optimized", true
	default:
		return "", false
	}
}

type XRCompositionLayer struct {
	XRLayer
}

func (this XRCompositionLayer) Once() XRCompositionLayer {
	this.ref.Once()
	return this
}

func (this XRCompositionLayer) Ref() js.Ref {
	return this.XRLayer.Ref()
}

func (this XRCompositionLayer) FromRef(ref js.Ref) XRCompositionLayer {
	this.XRLayer = this.XRLayer.FromRef(ref)
	return this
}

func (this XRCompositionLayer) Free() {
	this.ref.Free()
}

// Layout returns the value of property "XRCompositionLayer.layout".
//
// It returns ok=false if there is no such property.
func (this XRCompositionLayer) Layout() (ret XRLayerLayout, ok bool) {
	ok = js.True == bindings.GetXRCompositionLayerLayout(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BlendTextureSourceAlpha returns the value of property "XRCompositionLayer.blendTextureSourceAlpha".
//
// It returns ok=false if there is no such property.
func (this XRCompositionLayer) BlendTextureSourceAlpha() (ret bool, ok bool) {
	ok = js.True == bindings.GetXRCompositionLayerBlendTextureSourceAlpha(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBlendTextureSourceAlpha sets the value of property "XRCompositionLayer.blendTextureSourceAlpha" to val.
//
// It returns false if the property cannot be set.
func (this XRCompositionLayer) SetBlendTextureSourceAlpha(val bool) bool {
	return js.True == bindings.SetXRCompositionLayerBlendTextureSourceAlpha(
		this.ref,
		js.Bool(bool(val)),
	)
}

// ForceMonoPresentation returns the value of property "XRCompositionLayer.forceMonoPresentation".
//
// It returns ok=false if there is no such property.
func (this XRCompositionLayer) ForceMonoPresentation() (ret bool, ok bool) {
	ok = js.True == bindings.GetXRCompositionLayerForceMonoPresentation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetForceMonoPresentation sets the value of property "XRCompositionLayer.forceMonoPresentation" to val.
//
// It returns false if the property cannot be set.
func (this XRCompositionLayer) SetForceMonoPresentation(val bool) bool {
	return js.True == bindings.SetXRCompositionLayerForceMonoPresentation(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Opacity returns the value of property "XRCompositionLayer.opacity".
//
// It returns ok=false if there is no such property.
func (this XRCompositionLayer) Opacity() (ret float32, ok bool) {
	ok = js.True == bindings.GetXRCompositionLayerOpacity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetOpacity sets the value of property "XRCompositionLayer.opacity" to val.
//
// It returns false if the property cannot be set.
func (this XRCompositionLayer) SetOpacity(val float32) bool {
	return js.True == bindings.SetXRCompositionLayerOpacity(
		this.ref,
		float32(val),
	)
}

// MipLevels returns the value of property "XRCompositionLayer.mipLevels".
//
// It returns ok=false if there is no such property.
func (this XRCompositionLayer) MipLevels() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRCompositionLayerMipLevels(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Quality returns the value of property "XRCompositionLayer.quality".
//
// It returns ok=false if there is no such property.
func (this XRCompositionLayer) Quality() (ret XRLayerQuality, ok bool) {
	ok = js.True == bindings.GetXRCompositionLayerQuality(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetQuality sets the value of property "XRCompositionLayer.quality" to val.
//
// It returns false if the property cannot be set.
func (this XRCompositionLayer) SetQuality(val XRLayerQuality) bool {
	return js.True == bindings.SetXRCompositionLayerQuality(
		this.ref,
		uint32(val),
	)
}

// NeedsRedraw returns the value of property "XRCompositionLayer.needsRedraw".
//
// It returns ok=false if there is no such property.
func (this XRCompositionLayer) NeedsRedraw() (ret bool, ok bool) {
	ok = js.True == bindings.GetXRCompositionLayerNeedsRedraw(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncDestroy returns true if the method "XRCompositionLayer.destroy" exists.
func (this XRCompositionLayer) HasFuncDestroy() bool {
	return js.True == bindings.HasFuncXRCompositionLayerDestroy(
		this.ref,
	)
}

// FuncDestroy returns the method "XRCompositionLayer.destroy".
func (this XRCompositionLayer) FuncDestroy() (fn js.Func[func()]) {
	bindings.FuncXRCompositionLayerDestroy(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Destroy calls the method "XRCompositionLayer.destroy".
func (this XRCompositionLayer) Destroy() (ret js.Void) {
	bindings.CallXRCompositionLayerDestroy(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDestroy calls the method "XRCompositionLayer.destroy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRCompositionLayer) TryDestroy() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRCompositionLayerDestroy(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type XRCubeLayer struct {
	XRCompositionLayer
}

func (this XRCubeLayer) Once() XRCubeLayer {
	this.ref.Once()
	return this
}

func (this XRCubeLayer) Ref() js.Ref {
	return this.XRCompositionLayer.Ref()
}

func (this XRCubeLayer) FromRef(ref js.Ref) XRCubeLayer {
	this.XRCompositionLayer = this.XRCompositionLayer.FromRef(ref)
	return this
}

func (this XRCubeLayer) Free() {
	this.ref.Free()
}

// Space returns the value of property "XRCubeLayer.space".
//
// It returns ok=false if there is no such property.
func (this XRCubeLayer) Space() (ret XRSpace, ok bool) {
	ok = js.True == bindings.GetXRCubeLayerSpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSpace sets the value of property "XRCubeLayer.space" to val.
//
// It returns false if the property cannot be set.
func (this XRCubeLayer) SetSpace(val XRSpace) bool {
	return js.True == bindings.SetXRCubeLayerSpace(
		this.ref,
		val.Ref(),
	)
}

// Orientation returns the value of property "XRCubeLayer.orientation".
//
// It returns ok=false if there is no such property.
func (this XRCubeLayer) Orientation() (ret DOMPointReadOnly, ok bool) {
	ok = js.True == bindings.GetXRCubeLayerOrientation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetOrientation sets the value of property "XRCubeLayer.orientation" to val.
//
// It returns false if the property cannot be set.
func (this XRCubeLayer) SetOrientation(val DOMPointReadOnly) bool {
	return js.True == bindings.SetXRCubeLayerOrientation(
		this.ref,
		val.Ref(),
	)
}

type XRCubeLayerInit struct {
	// Orientation is "XRCubeLayerInit.orientation"
	//
	// Optional
	Orientation DOMPointReadOnly
	// Space is "XRCubeLayerInit.space"
	//
	// Required
	Space XRSpace
	// ColorFormat is "XRCubeLayerInit.colorFormat"
	//
	// Optional, defaults to 0x1908.
	//
	// NOTE: FFI_USE_ColorFormat MUST be set to true to make this field effective.
	ColorFormat GLenum
	// DepthFormat is "XRCubeLayerInit.depthFormat"
	//
	// Optional
	//
	// NOTE: FFI_USE_DepthFormat MUST be set to true to make this field effective.
	DepthFormat GLenum
	// MipLevels is "XRCubeLayerInit.mipLevels"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_MipLevels MUST be set to true to make this field effective.
	MipLevels uint32
	// ViewPixelWidth is "XRCubeLayerInit.viewPixelWidth"
	//
	// Required
	ViewPixelWidth uint32
	// ViewPixelHeight is "XRCubeLayerInit.viewPixelHeight"
	//
	// Required
	ViewPixelHeight uint32
	// Layout is "XRCubeLayerInit.layout"
	//
	// Optional, defaults to "mono".
	Layout XRLayerLayout
	// IsStatic is "XRCubeLayerInit.isStatic"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IsStatic MUST be set to true to make this field effective.
	IsStatic bool
	// ClearOnAccess is "XRCubeLayerInit.clearOnAccess"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_ClearOnAccess MUST be set to true to make this field effective.
	ClearOnAccess bool

	FFI_USE_ColorFormat   bool // for ColorFormat.
	FFI_USE_DepthFormat   bool // for DepthFormat.
	FFI_USE_MipLevels     bool // for MipLevels.
	FFI_USE_IsStatic      bool // for IsStatic.
	FFI_USE_ClearOnAccess bool // for ClearOnAccess.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRCubeLayerInit with all fields set.
func (p XRCubeLayerInit) FromRef(ref js.Ref) XRCubeLayerInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRCubeLayerInit in the application heap.
func (p XRCubeLayerInit) New() js.Ref {
	return bindings.XRCubeLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRCubeLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRCubeLayerInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRCubeLayerInit) Update(ref js.Ref) {
	bindings.XRCubeLayerInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRCubeLayerInit) FreeMembers(recursive bool) {
	js.Free(
		p.Orientation.Ref(),
		p.Space.Ref(),
	)
	p.Orientation = p.Orientation.FromRef(js.Undefined)
	p.Space = p.Space.FromRef(js.Undefined)
}

type XRCylinderLayer struct {
	XRCompositionLayer
}

func (this XRCylinderLayer) Once() XRCylinderLayer {
	this.ref.Once()
	return this
}

func (this XRCylinderLayer) Ref() js.Ref {
	return this.XRCompositionLayer.Ref()
}

func (this XRCylinderLayer) FromRef(ref js.Ref) XRCylinderLayer {
	this.XRCompositionLayer = this.XRCompositionLayer.FromRef(ref)
	return this
}

func (this XRCylinderLayer) Free() {
	this.ref.Free()
}

// Space returns the value of property "XRCylinderLayer.space".
//
// It returns ok=false if there is no such property.
func (this XRCylinderLayer) Space() (ret XRSpace, ok bool) {
	ok = js.True == bindings.GetXRCylinderLayerSpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSpace sets the value of property "XRCylinderLayer.space" to val.
//
// It returns false if the property cannot be set.
func (this XRCylinderLayer) SetSpace(val XRSpace) bool {
	return js.True == bindings.SetXRCylinderLayerSpace(
		this.ref,
		val.Ref(),
	)
}

// Transform returns the value of property "XRCylinderLayer.transform".
//
// It returns ok=false if there is no such property.
func (this XRCylinderLayer) Transform() (ret XRRigidTransform, ok bool) {
	ok = js.True == bindings.GetXRCylinderLayerTransform(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTransform sets the value of property "XRCylinderLayer.transform" to val.
//
// It returns false if the property cannot be set.
func (this XRCylinderLayer) SetTransform(val XRRigidTransform) bool {
	return js.True == bindings.SetXRCylinderLayerTransform(
		this.ref,
		val.Ref(),
	)
}

// Radius returns the value of property "XRCylinderLayer.radius".
//
// It returns ok=false if there is no such property.
func (this XRCylinderLayer) Radius() (ret float32, ok bool) {
	ok = js.True == bindings.GetXRCylinderLayerRadius(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRadius sets the value of property "XRCylinderLayer.radius" to val.
//
// It returns false if the property cannot be set.
func (this XRCylinderLayer) SetRadius(val float32) bool {
	return js.True == bindings.SetXRCylinderLayerRadius(
		this.ref,
		float32(val),
	)
}

// CentralAngle returns the value of property "XRCylinderLayer.centralAngle".
//
// It returns ok=false if there is no such property.
func (this XRCylinderLayer) CentralAngle() (ret float32, ok bool) {
	ok = js.True == bindings.GetXRCylinderLayerCentralAngle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCentralAngle sets the value of property "XRCylinderLayer.centralAngle" to val.
//
// It returns false if the property cannot be set.
func (this XRCylinderLayer) SetCentralAngle(val float32) bool {
	return js.True == bindings.SetXRCylinderLayerCentralAngle(
		this.ref,
		float32(val),
	)
}

// AspectRatio returns the value of property "XRCylinderLayer.aspectRatio".
//
// It returns ok=false if there is no such property.
func (this XRCylinderLayer) AspectRatio() (ret float32, ok bool) {
	ok = js.True == bindings.GetXRCylinderLayerAspectRatio(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAspectRatio sets the value of property "XRCylinderLayer.aspectRatio" to val.
//
// It returns false if the property cannot be set.
func (this XRCylinderLayer) SetAspectRatio(val float32) bool {
	return js.True == bindings.SetXRCylinderLayerAspectRatio(
		this.ref,
		float32(val),
	)
}

type XRTextureType uint32

const (
	_ XRTextureType = iota

	XRTextureType_TEXTURE
	XRTextureType_TEXTURE_ARRAY
)

func (XRTextureType) FromRef(str js.Ref) XRTextureType {
	return XRTextureType(bindings.ConstOfXRTextureType(str))
}

func (x XRTextureType) String() (string, bool) {
	switch x {
	case XRTextureType_TEXTURE:
		return "texture", true
	case XRTextureType_TEXTURE_ARRAY:
		return "texture-array", true
	default:
		return "", false
	}
}

type XRCylinderLayerInit struct {
	// TextureType is "XRCylinderLayerInit.textureType"
	//
	// Optional, defaults to "texture".
	TextureType XRTextureType
	// Transform is "XRCylinderLayerInit.transform"
	//
	// Optional
	Transform XRRigidTransform
	// Radius is "XRCylinderLayerInit.radius"
	//
	// Optional, defaults to 2.0.
	//
	// NOTE: FFI_USE_Radius MUST be set to true to make this field effective.
	Radius float32
	// CentralAngle is "XRCylinderLayerInit.centralAngle"
	//
	// Optional, defaults to 0.78539.
	//
	// NOTE: FFI_USE_CentralAngle MUST be set to true to make this field effective.
	CentralAngle float32
	// AspectRatio is "XRCylinderLayerInit.aspectRatio"
	//
	// Optional, defaults to 2.0.
	//
	// NOTE: FFI_USE_AspectRatio MUST be set to true to make this field effective.
	AspectRatio float32
	// Space is "XRCylinderLayerInit.space"
	//
	// Required
	Space XRSpace
	// ColorFormat is "XRCylinderLayerInit.colorFormat"
	//
	// Optional, defaults to 0x1908.
	//
	// NOTE: FFI_USE_ColorFormat MUST be set to true to make this field effective.
	ColorFormat GLenum
	// DepthFormat is "XRCylinderLayerInit.depthFormat"
	//
	// Optional
	//
	// NOTE: FFI_USE_DepthFormat MUST be set to true to make this field effective.
	DepthFormat GLenum
	// MipLevels is "XRCylinderLayerInit.mipLevels"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_MipLevels MUST be set to true to make this field effective.
	MipLevels uint32
	// ViewPixelWidth is "XRCylinderLayerInit.viewPixelWidth"
	//
	// Required
	ViewPixelWidth uint32
	// ViewPixelHeight is "XRCylinderLayerInit.viewPixelHeight"
	//
	// Required
	ViewPixelHeight uint32
	// Layout is "XRCylinderLayerInit.layout"
	//
	// Optional, defaults to "mono".
	Layout XRLayerLayout
	// IsStatic is "XRCylinderLayerInit.isStatic"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IsStatic MUST be set to true to make this field effective.
	IsStatic bool
	// ClearOnAccess is "XRCylinderLayerInit.clearOnAccess"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_ClearOnAccess MUST be set to true to make this field effective.
	ClearOnAccess bool

	FFI_USE_Radius        bool // for Radius.
	FFI_USE_CentralAngle  bool // for CentralAngle.
	FFI_USE_AspectRatio   bool // for AspectRatio.
	FFI_USE_ColorFormat   bool // for ColorFormat.
	FFI_USE_DepthFormat   bool // for DepthFormat.
	FFI_USE_MipLevels     bool // for MipLevels.
	FFI_USE_IsStatic      bool // for IsStatic.
	FFI_USE_ClearOnAccess bool // for ClearOnAccess.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRCylinderLayerInit with all fields set.
func (p XRCylinderLayerInit) FromRef(ref js.Ref) XRCylinderLayerInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRCylinderLayerInit in the application heap.
func (p XRCylinderLayerInit) New() js.Ref {
	return bindings.XRCylinderLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRCylinderLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRCylinderLayerInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRCylinderLayerInit) Update(ref js.Ref) {
	bindings.XRCylinderLayerInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRCylinderLayerInit) FreeMembers(recursive bool) {
	js.Free(
		p.Transform.Ref(),
		p.Space.Ref(),
	)
	p.Transform = p.Transform.FromRef(js.Undefined)
	p.Space = p.Space.FromRef(js.Undefined)
}

type XRDepthInformation struct {
	ref js.Ref
}

func (this XRDepthInformation) Once() XRDepthInformation {
	this.ref.Once()
	return this
}

func (this XRDepthInformation) Ref() js.Ref {
	return this.ref
}

func (this XRDepthInformation) FromRef(ref js.Ref) XRDepthInformation {
	this.ref = ref
	return this
}

func (this XRDepthInformation) Free() {
	this.ref.Free()
}

// Width returns the value of property "XRDepthInformation.width".
//
// It returns ok=false if there is no such property.
func (this XRDepthInformation) Width() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRDepthInformationWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "XRDepthInformation.height".
//
// It returns ok=false if there is no such property.
func (this XRDepthInformation) Height() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRDepthInformationHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NormDepthBufferFromNormView returns the value of property "XRDepthInformation.normDepthBufferFromNormView".
//
// It returns ok=false if there is no such property.
func (this XRDepthInformation) NormDepthBufferFromNormView() (ret XRRigidTransform, ok bool) {
	ok = js.True == bindings.GetXRDepthInformationNormDepthBufferFromNormView(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RawValueToMeters returns the value of property "XRDepthInformation.rawValueToMeters".
//
// It returns ok=false if there is no such property.
func (this XRDepthInformation) RawValueToMeters() (ret float32, ok bool) {
	ok = js.True == bindings.GetXRDepthInformationRawValueToMeters(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XREquirectLayer struct {
	XRCompositionLayer
}

func (this XREquirectLayer) Once() XREquirectLayer {
	this.ref.Once()
	return this
}

func (this XREquirectLayer) Ref() js.Ref {
	return this.XRCompositionLayer.Ref()
}

func (this XREquirectLayer) FromRef(ref js.Ref) XREquirectLayer {
	this.XRCompositionLayer = this.XRCompositionLayer.FromRef(ref)
	return this
}

func (this XREquirectLayer) Free() {
	this.ref.Free()
}

// Space returns the value of property "XREquirectLayer.space".
//
// It returns ok=false if there is no such property.
func (this XREquirectLayer) Space() (ret XRSpace, ok bool) {
	ok = js.True == bindings.GetXREquirectLayerSpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSpace sets the value of property "XREquirectLayer.space" to val.
//
// It returns false if the property cannot be set.
func (this XREquirectLayer) SetSpace(val XRSpace) bool {
	return js.True == bindings.SetXREquirectLayerSpace(
		this.ref,
		val.Ref(),
	)
}

// Transform returns the value of property "XREquirectLayer.transform".
//
// It returns ok=false if there is no such property.
func (this XREquirectLayer) Transform() (ret XRRigidTransform, ok bool) {
	ok = js.True == bindings.GetXREquirectLayerTransform(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTransform sets the value of property "XREquirectLayer.transform" to val.
//
// It returns false if the property cannot be set.
func (this XREquirectLayer) SetTransform(val XRRigidTransform) bool {
	return js.True == bindings.SetXREquirectLayerTransform(
		this.ref,
		val.Ref(),
	)
}

// Radius returns the value of property "XREquirectLayer.radius".
//
// It returns ok=false if there is no such property.
func (this XREquirectLayer) Radius() (ret float32, ok bool) {
	ok = js.True == bindings.GetXREquirectLayerRadius(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRadius sets the value of property "XREquirectLayer.radius" to val.
//
// It returns false if the property cannot be set.
func (this XREquirectLayer) SetRadius(val float32) bool {
	return js.True == bindings.SetXREquirectLayerRadius(
		this.ref,
		float32(val),
	)
}

// CentralHorizontalAngle returns the value of property "XREquirectLayer.centralHorizontalAngle".
//
// It returns ok=false if there is no such property.
func (this XREquirectLayer) CentralHorizontalAngle() (ret float32, ok bool) {
	ok = js.True == bindings.GetXREquirectLayerCentralHorizontalAngle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCentralHorizontalAngle sets the value of property "XREquirectLayer.centralHorizontalAngle" to val.
//
// It returns false if the property cannot be set.
func (this XREquirectLayer) SetCentralHorizontalAngle(val float32) bool {
	return js.True == bindings.SetXREquirectLayerCentralHorizontalAngle(
		this.ref,
		float32(val),
	)
}

// UpperVerticalAngle returns the value of property "XREquirectLayer.upperVerticalAngle".
//
// It returns ok=false if there is no such property.
func (this XREquirectLayer) UpperVerticalAngle() (ret float32, ok bool) {
	ok = js.True == bindings.GetXREquirectLayerUpperVerticalAngle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUpperVerticalAngle sets the value of property "XREquirectLayer.upperVerticalAngle" to val.
//
// It returns false if the property cannot be set.
func (this XREquirectLayer) SetUpperVerticalAngle(val float32) bool {
	return js.True == bindings.SetXREquirectLayerUpperVerticalAngle(
		this.ref,
		float32(val),
	)
}

// LowerVerticalAngle returns the value of property "XREquirectLayer.lowerVerticalAngle".
//
// It returns ok=false if there is no such property.
func (this XREquirectLayer) LowerVerticalAngle() (ret float32, ok bool) {
	ok = js.True == bindings.GetXREquirectLayerLowerVerticalAngle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLowerVerticalAngle sets the value of property "XREquirectLayer.lowerVerticalAngle" to val.
//
// It returns false if the property cannot be set.
func (this XREquirectLayer) SetLowerVerticalAngle(val float32) bool {
	return js.True == bindings.SetXREquirectLayerLowerVerticalAngle(
		this.ref,
		float32(val),
	)
}

type XREquirectLayerInit struct {
	// TextureType is "XREquirectLayerInit.textureType"
	//
	// Optional, defaults to "texture".
	TextureType XRTextureType
	// Transform is "XREquirectLayerInit.transform"
	//
	// Optional
	Transform XRRigidTransform
	// Radius is "XREquirectLayerInit.radius"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Radius MUST be set to true to make this field effective.
	Radius float32
	// CentralHorizontalAngle is "XREquirectLayerInit.centralHorizontalAngle"
	//
	// Optional, defaults to 6.28318.
	//
	// NOTE: FFI_USE_CentralHorizontalAngle MUST be set to true to make this field effective.
	CentralHorizontalAngle float32
	// UpperVerticalAngle is "XREquirectLayerInit.upperVerticalAngle"
	//
	// Optional, defaults to 1.570795.
	//
	// NOTE: FFI_USE_UpperVerticalAngle MUST be set to true to make this field effective.
	UpperVerticalAngle float32
	// LowerVerticalAngle is "XREquirectLayerInit.lowerVerticalAngle"
	//
	// Optional, defaults to -1.570795.
	//
	// NOTE: FFI_USE_LowerVerticalAngle MUST be set to true to make this field effective.
	LowerVerticalAngle float32
	// Space is "XREquirectLayerInit.space"
	//
	// Required
	Space XRSpace
	// ColorFormat is "XREquirectLayerInit.colorFormat"
	//
	// Optional, defaults to 0x1908.
	//
	// NOTE: FFI_USE_ColorFormat MUST be set to true to make this field effective.
	ColorFormat GLenum
	// DepthFormat is "XREquirectLayerInit.depthFormat"
	//
	// Optional
	//
	// NOTE: FFI_USE_DepthFormat MUST be set to true to make this field effective.
	DepthFormat GLenum
	// MipLevels is "XREquirectLayerInit.mipLevels"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_MipLevels MUST be set to true to make this field effective.
	MipLevels uint32
	// ViewPixelWidth is "XREquirectLayerInit.viewPixelWidth"
	//
	// Required
	ViewPixelWidth uint32
	// ViewPixelHeight is "XREquirectLayerInit.viewPixelHeight"
	//
	// Required
	ViewPixelHeight uint32
	// Layout is "XREquirectLayerInit.layout"
	//
	// Optional, defaults to "mono".
	Layout XRLayerLayout
	// IsStatic is "XREquirectLayerInit.isStatic"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IsStatic MUST be set to true to make this field effective.
	IsStatic bool
	// ClearOnAccess is "XREquirectLayerInit.clearOnAccess"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_ClearOnAccess MUST be set to true to make this field effective.
	ClearOnAccess bool

	FFI_USE_Radius                 bool // for Radius.
	FFI_USE_CentralHorizontalAngle bool // for CentralHorizontalAngle.
	FFI_USE_UpperVerticalAngle     bool // for UpperVerticalAngle.
	FFI_USE_LowerVerticalAngle     bool // for LowerVerticalAngle.
	FFI_USE_ColorFormat            bool // for ColorFormat.
	FFI_USE_DepthFormat            bool // for DepthFormat.
	FFI_USE_MipLevels              bool // for MipLevels.
	FFI_USE_IsStatic               bool // for IsStatic.
	FFI_USE_ClearOnAccess          bool // for ClearOnAccess.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XREquirectLayerInit with all fields set.
func (p XREquirectLayerInit) FromRef(ref js.Ref) XREquirectLayerInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XREquirectLayerInit in the application heap.
func (p XREquirectLayerInit) New() js.Ref {
	return bindings.XREquirectLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XREquirectLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XREquirectLayerInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XREquirectLayerInit) Update(ref js.Ref) {
	bindings.XREquirectLayerInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XREquirectLayerInit) FreeMembers(recursive bool) {
	js.Free(
		p.Transform.Ref(),
		p.Space.Ref(),
	)
	p.Transform = p.Transform.FromRef(js.Undefined)
	p.Space = p.Space.FromRef(js.Undefined)
}

type XRInputSourceEventInit struct {
	// Frame is "XRInputSourceEventInit.frame"
	//
	// Required
	Frame XRFrame
	// InputSource is "XRInputSourceEventInit.inputSource"
	//
	// Required
	InputSource XRInputSource
	// Bubbles is "XRInputSourceEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "XRInputSourceEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "XRInputSourceEventInit.composed"
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

// FromRef calls UpdateFrom and returns a XRInputSourceEventInit with all fields set.
func (p XRInputSourceEventInit) FromRef(ref js.Ref) XRInputSourceEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRInputSourceEventInit in the application heap.
func (p XRInputSourceEventInit) New() js.Ref {
	return bindings.XRInputSourceEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRInputSourceEventInit) UpdateFrom(ref js.Ref) {
	bindings.XRInputSourceEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRInputSourceEventInit) Update(ref js.Ref) {
	bindings.XRInputSourceEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRInputSourceEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Frame.Ref(),
		p.InputSource.Ref(),
	)
	p.Frame = p.Frame.FromRef(js.Undefined)
	p.InputSource = p.InputSource.FromRef(js.Undefined)
}

func NewXRInputSourceEvent(typ js.String, eventInitDict XRInputSourceEventInit) (ret XRInputSourceEvent) {
	ret.ref = bindings.NewXRInputSourceEventByXRInputSourceEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type XRInputSourceEvent struct {
	Event
}

func (this XRInputSourceEvent) Once() XRInputSourceEvent {
	this.ref.Once()
	return this
}

func (this XRInputSourceEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this XRInputSourceEvent) FromRef(ref js.Ref) XRInputSourceEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this XRInputSourceEvent) Free() {
	this.ref.Free()
}

// Frame returns the value of property "XRInputSourceEvent.frame".
//
// It returns ok=false if there is no such property.
func (this XRInputSourceEvent) Frame() (ret XRFrame, ok bool) {
	ok = js.True == bindings.GetXRInputSourceEventFrame(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InputSource returns the value of property "XRInputSourceEvent.inputSource".
//
// It returns ok=false if there is no such property.
func (this XRInputSourceEvent) InputSource() (ret XRInputSource, ok bool) {
	ok = js.True == bindings.GetXRInputSourceEventInputSource(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRInputSourcesChangeEventInit struct {
	// Session is "XRInputSourcesChangeEventInit.session"
	//
	// Required
	Session XRSession
	// Added is "XRInputSourcesChangeEventInit.added"
	//
	// Required
	Added js.FrozenArray[XRInputSource]
	// Removed is "XRInputSourcesChangeEventInit.removed"
	//
	// Required
	Removed js.FrozenArray[XRInputSource]
	// Bubbles is "XRInputSourcesChangeEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "XRInputSourcesChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "XRInputSourcesChangeEventInit.composed"
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

// FromRef calls UpdateFrom and returns a XRInputSourcesChangeEventInit with all fields set.
func (p XRInputSourcesChangeEventInit) FromRef(ref js.Ref) XRInputSourcesChangeEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRInputSourcesChangeEventInit in the application heap.
func (p XRInputSourcesChangeEventInit) New() js.Ref {
	return bindings.XRInputSourcesChangeEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRInputSourcesChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.XRInputSourcesChangeEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRInputSourcesChangeEventInit) Update(ref js.Ref) {
	bindings.XRInputSourcesChangeEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRInputSourcesChangeEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Session.Ref(),
		p.Added.Ref(),
		p.Removed.Ref(),
	)
	p.Session = p.Session.FromRef(js.Undefined)
	p.Added = p.Added.FromRef(js.Undefined)
	p.Removed = p.Removed.FromRef(js.Undefined)
}

func NewXRInputSourcesChangeEvent(typ js.String, eventInitDict XRInputSourcesChangeEventInit) (ret XRInputSourcesChangeEvent) {
	ret.ref = bindings.NewXRInputSourcesChangeEventByXRInputSourcesChangeEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type XRInputSourcesChangeEvent struct {
	Event
}

func (this XRInputSourcesChangeEvent) Once() XRInputSourcesChangeEvent {
	this.ref.Once()
	return this
}

func (this XRInputSourcesChangeEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this XRInputSourcesChangeEvent) FromRef(ref js.Ref) XRInputSourcesChangeEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this XRInputSourcesChangeEvent) Free() {
	this.ref.Free()
}

// Session returns the value of property "XRInputSourcesChangeEvent.session".
//
// It returns ok=false if there is no such property.
func (this XRInputSourcesChangeEvent) Session() (ret XRSession, ok bool) {
	ok = js.True == bindings.GetXRInputSourcesChangeEventSession(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Added returns the value of property "XRInputSourcesChangeEvent.added".
//
// It returns ok=false if there is no such property.
func (this XRInputSourcesChangeEvent) Added() (ret js.FrozenArray[XRInputSource], ok bool) {
	ok = js.True == bindings.GetXRInputSourcesChangeEventAdded(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Removed returns the value of property "XRInputSourcesChangeEvent.removed".
//
// It returns ok=false if there is no such property.
func (this XRInputSourcesChangeEvent) Removed() (ret js.FrozenArray[XRInputSource], ok bool) {
	ok = js.True == bindings.GetXRInputSourcesChangeEventRemoved(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRLayerEventInit struct {
	// Layer is "XRLayerEventInit.layer"
	//
	// Required
	Layer XRLayer
	// Bubbles is "XRLayerEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "XRLayerEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "XRLayerEventInit.composed"
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

// FromRef calls UpdateFrom and returns a XRLayerEventInit with all fields set.
func (p XRLayerEventInit) FromRef(ref js.Ref) XRLayerEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRLayerEventInit in the application heap.
func (p XRLayerEventInit) New() js.Ref {
	return bindings.XRLayerEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRLayerEventInit) UpdateFrom(ref js.Ref) {
	bindings.XRLayerEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRLayerEventInit) Update(ref js.Ref) {
	bindings.XRLayerEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRLayerEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Layer.Ref(),
	)
	p.Layer = p.Layer.FromRef(js.Undefined)
}

func NewXRLayerEvent(typ js.String, eventInitDict XRLayerEventInit) (ret XRLayerEvent) {
	ret.ref = bindings.NewXRLayerEventByXRLayerEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type XRLayerEvent struct {
	Event
}

func (this XRLayerEvent) Once() XRLayerEvent {
	this.ref.Once()
	return this
}

func (this XRLayerEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this XRLayerEvent) FromRef(ref js.Ref) XRLayerEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this XRLayerEvent) Free() {
	this.ref.Free()
}

// Layer returns the value of property "XRLayerEvent.layer".
//
// It returns ok=false if there is no such property.
func (this XRLayerEvent) Layer() (ret XRLayer, ok bool) {
	ok = js.True == bindings.GetXRLayerEventLayer(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRLayerInit struct {
	// Space is "XRLayerInit.space"
	//
	// Required
	Space XRSpace
	// ColorFormat is "XRLayerInit.colorFormat"
	//
	// Optional, defaults to 0x1908.
	//
	// NOTE: FFI_USE_ColorFormat MUST be set to true to make this field effective.
	ColorFormat GLenum
	// DepthFormat is "XRLayerInit.depthFormat"
	//
	// Optional
	//
	// NOTE: FFI_USE_DepthFormat MUST be set to true to make this field effective.
	DepthFormat GLenum
	// MipLevels is "XRLayerInit.mipLevels"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_MipLevels MUST be set to true to make this field effective.
	MipLevels uint32
	// ViewPixelWidth is "XRLayerInit.viewPixelWidth"
	//
	// Required
	ViewPixelWidth uint32
	// ViewPixelHeight is "XRLayerInit.viewPixelHeight"
	//
	// Required
	ViewPixelHeight uint32
	// Layout is "XRLayerInit.layout"
	//
	// Optional, defaults to "mono".
	Layout XRLayerLayout
	// IsStatic is "XRLayerInit.isStatic"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IsStatic MUST be set to true to make this field effective.
	IsStatic bool
	// ClearOnAccess is "XRLayerInit.clearOnAccess"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_ClearOnAccess MUST be set to true to make this field effective.
	ClearOnAccess bool

	FFI_USE_ColorFormat   bool // for ColorFormat.
	FFI_USE_DepthFormat   bool // for DepthFormat.
	FFI_USE_MipLevels     bool // for MipLevels.
	FFI_USE_IsStatic      bool // for IsStatic.
	FFI_USE_ClearOnAccess bool // for ClearOnAccess.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRLayerInit with all fields set.
func (p XRLayerInit) FromRef(ref js.Ref) XRLayerInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRLayerInit in the application heap.
func (p XRLayerInit) New() js.Ref {
	return bindings.XRLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRLayerInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRLayerInit) Update(ref js.Ref) {
	bindings.XRLayerInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRLayerInit) FreeMembers(recursive bool) {
	js.Free(
		p.Space.Ref(),
	)
	p.Space = p.Space.FromRef(js.Undefined)
}

type XRQuadLayer struct {
	XRCompositionLayer
}

func (this XRQuadLayer) Once() XRQuadLayer {
	this.ref.Once()
	return this
}

func (this XRQuadLayer) Ref() js.Ref {
	return this.XRCompositionLayer.Ref()
}

func (this XRQuadLayer) FromRef(ref js.Ref) XRQuadLayer {
	this.XRCompositionLayer = this.XRCompositionLayer.FromRef(ref)
	return this
}

func (this XRQuadLayer) Free() {
	this.ref.Free()
}

// Space returns the value of property "XRQuadLayer.space".
//
// It returns ok=false if there is no such property.
func (this XRQuadLayer) Space() (ret XRSpace, ok bool) {
	ok = js.True == bindings.GetXRQuadLayerSpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSpace sets the value of property "XRQuadLayer.space" to val.
//
// It returns false if the property cannot be set.
func (this XRQuadLayer) SetSpace(val XRSpace) bool {
	return js.True == bindings.SetXRQuadLayerSpace(
		this.ref,
		val.Ref(),
	)
}

// Transform returns the value of property "XRQuadLayer.transform".
//
// It returns ok=false if there is no such property.
func (this XRQuadLayer) Transform() (ret XRRigidTransform, ok bool) {
	ok = js.True == bindings.GetXRQuadLayerTransform(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTransform sets the value of property "XRQuadLayer.transform" to val.
//
// It returns false if the property cannot be set.
func (this XRQuadLayer) SetTransform(val XRRigidTransform) bool {
	return js.True == bindings.SetXRQuadLayerTransform(
		this.ref,
		val.Ref(),
	)
}

// Width returns the value of property "XRQuadLayer.width".
//
// It returns ok=false if there is no such property.
func (this XRQuadLayer) Width() (ret float32, ok bool) {
	ok = js.True == bindings.GetXRQuadLayerWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "XRQuadLayer.width" to val.
//
// It returns false if the property cannot be set.
func (this XRQuadLayer) SetWidth(val float32) bool {
	return js.True == bindings.SetXRQuadLayerWidth(
		this.ref,
		float32(val),
	)
}

// Height returns the value of property "XRQuadLayer.height".
//
// It returns ok=false if there is no such property.
func (this XRQuadLayer) Height() (ret float32, ok bool) {
	ok = js.True == bindings.GetXRQuadLayerHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "XRQuadLayer.height" to val.
//
// It returns false if the property cannot be set.
func (this XRQuadLayer) SetHeight(val float32) bool {
	return js.True == bindings.SetXRQuadLayerHeight(
		this.ref,
		float32(val),
	)
}

type XRMediaQuadLayerInit struct {
	// Transform is "XRMediaQuadLayerInit.transform"
	//
	// Optional
	Transform XRRigidTransform
	// Width is "XRMediaQuadLayerInit.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width float32
	// Height is "XRMediaQuadLayerInit.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height float32
	// Space is "XRMediaQuadLayerInit.space"
	//
	// Required
	Space XRSpace
	// Layout is "XRMediaQuadLayerInit.layout"
	//
	// Optional, defaults to "mono".
	Layout XRLayerLayout
	// InvertStereo is "XRMediaQuadLayerInit.invertStereo"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_InvertStereo MUST be set to true to make this field effective.
	InvertStereo bool

	FFI_USE_Width        bool // for Width.
	FFI_USE_Height       bool // for Height.
	FFI_USE_InvertStereo bool // for InvertStereo.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRMediaQuadLayerInit with all fields set.
func (p XRMediaQuadLayerInit) FromRef(ref js.Ref) XRMediaQuadLayerInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRMediaQuadLayerInit in the application heap.
func (p XRMediaQuadLayerInit) New() js.Ref {
	return bindings.XRMediaQuadLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRMediaQuadLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRMediaQuadLayerInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRMediaQuadLayerInit) Update(ref js.Ref) {
	bindings.XRMediaQuadLayerInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRMediaQuadLayerInit) FreeMembers(recursive bool) {
	js.Free(
		p.Transform.Ref(),
		p.Space.Ref(),
	)
	p.Transform = p.Transform.FromRef(js.Undefined)
	p.Space = p.Space.FromRef(js.Undefined)
}

type XRMediaCylinderLayerInit struct {
	// Transform is "XRMediaCylinderLayerInit.transform"
	//
	// Optional
	Transform XRRigidTransform
	// Radius is "XRMediaCylinderLayerInit.radius"
	//
	// Optional, defaults to 2.0.
	//
	// NOTE: FFI_USE_Radius MUST be set to true to make this field effective.
	Radius float32
	// CentralAngle is "XRMediaCylinderLayerInit.centralAngle"
	//
	// Optional, defaults to 0.78539.
	//
	// NOTE: FFI_USE_CentralAngle MUST be set to true to make this field effective.
	CentralAngle float32
	// AspectRatio is "XRMediaCylinderLayerInit.aspectRatio"
	//
	// Optional
	//
	// NOTE: FFI_USE_AspectRatio MUST be set to true to make this field effective.
	AspectRatio float32
	// Space is "XRMediaCylinderLayerInit.space"
	//
	// Required
	Space XRSpace
	// Layout is "XRMediaCylinderLayerInit.layout"
	//
	// Optional, defaults to "mono".
	Layout XRLayerLayout
	// InvertStereo is "XRMediaCylinderLayerInit.invertStereo"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_InvertStereo MUST be set to true to make this field effective.
	InvertStereo bool

	FFI_USE_Radius       bool // for Radius.
	FFI_USE_CentralAngle bool // for CentralAngle.
	FFI_USE_AspectRatio  bool // for AspectRatio.
	FFI_USE_InvertStereo bool // for InvertStereo.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRMediaCylinderLayerInit with all fields set.
func (p XRMediaCylinderLayerInit) FromRef(ref js.Ref) XRMediaCylinderLayerInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRMediaCylinderLayerInit in the application heap.
func (p XRMediaCylinderLayerInit) New() js.Ref {
	return bindings.XRMediaCylinderLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRMediaCylinderLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRMediaCylinderLayerInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRMediaCylinderLayerInit) Update(ref js.Ref) {
	bindings.XRMediaCylinderLayerInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRMediaCylinderLayerInit) FreeMembers(recursive bool) {
	js.Free(
		p.Transform.Ref(),
		p.Space.Ref(),
	)
	p.Transform = p.Transform.FromRef(js.Undefined)
	p.Space = p.Space.FromRef(js.Undefined)
}

type XRMediaEquirectLayerInit struct {
	// Transform is "XRMediaEquirectLayerInit.transform"
	//
	// Optional
	Transform XRRigidTransform
	// Radius is "XRMediaEquirectLayerInit.radius"
	//
	// Optional, defaults to 0.0.
	//
	// NOTE: FFI_USE_Radius MUST be set to true to make this field effective.
	Radius float32
	// CentralHorizontalAngle is "XRMediaEquirectLayerInit.centralHorizontalAngle"
	//
	// Optional, defaults to 6.28318.
	//
	// NOTE: FFI_USE_CentralHorizontalAngle MUST be set to true to make this field effective.
	CentralHorizontalAngle float32
	// UpperVerticalAngle is "XRMediaEquirectLayerInit.upperVerticalAngle"
	//
	// Optional, defaults to 1.570795.
	//
	// NOTE: FFI_USE_UpperVerticalAngle MUST be set to true to make this field effective.
	UpperVerticalAngle float32
	// LowerVerticalAngle is "XRMediaEquirectLayerInit.lowerVerticalAngle"
	//
	// Optional, defaults to -1.570795.
	//
	// NOTE: FFI_USE_LowerVerticalAngle MUST be set to true to make this field effective.
	LowerVerticalAngle float32
	// Space is "XRMediaEquirectLayerInit.space"
	//
	// Required
	Space XRSpace
	// Layout is "XRMediaEquirectLayerInit.layout"
	//
	// Optional, defaults to "mono".
	Layout XRLayerLayout
	// InvertStereo is "XRMediaEquirectLayerInit.invertStereo"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_InvertStereo MUST be set to true to make this field effective.
	InvertStereo bool

	FFI_USE_Radius                 bool // for Radius.
	FFI_USE_CentralHorizontalAngle bool // for CentralHorizontalAngle.
	FFI_USE_UpperVerticalAngle     bool // for UpperVerticalAngle.
	FFI_USE_LowerVerticalAngle     bool // for LowerVerticalAngle.
	FFI_USE_InvertStereo           bool // for InvertStereo.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRMediaEquirectLayerInit with all fields set.
func (p XRMediaEquirectLayerInit) FromRef(ref js.Ref) XRMediaEquirectLayerInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRMediaEquirectLayerInit in the application heap.
func (p XRMediaEquirectLayerInit) New() js.Ref {
	return bindings.XRMediaEquirectLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRMediaEquirectLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRMediaEquirectLayerInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRMediaEquirectLayerInit) Update(ref js.Ref) {
	bindings.XRMediaEquirectLayerInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRMediaEquirectLayerInit) FreeMembers(recursive bool) {
	js.Free(
		p.Transform.Ref(),
		p.Space.Ref(),
	)
	p.Transform = p.Transform.FromRef(js.Undefined)
	p.Space = p.Space.FromRef(js.Undefined)
}

func NewXRMediaBinding(session XRSession) (ret XRMediaBinding) {
	ret.ref = bindings.NewXRMediaBindingByXRMediaBinding(
		session.Ref())
	return
}

type XRMediaBinding struct {
	ref js.Ref
}

func (this XRMediaBinding) Once() XRMediaBinding {
	this.ref.Once()
	return this
}

func (this XRMediaBinding) Ref() js.Ref {
	return this.ref
}

func (this XRMediaBinding) FromRef(ref js.Ref) XRMediaBinding {
	this.ref = ref
	return this
}

func (this XRMediaBinding) Free() {
	this.ref.Free()
}

// HasFuncCreateQuadLayer returns true if the method "XRMediaBinding.createQuadLayer" exists.
func (this XRMediaBinding) HasFuncCreateQuadLayer() bool {
	return js.True == bindings.HasFuncXRMediaBindingCreateQuadLayer(
		this.ref,
	)
}

// FuncCreateQuadLayer returns the method "XRMediaBinding.createQuadLayer".
func (this XRMediaBinding) FuncCreateQuadLayer() (fn js.Func[func(video HTMLVideoElement, init XRMediaQuadLayerInit) XRQuadLayer]) {
	bindings.FuncXRMediaBindingCreateQuadLayer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateQuadLayer calls the method "XRMediaBinding.createQuadLayer".
func (this XRMediaBinding) CreateQuadLayer(video HTMLVideoElement, init XRMediaQuadLayerInit) (ret XRQuadLayer) {
	bindings.CallXRMediaBindingCreateQuadLayer(
		this.ref, js.Pointer(&ret),
		video.Ref(),
		js.Pointer(&init),
	)

	return
}

// TryCreateQuadLayer calls the method "XRMediaBinding.createQuadLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRMediaBinding) TryCreateQuadLayer(video HTMLVideoElement, init XRMediaQuadLayerInit) (ret XRQuadLayer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRMediaBindingCreateQuadLayer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		video.Ref(),
		js.Pointer(&init),
	)

	return
}

// HasFuncCreateQuadLayer1 returns true if the method "XRMediaBinding.createQuadLayer" exists.
func (this XRMediaBinding) HasFuncCreateQuadLayer1() bool {
	return js.True == bindings.HasFuncXRMediaBindingCreateQuadLayer1(
		this.ref,
	)
}

// FuncCreateQuadLayer1 returns the method "XRMediaBinding.createQuadLayer".
func (this XRMediaBinding) FuncCreateQuadLayer1() (fn js.Func[func(video HTMLVideoElement) XRQuadLayer]) {
	bindings.FuncXRMediaBindingCreateQuadLayer1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateQuadLayer1 calls the method "XRMediaBinding.createQuadLayer".
func (this XRMediaBinding) CreateQuadLayer1(video HTMLVideoElement) (ret XRQuadLayer) {
	bindings.CallXRMediaBindingCreateQuadLayer1(
		this.ref, js.Pointer(&ret),
		video.Ref(),
	)

	return
}

// TryCreateQuadLayer1 calls the method "XRMediaBinding.createQuadLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRMediaBinding) TryCreateQuadLayer1(video HTMLVideoElement) (ret XRQuadLayer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRMediaBindingCreateQuadLayer1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		video.Ref(),
	)

	return
}

// HasFuncCreateCylinderLayer returns true if the method "XRMediaBinding.createCylinderLayer" exists.
func (this XRMediaBinding) HasFuncCreateCylinderLayer() bool {
	return js.True == bindings.HasFuncXRMediaBindingCreateCylinderLayer(
		this.ref,
	)
}

// FuncCreateCylinderLayer returns the method "XRMediaBinding.createCylinderLayer".
func (this XRMediaBinding) FuncCreateCylinderLayer() (fn js.Func[func(video HTMLVideoElement, init XRMediaCylinderLayerInit) XRCylinderLayer]) {
	bindings.FuncXRMediaBindingCreateCylinderLayer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateCylinderLayer calls the method "XRMediaBinding.createCylinderLayer".
func (this XRMediaBinding) CreateCylinderLayer(video HTMLVideoElement, init XRMediaCylinderLayerInit) (ret XRCylinderLayer) {
	bindings.CallXRMediaBindingCreateCylinderLayer(
		this.ref, js.Pointer(&ret),
		video.Ref(),
		js.Pointer(&init),
	)

	return
}

// TryCreateCylinderLayer calls the method "XRMediaBinding.createCylinderLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRMediaBinding) TryCreateCylinderLayer(video HTMLVideoElement, init XRMediaCylinderLayerInit) (ret XRCylinderLayer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRMediaBindingCreateCylinderLayer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		video.Ref(),
		js.Pointer(&init),
	)

	return
}

// HasFuncCreateCylinderLayer1 returns true if the method "XRMediaBinding.createCylinderLayer" exists.
func (this XRMediaBinding) HasFuncCreateCylinderLayer1() bool {
	return js.True == bindings.HasFuncXRMediaBindingCreateCylinderLayer1(
		this.ref,
	)
}

// FuncCreateCylinderLayer1 returns the method "XRMediaBinding.createCylinderLayer".
func (this XRMediaBinding) FuncCreateCylinderLayer1() (fn js.Func[func(video HTMLVideoElement) XRCylinderLayer]) {
	bindings.FuncXRMediaBindingCreateCylinderLayer1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateCylinderLayer1 calls the method "XRMediaBinding.createCylinderLayer".
func (this XRMediaBinding) CreateCylinderLayer1(video HTMLVideoElement) (ret XRCylinderLayer) {
	bindings.CallXRMediaBindingCreateCylinderLayer1(
		this.ref, js.Pointer(&ret),
		video.Ref(),
	)

	return
}

// TryCreateCylinderLayer1 calls the method "XRMediaBinding.createCylinderLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRMediaBinding) TryCreateCylinderLayer1(video HTMLVideoElement) (ret XRCylinderLayer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRMediaBindingCreateCylinderLayer1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		video.Ref(),
	)

	return
}

// HasFuncCreateEquirectLayer returns true if the method "XRMediaBinding.createEquirectLayer" exists.
func (this XRMediaBinding) HasFuncCreateEquirectLayer() bool {
	return js.True == bindings.HasFuncXRMediaBindingCreateEquirectLayer(
		this.ref,
	)
}

// FuncCreateEquirectLayer returns the method "XRMediaBinding.createEquirectLayer".
func (this XRMediaBinding) FuncCreateEquirectLayer() (fn js.Func[func(video HTMLVideoElement, init XRMediaEquirectLayerInit) XREquirectLayer]) {
	bindings.FuncXRMediaBindingCreateEquirectLayer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateEquirectLayer calls the method "XRMediaBinding.createEquirectLayer".
func (this XRMediaBinding) CreateEquirectLayer(video HTMLVideoElement, init XRMediaEquirectLayerInit) (ret XREquirectLayer) {
	bindings.CallXRMediaBindingCreateEquirectLayer(
		this.ref, js.Pointer(&ret),
		video.Ref(),
		js.Pointer(&init),
	)

	return
}

// TryCreateEquirectLayer calls the method "XRMediaBinding.createEquirectLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRMediaBinding) TryCreateEquirectLayer(video HTMLVideoElement, init XRMediaEquirectLayerInit) (ret XREquirectLayer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRMediaBindingCreateEquirectLayer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		video.Ref(),
		js.Pointer(&init),
	)

	return
}

// HasFuncCreateEquirectLayer1 returns true if the method "XRMediaBinding.createEquirectLayer" exists.
func (this XRMediaBinding) HasFuncCreateEquirectLayer1() bool {
	return js.True == bindings.HasFuncXRMediaBindingCreateEquirectLayer1(
		this.ref,
	)
}

// FuncCreateEquirectLayer1 returns the method "XRMediaBinding.createEquirectLayer".
func (this XRMediaBinding) FuncCreateEquirectLayer1() (fn js.Func[func(video HTMLVideoElement) XREquirectLayer]) {
	bindings.FuncXRMediaBindingCreateEquirectLayer1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateEquirectLayer1 calls the method "XRMediaBinding.createEquirectLayer".
func (this XRMediaBinding) CreateEquirectLayer1(video HTMLVideoElement) (ret XREquirectLayer) {
	bindings.CallXRMediaBindingCreateEquirectLayer1(
		this.ref, js.Pointer(&ret),
		video.Ref(),
	)

	return
}

// TryCreateEquirectLayer1 calls the method "XRMediaBinding.createEquirectLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRMediaBinding) TryCreateEquirectLayer1(video HTMLVideoElement) (ret XREquirectLayer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRMediaBindingCreateEquirectLayer1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		video.Ref(),
	)

	return
}

type XRMediaLayerInit struct {
	// Space is "XRMediaLayerInit.space"
	//
	// Required
	Space XRSpace
	// Layout is "XRMediaLayerInit.layout"
	//
	// Optional, defaults to "mono".
	Layout XRLayerLayout
	// InvertStereo is "XRMediaLayerInit.invertStereo"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_InvertStereo MUST be set to true to make this field effective.
	InvertStereo bool

	FFI_USE_InvertStereo bool // for InvertStereo.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRMediaLayerInit with all fields set.
func (p XRMediaLayerInit) FromRef(ref js.Ref) XRMediaLayerInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRMediaLayerInit in the application heap.
func (p XRMediaLayerInit) New() js.Ref {
	return bindings.XRMediaLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRMediaLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRMediaLayerInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRMediaLayerInit) Update(ref js.Ref) {
	bindings.XRMediaLayerInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRMediaLayerInit) FreeMembers(recursive bool) {
	js.Free(
		p.Space.Ref(),
	)
	p.Space = p.Space.FromRef(js.Undefined)
}

type XRMesh struct {
	ref js.Ref
}

func (this XRMesh) Once() XRMesh {
	this.ref.Once()
	return this
}

func (this XRMesh) Ref() js.Ref {
	return this.ref
}

func (this XRMesh) FromRef(ref js.Ref) XRMesh {
	this.ref = ref
	return this
}

func (this XRMesh) Free() {
	this.ref.Free()
}

// MeshSpace returns the value of property "XRMesh.meshSpace".
//
// It returns ok=false if there is no such property.
func (this XRMesh) MeshSpace() (ret XRSpace, ok bool) {
	ok = js.True == bindings.GetXRMeshMeshSpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Vertices returns the value of property "XRMesh.vertices".
//
// It returns ok=false if there is no such property.
func (this XRMesh) Vertices() (ret js.FrozenArray[js.TypedArray[float32]], ok bool) {
	ok = js.True == bindings.GetXRMeshVertices(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Indices returns the value of property "XRMesh.indices".
//
// It returns ok=false if there is no such property.
func (this XRMesh) Indices() (ret js.TypedArray[uint32], ok bool) {
	ok = js.True == bindings.GetXRMeshIndices(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LastChangedTime returns the value of property "XRMesh.lastChangedTime".
//
// It returns ok=false if there is no such property.
func (this XRMesh) LastChangedTime() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetXRMeshLastChangedTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SemanticLabel returns the value of property "XRMesh.semanticLabel".
//
// It returns ok=false if there is no such property.
func (this XRMesh) SemanticLabel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetXRMeshSemanticLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRPermissionDescriptor struct {
	// Mode is "XRPermissionDescriptor.mode"
	//
	// Optional
	Mode XRSessionMode
	// RequiredFeatures is "XRPermissionDescriptor.requiredFeatures"
	//
	// Optional
	RequiredFeatures js.Array[js.String]
	// OptionalFeatures is "XRPermissionDescriptor.optionalFeatures"
	//
	// Optional
	OptionalFeatures js.Array[js.String]
	// Name is "XRPermissionDescriptor.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRPermissionDescriptor with all fields set.
func (p XRPermissionDescriptor) FromRef(ref js.Ref) XRPermissionDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRPermissionDescriptor in the application heap.
func (p XRPermissionDescriptor) New() js.Ref {
	return bindings.XRPermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.XRPermissionDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRPermissionDescriptor) Update(ref js.Ref) {
	bindings.XRPermissionDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRPermissionDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.RequiredFeatures.Ref(),
		p.OptionalFeatures.Ref(),
		p.Name.Ref(),
	)
	p.RequiredFeatures = p.RequiredFeatures.FromRef(js.Undefined)
	p.OptionalFeatures = p.OptionalFeatures.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

type XRPermissionStatus struct {
	PermissionStatus
}

func (this XRPermissionStatus) Once() XRPermissionStatus {
	this.ref.Once()
	return this
}

func (this XRPermissionStatus) Ref() js.Ref {
	return this.PermissionStatus.Ref()
}

func (this XRPermissionStatus) FromRef(ref js.Ref) XRPermissionStatus {
	this.PermissionStatus = this.PermissionStatus.FromRef(ref)
	return this
}

func (this XRPermissionStatus) Free() {
	this.ref.Free()
}

// Granted returns the value of property "XRPermissionStatus.granted".
//
// It returns ok=false if there is no such property.
func (this XRPermissionStatus) Granted() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetXRPermissionStatusGranted(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetGranted sets the value of property "XRPermissionStatus.granted" to val.
//
// It returns false if the property cannot be set.
func (this XRPermissionStatus) SetGranted(val js.FrozenArray[js.String]) bool {
	return js.True == bindings.SetXRPermissionStatusGranted(
		this.ref,
		val.Ref(),
	)
}

type XRProjectionLayer struct {
	XRCompositionLayer
}

func (this XRProjectionLayer) Once() XRProjectionLayer {
	this.ref.Once()
	return this
}

func (this XRProjectionLayer) Ref() js.Ref {
	return this.XRCompositionLayer.Ref()
}

func (this XRProjectionLayer) FromRef(ref js.Ref) XRProjectionLayer {
	this.XRCompositionLayer = this.XRCompositionLayer.FromRef(ref)
	return this
}

func (this XRProjectionLayer) Free() {
	this.ref.Free()
}

// TextureWidth returns the value of property "XRProjectionLayer.textureWidth".
//
// It returns ok=false if there is no such property.
func (this XRProjectionLayer) TextureWidth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRProjectionLayerTextureWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TextureHeight returns the value of property "XRProjectionLayer.textureHeight".
//
// It returns ok=false if there is no such property.
func (this XRProjectionLayer) TextureHeight() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRProjectionLayerTextureHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TextureArrayLength returns the value of property "XRProjectionLayer.textureArrayLength".
//
// It returns ok=false if there is no such property.
func (this XRProjectionLayer) TextureArrayLength() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRProjectionLayerTextureArrayLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IgnoreDepthValues returns the value of property "XRProjectionLayer.ignoreDepthValues".
//
// It returns ok=false if there is no such property.
func (this XRProjectionLayer) IgnoreDepthValues() (ret bool, ok bool) {
	ok = js.True == bindings.GetXRProjectionLayerIgnoreDepthValues(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FixedFoveation returns the value of property "XRProjectionLayer.fixedFoveation".
//
// It returns ok=false if there is no such property.
func (this XRProjectionLayer) FixedFoveation() (ret float32, ok bool) {
	ok = js.True == bindings.GetXRProjectionLayerFixedFoveation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFixedFoveation sets the value of property "XRProjectionLayer.fixedFoveation" to val.
//
// It returns false if the property cannot be set.
func (this XRProjectionLayer) SetFixedFoveation(val float32) bool {
	return js.True == bindings.SetXRProjectionLayerFixedFoveation(
		this.ref,
		float32(val),
	)
}

// DeltaPose returns the value of property "XRProjectionLayer.deltaPose".
//
// It returns ok=false if there is no such property.
func (this XRProjectionLayer) DeltaPose() (ret XRRigidTransform, ok bool) {
	ok = js.True == bindings.GetXRProjectionLayerDeltaPose(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDeltaPose sets the value of property "XRProjectionLayer.deltaPose" to val.
//
// It returns false if the property cannot be set.
func (this XRProjectionLayer) SetDeltaPose(val XRRigidTransform) bool {
	return js.True == bindings.SetXRProjectionLayerDeltaPose(
		this.ref,
		val.Ref(),
	)
}

type XRProjectionLayerInit struct {
	// TextureType is "XRProjectionLayerInit.textureType"
	//
	// Optional, defaults to "texture".
	TextureType XRTextureType
	// ColorFormat is "XRProjectionLayerInit.colorFormat"
	//
	// Optional, defaults to 0x1908.
	//
	// NOTE: FFI_USE_ColorFormat MUST be set to true to make this field effective.
	ColorFormat GLenum
	// DepthFormat is "XRProjectionLayerInit.depthFormat"
	//
	// Optional, defaults to 0x1902.
	//
	// NOTE: FFI_USE_DepthFormat MUST be set to true to make this field effective.
	DepthFormat GLenum
	// ScaleFactor is "XRProjectionLayerInit.scaleFactor"
	//
	// Optional, defaults to 1.0.
	//
	// NOTE: FFI_USE_ScaleFactor MUST be set to true to make this field effective.
	ScaleFactor float64
	// ClearOnAccess is "XRProjectionLayerInit.clearOnAccess"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_ClearOnAccess MUST be set to true to make this field effective.
	ClearOnAccess bool

	FFI_USE_ColorFormat   bool // for ColorFormat.
	FFI_USE_DepthFormat   bool // for DepthFormat.
	FFI_USE_ScaleFactor   bool // for ScaleFactor.
	FFI_USE_ClearOnAccess bool // for ClearOnAccess.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRProjectionLayerInit with all fields set.
func (p XRProjectionLayerInit) FromRef(ref js.Ref) XRProjectionLayerInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRProjectionLayerInit in the application heap.
func (p XRProjectionLayerInit) New() js.Ref {
	return bindings.XRProjectionLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRProjectionLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRProjectionLayerInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRProjectionLayerInit) Update(ref js.Ref) {
	bindings.XRProjectionLayerInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRProjectionLayerInit) FreeMembers(recursive bool) {
}

type XRQuadLayerInit struct {
	// TextureType is "XRQuadLayerInit.textureType"
	//
	// Optional, defaults to "texture".
	TextureType XRTextureType
	// Transform is "XRQuadLayerInit.transform"
	//
	// Optional
	Transform XRRigidTransform
	// Width is "XRQuadLayerInit.width"
	//
	// Optional, defaults to 1.0.
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width float32
	// Height is "XRQuadLayerInit.height"
	//
	// Optional, defaults to 1.0.
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height float32
	// Space is "XRQuadLayerInit.space"
	//
	// Required
	Space XRSpace
	// ColorFormat is "XRQuadLayerInit.colorFormat"
	//
	// Optional, defaults to 0x1908.
	//
	// NOTE: FFI_USE_ColorFormat MUST be set to true to make this field effective.
	ColorFormat GLenum
	// DepthFormat is "XRQuadLayerInit.depthFormat"
	//
	// Optional
	//
	// NOTE: FFI_USE_DepthFormat MUST be set to true to make this field effective.
	DepthFormat GLenum
	// MipLevels is "XRQuadLayerInit.mipLevels"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_MipLevels MUST be set to true to make this field effective.
	MipLevels uint32
	// ViewPixelWidth is "XRQuadLayerInit.viewPixelWidth"
	//
	// Required
	ViewPixelWidth uint32
	// ViewPixelHeight is "XRQuadLayerInit.viewPixelHeight"
	//
	// Required
	ViewPixelHeight uint32
	// Layout is "XRQuadLayerInit.layout"
	//
	// Optional, defaults to "mono".
	Layout XRLayerLayout
	// IsStatic is "XRQuadLayerInit.isStatic"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IsStatic MUST be set to true to make this field effective.
	IsStatic bool
	// ClearOnAccess is "XRQuadLayerInit.clearOnAccess"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_ClearOnAccess MUST be set to true to make this field effective.
	ClearOnAccess bool

	FFI_USE_Width         bool // for Width.
	FFI_USE_Height        bool // for Height.
	FFI_USE_ColorFormat   bool // for ColorFormat.
	FFI_USE_DepthFormat   bool // for DepthFormat.
	FFI_USE_MipLevels     bool // for MipLevels.
	FFI_USE_IsStatic      bool // for IsStatic.
	FFI_USE_ClearOnAccess bool // for ClearOnAccess.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRQuadLayerInit with all fields set.
func (p XRQuadLayerInit) FromRef(ref js.Ref) XRQuadLayerInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRQuadLayerInit in the application heap.
func (p XRQuadLayerInit) New() js.Ref {
	return bindings.XRQuadLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRQuadLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRQuadLayerInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRQuadLayerInit) Update(ref js.Ref) {
	bindings.XRQuadLayerInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRQuadLayerInit) FreeMembers(recursive bool) {
	js.Free(
		p.Transform.Ref(),
		p.Space.Ref(),
	)
	p.Transform = p.Transform.FromRef(js.Undefined)
	p.Space = p.Space.FromRef(js.Undefined)
}

type XRReferenceSpaceEventInit struct {
	// ReferenceSpace is "XRReferenceSpaceEventInit.referenceSpace"
	//
	// Required
	ReferenceSpace XRReferenceSpace
	// Transform is "XRReferenceSpaceEventInit.transform"
	//
	// Optional, defaults to null.
	Transform XRRigidTransform
	// Bubbles is "XRReferenceSpaceEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "XRReferenceSpaceEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "XRReferenceSpaceEventInit.composed"
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

// FromRef calls UpdateFrom and returns a XRReferenceSpaceEventInit with all fields set.
func (p XRReferenceSpaceEventInit) FromRef(ref js.Ref) XRReferenceSpaceEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRReferenceSpaceEventInit in the application heap.
func (p XRReferenceSpaceEventInit) New() js.Ref {
	return bindings.XRReferenceSpaceEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRReferenceSpaceEventInit) UpdateFrom(ref js.Ref) {
	bindings.XRReferenceSpaceEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRReferenceSpaceEventInit) Update(ref js.Ref) {
	bindings.XRReferenceSpaceEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRReferenceSpaceEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.ReferenceSpace.Ref(),
		p.Transform.Ref(),
	)
	p.ReferenceSpace = p.ReferenceSpace.FromRef(js.Undefined)
	p.Transform = p.Transform.FromRef(js.Undefined)
}

func NewXRReferenceSpaceEvent(typ js.String, eventInitDict XRReferenceSpaceEventInit) (ret XRReferenceSpaceEvent) {
	ret.ref = bindings.NewXRReferenceSpaceEventByXRReferenceSpaceEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type XRReferenceSpaceEvent struct {
	Event
}

func (this XRReferenceSpaceEvent) Once() XRReferenceSpaceEvent {
	this.ref.Once()
	return this
}

func (this XRReferenceSpaceEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this XRReferenceSpaceEvent) FromRef(ref js.Ref) XRReferenceSpaceEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this XRReferenceSpaceEvent) Free() {
	this.ref.Free()
}

// ReferenceSpace returns the value of property "XRReferenceSpaceEvent.referenceSpace".
//
// It returns ok=false if there is no such property.
func (this XRReferenceSpaceEvent) ReferenceSpace() (ret XRReferenceSpace, ok bool) {
	ok = js.True == bindings.GetXRReferenceSpaceEventReferenceSpace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Transform returns the value of property "XRReferenceSpaceEvent.transform".
//
// It returns ok=false if there is no such property.
func (this XRReferenceSpaceEvent) Transform() (ret XRRigidTransform, ok bool) {
	ok = js.True == bindings.GetXRReferenceSpaceEventTransform(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRSessionEventInit struct {
	// Session is "XRSessionEventInit.session"
	//
	// Required
	Session XRSession
	// Bubbles is "XRSessionEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "XRSessionEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "XRSessionEventInit.composed"
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

// FromRef calls UpdateFrom and returns a XRSessionEventInit with all fields set.
func (p XRSessionEventInit) FromRef(ref js.Ref) XRSessionEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRSessionEventInit in the application heap.
func (p XRSessionEventInit) New() js.Ref {
	return bindings.XRSessionEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRSessionEventInit) UpdateFrom(ref js.Ref) {
	bindings.XRSessionEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRSessionEventInit) Update(ref js.Ref) {
	bindings.XRSessionEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRSessionEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Session.Ref(),
	)
	p.Session = p.Session.FromRef(js.Undefined)
}

func NewXRSessionEvent(typ js.String, eventInitDict XRSessionEventInit) (ret XRSessionEvent) {
	ret.ref = bindings.NewXRSessionEventByXRSessionEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type XRSessionEvent struct {
	Event
}

func (this XRSessionEvent) Once() XRSessionEvent {
	this.ref.Once()
	return this
}

func (this XRSessionEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this XRSessionEvent) FromRef(ref js.Ref) XRSessionEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this XRSessionEvent) Free() {
	this.ref.Free()
}

// Session returns the value of property "XRSessionEvent.session".
//
// It returns ok=false if there is no such property.
func (this XRSessionEvent) Session() (ret XRSession, ok bool) {
	ok = js.True == bindings.GetXRSessionEventSession(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRSessionSupportedPermissionDescriptor struct {
	// Mode is "XRSessionSupportedPermissionDescriptor.mode"
	//
	// Optional
	Mode XRSessionMode
	// Name is "XRSessionSupportedPermissionDescriptor.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRSessionSupportedPermissionDescriptor with all fields set.
func (p XRSessionSupportedPermissionDescriptor) FromRef(ref js.Ref) XRSessionSupportedPermissionDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new XRSessionSupportedPermissionDescriptor in the application heap.
func (p XRSessionSupportedPermissionDescriptor) New() js.Ref {
	return bindings.XRSessionSupportedPermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *XRSessionSupportedPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.XRSessionSupportedPermissionDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *XRSessionSupportedPermissionDescriptor) Update(ref js.Ref) {
	bindings.XRSessionSupportedPermissionDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *XRSessionSupportedPermissionDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
}

type XRSubImage struct {
	ref js.Ref
}

func (this XRSubImage) Once() XRSubImage {
	this.ref.Once()
	return this
}

func (this XRSubImage) Ref() js.Ref {
	return this.ref
}

func (this XRSubImage) FromRef(ref js.Ref) XRSubImage {
	this.ref = ref
	return this
}

func (this XRSubImage) Free() {
	this.ref.Free()
}

// Viewport returns the value of property "XRSubImage.viewport".
//
// It returns ok=false if there is no such property.
func (this XRSubImage) Viewport() (ret XRViewport, ok bool) {
	ok = js.True == bindings.GetXRSubImageViewport(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRWebGLSubImage struct {
	XRSubImage
}

func (this XRWebGLSubImage) Once() XRWebGLSubImage {
	this.ref.Once()
	return this
}

func (this XRWebGLSubImage) Ref() js.Ref {
	return this.XRSubImage.Ref()
}

func (this XRWebGLSubImage) FromRef(ref js.Ref) XRWebGLSubImage {
	this.XRSubImage = this.XRSubImage.FromRef(ref)
	return this
}

func (this XRWebGLSubImage) Free() {
	this.ref.Free()
}

// ColorTexture returns the value of property "XRWebGLSubImage.colorTexture".
//
// It returns ok=false if there is no such property.
func (this XRWebGLSubImage) ColorTexture() (ret WebGLTexture, ok bool) {
	ok = js.True == bindings.GetXRWebGLSubImageColorTexture(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DepthStencilTexture returns the value of property "XRWebGLSubImage.depthStencilTexture".
//
// It returns ok=false if there is no such property.
func (this XRWebGLSubImage) DepthStencilTexture() (ret WebGLTexture, ok bool) {
	ok = js.True == bindings.GetXRWebGLSubImageDepthStencilTexture(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MotionVectorTexture returns the value of property "XRWebGLSubImage.motionVectorTexture".
//
// It returns ok=false if there is no such property.
func (this XRWebGLSubImage) MotionVectorTexture() (ret WebGLTexture, ok bool) {
	ok = js.True == bindings.GetXRWebGLSubImageMotionVectorTexture(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ImageIndex returns the value of property "XRWebGLSubImage.imageIndex".
//
// It returns ok=false if there is no such property.
func (this XRWebGLSubImage) ImageIndex() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRWebGLSubImageImageIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ColorTextureWidth returns the value of property "XRWebGLSubImage.colorTextureWidth".
//
// It returns ok=false if there is no such property.
func (this XRWebGLSubImage) ColorTextureWidth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRWebGLSubImageColorTextureWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ColorTextureHeight returns the value of property "XRWebGLSubImage.colorTextureHeight".
//
// It returns ok=false if there is no such property.
func (this XRWebGLSubImage) ColorTextureHeight() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRWebGLSubImageColorTextureHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DepthStencilTextureWidth returns the value of property "XRWebGLSubImage.depthStencilTextureWidth".
//
// It returns ok=false if there is no such property.
func (this XRWebGLSubImage) DepthStencilTextureWidth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRWebGLSubImageDepthStencilTextureWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DepthStencilTextureHeight returns the value of property "XRWebGLSubImage.depthStencilTextureHeight".
//
// It returns ok=false if there is no such property.
func (this XRWebGLSubImage) DepthStencilTextureHeight() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRWebGLSubImageDepthStencilTextureHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MotionVectorTextureWidth returns the value of property "XRWebGLSubImage.motionVectorTextureWidth".
//
// It returns ok=false if there is no such property.
func (this XRWebGLSubImage) MotionVectorTextureWidth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRWebGLSubImageMotionVectorTextureWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MotionVectorTextureHeight returns the value of property "XRWebGLSubImage.motionVectorTextureHeight".
//
// It returns ok=false if there is no such property.
func (this XRWebGLSubImage) MotionVectorTextureHeight() (ret uint32, ok bool) {
	ok = js.True == bindings.GetXRWebGLSubImageMotionVectorTextureHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

type XRWebGLDepthInformation struct {
	XRDepthInformation
}

func (this XRWebGLDepthInformation) Once() XRWebGLDepthInformation {
	this.ref.Once()
	return this
}

func (this XRWebGLDepthInformation) Ref() js.Ref {
	return this.XRDepthInformation.Ref()
}

func (this XRWebGLDepthInformation) FromRef(ref js.Ref) XRWebGLDepthInformation {
	this.XRDepthInformation = this.XRDepthInformation.FromRef(ref)
	return this
}

func (this XRWebGLDepthInformation) Free() {
	this.ref.Free()
}

// Texture returns the value of property "XRWebGLDepthInformation.texture".
//
// It returns ok=false if there is no such property.
func (this XRWebGLDepthInformation) Texture() (ret WebGLTexture, ok bool) {
	ok = js.True == bindings.GetXRWebGLDepthInformationTexture(
		this.ref, js.Pointer(&ret),
	)
	return
}

func NewXRWebGLBinding(session XRSession, context XRWebGLRenderingContext) (ret XRWebGLBinding) {
	ret.ref = bindings.NewXRWebGLBindingByXRWebGLBinding(
		session.Ref(),
		context.Ref())
	return
}

type XRWebGLBinding struct {
	ref js.Ref
}

func (this XRWebGLBinding) Once() XRWebGLBinding {
	this.ref.Once()
	return this
}

func (this XRWebGLBinding) Ref() js.Ref {
	return this.ref
}

func (this XRWebGLBinding) FromRef(ref js.Ref) XRWebGLBinding {
	this.ref = ref
	return this
}

func (this XRWebGLBinding) Free() {
	this.ref.Free()
}

// NativeProjectionScaleFactor returns the value of property "XRWebGLBinding.nativeProjectionScaleFactor".
//
// It returns ok=false if there is no such property.
func (this XRWebGLBinding) NativeProjectionScaleFactor() (ret float64, ok bool) {
	ok = js.True == bindings.GetXRWebGLBindingNativeProjectionScaleFactor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UsesDepthValues returns the value of property "XRWebGLBinding.usesDepthValues".
//
// It returns ok=false if there is no such property.
func (this XRWebGLBinding) UsesDepthValues() (ret bool, ok bool) {
	ok = js.True == bindings.GetXRWebGLBindingUsesDepthValues(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncCreateProjectionLayer returns true if the method "XRWebGLBinding.createProjectionLayer" exists.
func (this XRWebGLBinding) HasFuncCreateProjectionLayer() bool {
	return js.True == bindings.HasFuncXRWebGLBindingCreateProjectionLayer(
		this.ref,
	)
}

// FuncCreateProjectionLayer returns the method "XRWebGLBinding.createProjectionLayer".
func (this XRWebGLBinding) FuncCreateProjectionLayer() (fn js.Func[func(init XRProjectionLayerInit) XRProjectionLayer]) {
	bindings.FuncXRWebGLBindingCreateProjectionLayer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateProjectionLayer calls the method "XRWebGLBinding.createProjectionLayer".
func (this XRWebGLBinding) CreateProjectionLayer(init XRProjectionLayerInit) (ret XRProjectionLayer) {
	bindings.CallXRWebGLBindingCreateProjectionLayer(
		this.ref, js.Pointer(&ret),
		js.Pointer(&init),
	)

	return
}

// TryCreateProjectionLayer calls the method "XRWebGLBinding.createProjectionLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLBinding) TryCreateProjectionLayer(init XRProjectionLayerInit) (ret XRProjectionLayer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLBindingCreateProjectionLayer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&init),
	)

	return
}

// HasFuncCreateProjectionLayer1 returns true if the method "XRWebGLBinding.createProjectionLayer" exists.
func (this XRWebGLBinding) HasFuncCreateProjectionLayer1() bool {
	return js.True == bindings.HasFuncXRWebGLBindingCreateProjectionLayer1(
		this.ref,
	)
}

// FuncCreateProjectionLayer1 returns the method "XRWebGLBinding.createProjectionLayer".
func (this XRWebGLBinding) FuncCreateProjectionLayer1() (fn js.Func[func() XRProjectionLayer]) {
	bindings.FuncXRWebGLBindingCreateProjectionLayer1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateProjectionLayer1 calls the method "XRWebGLBinding.createProjectionLayer".
func (this XRWebGLBinding) CreateProjectionLayer1() (ret XRProjectionLayer) {
	bindings.CallXRWebGLBindingCreateProjectionLayer1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateProjectionLayer1 calls the method "XRWebGLBinding.createProjectionLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLBinding) TryCreateProjectionLayer1() (ret XRProjectionLayer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLBindingCreateProjectionLayer1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateQuadLayer returns true if the method "XRWebGLBinding.createQuadLayer" exists.
func (this XRWebGLBinding) HasFuncCreateQuadLayer() bool {
	return js.True == bindings.HasFuncXRWebGLBindingCreateQuadLayer(
		this.ref,
	)
}

// FuncCreateQuadLayer returns the method "XRWebGLBinding.createQuadLayer".
func (this XRWebGLBinding) FuncCreateQuadLayer() (fn js.Func[func(init XRQuadLayerInit) XRQuadLayer]) {
	bindings.FuncXRWebGLBindingCreateQuadLayer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateQuadLayer calls the method "XRWebGLBinding.createQuadLayer".
func (this XRWebGLBinding) CreateQuadLayer(init XRQuadLayerInit) (ret XRQuadLayer) {
	bindings.CallXRWebGLBindingCreateQuadLayer(
		this.ref, js.Pointer(&ret),
		js.Pointer(&init),
	)

	return
}

// TryCreateQuadLayer calls the method "XRWebGLBinding.createQuadLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLBinding) TryCreateQuadLayer(init XRQuadLayerInit) (ret XRQuadLayer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLBindingCreateQuadLayer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&init),
	)

	return
}

// HasFuncCreateQuadLayer1 returns true if the method "XRWebGLBinding.createQuadLayer" exists.
func (this XRWebGLBinding) HasFuncCreateQuadLayer1() bool {
	return js.True == bindings.HasFuncXRWebGLBindingCreateQuadLayer1(
		this.ref,
	)
}

// FuncCreateQuadLayer1 returns the method "XRWebGLBinding.createQuadLayer".
func (this XRWebGLBinding) FuncCreateQuadLayer1() (fn js.Func[func() XRQuadLayer]) {
	bindings.FuncXRWebGLBindingCreateQuadLayer1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateQuadLayer1 calls the method "XRWebGLBinding.createQuadLayer".
func (this XRWebGLBinding) CreateQuadLayer1() (ret XRQuadLayer) {
	bindings.CallXRWebGLBindingCreateQuadLayer1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateQuadLayer1 calls the method "XRWebGLBinding.createQuadLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLBinding) TryCreateQuadLayer1() (ret XRQuadLayer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLBindingCreateQuadLayer1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateCylinderLayer returns true if the method "XRWebGLBinding.createCylinderLayer" exists.
func (this XRWebGLBinding) HasFuncCreateCylinderLayer() bool {
	return js.True == bindings.HasFuncXRWebGLBindingCreateCylinderLayer(
		this.ref,
	)
}

// FuncCreateCylinderLayer returns the method "XRWebGLBinding.createCylinderLayer".
func (this XRWebGLBinding) FuncCreateCylinderLayer() (fn js.Func[func(init XRCylinderLayerInit) XRCylinderLayer]) {
	bindings.FuncXRWebGLBindingCreateCylinderLayer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateCylinderLayer calls the method "XRWebGLBinding.createCylinderLayer".
func (this XRWebGLBinding) CreateCylinderLayer(init XRCylinderLayerInit) (ret XRCylinderLayer) {
	bindings.CallXRWebGLBindingCreateCylinderLayer(
		this.ref, js.Pointer(&ret),
		js.Pointer(&init),
	)

	return
}

// TryCreateCylinderLayer calls the method "XRWebGLBinding.createCylinderLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLBinding) TryCreateCylinderLayer(init XRCylinderLayerInit) (ret XRCylinderLayer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLBindingCreateCylinderLayer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&init),
	)

	return
}

// HasFuncCreateCylinderLayer1 returns true if the method "XRWebGLBinding.createCylinderLayer" exists.
func (this XRWebGLBinding) HasFuncCreateCylinderLayer1() bool {
	return js.True == bindings.HasFuncXRWebGLBindingCreateCylinderLayer1(
		this.ref,
	)
}

// FuncCreateCylinderLayer1 returns the method "XRWebGLBinding.createCylinderLayer".
func (this XRWebGLBinding) FuncCreateCylinderLayer1() (fn js.Func[func() XRCylinderLayer]) {
	bindings.FuncXRWebGLBindingCreateCylinderLayer1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateCylinderLayer1 calls the method "XRWebGLBinding.createCylinderLayer".
func (this XRWebGLBinding) CreateCylinderLayer1() (ret XRCylinderLayer) {
	bindings.CallXRWebGLBindingCreateCylinderLayer1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateCylinderLayer1 calls the method "XRWebGLBinding.createCylinderLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLBinding) TryCreateCylinderLayer1() (ret XRCylinderLayer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLBindingCreateCylinderLayer1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateEquirectLayer returns true if the method "XRWebGLBinding.createEquirectLayer" exists.
func (this XRWebGLBinding) HasFuncCreateEquirectLayer() bool {
	return js.True == bindings.HasFuncXRWebGLBindingCreateEquirectLayer(
		this.ref,
	)
}

// FuncCreateEquirectLayer returns the method "XRWebGLBinding.createEquirectLayer".
func (this XRWebGLBinding) FuncCreateEquirectLayer() (fn js.Func[func(init XREquirectLayerInit) XREquirectLayer]) {
	bindings.FuncXRWebGLBindingCreateEquirectLayer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateEquirectLayer calls the method "XRWebGLBinding.createEquirectLayer".
func (this XRWebGLBinding) CreateEquirectLayer(init XREquirectLayerInit) (ret XREquirectLayer) {
	bindings.CallXRWebGLBindingCreateEquirectLayer(
		this.ref, js.Pointer(&ret),
		js.Pointer(&init),
	)

	return
}

// TryCreateEquirectLayer calls the method "XRWebGLBinding.createEquirectLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLBinding) TryCreateEquirectLayer(init XREquirectLayerInit) (ret XREquirectLayer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLBindingCreateEquirectLayer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&init),
	)

	return
}

// HasFuncCreateEquirectLayer1 returns true if the method "XRWebGLBinding.createEquirectLayer" exists.
func (this XRWebGLBinding) HasFuncCreateEquirectLayer1() bool {
	return js.True == bindings.HasFuncXRWebGLBindingCreateEquirectLayer1(
		this.ref,
	)
}

// FuncCreateEquirectLayer1 returns the method "XRWebGLBinding.createEquirectLayer".
func (this XRWebGLBinding) FuncCreateEquirectLayer1() (fn js.Func[func() XREquirectLayer]) {
	bindings.FuncXRWebGLBindingCreateEquirectLayer1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateEquirectLayer1 calls the method "XRWebGLBinding.createEquirectLayer".
func (this XRWebGLBinding) CreateEquirectLayer1() (ret XREquirectLayer) {
	bindings.CallXRWebGLBindingCreateEquirectLayer1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateEquirectLayer1 calls the method "XRWebGLBinding.createEquirectLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLBinding) TryCreateEquirectLayer1() (ret XREquirectLayer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLBindingCreateEquirectLayer1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateCubeLayer returns true if the method "XRWebGLBinding.createCubeLayer" exists.
func (this XRWebGLBinding) HasFuncCreateCubeLayer() bool {
	return js.True == bindings.HasFuncXRWebGLBindingCreateCubeLayer(
		this.ref,
	)
}

// FuncCreateCubeLayer returns the method "XRWebGLBinding.createCubeLayer".
func (this XRWebGLBinding) FuncCreateCubeLayer() (fn js.Func[func(init XRCubeLayerInit) XRCubeLayer]) {
	bindings.FuncXRWebGLBindingCreateCubeLayer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateCubeLayer calls the method "XRWebGLBinding.createCubeLayer".
func (this XRWebGLBinding) CreateCubeLayer(init XRCubeLayerInit) (ret XRCubeLayer) {
	bindings.CallXRWebGLBindingCreateCubeLayer(
		this.ref, js.Pointer(&ret),
		js.Pointer(&init),
	)

	return
}

// TryCreateCubeLayer calls the method "XRWebGLBinding.createCubeLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLBinding) TryCreateCubeLayer(init XRCubeLayerInit) (ret XRCubeLayer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLBindingCreateCubeLayer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&init),
	)

	return
}

// HasFuncCreateCubeLayer1 returns true if the method "XRWebGLBinding.createCubeLayer" exists.
func (this XRWebGLBinding) HasFuncCreateCubeLayer1() bool {
	return js.True == bindings.HasFuncXRWebGLBindingCreateCubeLayer1(
		this.ref,
	)
}

// FuncCreateCubeLayer1 returns the method "XRWebGLBinding.createCubeLayer".
func (this XRWebGLBinding) FuncCreateCubeLayer1() (fn js.Func[func() XRCubeLayer]) {
	bindings.FuncXRWebGLBindingCreateCubeLayer1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateCubeLayer1 calls the method "XRWebGLBinding.createCubeLayer".
func (this XRWebGLBinding) CreateCubeLayer1() (ret XRCubeLayer) {
	bindings.CallXRWebGLBindingCreateCubeLayer1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateCubeLayer1 calls the method "XRWebGLBinding.createCubeLayer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLBinding) TryCreateCubeLayer1() (ret XRCubeLayer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLBindingCreateCubeLayer1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetSubImage returns true if the method "XRWebGLBinding.getSubImage" exists.
func (this XRWebGLBinding) HasFuncGetSubImage() bool {
	return js.True == bindings.HasFuncXRWebGLBindingGetSubImage(
		this.ref,
	)
}

// FuncGetSubImage returns the method "XRWebGLBinding.getSubImage".
func (this XRWebGLBinding) FuncGetSubImage() (fn js.Func[func(layer XRCompositionLayer, frame XRFrame, eye XREye) XRWebGLSubImage]) {
	bindings.FuncXRWebGLBindingGetSubImage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetSubImage calls the method "XRWebGLBinding.getSubImage".
func (this XRWebGLBinding) GetSubImage(layer XRCompositionLayer, frame XRFrame, eye XREye) (ret XRWebGLSubImage) {
	bindings.CallXRWebGLBindingGetSubImage(
		this.ref, js.Pointer(&ret),
		layer.Ref(),
		frame.Ref(),
		uint32(eye),
	)

	return
}

// TryGetSubImage calls the method "XRWebGLBinding.getSubImage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLBinding) TryGetSubImage(layer XRCompositionLayer, frame XRFrame, eye XREye) (ret XRWebGLSubImage, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLBindingGetSubImage(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		layer.Ref(),
		frame.Ref(),
		uint32(eye),
	)

	return
}

// HasFuncGetSubImage1 returns true if the method "XRWebGLBinding.getSubImage" exists.
func (this XRWebGLBinding) HasFuncGetSubImage1() bool {
	return js.True == bindings.HasFuncXRWebGLBindingGetSubImage1(
		this.ref,
	)
}

// FuncGetSubImage1 returns the method "XRWebGLBinding.getSubImage".
func (this XRWebGLBinding) FuncGetSubImage1() (fn js.Func[func(layer XRCompositionLayer, frame XRFrame) XRWebGLSubImage]) {
	bindings.FuncXRWebGLBindingGetSubImage1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetSubImage1 calls the method "XRWebGLBinding.getSubImage".
func (this XRWebGLBinding) GetSubImage1(layer XRCompositionLayer, frame XRFrame) (ret XRWebGLSubImage) {
	bindings.CallXRWebGLBindingGetSubImage1(
		this.ref, js.Pointer(&ret),
		layer.Ref(),
		frame.Ref(),
	)

	return
}

// TryGetSubImage1 calls the method "XRWebGLBinding.getSubImage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLBinding) TryGetSubImage1(layer XRCompositionLayer, frame XRFrame) (ret XRWebGLSubImage, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLBindingGetSubImage1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		layer.Ref(),
		frame.Ref(),
	)

	return
}

// HasFuncGetViewSubImage returns true if the method "XRWebGLBinding.getViewSubImage" exists.
func (this XRWebGLBinding) HasFuncGetViewSubImage() bool {
	return js.True == bindings.HasFuncXRWebGLBindingGetViewSubImage(
		this.ref,
	)
}

// FuncGetViewSubImage returns the method "XRWebGLBinding.getViewSubImage".
func (this XRWebGLBinding) FuncGetViewSubImage() (fn js.Func[func(layer XRProjectionLayer, view XRView) XRWebGLSubImage]) {
	bindings.FuncXRWebGLBindingGetViewSubImage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetViewSubImage calls the method "XRWebGLBinding.getViewSubImage".
func (this XRWebGLBinding) GetViewSubImage(layer XRProjectionLayer, view XRView) (ret XRWebGLSubImage) {
	bindings.CallXRWebGLBindingGetViewSubImage(
		this.ref, js.Pointer(&ret),
		layer.Ref(),
		view.Ref(),
	)

	return
}

// TryGetViewSubImage calls the method "XRWebGLBinding.getViewSubImage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLBinding) TryGetViewSubImage(layer XRProjectionLayer, view XRView) (ret XRWebGLSubImage, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLBindingGetViewSubImage(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		layer.Ref(),
		view.Ref(),
	)

	return
}

// HasFuncGetCameraImage returns true if the method "XRWebGLBinding.getCameraImage" exists.
func (this XRWebGLBinding) HasFuncGetCameraImage() bool {
	return js.True == bindings.HasFuncXRWebGLBindingGetCameraImage(
		this.ref,
	)
}

// FuncGetCameraImage returns the method "XRWebGLBinding.getCameraImage".
func (this XRWebGLBinding) FuncGetCameraImage() (fn js.Func[func(camera XRCamera) WebGLTexture]) {
	bindings.FuncXRWebGLBindingGetCameraImage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetCameraImage calls the method "XRWebGLBinding.getCameraImage".
func (this XRWebGLBinding) GetCameraImage(camera XRCamera) (ret WebGLTexture) {
	bindings.CallXRWebGLBindingGetCameraImage(
		this.ref, js.Pointer(&ret),
		camera.Ref(),
	)

	return
}

// TryGetCameraImage calls the method "XRWebGLBinding.getCameraImage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLBinding) TryGetCameraImage(camera XRCamera) (ret WebGLTexture, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLBindingGetCameraImage(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		camera.Ref(),
	)

	return
}

// HasFuncGetDepthInformation returns true if the method "XRWebGLBinding.getDepthInformation" exists.
func (this XRWebGLBinding) HasFuncGetDepthInformation() bool {
	return js.True == bindings.HasFuncXRWebGLBindingGetDepthInformation(
		this.ref,
	)
}

// FuncGetDepthInformation returns the method "XRWebGLBinding.getDepthInformation".
func (this XRWebGLBinding) FuncGetDepthInformation() (fn js.Func[func(view XRView) XRWebGLDepthInformation]) {
	bindings.FuncXRWebGLBindingGetDepthInformation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDepthInformation calls the method "XRWebGLBinding.getDepthInformation".
func (this XRWebGLBinding) GetDepthInformation(view XRView) (ret XRWebGLDepthInformation) {
	bindings.CallXRWebGLBindingGetDepthInformation(
		this.ref, js.Pointer(&ret),
		view.Ref(),
	)

	return
}

// TryGetDepthInformation calls the method "XRWebGLBinding.getDepthInformation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLBinding) TryGetDepthInformation(view XRView) (ret XRWebGLDepthInformation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLBindingGetDepthInformation(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		view.Ref(),
	)

	return
}

// HasFuncGetReflectionCubeMap returns true if the method "XRWebGLBinding.getReflectionCubeMap" exists.
func (this XRWebGLBinding) HasFuncGetReflectionCubeMap() bool {
	return js.True == bindings.HasFuncXRWebGLBindingGetReflectionCubeMap(
		this.ref,
	)
}

// FuncGetReflectionCubeMap returns the method "XRWebGLBinding.getReflectionCubeMap".
func (this XRWebGLBinding) FuncGetReflectionCubeMap() (fn js.Func[func(lightProbe XRLightProbe) WebGLTexture]) {
	bindings.FuncXRWebGLBindingGetReflectionCubeMap(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetReflectionCubeMap calls the method "XRWebGLBinding.getReflectionCubeMap".
func (this XRWebGLBinding) GetReflectionCubeMap(lightProbe XRLightProbe) (ret WebGLTexture) {
	bindings.CallXRWebGLBindingGetReflectionCubeMap(
		this.ref, js.Pointer(&ret),
		lightProbe.Ref(),
	)

	return
}

// TryGetReflectionCubeMap calls the method "XRWebGLBinding.getReflectionCubeMap"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XRWebGLBinding) TryGetReflectionCubeMap(lightProbe XRLightProbe) (ret WebGLTexture, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXRWebGLBindingGetReflectionCubeMap(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		lightProbe.Ref(),
	)

	return
}

type XSLTProcessor struct {
	ref js.Ref
}

func (this XSLTProcessor) Once() XSLTProcessor {
	this.ref.Once()
	return this
}

func (this XSLTProcessor) Ref() js.Ref {
	return this.ref
}

func (this XSLTProcessor) FromRef(ref js.Ref) XSLTProcessor {
	this.ref = ref
	return this
}

func (this XSLTProcessor) Free() {
	this.ref.Free()
}

// HasFuncImportStylesheet returns true if the method "XSLTProcessor.importStylesheet" exists.
func (this XSLTProcessor) HasFuncImportStylesheet() bool {
	return js.True == bindings.HasFuncXSLTProcessorImportStylesheet(
		this.ref,
	)
}

// FuncImportStylesheet returns the method "XSLTProcessor.importStylesheet".
func (this XSLTProcessor) FuncImportStylesheet() (fn js.Func[func(style Node)]) {
	bindings.FuncXSLTProcessorImportStylesheet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ImportStylesheet calls the method "XSLTProcessor.importStylesheet".
func (this XSLTProcessor) ImportStylesheet(style Node) (ret js.Void) {
	bindings.CallXSLTProcessorImportStylesheet(
		this.ref, js.Pointer(&ret),
		style.Ref(),
	)

	return
}

// TryImportStylesheet calls the method "XSLTProcessor.importStylesheet"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XSLTProcessor) TryImportStylesheet(style Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXSLTProcessorImportStylesheet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		style.Ref(),
	)

	return
}

// HasFuncTransformToFragment returns true if the method "XSLTProcessor.transformToFragment" exists.
func (this XSLTProcessor) HasFuncTransformToFragment() bool {
	return js.True == bindings.HasFuncXSLTProcessorTransformToFragment(
		this.ref,
	)
}

// FuncTransformToFragment returns the method "XSLTProcessor.transformToFragment".
func (this XSLTProcessor) FuncTransformToFragment() (fn js.Func[func(source Node, output Document) DocumentFragment]) {
	bindings.FuncXSLTProcessorTransformToFragment(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TransformToFragment calls the method "XSLTProcessor.transformToFragment".
func (this XSLTProcessor) TransformToFragment(source Node, output Document) (ret DocumentFragment) {
	bindings.CallXSLTProcessorTransformToFragment(
		this.ref, js.Pointer(&ret),
		source.Ref(),
		output.Ref(),
	)

	return
}

// TryTransformToFragment calls the method "XSLTProcessor.transformToFragment"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XSLTProcessor) TryTransformToFragment(source Node, output Document) (ret DocumentFragment, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXSLTProcessorTransformToFragment(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		source.Ref(),
		output.Ref(),
	)

	return
}

// HasFuncTransformToDocument returns true if the method "XSLTProcessor.transformToDocument" exists.
func (this XSLTProcessor) HasFuncTransformToDocument() bool {
	return js.True == bindings.HasFuncXSLTProcessorTransformToDocument(
		this.ref,
	)
}

// FuncTransformToDocument returns the method "XSLTProcessor.transformToDocument".
func (this XSLTProcessor) FuncTransformToDocument() (fn js.Func[func(source Node) Document]) {
	bindings.FuncXSLTProcessorTransformToDocument(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TransformToDocument calls the method "XSLTProcessor.transformToDocument".
func (this XSLTProcessor) TransformToDocument(source Node) (ret Document) {
	bindings.CallXSLTProcessorTransformToDocument(
		this.ref, js.Pointer(&ret),
		source.Ref(),
	)

	return
}

// TryTransformToDocument calls the method "XSLTProcessor.transformToDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XSLTProcessor) TryTransformToDocument(source Node) (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXSLTProcessorTransformToDocument(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		source.Ref(),
	)

	return
}

// HasFuncSetParameter returns true if the method "XSLTProcessor.setParameter" exists.
func (this XSLTProcessor) HasFuncSetParameter() bool {
	return js.True == bindings.HasFuncXSLTProcessorSetParameter(
		this.ref,
	)
}

// FuncSetParameter returns the method "XSLTProcessor.setParameter".
func (this XSLTProcessor) FuncSetParameter() (fn js.Func[func(namespaceURI js.String, localName js.String, value js.Any)]) {
	bindings.FuncXSLTProcessorSetParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetParameter calls the method "XSLTProcessor.setParameter".
func (this XSLTProcessor) SetParameter(namespaceURI js.String, localName js.String, value js.Any) (ret js.Void) {
	bindings.CallXSLTProcessorSetParameter(
		this.ref, js.Pointer(&ret),
		namespaceURI.Ref(),
		localName.Ref(),
		value.Ref(),
	)

	return
}

// TrySetParameter calls the method "XSLTProcessor.setParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XSLTProcessor) TrySetParameter(namespaceURI js.String, localName js.String, value js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXSLTProcessorSetParameter(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespaceURI.Ref(),
		localName.Ref(),
		value.Ref(),
	)

	return
}

// HasFuncGetParameter returns true if the method "XSLTProcessor.getParameter" exists.
func (this XSLTProcessor) HasFuncGetParameter() bool {
	return js.True == bindings.HasFuncXSLTProcessorGetParameter(
		this.ref,
	)
}

// FuncGetParameter returns the method "XSLTProcessor.getParameter".
func (this XSLTProcessor) FuncGetParameter() (fn js.Func[func(namespaceURI js.String, localName js.String) js.Any]) {
	bindings.FuncXSLTProcessorGetParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetParameter calls the method "XSLTProcessor.getParameter".
func (this XSLTProcessor) GetParameter(namespaceURI js.String, localName js.String) (ret js.Any) {
	bindings.CallXSLTProcessorGetParameter(
		this.ref, js.Pointer(&ret),
		namespaceURI.Ref(),
		localName.Ref(),
	)

	return
}

// TryGetParameter calls the method "XSLTProcessor.getParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XSLTProcessor) TryGetParameter(namespaceURI js.String, localName js.String) (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXSLTProcessorGetParameter(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespaceURI.Ref(),
		localName.Ref(),
	)

	return
}

// HasFuncRemoveParameter returns true if the method "XSLTProcessor.removeParameter" exists.
func (this XSLTProcessor) HasFuncRemoveParameter() bool {
	return js.True == bindings.HasFuncXSLTProcessorRemoveParameter(
		this.ref,
	)
}

// FuncRemoveParameter returns the method "XSLTProcessor.removeParameter".
func (this XSLTProcessor) FuncRemoveParameter() (fn js.Func[func(namespaceURI js.String, localName js.String)]) {
	bindings.FuncXSLTProcessorRemoveParameter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveParameter calls the method "XSLTProcessor.removeParameter".
func (this XSLTProcessor) RemoveParameter(namespaceURI js.String, localName js.String) (ret js.Void) {
	bindings.CallXSLTProcessorRemoveParameter(
		this.ref, js.Pointer(&ret),
		namespaceURI.Ref(),
		localName.Ref(),
	)

	return
}

// TryRemoveParameter calls the method "XSLTProcessor.removeParameter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XSLTProcessor) TryRemoveParameter(namespaceURI js.String, localName js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXSLTProcessorRemoveParameter(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		namespaceURI.Ref(),
		localName.Ref(),
	)

	return
}

// HasFuncClearParameters returns true if the method "XSLTProcessor.clearParameters" exists.
func (this XSLTProcessor) HasFuncClearParameters() bool {
	return js.True == bindings.HasFuncXSLTProcessorClearParameters(
		this.ref,
	)
}

// FuncClearParameters returns the method "XSLTProcessor.clearParameters".
func (this XSLTProcessor) FuncClearParameters() (fn js.Func[func()]) {
	bindings.FuncXSLTProcessorClearParameters(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearParameters calls the method "XSLTProcessor.clearParameters".
func (this XSLTProcessor) ClearParameters() (ret js.Void) {
	bindings.CallXSLTProcessorClearParameters(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClearParameters calls the method "XSLTProcessor.clearParameters"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XSLTProcessor) TryClearParameters() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXSLTProcessorClearParameters(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReset returns true if the method "XSLTProcessor.reset" exists.
func (this XSLTProcessor) HasFuncReset() bool {
	return js.True == bindings.HasFuncXSLTProcessorReset(
		this.ref,
	)
}

// FuncReset returns the method "XSLTProcessor.reset".
func (this XSLTProcessor) FuncReset() (fn js.Func[func()]) {
	bindings.FuncXSLTProcessorReset(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Reset calls the method "XSLTProcessor.reset".
func (this XSLTProcessor) Reset() (ret js.Void) {
	bindings.CallXSLTProcessorReset(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReset calls the method "XSLTProcessor.reset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this XSLTProcessor) TryReset() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryXSLTProcessorReset(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type Console struct{}

// HasFuncAssert returns ture if the function "console.assert" exists.
func (Console) HasFuncAssert() bool {
	return js.True == bindings.HasFuncConsoleAssert()
}

// FuncAssert returns the function "console.assert".
func (Console) FuncAssert() (fn js.Func[func(condition bool, data ...js.Any)]) {
	bindings.FuncConsoleAssert(
		js.Pointer(&fn),
	)
	return
}

// Assert calls the function "console.assert".
func (Console) Assert(condition bool, data ...js.Any) (ret js.Void) {
	bindings.CallConsoleAssert(
		js.Pointer(&ret),
		js.Bool(bool(condition)),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// TryAssert calls the function "console.assert"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryAssert(condition bool, data ...js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleAssert(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(condition)),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// HasFuncAssert1 returns ture if the function "console.assert" exists.
func (Console) HasFuncAssert1() bool {
	return js.True == bindings.HasFuncConsoleAssert1()
}

// FuncAssert1 returns the function "console.assert".
func (Console) FuncAssert1() (fn js.Func[func()]) {
	bindings.FuncConsoleAssert1(
		js.Pointer(&fn),
	)
	return
}

// Assert1 calls the function "console.assert".
func (Console) Assert1() (ret js.Void) {
	bindings.CallConsoleAssert1(
		js.Pointer(&ret),
	)
	return
}

// TryAssert1 calls the function "console.assert"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryAssert1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleAssert1(
		js.Pointer(&ret), js.Pointer(&exception),
	)
	return
}

// HasFuncClear returns ture if the function "console.clear" exists.
func (Console) HasFuncClear() bool {
	return js.True == bindings.HasFuncConsoleClear()
}

// FuncClear returns the function "console.clear".
func (Console) FuncClear() (fn js.Func[func()]) {
	bindings.FuncConsoleClear(
		js.Pointer(&fn),
	)
	return
}

// Clear calls the function "console.clear".
func (Console) Clear() (ret js.Void) {
	bindings.CallConsoleClear(
		js.Pointer(&ret),
	)
	return
}

// TryClear calls the function "console.clear"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryClear() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleClear(
		js.Pointer(&ret), js.Pointer(&exception),
	)
	return
}

// HasFuncDebug returns ture if the function "console.debug" exists.
func (Console) HasFuncDebug() bool {
	return js.True == bindings.HasFuncConsoleDebug()
}

// FuncDebug returns the function "console.debug".
func (Console) FuncDebug() (fn js.Func[func(data ...js.Any)]) {
	bindings.FuncConsoleDebug(
		js.Pointer(&fn),
	)
	return
}

// Debug calls the function "console.debug".
func (Console) Debug(data ...js.Any) (ret js.Void) {
	bindings.CallConsoleDebug(
		js.Pointer(&ret),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// TryDebug calls the function "console.debug"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryDebug(data ...js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleDebug(
		js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// HasFuncError returns ture if the function "console.error" exists.
func (Console) HasFuncError() bool {
	return js.True == bindings.HasFuncConsoleError()
}

// FuncError returns the function "console.error".
func (Console) FuncError() (fn js.Func[func(data ...js.Any)]) {
	bindings.FuncConsoleError(
		js.Pointer(&fn),
	)
	return
}

// Error calls the function "console.error".
func (Console) Error(data ...js.Any) (ret js.Void) {
	bindings.CallConsoleError(
		js.Pointer(&ret),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// TryError calls the function "console.error"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryError(data ...js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleError(
		js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// HasFuncInfo returns ture if the function "console.info" exists.
func (Console) HasFuncInfo() bool {
	return js.True == bindings.HasFuncConsoleInfo()
}

// FuncInfo returns the function "console.info".
func (Console) FuncInfo() (fn js.Func[func(data ...js.Any)]) {
	bindings.FuncConsoleInfo(
		js.Pointer(&fn),
	)
	return
}

// Info calls the function "console.info".
func (Console) Info(data ...js.Any) (ret js.Void) {
	bindings.CallConsoleInfo(
		js.Pointer(&ret),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// TryInfo calls the function "console.info"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryInfo(data ...js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// HasFuncLog returns ture if the function "console.log" exists.
func (Console) HasFuncLog() bool {
	return js.True == bindings.HasFuncConsoleLog()
}

// FuncLog returns the function "console.log".
func (Console) FuncLog() (fn js.Func[func(data ...js.Any)]) {
	bindings.FuncConsoleLog(
		js.Pointer(&fn),
	)
	return
}

// Log calls the function "console.log".
func (Console) Log(data ...js.Any) (ret js.Void) {
	bindings.CallConsoleLog(
		js.Pointer(&ret),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// TryLog calls the function "console.log"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryLog(data ...js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleLog(
		js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// HasFuncTable returns ture if the function "console.table" exists.
func (Console) HasFuncTable() bool {
	return js.True == bindings.HasFuncConsoleTable()
}

// FuncTable returns the function "console.table".
func (Console) FuncTable() (fn js.Func[func(tabularData js.Any, properties js.Array[js.String])]) {
	bindings.FuncConsoleTable(
		js.Pointer(&fn),
	)
	return
}

// Table calls the function "console.table".
func (Console) Table(tabularData js.Any, properties js.Array[js.String]) (ret js.Void) {
	bindings.CallConsoleTable(
		js.Pointer(&ret),
		tabularData.Ref(),
		properties.Ref(),
	)
	return
}

// TryTable calls the function "console.table"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryTable(tabularData js.Any, properties js.Array[js.String]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleTable(
		js.Pointer(&ret), js.Pointer(&exception),
		tabularData.Ref(),
		properties.Ref(),
	)
	return
}

// HasFuncTable1 returns ture if the function "console.table" exists.
func (Console) HasFuncTable1() bool {
	return js.True == bindings.HasFuncConsoleTable1()
}

// FuncTable1 returns the function "console.table".
func (Console) FuncTable1() (fn js.Func[func(tabularData js.Any)]) {
	bindings.FuncConsoleTable1(
		js.Pointer(&fn),
	)
	return
}

// Table1 calls the function "console.table".
func (Console) Table1(tabularData js.Any) (ret js.Void) {
	bindings.CallConsoleTable1(
		js.Pointer(&ret),
		tabularData.Ref(),
	)
	return
}

// TryTable1 calls the function "console.table"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryTable1(tabularData js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleTable1(
		js.Pointer(&ret), js.Pointer(&exception),
		tabularData.Ref(),
	)
	return
}

// HasFuncTable2 returns ture if the function "console.table" exists.
func (Console) HasFuncTable2() bool {
	return js.True == bindings.HasFuncConsoleTable2()
}

// FuncTable2 returns the function "console.table".
func (Console) FuncTable2() (fn js.Func[func()]) {
	bindings.FuncConsoleTable2(
		js.Pointer(&fn),
	)
	return
}

// Table2 calls the function "console.table".
func (Console) Table2() (ret js.Void) {
	bindings.CallConsoleTable2(
		js.Pointer(&ret),
	)
	return
}

// TryTable2 calls the function "console.table"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryTable2() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleTable2(
		js.Pointer(&ret), js.Pointer(&exception),
	)
	return
}

// HasFuncTrace returns ture if the function "console.trace" exists.
func (Console) HasFuncTrace() bool {
	return js.True == bindings.HasFuncConsoleTrace()
}

// FuncTrace returns the function "console.trace".
func (Console) FuncTrace() (fn js.Func[func(data ...js.Any)]) {
	bindings.FuncConsoleTrace(
		js.Pointer(&fn),
	)
	return
}

// Trace calls the function "console.trace".
func (Console) Trace(data ...js.Any) (ret js.Void) {
	bindings.CallConsoleTrace(
		js.Pointer(&ret),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// TryTrace calls the function "console.trace"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryTrace(data ...js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleTrace(
		js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// HasFuncWarn returns ture if the function "console.warn" exists.
func (Console) HasFuncWarn() bool {
	return js.True == bindings.HasFuncConsoleWarn()
}

// FuncWarn returns the function "console.warn".
func (Console) FuncWarn() (fn js.Func[func(data ...js.Any)]) {
	bindings.FuncConsoleWarn(
		js.Pointer(&fn),
	)
	return
}

// Warn calls the function "console.warn".
func (Console) Warn(data ...js.Any) (ret js.Void) {
	bindings.CallConsoleWarn(
		js.Pointer(&ret),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// TryWarn calls the function "console.warn"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryWarn(data ...js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleWarn(
		js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// HasFuncDir returns ture if the function "console.dir" exists.
func (Console) HasFuncDir() bool {
	return js.True == bindings.HasFuncConsoleDir()
}

// FuncDir returns the function "console.dir".
func (Console) FuncDir() (fn js.Func[func(item js.Any, options js.Object)]) {
	bindings.FuncConsoleDir(
		js.Pointer(&fn),
	)
	return
}

// Dir calls the function "console.dir".
func (Console) Dir(item js.Any, options js.Object) (ret js.Void) {
	bindings.CallConsoleDir(
		js.Pointer(&ret),
		item.Ref(),
		options.Ref(),
	)
	return
}

// TryDir calls the function "console.dir"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryDir(item js.Any, options js.Object) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleDir(
		js.Pointer(&ret), js.Pointer(&exception),
		item.Ref(),
		options.Ref(),
	)
	return
}

// HasFuncDir1 returns ture if the function "console.dir" exists.
func (Console) HasFuncDir1() bool {
	return js.True == bindings.HasFuncConsoleDir1()
}

// FuncDir1 returns the function "console.dir".
func (Console) FuncDir1() (fn js.Func[func(item js.Any)]) {
	bindings.FuncConsoleDir1(
		js.Pointer(&fn),
	)
	return
}

// Dir1 calls the function "console.dir".
func (Console) Dir1(item js.Any) (ret js.Void) {
	bindings.CallConsoleDir1(
		js.Pointer(&ret),
		item.Ref(),
	)
	return
}

// TryDir1 calls the function "console.dir"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryDir1(item js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleDir1(
		js.Pointer(&ret), js.Pointer(&exception),
		item.Ref(),
	)
	return
}

// HasFuncDir2 returns ture if the function "console.dir" exists.
func (Console) HasFuncDir2() bool {
	return js.True == bindings.HasFuncConsoleDir2()
}

// FuncDir2 returns the function "console.dir".
func (Console) FuncDir2() (fn js.Func[func()]) {
	bindings.FuncConsoleDir2(
		js.Pointer(&fn),
	)
	return
}

// Dir2 calls the function "console.dir".
func (Console) Dir2() (ret js.Void) {
	bindings.CallConsoleDir2(
		js.Pointer(&ret),
	)
	return
}

// TryDir2 calls the function "console.dir"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryDir2() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleDir2(
		js.Pointer(&ret), js.Pointer(&exception),
	)
	return
}

// HasFuncDirxml returns ture if the function "console.dirxml" exists.
func (Console) HasFuncDirxml() bool {
	return js.True == bindings.HasFuncConsoleDirxml()
}

// FuncDirxml returns the function "console.dirxml".
func (Console) FuncDirxml() (fn js.Func[func(data ...js.Any)]) {
	bindings.FuncConsoleDirxml(
		js.Pointer(&fn),
	)
	return
}

// Dirxml calls the function "console.dirxml".
func (Console) Dirxml(data ...js.Any) (ret js.Void) {
	bindings.CallConsoleDirxml(
		js.Pointer(&ret),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// TryDirxml calls the function "console.dirxml"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryDirxml(data ...js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleDirxml(
		js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// HasFuncCount returns ture if the function "console.count" exists.
func (Console) HasFuncCount() bool {
	return js.True == bindings.HasFuncConsoleCount()
}

// FuncCount returns the function "console.count".
func (Console) FuncCount() (fn js.Func[func(label js.String)]) {
	bindings.FuncConsoleCount(
		js.Pointer(&fn),
	)
	return
}

// Count calls the function "console.count".
func (Console) Count(label js.String) (ret js.Void) {
	bindings.CallConsoleCount(
		js.Pointer(&ret),
		label.Ref(),
	)
	return
}

// TryCount calls the function "console.count"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryCount(label js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleCount(
		js.Pointer(&ret), js.Pointer(&exception),
		label.Ref(),
	)
	return
}

// HasFuncCount1 returns ture if the function "console.count" exists.
func (Console) HasFuncCount1() bool {
	return js.True == bindings.HasFuncConsoleCount1()
}

// FuncCount1 returns the function "console.count".
func (Console) FuncCount1() (fn js.Func[func()]) {
	bindings.FuncConsoleCount1(
		js.Pointer(&fn),
	)
	return
}

// Count1 calls the function "console.count".
func (Console) Count1() (ret js.Void) {
	bindings.CallConsoleCount1(
		js.Pointer(&ret),
	)
	return
}

// TryCount1 calls the function "console.count"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryCount1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleCount1(
		js.Pointer(&ret), js.Pointer(&exception),
	)
	return
}

// HasFuncCountReset returns ture if the function "console.countReset" exists.
func (Console) HasFuncCountReset() bool {
	return js.True == bindings.HasFuncConsoleCountReset()
}

// FuncCountReset returns the function "console.countReset".
func (Console) FuncCountReset() (fn js.Func[func(label js.String)]) {
	bindings.FuncConsoleCountReset(
		js.Pointer(&fn),
	)
	return
}

// CountReset calls the function "console.countReset".
func (Console) CountReset(label js.String) (ret js.Void) {
	bindings.CallConsoleCountReset(
		js.Pointer(&ret),
		label.Ref(),
	)
	return
}

// TryCountReset calls the function "console.countReset"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryCountReset(label js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleCountReset(
		js.Pointer(&ret), js.Pointer(&exception),
		label.Ref(),
	)
	return
}

// HasFuncCountReset1 returns ture if the function "console.countReset" exists.
func (Console) HasFuncCountReset1() bool {
	return js.True == bindings.HasFuncConsoleCountReset1()
}

// FuncCountReset1 returns the function "console.countReset".
func (Console) FuncCountReset1() (fn js.Func[func()]) {
	bindings.FuncConsoleCountReset1(
		js.Pointer(&fn),
	)
	return
}

// CountReset1 calls the function "console.countReset".
func (Console) CountReset1() (ret js.Void) {
	bindings.CallConsoleCountReset1(
		js.Pointer(&ret),
	)
	return
}

// TryCountReset1 calls the function "console.countReset"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryCountReset1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleCountReset1(
		js.Pointer(&ret), js.Pointer(&exception),
	)
	return
}

// HasFuncGroup returns ture if the function "console.group" exists.
func (Console) HasFuncGroup() bool {
	return js.True == bindings.HasFuncConsoleGroup()
}

// FuncGroup returns the function "console.group".
func (Console) FuncGroup() (fn js.Func[func(data ...js.Any)]) {
	bindings.FuncConsoleGroup(
		js.Pointer(&fn),
	)
	return
}

// Group calls the function "console.group".
func (Console) Group(data ...js.Any) (ret js.Void) {
	bindings.CallConsoleGroup(
		js.Pointer(&ret),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// TryGroup calls the function "console.group"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryGroup(data ...js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleGroup(
		js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// HasFuncGroupCollapsed returns ture if the function "console.groupCollapsed" exists.
func (Console) HasFuncGroupCollapsed() bool {
	return js.True == bindings.HasFuncConsoleGroupCollapsed()
}

// FuncGroupCollapsed returns the function "console.groupCollapsed".
func (Console) FuncGroupCollapsed() (fn js.Func[func(data ...js.Any)]) {
	bindings.FuncConsoleGroupCollapsed(
		js.Pointer(&fn),
	)
	return
}

// GroupCollapsed calls the function "console.groupCollapsed".
func (Console) GroupCollapsed(data ...js.Any) (ret js.Void) {
	bindings.CallConsoleGroupCollapsed(
		js.Pointer(&ret),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// TryGroupCollapsed calls the function "console.groupCollapsed"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryGroupCollapsed(data ...js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleGroupCollapsed(
		js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// HasFuncGroupEnd returns ture if the function "console.groupEnd" exists.
func (Console) HasFuncGroupEnd() bool {
	return js.True == bindings.HasFuncConsoleGroupEnd()
}

// FuncGroupEnd returns the function "console.groupEnd".
func (Console) FuncGroupEnd() (fn js.Func[func()]) {
	bindings.FuncConsoleGroupEnd(
		js.Pointer(&fn),
	)
	return
}

// GroupEnd calls the function "console.groupEnd".
func (Console) GroupEnd() (ret js.Void) {
	bindings.CallConsoleGroupEnd(
		js.Pointer(&ret),
	)
	return
}

// TryGroupEnd calls the function "console.groupEnd"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryGroupEnd() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleGroupEnd(
		js.Pointer(&ret), js.Pointer(&exception),
	)
	return
}

// HasFuncTime returns ture if the function "console.time" exists.
func (Console) HasFuncTime() bool {
	return js.True == bindings.HasFuncConsoleTime()
}

// FuncTime returns the function "console.time".
func (Console) FuncTime() (fn js.Func[func(label js.String)]) {
	bindings.FuncConsoleTime(
		js.Pointer(&fn),
	)
	return
}

// Time calls the function "console.time".
func (Console) Time(label js.String) (ret js.Void) {
	bindings.CallConsoleTime(
		js.Pointer(&ret),
		label.Ref(),
	)
	return
}

// TryTime calls the function "console.time"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryTime(label js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleTime(
		js.Pointer(&ret), js.Pointer(&exception),
		label.Ref(),
	)
	return
}

// HasFuncTime1 returns ture if the function "console.time" exists.
func (Console) HasFuncTime1() bool {
	return js.True == bindings.HasFuncConsoleTime1()
}

// FuncTime1 returns the function "console.time".
func (Console) FuncTime1() (fn js.Func[func()]) {
	bindings.FuncConsoleTime1(
		js.Pointer(&fn),
	)
	return
}

// Time1 calls the function "console.time".
func (Console) Time1() (ret js.Void) {
	bindings.CallConsoleTime1(
		js.Pointer(&ret),
	)
	return
}

// TryTime1 calls the function "console.time"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryTime1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleTime1(
		js.Pointer(&ret), js.Pointer(&exception),
	)
	return
}

// HasFuncTimeLog returns ture if the function "console.timeLog" exists.
func (Console) HasFuncTimeLog() bool {
	return js.True == bindings.HasFuncConsoleTimeLog()
}

// FuncTimeLog returns the function "console.timeLog".
func (Console) FuncTimeLog() (fn js.Func[func(label js.String, data ...js.Any)]) {
	bindings.FuncConsoleTimeLog(
		js.Pointer(&fn),
	)
	return
}

// TimeLog calls the function "console.timeLog".
func (Console) TimeLog(label js.String, data ...js.Any) (ret js.Void) {
	bindings.CallConsoleTimeLog(
		js.Pointer(&ret),
		label.Ref(),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// TryTimeLog calls the function "console.timeLog"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryTimeLog(label js.String, data ...js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleTimeLog(
		js.Pointer(&ret), js.Pointer(&exception),
		label.Ref(),
		js.SliceData(data),
		js.SizeU(len(data)),
	)
	return
}

// HasFuncTimeLog1 returns ture if the function "console.timeLog" exists.
func (Console) HasFuncTimeLog1() bool {
	return js.True == bindings.HasFuncConsoleTimeLog1()
}

// FuncTimeLog1 returns the function "console.timeLog".
func (Console) FuncTimeLog1() (fn js.Func[func()]) {
	bindings.FuncConsoleTimeLog1(
		js.Pointer(&fn),
	)
	return
}

// TimeLog1 calls the function "console.timeLog".
func (Console) TimeLog1() (ret js.Void) {
	bindings.CallConsoleTimeLog1(
		js.Pointer(&ret),
	)
	return
}

// TryTimeLog1 calls the function "console.timeLog"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryTimeLog1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleTimeLog1(
		js.Pointer(&ret), js.Pointer(&exception),
	)
	return
}

// HasFuncTimeEnd returns ture if the function "console.timeEnd" exists.
func (Console) HasFuncTimeEnd() bool {
	return js.True == bindings.HasFuncConsoleTimeEnd()
}

// FuncTimeEnd returns the function "console.timeEnd".
func (Console) FuncTimeEnd() (fn js.Func[func(label js.String)]) {
	bindings.FuncConsoleTimeEnd(
		js.Pointer(&fn),
	)
	return
}

// TimeEnd calls the function "console.timeEnd".
func (Console) TimeEnd(label js.String) (ret js.Void) {
	bindings.CallConsoleTimeEnd(
		js.Pointer(&ret),
		label.Ref(),
	)
	return
}

// TryTimeEnd calls the function "console.timeEnd"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryTimeEnd(label js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleTimeEnd(
		js.Pointer(&ret), js.Pointer(&exception),
		label.Ref(),
	)
	return
}

// HasFuncTimeEnd1 returns ture if the function "console.timeEnd" exists.
func (Console) HasFuncTimeEnd1() bool {
	return js.True == bindings.HasFuncConsoleTimeEnd1()
}

// FuncTimeEnd1 returns the function "console.timeEnd".
func (Console) FuncTimeEnd1() (fn js.Func[func()]) {
	bindings.FuncConsoleTimeEnd1(
		js.Pointer(&fn),
	)
	return
}

// TimeEnd1 calls the function "console.timeEnd".
func (Console) TimeEnd1() (ret js.Void) {
	bindings.CallConsoleTimeEnd1(
		js.Pointer(&ret),
	)
	return
}

// TryTimeEnd1 calls the function "console.timeEnd"
// in a try/catch block and returns (_, err, ok = true) when it went through
// the catch clause.
func (Console) TryTimeEnd1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConsoleTimeEnd1(
		js.Pointer(&ret), js.Pointer(&exception),
	)
	return
}

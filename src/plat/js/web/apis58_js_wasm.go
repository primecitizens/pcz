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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Layout returns the value of property "XRCompositionLayer.layout".
//
// The returned bool will be false if there is no such property.
func (this XRCompositionLayer) Layout() (XRLayerLayout, bool) {
	var _ok bool
	_ret := bindings.GetXRCompositionLayerLayout(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRLayerLayout(_ret), _ok
}

// BlendTextureSourceAlpha returns the value of property "XRCompositionLayer.blendTextureSourceAlpha".
//
// The returned bool will be false if there is no such property.
func (this XRCompositionLayer) BlendTextureSourceAlpha() (bool, bool) {
	var _ok bool
	_ret := bindings.GetXRCompositionLayerBlendTextureSourceAlpha(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// BlendTextureSourceAlpha sets the value of property "XRCompositionLayer.blendTextureSourceAlpha" to val.
//
// It returns false if the property cannot be set.
func (this XRCompositionLayer) SetBlendTextureSourceAlpha(val bool) bool {
	return js.True == bindings.SetXRCompositionLayerBlendTextureSourceAlpha(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// ForceMonoPresentation returns the value of property "XRCompositionLayer.forceMonoPresentation".
//
// The returned bool will be false if there is no such property.
func (this XRCompositionLayer) ForceMonoPresentation() (bool, bool) {
	var _ok bool
	_ret := bindings.GetXRCompositionLayerForceMonoPresentation(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ForceMonoPresentation sets the value of property "XRCompositionLayer.forceMonoPresentation" to val.
//
// It returns false if the property cannot be set.
func (this XRCompositionLayer) SetForceMonoPresentation(val bool) bool {
	return js.True == bindings.SetXRCompositionLayerForceMonoPresentation(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Opacity returns the value of property "XRCompositionLayer.opacity".
//
// The returned bool will be false if there is no such property.
func (this XRCompositionLayer) Opacity() (float32, bool) {
	var _ok bool
	_ret := bindings.GetXRCompositionLayerOpacity(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Opacity sets the value of property "XRCompositionLayer.opacity" to val.
//
// It returns false if the property cannot be set.
func (this XRCompositionLayer) SetOpacity(val float32) bool {
	return js.True == bindings.SetXRCompositionLayerOpacity(
		this.Ref(),
		float32(val),
	)
}

// MipLevels returns the value of property "XRCompositionLayer.mipLevels".
//
// The returned bool will be false if there is no such property.
func (this XRCompositionLayer) MipLevels() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRCompositionLayerMipLevels(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Quality returns the value of property "XRCompositionLayer.quality".
//
// The returned bool will be false if there is no such property.
func (this XRCompositionLayer) Quality() (XRLayerQuality, bool) {
	var _ok bool
	_ret := bindings.GetXRCompositionLayerQuality(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRLayerQuality(_ret), _ok
}

// Quality sets the value of property "XRCompositionLayer.quality" to val.
//
// It returns false if the property cannot be set.
func (this XRCompositionLayer) SetQuality(val XRLayerQuality) bool {
	return js.True == bindings.SetXRCompositionLayerQuality(
		this.Ref(),
		uint32(val),
	)
}

// NeedsRedraw returns the value of property "XRCompositionLayer.needsRedraw".
//
// The returned bool will be false if there is no such property.
func (this XRCompositionLayer) NeedsRedraw() (bool, bool) {
	var _ok bool
	_ret := bindings.GetXRCompositionLayerNeedsRedraw(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Destroy calls the method "XRCompositionLayer.destroy".
//
// The returned bool will be false if there is no such method.
func (this XRCompositionLayer) Destroy() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXRCompositionLayerDestroy(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DestroyFunc returns the method "XRCompositionLayer.destroy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRCompositionLayer) DestroyFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.XRCompositionLayerDestroyFunc(
			this.Ref(),
		),
	)
}

type XRCubeLayer struct {
	XRCompositionLayer
}

func (this XRCubeLayer) Once() XRCubeLayer {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Space returns the value of property "XRCubeLayer.space".
//
// The returned bool will be false if there is no such property.
func (this XRCubeLayer) Space() (XRSpace, bool) {
	var _ok bool
	_ret := bindings.GetXRCubeLayerSpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRSpace{}.FromRef(_ret), _ok
}

// Space sets the value of property "XRCubeLayer.space" to val.
//
// It returns false if the property cannot be set.
func (this XRCubeLayer) SetSpace(val XRSpace) bool {
	return js.True == bindings.SetXRCubeLayerSpace(
		this.Ref(),
		val.Ref(),
	)
}

// Orientation returns the value of property "XRCubeLayer.orientation".
//
// The returned bool will be false if there is no such property.
func (this XRCubeLayer) Orientation() (DOMPointReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetXRCubeLayerOrientation(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMPointReadOnly{}.FromRef(_ret), _ok
}

// Orientation sets the value of property "XRCubeLayer.orientation" to val.
//
// It returns false if the property cannot be set.
func (this XRCubeLayer) SetOrientation(val DOMPointReadOnly) bool {
	return js.True == bindings.SetXRCubeLayerOrientation(
		this.Ref(),
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
	ColorFormat GLenum
	// DepthFormat is "XRCubeLayerInit.depthFormat"
	//
	// Optional
	DepthFormat GLenum
	// MipLevels is "XRCubeLayerInit.mipLevels"
	//
	// Optional, defaults to 1.
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
	IsStatic bool
	// ClearOnAccess is "XRCubeLayerInit.clearOnAccess"
	//
	// Optional, defaults to true.
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

// New creates a new {0x140004cc0e0 XRCubeLayerInit XRCubeLayerInit [// XRCubeLayerInit] [0x14000aee640 0x14000aee6e0 0x14000aee780 0x14000aee960 0x14000aeeaa0 0x14000aeebe0 0x14000aeec80 0x14000aeed20 0x14000aeedc0 0x14000aeef00 0x14000aee820 0x14000aeea00 0x14000aeeb40 0x14000aeee60 0x14000aeefa0] 0x14000ad4e70 {0 0}} in the application heap.
func (p XRCubeLayerInit) New() js.Ref {
	return bindings.XRCubeLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRCubeLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRCubeLayerInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRCubeLayerInit) Update(ref js.Ref) {
	bindings.XRCubeLayerInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type XRCylinderLayer struct {
	XRCompositionLayer
}

func (this XRCylinderLayer) Once() XRCylinderLayer {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Space returns the value of property "XRCylinderLayer.space".
//
// The returned bool will be false if there is no such property.
func (this XRCylinderLayer) Space() (XRSpace, bool) {
	var _ok bool
	_ret := bindings.GetXRCylinderLayerSpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRSpace{}.FromRef(_ret), _ok
}

// Space sets the value of property "XRCylinderLayer.space" to val.
//
// It returns false if the property cannot be set.
func (this XRCylinderLayer) SetSpace(val XRSpace) bool {
	return js.True == bindings.SetXRCylinderLayerSpace(
		this.Ref(),
		val.Ref(),
	)
}

// Transform returns the value of property "XRCylinderLayer.transform".
//
// The returned bool will be false if there is no such property.
func (this XRCylinderLayer) Transform() (XRRigidTransform, bool) {
	var _ok bool
	_ret := bindings.GetXRCylinderLayerTransform(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRRigidTransform{}.FromRef(_ret), _ok
}

// Transform sets the value of property "XRCylinderLayer.transform" to val.
//
// It returns false if the property cannot be set.
func (this XRCylinderLayer) SetTransform(val XRRigidTransform) bool {
	return js.True == bindings.SetXRCylinderLayerTransform(
		this.Ref(),
		val.Ref(),
	)
}

// Radius returns the value of property "XRCylinderLayer.radius".
//
// The returned bool will be false if there is no such property.
func (this XRCylinderLayer) Radius() (float32, bool) {
	var _ok bool
	_ret := bindings.GetXRCylinderLayerRadius(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Radius sets the value of property "XRCylinderLayer.radius" to val.
//
// It returns false if the property cannot be set.
func (this XRCylinderLayer) SetRadius(val float32) bool {
	return js.True == bindings.SetXRCylinderLayerRadius(
		this.Ref(),
		float32(val),
	)
}

// CentralAngle returns the value of property "XRCylinderLayer.centralAngle".
//
// The returned bool will be false if there is no such property.
func (this XRCylinderLayer) CentralAngle() (float32, bool) {
	var _ok bool
	_ret := bindings.GetXRCylinderLayerCentralAngle(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// CentralAngle sets the value of property "XRCylinderLayer.centralAngle" to val.
//
// It returns false if the property cannot be set.
func (this XRCylinderLayer) SetCentralAngle(val float32) bool {
	return js.True == bindings.SetXRCylinderLayerCentralAngle(
		this.Ref(),
		float32(val),
	)
}

// AspectRatio returns the value of property "XRCylinderLayer.aspectRatio".
//
// The returned bool will be false if there is no such property.
func (this XRCylinderLayer) AspectRatio() (float32, bool) {
	var _ok bool
	_ret := bindings.GetXRCylinderLayerAspectRatio(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// AspectRatio sets the value of property "XRCylinderLayer.aspectRatio" to val.
//
// It returns false if the property cannot be set.
func (this XRCylinderLayer) SetAspectRatio(val float32) bool {
	return js.True == bindings.SetXRCylinderLayerAspectRatio(
		this.Ref(),
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
	Radius float32
	// CentralAngle is "XRCylinderLayerInit.centralAngle"
	//
	// Optional, defaults to 0.78539.
	CentralAngle float32
	// AspectRatio is "XRCylinderLayerInit.aspectRatio"
	//
	// Optional, defaults to 2.0.
	AspectRatio float32
	// Space is "XRCylinderLayerInit.space"
	//
	// Required
	Space XRSpace
	// ColorFormat is "XRCylinderLayerInit.colorFormat"
	//
	// Optional, defaults to 0x1908.
	ColorFormat GLenum
	// DepthFormat is "XRCylinderLayerInit.depthFormat"
	//
	// Optional
	DepthFormat GLenum
	// MipLevels is "XRCylinderLayerInit.mipLevels"
	//
	// Optional, defaults to 1.
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
	IsStatic bool
	// ClearOnAccess is "XRCylinderLayerInit.clearOnAccess"
	//
	// Optional, defaults to true.
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

// New creates a new {0x140004cc0e0 XRCylinderLayerInit XRCylinderLayerInit [// XRCylinderLayerInit] [0x14000aef0e0 0x14000aef180 0x14000aef220 0x14000aef360 0x14000aef4a0 0x14000aef5e0 0x14000aef680 0x14000aef7c0 0x14000aef900 0x14000aefa40 0x14000aefae0 0x14000aefb80 0x14000aefc20 0x14000aefd60 0x14000aef2c0 0x14000aef400 0x14000aef540 0x14000aef720 0x14000aef860 0x14000aef9a0 0x14000aefcc0 0x14000aefe00] 0x14000ad4f00 {0 0}} in the application heap.
func (p XRCylinderLayerInit) New() js.Ref {
	return bindings.XRCylinderLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRCylinderLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRCylinderLayerInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRCylinderLayerInit) Update(ref js.Ref) {
	bindings.XRCylinderLayerInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type XRDepthInformation struct {
	ref js.Ref
}

func (this XRDepthInformation) Once() XRDepthInformation {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Width returns the value of property "XRDepthInformation.width".
//
// The returned bool will be false if there is no such property.
func (this XRDepthInformation) Width() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRDepthInformationWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Height returns the value of property "XRDepthInformation.height".
//
// The returned bool will be false if there is no such property.
func (this XRDepthInformation) Height() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRDepthInformationHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// NormDepthBufferFromNormView returns the value of property "XRDepthInformation.normDepthBufferFromNormView".
//
// The returned bool will be false if there is no such property.
func (this XRDepthInformation) NormDepthBufferFromNormView() (XRRigidTransform, bool) {
	var _ok bool
	_ret := bindings.GetXRDepthInformationNormDepthBufferFromNormView(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRRigidTransform{}.FromRef(_ret), _ok
}

// RawValueToMeters returns the value of property "XRDepthInformation.rawValueToMeters".
//
// The returned bool will be false if there is no such property.
func (this XRDepthInformation) RawValueToMeters() (float32, bool) {
	var _ok bool
	_ret := bindings.GetXRDepthInformationRawValueToMeters(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

type XREquirectLayer struct {
	XRCompositionLayer
}

func (this XREquirectLayer) Once() XREquirectLayer {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Space returns the value of property "XREquirectLayer.space".
//
// The returned bool will be false if there is no such property.
func (this XREquirectLayer) Space() (XRSpace, bool) {
	var _ok bool
	_ret := bindings.GetXREquirectLayerSpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRSpace{}.FromRef(_ret), _ok
}

// Space sets the value of property "XREquirectLayer.space" to val.
//
// It returns false if the property cannot be set.
func (this XREquirectLayer) SetSpace(val XRSpace) bool {
	return js.True == bindings.SetXREquirectLayerSpace(
		this.Ref(),
		val.Ref(),
	)
}

// Transform returns the value of property "XREquirectLayer.transform".
//
// The returned bool will be false if there is no such property.
func (this XREquirectLayer) Transform() (XRRigidTransform, bool) {
	var _ok bool
	_ret := bindings.GetXREquirectLayerTransform(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRRigidTransform{}.FromRef(_ret), _ok
}

// Transform sets the value of property "XREquirectLayer.transform" to val.
//
// It returns false if the property cannot be set.
func (this XREquirectLayer) SetTransform(val XRRigidTransform) bool {
	return js.True == bindings.SetXREquirectLayerTransform(
		this.Ref(),
		val.Ref(),
	)
}

// Radius returns the value of property "XREquirectLayer.radius".
//
// The returned bool will be false if there is no such property.
func (this XREquirectLayer) Radius() (float32, bool) {
	var _ok bool
	_ret := bindings.GetXREquirectLayerRadius(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Radius sets the value of property "XREquirectLayer.radius" to val.
//
// It returns false if the property cannot be set.
func (this XREquirectLayer) SetRadius(val float32) bool {
	return js.True == bindings.SetXREquirectLayerRadius(
		this.Ref(),
		float32(val),
	)
}

// CentralHorizontalAngle returns the value of property "XREquirectLayer.centralHorizontalAngle".
//
// The returned bool will be false if there is no such property.
func (this XREquirectLayer) CentralHorizontalAngle() (float32, bool) {
	var _ok bool
	_ret := bindings.GetXREquirectLayerCentralHorizontalAngle(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// CentralHorizontalAngle sets the value of property "XREquirectLayer.centralHorizontalAngle" to val.
//
// It returns false if the property cannot be set.
func (this XREquirectLayer) SetCentralHorizontalAngle(val float32) bool {
	return js.True == bindings.SetXREquirectLayerCentralHorizontalAngle(
		this.Ref(),
		float32(val),
	)
}

// UpperVerticalAngle returns the value of property "XREquirectLayer.upperVerticalAngle".
//
// The returned bool will be false if there is no such property.
func (this XREquirectLayer) UpperVerticalAngle() (float32, bool) {
	var _ok bool
	_ret := bindings.GetXREquirectLayerUpperVerticalAngle(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// UpperVerticalAngle sets the value of property "XREquirectLayer.upperVerticalAngle" to val.
//
// It returns false if the property cannot be set.
func (this XREquirectLayer) SetUpperVerticalAngle(val float32) bool {
	return js.True == bindings.SetXREquirectLayerUpperVerticalAngle(
		this.Ref(),
		float32(val),
	)
}

// LowerVerticalAngle returns the value of property "XREquirectLayer.lowerVerticalAngle".
//
// The returned bool will be false if there is no such property.
func (this XREquirectLayer) LowerVerticalAngle() (float32, bool) {
	var _ok bool
	_ret := bindings.GetXREquirectLayerLowerVerticalAngle(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// LowerVerticalAngle sets the value of property "XREquirectLayer.lowerVerticalAngle" to val.
//
// It returns false if the property cannot be set.
func (this XREquirectLayer) SetLowerVerticalAngle(val float32) bool {
	return js.True == bindings.SetXREquirectLayerLowerVerticalAngle(
		this.Ref(),
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
	Radius float32
	// CentralHorizontalAngle is "XREquirectLayerInit.centralHorizontalAngle"
	//
	// Optional, defaults to 6.28318.
	CentralHorizontalAngle float32
	// UpperVerticalAngle is "XREquirectLayerInit.upperVerticalAngle"
	//
	// Optional, defaults to 1.570795.
	UpperVerticalAngle float32
	// LowerVerticalAngle is "XREquirectLayerInit.lowerVerticalAngle"
	//
	// Optional, defaults to -1.570795.
	LowerVerticalAngle float32
	// Space is "XREquirectLayerInit.space"
	//
	// Required
	Space XRSpace
	// ColorFormat is "XREquirectLayerInit.colorFormat"
	//
	// Optional, defaults to 0x1908.
	ColorFormat GLenum
	// DepthFormat is "XREquirectLayerInit.depthFormat"
	//
	// Optional
	DepthFormat GLenum
	// MipLevels is "XREquirectLayerInit.mipLevels"
	//
	// Optional, defaults to 1.
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
	IsStatic bool
	// ClearOnAccess is "XREquirectLayerInit.clearOnAccess"
	//
	// Optional, defaults to true.
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

// New creates a new {0x140004cc0e0 XREquirectLayerInit XREquirectLayerInit [// XREquirectLayerInit] [0x14000b0c000 0x14000b0c0a0 0x14000b0c140 0x14000b0c280 0x14000b0c3c0 0x14000b0c500 0x14000b0c640 0x14000b0c6e0 0x14000b0c820 0x14000b0c960 0x14000b0caa0 0x14000b0cb40 0x14000b0cbe0 0x14000b0cc80 0x14000b0cdc0 0x14000b0c1e0 0x14000b0c320 0x14000b0c460 0x14000b0c5a0 0x14000b0c780 0x14000b0c8c0 0x14000b0ca00 0x14000b0cd20 0x14000b0ce60] 0x14000ad4ff0 {0 0}} in the application heap.
func (p XREquirectLayerInit) New() js.Ref {
	return bindings.XREquirectLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XREquirectLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XREquirectLayerInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XREquirectLayerInit) Update(ref js.Ref) {
	bindings.XREquirectLayerInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	Bubbles bool
	// Cancelable is "XRInputSourceEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "XRInputSourceEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 XRInputSourceEventInit XRInputSourceEventInit [// XRInputSourceEventInit] [0x14000b0cf00 0x14000b0cfa0 0x14000b0d040 0x14000b0d180 0x14000b0d2c0 0x14000b0d0e0 0x14000b0d220 0x14000b0d360] 0x14000ad5068 {0 0}} in the application heap.
func (p XRInputSourceEventInit) New() js.Ref {
	return bindings.XRInputSourceEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRInputSourceEventInit) UpdateFrom(ref js.Ref) {
	bindings.XRInputSourceEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRInputSourceEventInit) Update(ref js.Ref) {
	bindings.XRInputSourceEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewXRInputSourceEvent(typ js.String, eventInitDict XRInputSourceEventInit) XRInputSourceEvent {
	return XRInputSourceEvent{}.FromRef(
		bindings.NewXRInputSourceEventByXRInputSourceEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type XRInputSourceEvent struct {
	Event
}

func (this XRInputSourceEvent) Once() XRInputSourceEvent {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Frame returns the value of property "XRInputSourceEvent.frame".
//
// The returned bool will be false if there is no such property.
func (this XRInputSourceEvent) Frame() (XRFrame, bool) {
	var _ok bool
	_ret := bindings.GetXRInputSourceEventFrame(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRFrame{}.FromRef(_ret), _ok
}

// InputSource returns the value of property "XRInputSourceEvent.inputSource".
//
// The returned bool will be false if there is no such property.
func (this XRInputSourceEvent) InputSource() (XRInputSource, bool) {
	var _ok bool
	_ret := bindings.GetXRInputSourceEventInputSource(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRInputSource{}.FromRef(_ret), _ok
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
	Bubbles bool
	// Cancelable is "XRInputSourcesChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "XRInputSourcesChangeEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 XRInputSourcesChangeEventInit XRInputSourcesChangeEventInit [// XRInputSourcesChangeEventInit] [0x14000b0d4a0 0x14000b0d540 0x14000b0d5e0 0x14000b0d680 0x14000b0d7c0 0x14000b0d900 0x14000b0d720 0x14000b0d860 0x14000b0d9a0] 0x14000ad5098 {0 0}} in the application heap.
func (p XRInputSourcesChangeEventInit) New() js.Ref {
	return bindings.XRInputSourcesChangeEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRInputSourcesChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.XRInputSourcesChangeEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRInputSourcesChangeEventInit) Update(ref js.Ref) {
	bindings.XRInputSourcesChangeEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewXRInputSourcesChangeEvent(typ js.String, eventInitDict XRInputSourcesChangeEventInit) XRInputSourcesChangeEvent {
	return XRInputSourcesChangeEvent{}.FromRef(
		bindings.NewXRInputSourcesChangeEventByXRInputSourcesChangeEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type XRInputSourcesChangeEvent struct {
	Event
}

func (this XRInputSourcesChangeEvent) Once() XRInputSourcesChangeEvent {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Session returns the value of property "XRInputSourcesChangeEvent.session".
//
// The returned bool will be false if there is no such property.
func (this XRInputSourcesChangeEvent) Session() (XRSession, bool) {
	var _ok bool
	_ret := bindings.GetXRInputSourcesChangeEventSession(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRSession{}.FromRef(_ret), _ok
}

// Added returns the value of property "XRInputSourcesChangeEvent.added".
//
// The returned bool will be false if there is no such property.
func (this XRInputSourcesChangeEvent) Added() (js.FrozenArray[XRInputSource], bool) {
	var _ok bool
	_ret := bindings.GetXRInputSourcesChangeEventAdded(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[XRInputSource]{}.FromRef(_ret), _ok
}

// Removed returns the value of property "XRInputSourcesChangeEvent.removed".
//
// The returned bool will be false if there is no such property.
func (this XRInputSourcesChangeEvent) Removed() (js.FrozenArray[XRInputSource], bool) {
	var _ok bool
	_ret := bindings.GetXRInputSourcesChangeEventRemoved(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[XRInputSource]{}.FromRef(_ret), _ok
}

type XRLayerEventInit struct {
	// Layer is "XRLayerEventInit.layer"
	//
	// Required
	Layer XRLayer
	// Bubbles is "XRLayerEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "XRLayerEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "XRLayerEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 XRLayerEventInit XRLayerEventInit [// XRLayerEventInit] [0x14000b0dae0 0x14000b0db80 0x14000b0dcc0 0x14000b0de00 0x14000b0dc20 0x14000b0dd60 0x14000b0dea0] 0x14000ad50e0 {0 0}} in the application heap.
func (p XRLayerEventInit) New() js.Ref {
	return bindings.XRLayerEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRLayerEventInit) UpdateFrom(ref js.Ref) {
	bindings.XRLayerEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRLayerEventInit) Update(ref js.Ref) {
	bindings.XRLayerEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewXRLayerEvent(typ js.String, eventInitDict XRLayerEventInit) XRLayerEvent {
	return XRLayerEvent{}.FromRef(
		bindings.NewXRLayerEventByXRLayerEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type XRLayerEvent struct {
	Event
}

func (this XRLayerEvent) Once() XRLayerEvent {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Layer returns the value of property "XRLayerEvent.layer".
//
// The returned bool will be false if there is no such property.
func (this XRLayerEvent) Layer() (XRLayer, bool) {
	var _ok bool
	_ret := bindings.GetXRLayerEventLayer(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRLayer{}.FromRef(_ret), _ok
}

type XRLayerInit struct {
	// Space is "XRLayerInit.space"
	//
	// Required
	Space XRSpace
	// ColorFormat is "XRLayerInit.colorFormat"
	//
	// Optional, defaults to 0x1908.
	ColorFormat GLenum
	// DepthFormat is "XRLayerInit.depthFormat"
	//
	// Optional
	DepthFormat GLenum
	// MipLevels is "XRLayerInit.mipLevels"
	//
	// Optional, defaults to 1.
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
	IsStatic bool
	// ClearOnAccess is "XRLayerInit.clearOnAccess"
	//
	// Optional, defaults to true.
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

// New creates a new {0x140004cc0e0 XRLayerInit XRLayerInit [// XRLayerInit] [0x14000b0df40 0x14000b0e000 0x14000b0e140 0x14000b0e280 0x14000b0e3c0 0x14000b0e460 0x14000b0e500 0x14000b0e5a0 0x14000b0e6e0 0x14000b0e0a0 0x14000b0e1e0 0x14000b0e320 0x14000b0e640 0x14000b0e780] 0x14000ad5110 {0 0}} in the application heap.
func (p XRLayerInit) New() js.Ref {
	return bindings.XRLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRLayerInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRLayerInit) Update(ref js.Ref) {
	bindings.XRLayerInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type XRQuadLayer struct {
	XRCompositionLayer
}

func (this XRQuadLayer) Once() XRQuadLayer {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Space returns the value of property "XRQuadLayer.space".
//
// The returned bool will be false if there is no such property.
func (this XRQuadLayer) Space() (XRSpace, bool) {
	var _ok bool
	_ret := bindings.GetXRQuadLayerSpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRSpace{}.FromRef(_ret), _ok
}

// Space sets the value of property "XRQuadLayer.space" to val.
//
// It returns false if the property cannot be set.
func (this XRQuadLayer) SetSpace(val XRSpace) bool {
	return js.True == bindings.SetXRQuadLayerSpace(
		this.Ref(),
		val.Ref(),
	)
}

// Transform returns the value of property "XRQuadLayer.transform".
//
// The returned bool will be false if there is no such property.
func (this XRQuadLayer) Transform() (XRRigidTransform, bool) {
	var _ok bool
	_ret := bindings.GetXRQuadLayerTransform(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRRigidTransform{}.FromRef(_ret), _ok
}

// Transform sets the value of property "XRQuadLayer.transform" to val.
//
// It returns false if the property cannot be set.
func (this XRQuadLayer) SetTransform(val XRRigidTransform) bool {
	return js.True == bindings.SetXRQuadLayerTransform(
		this.Ref(),
		val.Ref(),
	)
}

// Width returns the value of property "XRQuadLayer.width".
//
// The returned bool will be false if there is no such property.
func (this XRQuadLayer) Width() (float32, bool) {
	var _ok bool
	_ret := bindings.GetXRQuadLayerWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Width sets the value of property "XRQuadLayer.width" to val.
//
// It returns false if the property cannot be set.
func (this XRQuadLayer) SetWidth(val float32) bool {
	return js.True == bindings.SetXRQuadLayerWidth(
		this.Ref(),
		float32(val),
	)
}

// Height returns the value of property "XRQuadLayer.height".
//
// The returned bool will be false if there is no such property.
func (this XRQuadLayer) Height() (float32, bool) {
	var _ok bool
	_ret := bindings.GetXRQuadLayerHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Height sets the value of property "XRQuadLayer.height" to val.
//
// It returns false if the property cannot be set.
func (this XRQuadLayer) SetHeight(val float32) bool {
	return js.True == bindings.SetXRQuadLayerHeight(
		this.Ref(),
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
	Width float32
	// Height is "XRMediaQuadLayerInit.height"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 XRMediaQuadLayerInit XRMediaQuadLayerInit [// XRMediaQuadLayerInit] [0x14000b0e8c0 0x14000b0e960 0x14000b0eaa0 0x14000b0ebe0 0x14000b0ec80 0x14000b0ed20 0x14000b0ea00 0x14000b0eb40 0x14000b0edc0] 0x14000ad51b8 {0 0}} in the application heap.
func (p XRMediaQuadLayerInit) New() js.Ref {
	return bindings.XRMediaQuadLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRMediaQuadLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRMediaQuadLayerInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRMediaQuadLayerInit) Update(ref js.Ref) {
	bindings.XRMediaQuadLayerInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type XRMediaCylinderLayerInit struct {
	// Transform is "XRMediaCylinderLayerInit.transform"
	//
	// Optional
	Transform XRRigidTransform
	// Radius is "XRMediaCylinderLayerInit.radius"
	//
	// Optional, defaults to 2.0.
	Radius float32
	// CentralAngle is "XRMediaCylinderLayerInit.centralAngle"
	//
	// Optional, defaults to 0.78539.
	CentralAngle float32
	// AspectRatio is "XRMediaCylinderLayerInit.aspectRatio"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 XRMediaCylinderLayerInit XRMediaCylinderLayerInit [// XRMediaCylinderLayerInit] [0x14000b0ee60 0x14000b0ef00 0x14000b0f040 0x14000b0f180 0x14000b0f2c0 0x14000b0f360 0x14000b0f400 0x14000b0efa0 0x14000b0f0e0 0x14000b0f220 0x14000b0f4a0] 0x14000ad5200 {0 0}} in the application heap.
func (p XRMediaCylinderLayerInit) New() js.Ref {
	return bindings.XRMediaCylinderLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRMediaCylinderLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRMediaCylinderLayerInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRMediaCylinderLayerInit) Update(ref js.Ref) {
	bindings.XRMediaCylinderLayerInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type XRMediaEquirectLayerInit struct {
	// Transform is "XRMediaEquirectLayerInit.transform"
	//
	// Optional
	Transform XRRigidTransform
	// Radius is "XRMediaEquirectLayerInit.radius"
	//
	// Optional, defaults to 0.0.
	Radius float32
	// CentralHorizontalAngle is "XRMediaEquirectLayerInit.centralHorizontalAngle"
	//
	// Optional, defaults to 6.28318.
	CentralHorizontalAngle float32
	// UpperVerticalAngle is "XRMediaEquirectLayerInit.upperVerticalAngle"
	//
	// Optional, defaults to 1.570795.
	UpperVerticalAngle float32
	// LowerVerticalAngle is "XRMediaEquirectLayerInit.lowerVerticalAngle"
	//
	// Optional, defaults to -1.570795.
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

// New creates a new {0x140004cc0e0 XRMediaEquirectLayerInit XRMediaEquirectLayerInit [// XRMediaEquirectLayerInit] [0x14000b0f540 0x14000b0f5e0 0x14000b0f720 0x14000b0f860 0x14000b0f9a0 0x14000b0fae0 0x14000b0fb80 0x14000b0fc20 0x14000b0f680 0x14000b0f7c0 0x14000b0f900 0x14000b0fa40 0x14000b0fcc0] 0x14000ad5290 {0 0}} in the application heap.
func (p XRMediaEquirectLayerInit) New() js.Ref {
	return bindings.XRMediaEquirectLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRMediaEquirectLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRMediaEquirectLayerInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRMediaEquirectLayerInit) Update(ref js.Ref) {
	bindings.XRMediaEquirectLayerInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewXRMediaBinding(session XRSession) XRMediaBinding {
	return XRMediaBinding{}.FromRef(
		bindings.NewXRMediaBindingByXRMediaBinding(
			session.Ref()),
	)
}

type XRMediaBinding struct {
	ref js.Ref
}

func (this XRMediaBinding) Once() XRMediaBinding {
	this.Ref().Once()
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
	this.Ref().Free()
}

// CreateQuadLayer calls the method "XRMediaBinding.createQuadLayer".
//
// The returned bool will be false if there is no such method.
func (this XRMediaBinding) CreateQuadLayer(video HTMLVideoElement, init XRMediaQuadLayerInit) (XRQuadLayer, bool) {
	var _ok bool
	_ret := bindings.CallXRMediaBindingCreateQuadLayer(
		this.Ref(), js.Pointer(&_ok),
		video.Ref(),
		js.Pointer(&init),
	)

	return XRQuadLayer{}.FromRef(_ret), _ok
}

// CreateQuadLayerFunc returns the method "XRMediaBinding.createQuadLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRMediaBinding) CreateQuadLayerFunc() (fn js.Func[func(video HTMLVideoElement, init XRMediaQuadLayerInit) XRQuadLayer]) {
	return fn.FromRef(
		bindings.XRMediaBindingCreateQuadLayerFunc(
			this.Ref(),
		),
	)
}

// CreateQuadLayer1 calls the method "XRMediaBinding.createQuadLayer".
//
// The returned bool will be false if there is no such method.
func (this XRMediaBinding) CreateQuadLayer1(video HTMLVideoElement) (XRQuadLayer, bool) {
	var _ok bool
	_ret := bindings.CallXRMediaBindingCreateQuadLayer1(
		this.Ref(), js.Pointer(&_ok),
		video.Ref(),
	)

	return XRQuadLayer{}.FromRef(_ret), _ok
}

// CreateQuadLayer1Func returns the method "XRMediaBinding.createQuadLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRMediaBinding) CreateQuadLayer1Func() (fn js.Func[func(video HTMLVideoElement) XRQuadLayer]) {
	return fn.FromRef(
		bindings.XRMediaBindingCreateQuadLayer1Func(
			this.Ref(),
		),
	)
}

// CreateCylinderLayer calls the method "XRMediaBinding.createCylinderLayer".
//
// The returned bool will be false if there is no such method.
func (this XRMediaBinding) CreateCylinderLayer(video HTMLVideoElement, init XRMediaCylinderLayerInit) (XRCylinderLayer, bool) {
	var _ok bool
	_ret := bindings.CallXRMediaBindingCreateCylinderLayer(
		this.Ref(), js.Pointer(&_ok),
		video.Ref(),
		js.Pointer(&init),
	)

	return XRCylinderLayer{}.FromRef(_ret), _ok
}

// CreateCylinderLayerFunc returns the method "XRMediaBinding.createCylinderLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRMediaBinding) CreateCylinderLayerFunc() (fn js.Func[func(video HTMLVideoElement, init XRMediaCylinderLayerInit) XRCylinderLayer]) {
	return fn.FromRef(
		bindings.XRMediaBindingCreateCylinderLayerFunc(
			this.Ref(),
		),
	)
}

// CreateCylinderLayer1 calls the method "XRMediaBinding.createCylinderLayer".
//
// The returned bool will be false if there is no such method.
func (this XRMediaBinding) CreateCylinderLayer1(video HTMLVideoElement) (XRCylinderLayer, bool) {
	var _ok bool
	_ret := bindings.CallXRMediaBindingCreateCylinderLayer1(
		this.Ref(), js.Pointer(&_ok),
		video.Ref(),
	)

	return XRCylinderLayer{}.FromRef(_ret), _ok
}

// CreateCylinderLayer1Func returns the method "XRMediaBinding.createCylinderLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRMediaBinding) CreateCylinderLayer1Func() (fn js.Func[func(video HTMLVideoElement) XRCylinderLayer]) {
	return fn.FromRef(
		bindings.XRMediaBindingCreateCylinderLayer1Func(
			this.Ref(),
		),
	)
}

// CreateEquirectLayer calls the method "XRMediaBinding.createEquirectLayer".
//
// The returned bool will be false if there is no such method.
func (this XRMediaBinding) CreateEquirectLayer(video HTMLVideoElement, init XRMediaEquirectLayerInit) (XREquirectLayer, bool) {
	var _ok bool
	_ret := bindings.CallXRMediaBindingCreateEquirectLayer(
		this.Ref(), js.Pointer(&_ok),
		video.Ref(),
		js.Pointer(&init),
	)

	return XREquirectLayer{}.FromRef(_ret), _ok
}

// CreateEquirectLayerFunc returns the method "XRMediaBinding.createEquirectLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRMediaBinding) CreateEquirectLayerFunc() (fn js.Func[func(video HTMLVideoElement, init XRMediaEquirectLayerInit) XREquirectLayer]) {
	return fn.FromRef(
		bindings.XRMediaBindingCreateEquirectLayerFunc(
			this.Ref(),
		),
	)
}

// CreateEquirectLayer1 calls the method "XRMediaBinding.createEquirectLayer".
//
// The returned bool will be false if there is no such method.
func (this XRMediaBinding) CreateEquirectLayer1(video HTMLVideoElement) (XREquirectLayer, bool) {
	var _ok bool
	_ret := bindings.CallXRMediaBindingCreateEquirectLayer1(
		this.Ref(), js.Pointer(&_ok),
		video.Ref(),
	)

	return XREquirectLayer{}.FromRef(_ret), _ok
}

// CreateEquirectLayer1Func returns the method "XRMediaBinding.createEquirectLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRMediaBinding) CreateEquirectLayer1Func() (fn js.Func[func(video HTMLVideoElement) XREquirectLayer]) {
	return fn.FromRef(
		bindings.XRMediaBindingCreateEquirectLayer1Func(
			this.Ref(),
		),
	)
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
	InvertStereo bool

	FFI_USE_InvertStereo bool // for InvertStereo.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a XRMediaLayerInit with all fields set.
func (p XRMediaLayerInit) FromRef(ref js.Ref) XRMediaLayerInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 XRMediaLayerInit XRMediaLayerInit [// XRMediaLayerInit] [0x14000b0fd60 0x14000b0fe00 0x14000b0fea0 0x14000b0ff40] 0x14000ad52f0 {0 0}} in the application heap.
func (p XRMediaLayerInit) New() js.Ref {
	return bindings.XRMediaLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRMediaLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRMediaLayerInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRMediaLayerInit) Update(ref js.Ref) {
	bindings.XRMediaLayerInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type XRMesh struct {
	ref js.Ref
}

func (this XRMesh) Once() XRMesh {
	this.Ref().Once()
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
	this.Ref().Free()
}

// MeshSpace returns the value of property "XRMesh.meshSpace".
//
// The returned bool will be false if there is no such property.
func (this XRMesh) MeshSpace() (XRSpace, bool) {
	var _ok bool
	_ret := bindings.GetXRMeshMeshSpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRSpace{}.FromRef(_ret), _ok
}

// Vertices returns the value of property "XRMesh.vertices".
//
// The returned bool will be false if there is no such property.
func (this XRMesh) Vertices() (js.FrozenArray[js.TypedArray[float32]], bool) {
	var _ok bool
	_ret := bindings.GetXRMeshVertices(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[js.TypedArray[float32]]{}.FromRef(_ret), _ok
}

// Indices returns the value of property "XRMesh.indices".
//
// The returned bool will be false if there is no such property.
func (this XRMesh) Indices() (js.TypedArray[uint32], bool) {
	var _ok bool
	_ret := bindings.GetXRMeshIndices(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[uint32]{}.FromRef(_ret), _ok
}

// LastChangedTime returns the value of property "XRMesh.lastChangedTime".
//
// The returned bool will be false if there is no such property.
func (this XRMesh) LastChangedTime() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetXRMeshLastChangedTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// SemanticLabel returns the value of property "XRMesh.semanticLabel".
//
// The returned bool will be false if there is no such property.
func (this XRMesh) SemanticLabel() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetXRMeshSemanticLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
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

// New creates a new {0x140004cc0e0 XRPermissionDescriptor XRPermissionDescriptor [// XRPermissionDescriptor] [0x14000b160a0 0x14000b16140 0x14000b161e0 0x14000b16280] 0x14000ad5320 {0 0}} in the application heap.
func (p XRPermissionDescriptor) New() js.Ref {
	return bindings.XRPermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.XRPermissionDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRPermissionDescriptor) Update(ref js.Ref) {
	bindings.XRPermissionDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type XRPermissionStatus struct {
	PermissionStatus
}

func (this XRPermissionStatus) Once() XRPermissionStatus {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Granted returns the value of property "XRPermissionStatus.granted".
//
// The returned bool will be false if there is no such property.
func (this XRPermissionStatus) Granted() (js.FrozenArray[js.String], bool) {
	var _ok bool
	_ret := bindings.GetXRPermissionStatusGranted(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[js.String]{}.FromRef(_ret), _ok
}

// Granted sets the value of property "XRPermissionStatus.granted" to val.
//
// It returns false if the property cannot be set.
func (this XRPermissionStatus) SetGranted(val js.FrozenArray[js.String]) bool {
	return js.True == bindings.SetXRPermissionStatusGranted(
		this.Ref(),
		val.Ref(),
	)
}

type XRProjectionLayer struct {
	XRCompositionLayer
}

func (this XRProjectionLayer) Once() XRProjectionLayer {
	this.Ref().Once()
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
	this.Ref().Free()
}

// TextureWidth returns the value of property "XRProjectionLayer.textureWidth".
//
// The returned bool will be false if there is no such property.
func (this XRProjectionLayer) TextureWidth() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRProjectionLayerTextureWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// TextureHeight returns the value of property "XRProjectionLayer.textureHeight".
//
// The returned bool will be false if there is no such property.
func (this XRProjectionLayer) TextureHeight() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRProjectionLayerTextureHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// TextureArrayLength returns the value of property "XRProjectionLayer.textureArrayLength".
//
// The returned bool will be false if there is no such property.
func (this XRProjectionLayer) TextureArrayLength() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRProjectionLayerTextureArrayLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// IgnoreDepthValues returns the value of property "XRProjectionLayer.ignoreDepthValues".
//
// The returned bool will be false if there is no such property.
func (this XRProjectionLayer) IgnoreDepthValues() (bool, bool) {
	var _ok bool
	_ret := bindings.GetXRProjectionLayerIgnoreDepthValues(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// FixedFoveation returns the value of property "XRProjectionLayer.fixedFoveation".
//
// The returned bool will be false if there is no such property.
func (this XRProjectionLayer) FixedFoveation() (float32, bool) {
	var _ok bool
	_ret := bindings.GetXRProjectionLayerFixedFoveation(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// FixedFoveation sets the value of property "XRProjectionLayer.fixedFoveation" to val.
//
// It returns false if the property cannot be set.
func (this XRProjectionLayer) SetFixedFoveation(val float32) bool {
	return js.True == bindings.SetXRProjectionLayerFixedFoveation(
		this.Ref(),
		float32(val),
	)
}

// DeltaPose returns the value of property "XRProjectionLayer.deltaPose".
//
// The returned bool will be false if there is no such property.
func (this XRProjectionLayer) DeltaPose() (XRRigidTransform, bool) {
	var _ok bool
	_ret := bindings.GetXRProjectionLayerDeltaPose(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRRigidTransform{}.FromRef(_ret), _ok
}

// DeltaPose sets the value of property "XRProjectionLayer.deltaPose" to val.
//
// It returns false if the property cannot be set.
func (this XRProjectionLayer) SetDeltaPose(val XRRigidTransform) bool {
	return js.True == bindings.SetXRProjectionLayerDeltaPose(
		this.Ref(),
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
	ColorFormat GLenum
	// DepthFormat is "XRProjectionLayerInit.depthFormat"
	//
	// Optional, defaults to 0x1902.
	DepthFormat GLenum
	// ScaleFactor is "XRProjectionLayerInit.scaleFactor"
	//
	// Optional, defaults to 1.0.
	ScaleFactor float64
	// ClearOnAccess is "XRProjectionLayerInit.clearOnAccess"
	//
	// Optional, defaults to true.
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

// New creates a new {0x140004cc0e0 XRProjectionLayerInit XRProjectionLayerInit [// XRProjectionLayerInit] [0x14000b163c0 0x14000b16460 0x14000b165a0 0x14000b166e0 0x14000b16820 0x14000b16500 0x14000b16640 0x14000b16780 0x14000b168c0] 0x14000ad5380 {0 0}} in the application heap.
func (p XRProjectionLayerInit) New() js.Ref {
	return bindings.XRProjectionLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRProjectionLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRProjectionLayerInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRProjectionLayerInit) Update(ref js.Ref) {
	bindings.XRProjectionLayerInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	Width float32
	// Height is "XRQuadLayerInit.height"
	//
	// Optional, defaults to 1.0.
	Height float32
	// Space is "XRQuadLayerInit.space"
	//
	// Required
	Space XRSpace
	// ColorFormat is "XRQuadLayerInit.colorFormat"
	//
	// Optional, defaults to 0x1908.
	ColorFormat GLenum
	// DepthFormat is "XRQuadLayerInit.depthFormat"
	//
	// Optional
	DepthFormat GLenum
	// MipLevels is "XRQuadLayerInit.mipLevels"
	//
	// Optional, defaults to 1.
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
	IsStatic bool
	// ClearOnAccess is "XRQuadLayerInit.clearOnAccess"
	//
	// Optional, defaults to true.
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

// New creates a new {0x140004cc0e0 XRQuadLayerInit XRQuadLayerInit [// XRQuadLayerInit] [0x14000b16960 0x14000b16a00 0x14000b16aa0 0x14000b16be0 0x14000b16d20 0x14000b16dc0 0x14000b16f00 0x14000b17040 0x14000b17180 0x14000b17220 0x14000b172c0 0x14000b17360 0x14000b174a0 0x14000b16b40 0x14000b16c80 0x14000b16e60 0x14000b16fa0 0x14000b170e0 0x14000b17400 0x14000b17540] 0x14000ad5410 {0 0}} in the application heap.
func (p XRQuadLayerInit) New() js.Ref {
	return bindings.XRQuadLayerInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRQuadLayerInit) UpdateFrom(ref js.Ref) {
	bindings.XRQuadLayerInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRQuadLayerInit) Update(ref js.Ref) {
	bindings.XRQuadLayerInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	Bubbles bool
	// Cancelable is "XRReferenceSpaceEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "XRReferenceSpaceEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 XRReferenceSpaceEventInit XRReferenceSpaceEventInit [// XRReferenceSpaceEventInit] [0x14000b175e0 0x14000b17680 0x14000b17720 0x14000b17860 0x14000b179a0 0x14000b177c0 0x14000b17900 0x14000b17a40] 0x14000ad5488 {0 0}} in the application heap.
func (p XRReferenceSpaceEventInit) New() js.Ref {
	return bindings.XRReferenceSpaceEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRReferenceSpaceEventInit) UpdateFrom(ref js.Ref) {
	bindings.XRReferenceSpaceEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRReferenceSpaceEventInit) Update(ref js.Ref) {
	bindings.XRReferenceSpaceEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewXRReferenceSpaceEvent(typ js.String, eventInitDict XRReferenceSpaceEventInit) XRReferenceSpaceEvent {
	return XRReferenceSpaceEvent{}.FromRef(
		bindings.NewXRReferenceSpaceEventByXRReferenceSpaceEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type XRReferenceSpaceEvent struct {
	Event
}

func (this XRReferenceSpaceEvent) Once() XRReferenceSpaceEvent {
	this.Ref().Once()
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
	this.Ref().Free()
}

// ReferenceSpace returns the value of property "XRReferenceSpaceEvent.referenceSpace".
//
// The returned bool will be false if there is no such property.
func (this XRReferenceSpaceEvent) ReferenceSpace() (XRReferenceSpace, bool) {
	var _ok bool
	_ret := bindings.GetXRReferenceSpaceEventReferenceSpace(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRReferenceSpace{}.FromRef(_ret), _ok
}

// Transform returns the value of property "XRReferenceSpaceEvent.transform".
//
// The returned bool will be false if there is no such property.
func (this XRReferenceSpaceEvent) Transform() (XRRigidTransform, bool) {
	var _ok bool
	_ret := bindings.GetXRReferenceSpaceEventTransform(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRRigidTransform{}.FromRef(_ret), _ok
}

type XRSessionEventInit struct {
	// Session is "XRSessionEventInit.session"
	//
	// Required
	Session XRSession
	// Bubbles is "XRSessionEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "XRSessionEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "XRSessionEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 XRSessionEventInit XRSessionEventInit [// XRSessionEventInit] [0x14000b17b80 0x14000b17c20 0x14000b17d60 0x14000b17ea0 0x14000b17cc0 0x14000b17e00 0x14000b17f40] 0x14000ad54d0 {0 0}} in the application heap.
func (p XRSessionEventInit) New() js.Ref {
	return bindings.XRSessionEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRSessionEventInit) UpdateFrom(ref js.Ref) {
	bindings.XRSessionEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRSessionEventInit) Update(ref js.Ref) {
	bindings.XRSessionEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewXRSessionEvent(typ js.String, eventInitDict XRSessionEventInit) XRSessionEvent {
	return XRSessionEvent{}.FromRef(
		bindings.NewXRSessionEventByXRSessionEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type XRSessionEvent struct {
	Event
}

func (this XRSessionEvent) Once() XRSessionEvent {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Session returns the value of property "XRSessionEvent.session".
//
// The returned bool will be false if there is no such property.
func (this XRSessionEvent) Session() (XRSession, bool) {
	var _ok bool
	_ret := bindings.GetXRSessionEventSession(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRSession{}.FromRef(_ret), _ok
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

// New creates a new {0x140004cc0e0 XRSessionSupportedPermissionDescriptor XRSessionSupportedPermissionDescriptor [// XRSessionSupportedPermissionDescriptor] [0x14000b18000 0x14000b180a0] 0x14000ad5500 {0 0}} in the application heap.
func (p XRSessionSupportedPermissionDescriptor) New() js.Ref {
	return bindings.XRSessionSupportedPermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p XRSessionSupportedPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.XRSessionSupportedPermissionDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p XRSessionSupportedPermissionDescriptor) Update(ref js.Ref) {
	bindings.XRSessionSupportedPermissionDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type XRSubImage struct {
	ref js.Ref
}

func (this XRSubImage) Once() XRSubImage {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Viewport returns the value of property "XRSubImage.viewport".
//
// The returned bool will be false if there is no such property.
func (this XRSubImage) Viewport() (XRViewport, bool) {
	var _ok bool
	_ret := bindings.GetXRSubImageViewport(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRViewport{}.FromRef(_ret), _ok
}

type XRWebGLSubImage struct {
	XRSubImage
}

func (this XRWebGLSubImage) Once() XRWebGLSubImage {
	this.Ref().Once()
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
	this.Ref().Free()
}

// ColorTexture returns the value of property "XRWebGLSubImage.colorTexture".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLSubImage) ColorTexture() (WebGLTexture, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLSubImageColorTexture(
		this.Ref(), js.Pointer(&_ok),
	)
	return WebGLTexture{}.FromRef(_ret), _ok
}

// DepthStencilTexture returns the value of property "XRWebGLSubImage.depthStencilTexture".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLSubImage) DepthStencilTexture() (WebGLTexture, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLSubImageDepthStencilTexture(
		this.Ref(), js.Pointer(&_ok),
	)
	return WebGLTexture{}.FromRef(_ret), _ok
}

// MotionVectorTexture returns the value of property "XRWebGLSubImage.motionVectorTexture".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLSubImage) MotionVectorTexture() (WebGLTexture, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLSubImageMotionVectorTexture(
		this.Ref(), js.Pointer(&_ok),
	)
	return WebGLTexture{}.FromRef(_ret), _ok
}

// ImageIndex returns the value of property "XRWebGLSubImage.imageIndex".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLSubImage) ImageIndex() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLSubImageImageIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// ColorTextureWidth returns the value of property "XRWebGLSubImage.colorTextureWidth".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLSubImage) ColorTextureWidth() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLSubImageColorTextureWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// ColorTextureHeight returns the value of property "XRWebGLSubImage.colorTextureHeight".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLSubImage) ColorTextureHeight() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLSubImageColorTextureHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// DepthStencilTextureWidth returns the value of property "XRWebGLSubImage.depthStencilTextureWidth".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLSubImage) DepthStencilTextureWidth() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLSubImageDepthStencilTextureWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// DepthStencilTextureHeight returns the value of property "XRWebGLSubImage.depthStencilTextureHeight".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLSubImage) DepthStencilTextureHeight() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLSubImageDepthStencilTextureHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MotionVectorTextureWidth returns the value of property "XRWebGLSubImage.motionVectorTextureWidth".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLSubImage) MotionVectorTextureWidth() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLSubImageMotionVectorTextureWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MotionVectorTextureHeight returns the value of property "XRWebGLSubImage.motionVectorTextureHeight".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLSubImage) MotionVectorTextureHeight() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLSubImageMotionVectorTextureHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

type XRWebGLDepthInformation struct {
	XRDepthInformation
}

func (this XRWebGLDepthInformation) Once() XRWebGLDepthInformation {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Texture returns the value of property "XRWebGLDepthInformation.texture".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLDepthInformation) Texture() (WebGLTexture, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLDepthInformationTexture(
		this.Ref(), js.Pointer(&_ok),
	)
	return WebGLTexture{}.FromRef(_ret), _ok
}

func NewXRWebGLBinding(session XRSession, context XRWebGLRenderingContext) XRWebGLBinding {
	return XRWebGLBinding{}.FromRef(
		bindings.NewXRWebGLBindingByXRWebGLBinding(
			session.Ref(),
			context.Ref()),
	)
}

type XRWebGLBinding struct {
	ref js.Ref
}

func (this XRWebGLBinding) Once() XRWebGLBinding {
	this.Ref().Once()
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
	this.Ref().Free()
}

// NativeProjectionScaleFactor returns the value of property "XRWebGLBinding.nativeProjectionScaleFactor".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLBinding) NativeProjectionScaleFactor() (float64, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLBindingNativeProjectionScaleFactor(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// UsesDepthValues returns the value of property "XRWebGLBinding.usesDepthValues".
//
// The returned bool will be false if there is no such property.
func (this XRWebGLBinding) UsesDepthValues() (bool, bool) {
	var _ok bool
	_ret := bindings.GetXRWebGLBindingUsesDepthValues(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// CreateProjectionLayer calls the method "XRWebGLBinding.createProjectionLayer".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLBinding) CreateProjectionLayer(init XRProjectionLayerInit) (XRProjectionLayer, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLBindingCreateProjectionLayer(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&init),
	)

	return XRProjectionLayer{}.FromRef(_ret), _ok
}

// CreateProjectionLayerFunc returns the method "XRWebGLBinding.createProjectionLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLBinding) CreateProjectionLayerFunc() (fn js.Func[func(init XRProjectionLayerInit) XRProjectionLayer]) {
	return fn.FromRef(
		bindings.XRWebGLBindingCreateProjectionLayerFunc(
			this.Ref(),
		),
	)
}

// CreateProjectionLayer1 calls the method "XRWebGLBinding.createProjectionLayer".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLBinding) CreateProjectionLayer1() (XRProjectionLayer, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLBindingCreateProjectionLayer1(
		this.Ref(), js.Pointer(&_ok),
	)

	return XRProjectionLayer{}.FromRef(_ret), _ok
}

// CreateProjectionLayer1Func returns the method "XRWebGLBinding.createProjectionLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLBinding) CreateProjectionLayer1Func() (fn js.Func[func() XRProjectionLayer]) {
	return fn.FromRef(
		bindings.XRWebGLBindingCreateProjectionLayer1Func(
			this.Ref(),
		),
	)
}

// CreateQuadLayer calls the method "XRWebGLBinding.createQuadLayer".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLBinding) CreateQuadLayer(init XRQuadLayerInit) (XRQuadLayer, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLBindingCreateQuadLayer(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&init),
	)

	return XRQuadLayer{}.FromRef(_ret), _ok
}

// CreateQuadLayerFunc returns the method "XRWebGLBinding.createQuadLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLBinding) CreateQuadLayerFunc() (fn js.Func[func(init XRQuadLayerInit) XRQuadLayer]) {
	return fn.FromRef(
		bindings.XRWebGLBindingCreateQuadLayerFunc(
			this.Ref(),
		),
	)
}

// CreateQuadLayer1 calls the method "XRWebGLBinding.createQuadLayer".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLBinding) CreateQuadLayer1() (XRQuadLayer, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLBindingCreateQuadLayer1(
		this.Ref(), js.Pointer(&_ok),
	)

	return XRQuadLayer{}.FromRef(_ret), _ok
}

// CreateQuadLayer1Func returns the method "XRWebGLBinding.createQuadLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLBinding) CreateQuadLayer1Func() (fn js.Func[func() XRQuadLayer]) {
	return fn.FromRef(
		bindings.XRWebGLBindingCreateQuadLayer1Func(
			this.Ref(),
		),
	)
}

// CreateCylinderLayer calls the method "XRWebGLBinding.createCylinderLayer".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLBinding) CreateCylinderLayer(init XRCylinderLayerInit) (XRCylinderLayer, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLBindingCreateCylinderLayer(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&init),
	)

	return XRCylinderLayer{}.FromRef(_ret), _ok
}

// CreateCylinderLayerFunc returns the method "XRWebGLBinding.createCylinderLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLBinding) CreateCylinderLayerFunc() (fn js.Func[func(init XRCylinderLayerInit) XRCylinderLayer]) {
	return fn.FromRef(
		bindings.XRWebGLBindingCreateCylinderLayerFunc(
			this.Ref(),
		),
	)
}

// CreateCylinderLayer1 calls the method "XRWebGLBinding.createCylinderLayer".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLBinding) CreateCylinderLayer1() (XRCylinderLayer, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLBindingCreateCylinderLayer1(
		this.Ref(), js.Pointer(&_ok),
	)

	return XRCylinderLayer{}.FromRef(_ret), _ok
}

// CreateCylinderLayer1Func returns the method "XRWebGLBinding.createCylinderLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLBinding) CreateCylinderLayer1Func() (fn js.Func[func() XRCylinderLayer]) {
	return fn.FromRef(
		bindings.XRWebGLBindingCreateCylinderLayer1Func(
			this.Ref(),
		),
	)
}

// CreateEquirectLayer calls the method "XRWebGLBinding.createEquirectLayer".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLBinding) CreateEquirectLayer(init XREquirectLayerInit) (XREquirectLayer, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLBindingCreateEquirectLayer(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&init),
	)

	return XREquirectLayer{}.FromRef(_ret), _ok
}

// CreateEquirectLayerFunc returns the method "XRWebGLBinding.createEquirectLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLBinding) CreateEquirectLayerFunc() (fn js.Func[func(init XREquirectLayerInit) XREquirectLayer]) {
	return fn.FromRef(
		bindings.XRWebGLBindingCreateEquirectLayerFunc(
			this.Ref(),
		),
	)
}

// CreateEquirectLayer1 calls the method "XRWebGLBinding.createEquirectLayer".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLBinding) CreateEquirectLayer1() (XREquirectLayer, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLBindingCreateEquirectLayer1(
		this.Ref(), js.Pointer(&_ok),
	)

	return XREquirectLayer{}.FromRef(_ret), _ok
}

// CreateEquirectLayer1Func returns the method "XRWebGLBinding.createEquirectLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLBinding) CreateEquirectLayer1Func() (fn js.Func[func() XREquirectLayer]) {
	return fn.FromRef(
		bindings.XRWebGLBindingCreateEquirectLayer1Func(
			this.Ref(),
		),
	)
}

// CreateCubeLayer calls the method "XRWebGLBinding.createCubeLayer".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLBinding) CreateCubeLayer(init XRCubeLayerInit) (XRCubeLayer, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLBindingCreateCubeLayer(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&init),
	)

	return XRCubeLayer{}.FromRef(_ret), _ok
}

// CreateCubeLayerFunc returns the method "XRWebGLBinding.createCubeLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLBinding) CreateCubeLayerFunc() (fn js.Func[func(init XRCubeLayerInit) XRCubeLayer]) {
	return fn.FromRef(
		bindings.XRWebGLBindingCreateCubeLayerFunc(
			this.Ref(),
		),
	)
}

// CreateCubeLayer1 calls the method "XRWebGLBinding.createCubeLayer".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLBinding) CreateCubeLayer1() (XRCubeLayer, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLBindingCreateCubeLayer1(
		this.Ref(), js.Pointer(&_ok),
	)

	return XRCubeLayer{}.FromRef(_ret), _ok
}

// CreateCubeLayer1Func returns the method "XRWebGLBinding.createCubeLayer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLBinding) CreateCubeLayer1Func() (fn js.Func[func() XRCubeLayer]) {
	return fn.FromRef(
		bindings.XRWebGLBindingCreateCubeLayer1Func(
			this.Ref(),
		),
	)
}

// GetSubImage calls the method "XRWebGLBinding.getSubImage".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLBinding) GetSubImage(layer XRCompositionLayer, frame XRFrame, eye XREye) (XRWebGLSubImage, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLBindingGetSubImage(
		this.Ref(), js.Pointer(&_ok),
		layer.Ref(),
		frame.Ref(),
		uint32(eye),
	)

	return XRWebGLSubImage{}.FromRef(_ret), _ok
}

// GetSubImageFunc returns the method "XRWebGLBinding.getSubImage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLBinding) GetSubImageFunc() (fn js.Func[func(layer XRCompositionLayer, frame XRFrame, eye XREye) XRWebGLSubImage]) {
	return fn.FromRef(
		bindings.XRWebGLBindingGetSubImageFunc(
			this.Ref(),
		),
	)
}

// GetSubImage1 calls the method "XRWebGLBinding.getSubImage".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLBinding) GetSubImage1(layer XRCompositionLayer, frame XRFrame) (XRWebGLSubImage, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLBindingGetSubImage1(
		this.Ref(), js.Pointer(&_ok),
		layer.Ref(),
		frame.Ref(),
	)

	return XRWebGLSubImage{}.FromRef(_ret), _ok
}

// GetSubImage1Func returns the method "XRWebGLBinding.getSubImage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLBinding) GetSubImage1Func() (fn js.Func[func(layer XRCompositionLayer, frame XRFrame) XRWebGLSubImage]) {
	return fn.FromRef(
		bindings.XRWebGLBindingGetSubImage1Func(
			this.Ref(),
		),
	)
}

// GetViewSubImage calls the method "XRWebGLBinding.getViewSubImage".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLBinding) GetViewSubImage(layer XRProjectionLayer, view XRView) (XRWebGLSubImage, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLBindingGetViewSubImage(
		this.Ref(), js.Pointer(&_ok),
		layer.Ref(),
		view.Ref(),
	)

	return XRWebGLSubImage{}.FromRef(_ret), _ok
}

// GetViewSubImageFunc returns the method "XRWebGLBinding.getViewSubImage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLBinding) GetViewSubImageFunc() (fn js.Func[func(layer XRProjectionLayer, view XRView) XRWebGLSubImage]) {
	return fn.FromRef(
		bindings.XRWebGLBindingGetViewSubImageFunc(
			this.Ref(),
		),
	)
}

// GetCameraImage calls the method "XRWebGLBinding.getCameraImage".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLBinding) GetCameraImage(camera XRCamera) (WebGLTexture, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLBindingGetCameraImage(
		this.Ref(), js.Pointer(&_ok),
		camera.Ref(),
	)

	return WebGLTexture{}.FromRef(_ret), _ok
}

// GetCameraImageFunc returns the method "XRWebGLBinding.getCameraImage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLBinding) GetCameraImageFunc() (fn js.Func[func(camera XRCamera) WebGLTexture]) {
	return fn.FromRef(
		bindings.XRWebGLBindingGetCameraImageFunc(
			this.Ref(),
		),
	)
}

// GetDepthInformation calls the method "XRWebGLBinding.getDepthInformation".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLBinding) GetDepthInformation(view XRView) (XRWebGLDepthInformation, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLBindingGetDepthInformation(
		this.Ref(), js.Pointer(&_ok),
		view.Ref(),
	)

	return XRWebGLDepthInformation{}.FromRef(_ret), _ok
}

// GetDepthInformationFunc returns the method "XRWebGLBinding.getDepthInformation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLBinding) GetDepthInformationFunc() (fn js.Func[func(view XRView) XRWebGLDepthInformation]) {
	return fn.FromRef(
		bindings.XRWebGLBindingGetDepthInformationFunc(
			this.Ref(),
		),
	)
}

// GetReflectionCubeMap calls the method "XRWebGLBinding.getReflectionCubeMap".
//
// The returned bool will be false if there is no such method.
func (this XRWebGLBinding) GetReflectionCubeMap(lightProbe XRLightProbe) (WebGLTexture, bool) {
	var _ok bool
	_ret := bindings.CallXRWebGLBindingGetReflectionCubeMap(
		this.Ref(), js.Pointer(&_ok),
		lightProbe.Ref(),
	)

	return WebGLTexture{}.FromRef(_ret), _ok
}

// GetReflectionCubeMapFunc returns the method "XRWebGLBinding.getReflectionCubeMap".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XRWebGLBinding) GetReflectionCubeMapFunc() (fn js.Func[func(lightProbe XRLightProbe) WebGLTexture]) {
	return fn.FromRef(
		bindings.XRWebGLBindingGetReflectionCubeMapFunc(
			this.Ref(),
		),
	)
}

type XSLTProcessor struct {
	ref js.Ref
}

func (this XSLTProcessor) Once() XSLTProcessor {
	this.Ref().Once()
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
	this.Ref().Free()
}

// ImportStylesheet calls the method "XSLTProcessor.importStylesheet".
//
// The returned bool will be false if there is no such method.
func (this XSLTProcessor) ImportStylesheet(style Node) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXSLTProcessorImportStylesheet(
		this.Ref(), js.Pointer(&_ok),
		style.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ImportStylesheetFunc returns the method "XSLTProcessor.importStylesheet".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XSLTProcessor) ImportStylesheetFunc() (fn js.Func[func(style Node)]) {
	return fn.FromRef(
		bindings.XSLTProcessorImportStylesheetFunc(
			this.Ref(),
		),
	)
}

// TransformToFragment calls the method "XSLTProcessor.transformToFragment".
//
// The returned bool will be false if there is no such method.
func (this XSLTProcessor) TransformToFragment(source Node, output Document) (DocumentFragment, bool) {
	var _ok bool
	_ret := bindings.CallXSLTProcessorTransformToFragment(
		this.Ref(), js.Pointer(&_ok),
		source.Ref(),
		output.Ref(),
	)

	return DocumentFragment{}.FromRef(_ret), _ok
}

// TransformToFragmentFunc returns the method "XSLTProcessor.transformToFragment".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XSLTProcessor) TransformToFragmentFunc() (fn js.Func[func(source Node, output Document) DocumentFragment]) {
	return fn.FromRef(
		bindings.XSLTProcessorTransformToFragmentFunc(
			this.Ref(),
		),
	)
}

// TransformToDocument calls the method "XSLTProcessor.transformToDocument".
//
// The returned bool will be false if there is no such method.
func (this XSLTProcessor) TransformToDocument(source Node) (Document, bool) {
	var _ok bool
	_ret := bindings.CallXSLTProcessorTransformToDocument(
		this.Ref(), js.Pointer(&_ok),
		source.Ref(),
	)

	return Document{}.FromRef(_ret), _ok
}

// TransformToDocumentFunc returns the method "XSLTProcessor.transformToDocument".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XSLTProcessor) TransformToDocumentFunc() (fn js.Func[func(source Node) Document]) {
	return fn.FromRef(
		bindings.XSLTProcessorTransformToDocumentFunc(
			this.Ref(),
		),
	)
}

// SetParameter calls the method "XSLTProcessor.setParameter".
//
// The returned bool will be false if there is no such method.
func (this XSLTProcessor) SetParameter(namespaceURI js.String, localName js.String, value js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXSLTProcessorSetParameter(
		this.Ref(), js.Pointer(&_ok),
		namespaceURI.Ref(),
		localName.Ref(),
		value.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetParameterFunc returns the method "XSLTProcessor.setParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XSLTProcessor) SetParameterFunc() (fn js.Func[func(namespaceURI js.String, localName js.String, value js.Any)]) {
	return fn.FromRef(
		bindings.XSLTProcessorSetParameterFunc(
			this.Ref(),
		),
	)
}

// GetParameter calls the method "XSLTProcessor.getParameter".
//
// The returned bool will be false if there is no such method.
func (this XSLTProcessor) GetParameter(namespaceURI js.String, localName js.String) (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallXSLTProcessorGetParameter(
		this.Ref(), js.Pointer(&_ok),
		namespaceURI.Ref(),
		localName.Ref(),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetParameterFunc returns the method "XSLTProcessor.getParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XSLTProcessor) GetParameterFunc() (fn js.Func[func(namespaceURI js.String, localName js.String) js.Any]) {
	return fn.FromRef(
		bindings.XSLTProcessorGetParameterFunc(
			this.Ref(),
		),
	)
}

// RemoveParameter calls the method "XSLTProcessor.removeParameter".
//
// The returned bool will be false if there is no such method.
func (this XSLTProcessor) RemoveParameter(namespaceURI js.String, localName js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXSLTProcessorRemoveParameter(
		this.Ref(), js.Pointer(&_ok),
		namespaceURI.Ref(),
		localName.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveParameterFunc returns the method "XSLTProcessor.removeParameter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XSLTProcessor) RemoveParameterFunc() (fn js.Func[func(namespaceURI js.String, localName js.String)]) {
	return fn.FromRef(
		bindings.XSLTProcessorRemoveParameterFunc(
			this.Ref(),
		),
	)
}

// ClearParameters calls the method "XSLTProcessor.clearParameters".
//
// The returned bool will be false if there is no such method.
func (this XSLTProcessor) ClearParameters() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXSLTProcessorClearParameters(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearParametersFunc returns the method "XSLTProcessor.clearParameters".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XSLTProcessor) ClearParametersFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.XSLTProcessorClearParametersFunc(
			this.Ref(),
		),
	)
}

// Reset calls the method "XSLTProcessor.reset".
//
// The returned bool will be false if there is no such method.
func (this XSLTProcessor) Reset() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallXSLTProcessorReset(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResetFunc returns the method "XSLTProcessor.reset".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this XSLTProcessor) ResetFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.XSLTProcessorResetFunc(
			this.Ref(),
		),
	)
}

type Console struct{}

// Assert calls the function "console.assert".
//
// The returned bool will be false if there is no such method.
func (Console) Assert(condition bool, data ...js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleAssert(
		js.Pointer(&_ok),
		js.Bool(bool(condition)),
		js.SliceData(data),
		js.SizeU(len(data)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AssertFunc returns the function "console.assert".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) AssertFunc() (fn js.Func[func(condition bool, data ...js.Any)]) {
	return fn.FromRef(
		bindings.ConsoleAssertFunc(),
	)
}

// Assert1 calls the function "console.assert".
//
// The returned bool will be false if there is no such method.
func (Console) Assert1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleAssert1(
		js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Assert1Func returns the function "console.assert".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) Assert1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ConsoleAssert1Func(),
	)
}

// Clear calls the function "console.clear".
//
// The returned bool will be false if there is no such method.
func (Console) Clear() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleClear(
		js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearFunc returns the function "console.clear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) ClearFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ConsoleClearFunc(),
	)
}

// Debug calls the function "console.debug".
//
// The returned bool will be false if there is no such method.
func (Console) Debug(data ...js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleDebug(
		js.Pointer(&_ok),
		js.SliceData(data),
		js.SizeU(len(data)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DebugFunc returns the function "console.debug".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) DebugFunc() (fn js.Func[func(data ...js.Any)]) {
	return fn.FromRef(
		bindings.ConsoleDebugFunc(),
	)
}

// Error calls the function "console.error".
//
// The returned bool will be false if there is no such method.
func (Console) Error(data ...js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleError(
		js.Pointer(&_ok),
		js.SliceData(data),
		js.SizeU(len(data)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ErrorFunc returns the function "console.error".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) ErrorFunc() (fn js.Func[func(data ...js.Any)]) {
	return fn.FromRef(
		bindings.ConsoleErrorFunc(),
	)
}

// Info calls the function "console.info".
//
// The returned bool will be false if there is no such method.
func (Console) Info(data ...js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleInfo(
		js.Pointer(&_ok),
		js.SliceData(data),
		js.SizeU(len(data)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InfoFunc returns the function "console.info".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) InfoFunc() (fn js.Func[func(data ...js.Any)]) {
	return fn.FromRef(
		bindings.ConsoleInfoFunc(),
	)
}

// Log calls the function "console.log".
//
// The returned bool will be false if there is no such method.
func (Console) Log(data ...js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleLog(
		js.Pointer(&_ok),
		js.SliceData(data),
		js.SizeU(len(data)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// LogFunc returns the function "console.log".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) LogFunc() (fn js.Func[func(data ...js.Any)]) {
	return fn.FromRef(
		bindings.ConsoleLogFunc(),
	)
}

// Table calls the function "console.table".
//
// The returned bool will be false if there is no such method.
func (Console) Table(tabularData js.Any, properties js.Array[js.String]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleTable(
		js.Pointer(&_ok),
		tabularData.Ref(),
		properties.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TableFunc returns the function "console.table".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) TableFunc() (fn js.Func[func(tabularData js.Any, properties js.Array[js.String])]) {
	return fn.FromRef(
		bindings.ConsoleTableFunc(),
	)
}

// Table1 calls the function "console.table".
//
// The returned bool will be false if there is no such method.
func (Console) Table1(tabularData js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleTable1(
		js.Pointer(&_ok),
		tabularData.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Table1Func returns the function "console.table".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) Table1Func() (fn js.Func[func(tabularData js.Any)]) {
	return fn.FromRef(
		bindings.ConsoleTable1Func(),
	)
}

// Table2 calls the function "console.table".
//
// The returned bool will be false if there is no such method.
func (Console) Table2() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleTable2(
		js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Table2Func returns the function "console.table".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) Table2Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ConsoleTable2Func(),
	)
}

// Trace calls the function "console.trace".
//
// The returned bool will be false if there is no such method.
func (Console) Trace(data ...js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleTrace(
		js.Pointer(&_ok),
		js.SliceData(data),
		js.SizeU(len(data)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TraceFunc returns the function "console.trace".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) TraceFunc() (fn js.Func[func(data ...js.Any)]) {
	return fn.FromRef(
		bindings.ConsoleTraceFunc(),
	)
}

// Warn calls the function "console.warn".
//
// The returned bool will be false if there is no such method.
func (Console) Warn(data ...js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleWarn(
		js.Pointer(&_ok),
		js.SliceData(data),
		js.SizeU(len(data)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// WarnFunc returns the function "console.warn".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) WarnFunc() (fn js.Func[func(data ...js.Any)]) {
	return fn.FromRef(
		bindings.ConsoleWarnFunc(),
	)
}

// Dir calls the function "console.dir".
//
// The returned bool will be false if there is no such method.
func (Console) Dir(item js.Any, options js.Object) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleDir(
		js.Pointer(&_ok),
		item.Ref(),
		options.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DirFunc returns the function "console.dir".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) DirFunc() (fn js.Func[func(item js.Any, options js.Object)]) {
	return fn.FromRef(
		bindings.ConsoleDirFunc(),
	)
}

// Dir1 calls the function "console.dir".
//
// The returned bool will be false if there is no such method.
func (Console) Dir1(item js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleDir1(
		js.Pointer(&_ok),
		item.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Dir1Func returns the function "console.dir".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) Dir1Func() (fn js.Func[func(item js.Any)]) {
	return fn.FromRef(
		bindings.ConsoleDir1Func(),
	)
}

// Dir2 calls the function "console.dir".
//
// The returned bool will be false if there is no such method.
func (Console) Dir2() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleDir2(
		js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Dir2Func returns the function "console.dir".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) Dir2Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ConsoleDir2Func(),
	)
}

// Dirxml calls the function "console.dirxml".
//
// The returned bool will be false if there is no such method.
func (Console) Dirxml(data ...js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleDirxml(
		js.Pointer(&_ok),
		js.SliceData(data),
		js.SizeU(len(data)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DirxmlFunc returns the function "console.dirxml".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) DirxmlFunc() (fn js.Func[func(data ...js.Any)]) {
	return fn.FromRef(
		bindings.ConsoleDirxmlFunc(),
	)
}

// Count calls the function "console.count".
//
// The returned bool will be false if there is no such method.
func (Console) Count(label js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleCount(
		js.Pointer(&_ok),
		label.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CountFunc returns the function "console.count".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) CountFunc() (fn js.Func[func(label js.String)]) {
	return fn.FromRef(
		bindings.ConsoleCountFunc(),
	)
}

// Count1 calls the function "console.count".
//
// The returned bool will be false if there is no such method.
func (Console) Count1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleCount1(
		js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Count1Func returns the function "console.count".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) Count1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ConsoleCount1Func(),
	)
}

// CountReset calls the function "console.countReset".
//
// The returned bool will be false if there is no such method.
func (Console) CountReset(label js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleCountReset(
		js.Pointer(&_ok),
		label.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CountResetFunc returns the function "console.countReset".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) CountResetFunc() (fn js.Func[func(label js.String)]) {
	return fn.FromRef(
		bindings.ConsoleCountResetFunc(),
	)
}

// CountReset1 calls the function "console.countReset".
//
// The returned bool will be false if there is no such method.
func (Console) CountReset1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleCountReset1(
		js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CountReset1Func returns the function "console.countReset".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) CountReset1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ConsoleCountReset1Func(),
	)
}

// Group calls the function "console.group".
//
// The returned bool will be false if there is no such method.
func (Console) Group(data ...js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleGroup(
		js.Pointer(&_ok),
		js.SliceData(data),
		js.SizeU(len(data)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GroupFunc returns the function "console.group".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) GroupFunc() (fn js.Func[func(data ...js.Any)]) {
	return fn.FromRef(
		bindings.ConsoleGroupFunc(),
	)
}

// GroupCollapsed calls the function "console.groupCollapsed".
//
// The returned bool will be false if there is no such method.
func (Console) GroupCollapsed(data ...js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleGroupCollapsed(
		js.Pointer(&_ok),
		js.SliceData(data),
		js.SizeU(len(data)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GroupCollapsedFunc returns the function "console.groupCollapsed".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) GroupCollapsedFunc() (fn js.Func[func(data ...js.Any)]) {
	return fn.FromRef(
		bindings.ConsoleGroupCollapsedFunc(),
	)
}

// GroupEnd calls the function "console.groupEnd".
//
// The returned bool will be false if there is no such method.
func (Console) GroupEnd() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleGroupEnd(
		js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GroupEndFunc returns the function "console.groupEnd".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) GroupEndFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ConsoleGroupEndFunc(),
	)
}

// Time calls the function "console.time".
//
// The returned bool will be false if there is no such method.
func (Console) Time(label js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleTime(
		js.Pointer(&_ok),
		label.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TimeFunc returns the function "console.time".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) TimeFunc() (fn js.Func[func(label js.String)]) {
	return fn.FromRef(
		bindings.ConsoleTimeFunc(),
	)
}

// Time1 calls the function "console.time".
//
// The returned bool will be false if there is no such method.
func (Console) Time1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleTime1(
		js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Time1Func returns the function "console.time".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) Time1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ConsoleTime1Func(),
	)
}

// TimeLog calls the function "console.timeLog".
//
// The returned bool will be false if there is no such method.
func (Console) TimeLog(label js.String, data ...js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleTimeLog(
		js.Pointer(&_ok),
		label.Ref(),
		js.SliceData(data),
		js.SizeU(len(data)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TimeLogFunc returns the function "console.timeLog".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) TimeLogFunc() (fn js.Func[func(label js.String, data ...js.Any)]) {
	return fn.FromRef(
		bindings.ConsoleTimeLogFunc(),
	)
}

// TimeLog1 calls the function "console.timeLog".
//
// The returned bool will be false if there is no such method.
func (Console) TimeLog1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleTimeLog1(
		js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TimeLog1Func returns the function "console.timeLog".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) TimeLog1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ConsoleTimeLog1Func(),
	)
}

// TimeEnd calls the function "console.timeEnd".
//
// The returned bool will be false if there is no such method.
func (Console) TimeEnd(label js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleTimeEnd(
		js.Pointer(&_ok),
		label.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TimeEndFunc returns the function "console.timeEnd".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) TimeEndFunc() (fn js.Func[func(label js.String)]) {
	return fn.FromRef(
		bindings.ConsoleTimeEndFunc(),
	)
}

// TimeEnd1 calls the function "console.timeEnd".
//
// The returned bool will be false if there is no such method.
func (Console) TimeEnd1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallConsoleTimeEnd1(
		js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TimeEnd1Func returns the function "console.timeEnd".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Console) TimeEnd1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ConsoleTimeEnd1Func(),
	)
}

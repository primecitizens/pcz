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

type GPUExtent3DDict struct {
	// Width is "GPUExtent3DDict.width"
	//
	// Required
	Width GPUIntegerCoordinate
	// Height is "GPUExtent3DDict.height"
	//
	// Optional, defaults to 1.
	Height GPUIntegerCoordinate
	// DepthOrArrayLayers is "GPUExtent3DDict.depthOrArrayLayers"
	//
	// Optional, defaults to 1.
	DepthOrArrayLayers GPUIntegerCoordinate

	FFI_USE_Height             bool // for Height.
	FFI_USE_DepthOrArrayLayers bool // for DepthOrArrayLayers.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUExtent3DDict with all fields set.
func (p GPUExtent3DDict) FromRef(ref js.Ref) GPUExtent3DDict {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUExtent3DDict GPUExtent3DDict [// GPUExtent3DDict] [0x140003d00a0 0x140003d0140 0x140003d0280 0x140003d01e0 0x140003d0320] 0x1400037d8c0 {0 0}} in the application heap.
func (p GPUExtent3DDict) New() js.Ref {
	return bindings.GPUExtent3DDictJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUExtent3DDict) UpdateFrom(ref js.Ref) {
	bindings.GPUExtent3DDictJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUExtent3DDict) Update(ref js.Ref) {
	bindings.GPUExtent3DDictJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_ArrayGPUIntegerCoordinate_GPUExtent3DDict struct {
	ref js.Ref
}

func (x OneOf_ArrayGPUIntegerCoordinate_GPUExtent3DDict) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ArrayGPUIntegerCoordinate_GPUExtent3DDict) Free() {
	x.ref.Free()
}

func (x OneOf_ArrayGPUIntegerCoordinate_GPUExtent3DDict) FromRef(ref js.Ref) OneOf_ArrayGPUIntegerCoordinate_GPUExtent3DDict {
	return OneOf_ArrayGPUIntegerCoordinate_GPUExtent3DDict{
		ref: ref,
	}
}

func (x OneOf_ArrayGPUIntegerCoordinate_GPUExtent3DDict) ArrayGPUIntegerCoordinate() js.Array[GPUIntegerCoordinate] {
	return js.Array[GPUIntegerCoordinate]{}.FromRef(x.ref)
}

func (x OneOf_ArrayGPUIntegerCoordinate_GPUExtent3DDict) GPUExtent3DDict() GPUExtent3DDict {
	var ret GPUExtent3DDict
	ret.UpdateFrom(x.ref)
	return ret
}

type GPUExtent3D = OneOf_ArrayGPUIntegerCoordinate_GPUExtent3DDict

type GPUSize32 uint32

type GPUTextureUsageFlags uint32

type GPUTextureDescriptor struct {
	// Size is "GPUTextureDescriptor.size"
	//
	// Required
	Size GPUExtent3D
	// MipLevelCount is "GPUTextureDescriptor.mipLevelCount"
	//
	// Optional, defaults to 1.
	MipLevelCount GPUIntegerCoordinate
	// SampleCount is "GPUTextureDescriptor.sampleCount"
	//
	// Optional, defaults to 1.
	SampleCount GPUSize32
	// Dimension is "GPUTextureDescriptor.dimension"
	//
	// Optional, defaults to "2d".
	Dimension GPUTextureDimension
	// Format is "GPUTextureDescriptor.format"
	//
	// Required
	Format GPUTextureFormat
	// Usage is "GPUTextureDescriptor.usage"
	//
	// Required
	Usage GPUTextureUsageFlags
	// ViewFormats is "GPUTextureDescriptor.viewFormats"
	//
	// Optional, defaults to [].
	ViewFormats js.Array[GPUTextureFormat]
	// Label is "GPUTextureDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE_MipLevelCount bool // for MipLevelCount.
	FFI_USE_SampleCount   bool // for SampleCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUTextureDescriptor with all fields set.
func (p GPUTextureDescriptor) FromRef(ref js.Ref) GPUTextureDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUTextureDescriptor GPUTextureDescriptor [// GPUTextureDescriptor] [0x140003d03c0 0x140003d0460 0x140003d05a0 0x140003d06e0 0x140003d0780 0x140003d0820 0x140003d08c0 0x140003d0960 0x140003d0500 0x140003d0640] 0x1400037d878 {0 0}} in the application heap.
func (p GPUTextureDescriptor) New() js.Ref {
	return bindings.GPUTextureDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUTextureDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUTextureDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUTextureDescriptor) Update(ref js.Ref) {
	bindings.GPUTextureDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUSampler struct {
	ref js.Ref
}

func (this GPUSampler) Once() GPUSampler {
	this.Ref().Once()
	return this
}

func (this GPUSampler) Ref() js.Ref {
	return this.ref
}

func (this GPUSampler) FromRef(ref js.Ref) GPUSampler {
	this.ref = ref
	return this
}

func (this GPUSampler) Free() {
	this.Ref().Free()
}

// Label returns the value of property "GPUSampler.label".
//
// The returned bool will be false if there is no such property.
func (this GPUSampler) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUSamplerLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPUSampler.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUSampler) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUSamplerLabel(
		this.Ref(),
		val.Ref(),
	)
}

type GPUAddressMode uint32

const (
	_ GPUAddressMode = iota

	GPUAddressMode_CLAMP_TO_EDGE
	GPUAddressMode_REPEAT
	GPUAddressMode_MIRROR_REPEAT
)

func (GPUAddressMode) FromRef(str js.Ref) GPUAddressMode {
	return GPUAddressMode(bindings.ConstOfGPUAddressMode(str))
}

func (x GPUAddressMode) String() (string, bool) {
	switch x {
	case GPUAddressMode_CLAMP_TO_EDGE:
		return "clamp-to-edge", true
	case GPUAddressMode_REPEAT:
		return "repeat", true
	case GPUAddressMode_MIRROR_REPEAT:
		return "mirror-repeat", true
	default:
		return "", false
	}
}

type GPUFilterMode uint32

const (
	_ GPUFilterMode = iota

	GPUFilterMode_NEAREST
	GPUFilterMode_LINEAR
)

func (GPUFilterMode) FromRef(str js.Ref) GPUFilterMode {
	return GPUFilterMode(bindings.ConstOfGPUFilterMode(str))
}

func (x GPUFilterMode) String() (string, bool) {
	switch x {
	case GPUFilterMode_NEAREST:
		return "nearest", true
	case GPUFilterMode_LINEAR:
		return "linear", true
	default:
		return "", false
	}
}

type GPUMipmapFilterMode uint32

const (
	_ GPUMipmapFilterMode = iota

	GPUMipmapFilterMode_NEAREST
	GPUMipmapFilterMode_LINEAR
)

func (GPUMipmapFilterMode) FromRef(str js.Ref) GPUMipmapFilterMode {
	return GPUMipmapFilterMode(bindings.ConstOfGPUMipmapFilterMode(str))
}

func (x GPUMipmapFilterMode) String() (string, bool) {
	switch x {
	case GPUMipmapFilterMode_NEAREST:
		return "nearest", true
	case GPUMipmapFilterMode_LINEAR:
		return "linear", true
	default:
		return "", false
	}
}

type GPUCompareFunction uint32

const (
	_ GPUCompareFunction = iota

	GPUCompareFunction_NEVER
	GPUCompareFunction_LESS
	GPUCompareFunction_EQUAL
	GPUCompareFunction_LESS_EQUAL
	GPUCompareFunction_GREATER
	GPUCompareFunction_NOT_EQUAL
	GPUCompareFunction_GREATER_EQUAL
	GPUCompareFunction_ALWAYS
)

func (GPUCompareFunction) FromRef(str js.Ref) GPUCompareFunction {
	return GPUCompareFunction(bindings.ConstOfGPUCompareFunction(str))
}

func (x GPUCompareFunction) String() (string, bool) {
	switch x {
	case GPUCompareFunction_NEVER:
		return "never", true
	case GPUCompareFunction_LESS:
		return "less", true
	case GPUCompareFunction_EQUAL:
		return "equal", true
	case GPUCompareFunction_LESS_EQUAL:
		return "less-equal", true
	case GPUCompareFunction_GREATER:
		return "greater", true
	case GPUCompareFunction_NOT_EQUAL:
		return "not-equal", true
	case GPUCompareFunction_GREATER_EQUAL:
		return "greater-equal", true
	case GPUCompareFunction_ALWAYS:
		return "always", true
	default:
		return "", false
	}
}

type GPUSamplerDescriptor struct {
	// AddressModeU is "GPUSamplerDescriptor.addressModeU"
	//
	// Optional, defaults to "clamp-to-edge".
	AddressModeU GPUAddressMode
	// AddressModeV is "GPUSamplerDescriptor.addressModeV"
	//
	// Optional, defaults to "clamp-to-edge".
	AddressModeV GPUAddressMode
	// AddressModeW is "GPUSamplerDescriptor.addressModeW"
	//
	// Optional, defaults to "clamp-to-edge".
	AddressModeW GPUAddressMode
	// MagFilter is "GPUSamplerDescriptor.magFilter"
	//
	// Optional, defaults to "nearest".
	MagFilter GPUFilterMode
	// MinFilter is "GPUSamplerDescriptor.minFilter"
	//
	// Optional, defaults to "nearest".
	MinFilter GPUFilterMode
	// MipmapFilter is "GPUSamplerDescriptor.mipmapFilter"
	//
	// Optional, defaults to "nearest".
	MipmapFilter GPUMipmapFilterMode
	// LodMinClamp is "GPUSamplerDescriptor.lodMinClamp"
	//
	// Optional, defaults to 0.
	LodMinClamp float32
	// LodMaxClamp is "GPUSamplerDescriptor.lodMaxClamp"
	//
	// Optional, defaults to 32.
	LodMaxClamp float32
	// Compare is "GPUSamplerDescriptor.compare"
	//
	// Optional
	Compare GPUCompareFunction
	// MaxAnisotropy is "GPUSamplerDescriptor.maxAnisotropy"
	//
	// Optional, defaults to 1.
	MaxAnisotropy uint16
	// Label is "GPUSamplerDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE_LodMinClamp   bool // for LodMinClamp.
	FFI_USE_LodMaxClamp   bool // for LodMaxClamp.
	FFI_USE_MaxAnisotropy bool // for MaxAnisotropy.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUSamplerDescriptor with all fields set.
func (p GPUSamplerDescriptor) FromRef(ref js.Ref) GPUSamplerDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUSamplerDescriptor GPUSamplerDescriptor [// GPUSamplerDescriptor] [0x140003d0a00 0x140003d0aa0 0x140003d0b40 0x140003d0be0 0x140003d0c80 0x140003d0d20 0x140003d0dc0 0x140003d0f00 0x140003d1040 0x140003d10e0 0x140003d1220 0x140003d0e60 0x140003d0fa0 0x140003d1180] 0x1400037d938 {0 0}} in the application heap.
func (p GPUSamplerDescriptor) New() js.Ref {
	return bindings.GPUSamplerDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUSamplerDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUSamplerDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUSamplerDescriptor) Update(ref js.Ref) {
	bindings.GPUSamplerDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUExternalTexture struct {
	ref js.Ref
}

func (this GPUExternalTexture) Once() GPUExternalTexture {
	this.Ref().Once()
	return this
}

func (this GPUExternalTexture) Ref() js.Ref {
	return this.ref
}

func (this GPUExternalTexture) FromRef(ref js.Ref) GPUExternalTexture {
	this.ref = ref
	return this
}

func (this GPUExternalTexture) Free() {
	this.Ref().Free()
}

// Label returns the value of property "GPUExternalTexture.label".
//
// The returned bool will be false if there is no such property.
func (this GPUExternalTexture) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUExternalTextureLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPUExternalTexture.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUExternalTexture) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUExternalTextureLabel(
		this.Ref(),
		val.Ref(),
	)
}

type OneOf_HTMLVideoElement_VideoFrame struct {
	ref js.Ref
}

func (x OneOf_HTMLVideoElement_VideoFrame) Ref() js.Ref {
	return x.ref
}

func (x OneOf_HTMLVideoElement_VideoFrame) Free() {
	x.ref.Free()
}

func (x OneOf_HTMLVideoElement_VideoFrame) FromRef(ref js.Ref) OneOf_HTMLVideoElement_VideoFrame {
	return OneOf_HTMLVideoElement_VideoFrame{
		ref: ref,
	}
}

func (x OneOf_HTMLVideoElement_VideoFrame) HTMLVideoElement() HTMLVideoElement {
	return HTMLVideoElement{}.FromRef(x.ref)
}

func (x OneOf_HTMLVideoElement_VideoFrame) VideoFrame() VideoFrame {
	return VideoFrame{}.FromRef(x.ref)
}

type GPUExternalTextureDescriptor struct {
	// Source is "GPUExternalTextureDescriptor.source"
	//
	// Required
	Source OneOf_HTMLVideoElement_VideoFrame
	// ColorSpace is "GPUExternalTextureDescriptor.colorSpace"
	//
	// Optional, defaults to "srgb".
	ColorSpace PredefinedColorSpace
	// Label is "GPUExternalTextureDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUExternalTextureDescriptor with all fields set.
func (p GPUExternalTextureDescriptor) FromRef(ref js.Ref) GPUExternalTextureDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUExternalTextureDescriptor GPUExternalTextureDescriptor [// GPUExternalTextureDescriptor] [0x140003d12c0 0x140003d1360 0x140003d1400] 0x1400037da10 {0 0}} in the application heap.
func (p GPUExternalTextureDescriptor) New() js.Ref {
	return bindings.GPUExternalTextureDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUExternalTextureDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUExternalTextureDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUExternalTextureDescriptor) Update(ref js.Ref) {
	bindings.GPUExternalTextureDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUBindGroupLayout struct {
	ref js.Ref
}

func (this GPUBindGroupLayout) Once() GPUBindGroupLayout {
	this.Ref().Once()
	return this
}

func (this GPUBindGroupLayout) Ref() js.Ref {
	return this.ref
}

func (this GPUBindGroupLayout) FromRef(ref js.Ref) GPUBindGroupLayout {
	this.ref = ref
	return this
}

func (this GPUBindGroupLayout) Free() {
	this.Ref().Free()
}

// Label returns the value of property "GPUBindGroupLayout.label".
//
// The returned bool will be false if there is no such property.
func (this GPUBindGroupLayout) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUBindGroupLayoutLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPUBindGroupLayout.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUBindGroupLayout) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUBindGroupLayoutLabel(
		this.Ref(),
		val.Ref(),
	)
}

type GPUIndex32 uint32

type GPUShaderStageFlags uint32

type GPUBufferBindingType uint32

const (
	_ GPUBufferBindingType = iota

	GPUBufferBindingType_UNIFORM
	GPUBufferBindingType_STORAGE
	GPUBufferBindingType_READ_ONLY_STORAGE
)

func (GPUBufferBindingType) FromRef(str js.Ref) GPUBufferBindingType {
	return GPUBufferBindingType(bindings.ConstOfGPUBufferBindingType(str))
}

func (x GPUBufferBindingType) String() (string, bool) {
	switch x {
	case GPUBufferBindingType_UNIFORM:
		return "uniform", true
	case GPUBufferBindingType_STORAGE:
		return "storage", true
	case GPUBufferBindingType_READ_ONLY_STORAGE:
		return "read-only-storage", true
	default:
		return "", false
	}
}

type GPUBufferBindingLayout struct {
	// Type is "GPUBufferBindingLayout.type"
	//
	// Optional, defaults to "uniform".
	Type GPUBufferBindingType
	// HasDynamicOffset is "GPUBufferBindingLayout.hasDynamicOffset"
	//
	// Optional, defaults to false.
	HasDynamicOffset bool
	// MinBindingSize is "GPUBufferBindingLayout.minBindingSize"
	//
	// Optional, defaults to 0.
	MinBindingSize GPUSize64

	FFI_USE_HasDynamicOffset bool // for HasDynamicOffset.
	FFI_USE_MinBindingSize   bool // for MinBindingSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUBufferBindingLayout with all fields set.
func (p GPUBufferBindingLayout) FromRef(ref js.Ref) GPUBufferBindingLayout {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUBufferBindingLayout GPUBufferBindingLayout [// GPUBufferBindingLayout] [0x140003d15e0 0x140003d1680 0x140003d17c0 0x140003d1720 0x140003d1860] 0x1400037da70 {0 0}} in the application heap.
func (p GPUBufferBindingLayout) New() js.Ref {
	return bindings.GPUBufferBindingLayoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUBufferBindingLayout) UpdateFrom(ref js.Ref) {
	bindings.GPUBufferBindingLayoutJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUBufferBindingLayout) Update(ref js.Ref) {
	bindings.GPUBufferBindingLayoutJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUSamplerBindingType uint32

const (
	_ GPUSamplerBindingType = iota

	GPUSamplerBindingType_FILTERING
	GPUSamplerBindingType_NON_FILTERING
	GPUSamplerBindingType_COMPARISON
)

func (GPUSamplerBindingType) FromRef(str js.Ref) GPUSamplerBindingType {
	return GPUSamplerBindingType(bindings.ConstOfGPUSamplerBindingType(str))
}

func (x GPUSamplerBindingType) String() (string, bool) {
	switch x {
	case GPUSamplerBindingType_FILTERING:
		return "filtering", true
	case GPUSamplerBindingType_NON_FILTERING:
		return "non-filtering", true
	case GPUSamplerBindingType_COMPARISON:
		return "comparison", true
	default:
		return "", false
	}
}

type GPUSamplerBindingLayout struct {
	// Type is "GPUSamplerBindingLayout.type"
	//
	// Optional, defaults to "filtering".
	Type GPUSamplerBindingType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUSamplerBindingLayout with all fields set.
func (p GPUSamplerBindingLayout) FromRef(ref js.Ref) GPUSamplerBindingLayout {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUSamplerBindingLayout GPUSamplerBindingLayout [// GPUSamplerBindingLayout] [0x140003d19a0] 0x1400037db18 {0 0}} in the application heap.
func (p GPUSamplerBindingLayout) New() js.Ref {
	return bindings.GPUSamplerBindingLayoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUSamplerBindingLayout) UpdateFrom(ref js.Ref) {
	bindings.GPUSamplerBindingLayoutJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUSamplerBindingLayout) Update(ref js.Ref) {
	bindings.GPUSamplerBindingLayoutJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUTextureSampleType uint32

const (
	_ GPUTextureSampleType = iota

	GPUTextureSampleType_FLOAT
	GPUTextureSampleType_UNFILTERABLE_FLOAT
	GPUTextureSampleType_DEPTH
	GPUTextureSampleType_SINT
	GPUTextureSampleType_UINT
)

func (GPUTextureSampleType) FromRef(str js.Ref) GPUTextureSampleType {
	return GPUTextureSampleType(bindings.ConstOfGPUTextureSampleType(str))
}

func (x GPUTextureSampleType) String() (string, bool) {
	switch x {
	case GPUTextureSampleType_FLOAT:
		return "float", true
	case GPUTextureSampleType_UNFILTERABLE_FLOAT:
		return "unfilterable-float", true
	case GPUTextureSampleType_DEPTH:
		return "depth", true
	case GPUTextureSampleType_SINT:
		return "sint", true
	case GPUTextureSampleType_UINT:
		return "uint", true
	default:
		return "", false
	}
}

type GPUTextureBindingLayout struct {
	// SampleType is "GPUTextureBindingLayout.sampleType"
	//
	// Optional, defaults to "float".
	SampleType GPUTextureSampleType
	// ViewDimension is "GPUTextureBindingLayout.viewDimension"
	//
	// Optional, defaults to "2d".
	ViewDimension GPUTextureViewDimension
	// Multisampled is "GPUTextureBindingLayout.multisampled"
	//
	// Optional, defaults to false.
	Multisampled bool

	FFI_USE_Multisampled bool // for Multisampled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUTextureBindingLayout with all fields set.
func (p GPUTextureBindingLayout) FromRef(ref js.Ref) GPUTextureBindingLayout {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUTextureBindingLayout GPUTextureBindingLayout [// GPUTextureBindingLayout] [0x140003d1ae0 0x140003d1b80 0x140003d1c20 0x140003d1cc0] 0x1400037db60 {0 0}} in the application heap.
func (p GPUTextureBindingLayout) New() js.Ref {
	return bindings.GPUTextureBindingLayoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUTextureBindingLayout) UpdateFrom(ref js.Ref) {
	bindings.GPUTextureBindingLayoutJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUTextureBindingLayout) Update(ref js.Ref) {
	bindings.GPUTextureBindingLayoutJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUStorageTextureAccess uint32

const (
	_ GPUStorageTextureAccess = iota

	GPUStorageTextureAccess_WRITE_ONLY
)

func (GPUStorageTextureAccess) FromRef(str js.Ref) GPUStorageTextureAccess {
	return GPUStorageTextureAccess(bindings.ConstOfGPUStorageTextureAccess(str))
}

func (x GPUStorageTextureAccess) String() (string, bool) {
	switch x {
	case GPUStorageTextureAccess_WRITE_ONLY:
		return "write-only", true
	default:
		return "", false
	}
}

type GPUStorageTextureBindingLayout struct {
	// Access is "GPUStorageTextureBindingLayout.access"
	//
	// Optional, defaults to "write-only".
	Access GPUStorageTextureAccess
	// Format is "GPUStorageTextureBindingLayout.format"
	//
	// Required
	Format GPUTextureFormat
	// ViewDimension is "GPUStorageTextureBindingLayout.viewDimension"
	//
	// Optional, defaults to "2d".
	ViewDimension GPUTextureViewDimension

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUStorageTextureBindingLayout with all fields set.
func (p GPUStorageTextureBindingLayout) FromRef(ref js.Ref) GPUStorageTextureBindingLayout {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUStorageTextureBindingLayout GPUStorageTextureBindingLayout [// GPUStorageTextureBindingLayout] [0x140003d1e00 0x140003d1ea0 0x140003d1f40] 0x1400037dbd8 {0 0}} in the application heap.
func (p GPUStorageTextureBindingLayout) New() js.Ref {
	return bindings.GPUStorageTextureBindingLayoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUStorageTextureBindingLayout) UpdateFrom(ref js.Ref) {
	bindings.GPUStorageTextureBindingLayoutJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUStorageTextureBindingLayout) Update(ref js.Ref) {
	bindings.GPUStorageTextureBindingLayoutJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUExternalTextureBindingLayout struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUExternalTextureBindingLayout with all fields set.
func (p GPUExternalTextureBindingLayout) FromRef(ref js.Ref) GPUExternalTextureBindingLayout {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUExternalTextureBindingLayout GPUExternalTextureBindingLayout [// GPUExternalTextureBindingLayout] [] 0x1400037dbf0 {0 0}} in the application heap.
func (p GPUExternalTextureBindingLayout) New() js.Ref {
	return bindings.GPUExternalTextureBindingLayoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUExternalTextureBindingLayout) UpdateFrom(ref js.Ref) {
	bindings.GPUExternalTextureBindingLayoutJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUExternalTextureBindingLayout) Update(ref js.Ref) {
	bindings.GPUExternalTextureBindingLayoutJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUBindGroupLayoutEntry struct {
	// Binding is "GPUBindGroupLayoutEntry.binding"
	//
	// Required
	Binding GPUIndex32
	// Visibility is "GPUBindGroupLayoutEntry.visibility"
	//
	// Required
	Visibility GPUShaderStageFlags
	// Buffer is "GPUBindGroupLayoutEntry.buffer"
	//
	// Optional
	Buffer GPUBufferBindingLayout
	// Sampler is "GPUBindGroupLayoutEntry.sampler"
	//
	// Optional
	Sampler GPUSamplerBindingLayout
	// Texture is "GPUBindGroupLayoutEntry.texture"
	//
	// Optional
	Texture GPUTextureBindingLayout
	// StorageTexture is "GPUBindGroupLayoutEntry.storageTexture"
	//
	// Optional
	StorageTexture GPUStorageTextureBindingLayout
	// ExternalTexture is "GPUBindGroupLayoutEntry.externalTexture"
	//
	// Optional
	ExternalTexture GPUExternalTextureBindingLayout

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUBindGroupLayoutEntry with all fields set.
func (p GPUBindGroupLayoutEntry) FromRef(ref js.Ref) GPUBindGroupLayoutEntry {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUBindGroupLayoutEntry GPUBindGroupLayoutEntry [// GPUBindGroupLayoutEntry] [0x140003d14a0 0x140003d1540 0x140003d1900 0x140003d1a40 0x140003d1d60 0x140003d8000 0x140003d80a0] 0x1400037da58 {0 0}} in the application heap.
func (p GPUBindGroupLayoutEntry) New() js.Ref {
	return bindings.GPUBindGroupLayoutEntryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUBindGroupLayoutEntry) UpdateFrom(ref js.Ref) {
	bindings.GPUBindGroupLayoutEntryJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUBindGroupLayoutEntry) Update(ref js.Ref) {
	bindings.GPUBindGroupLayoutEntryJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUBindGroupLayoutDescriptor struct {
	// Entries is "GPUBindGroupLayoutDescriptor.entries"
	//
	// Required
	Entries js.Array[GPUBindGroupLayoutEntry]
	// Label is "GPUBindGroupLayoutDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUBindGroupLayoutDescriptor with all fields set.
func (p GPUBindGroupLayoutDescriptor) FromRef(ref js.Ref) GPUBindGroupLayoutDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUBindGroupLayoutDescriptor GPUBindGroupLayoutDescriptor [// GPUBindGroupLayoutDescriptor] [0x140003d8140 0x140003d81e0] 0x1400037da40 {0 0}} in the application heap.
func (p GPUBindGroupLayoutDescriptor) New() js.Ref {
	return bindings.GPUBindGroupLayoutDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUBindGroupLayoutDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUBindGroupLayoutDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUBindGroupLayoutDescriptor) Update(ref js.Ref) {
	bindings.GPUBindGroupLayoutDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUPipelineLayout struct {
	ref js.Ref
}

func (this GPUPipelineLayout) Once() GPUPipelineLayout {
	this.Ref().Once()
	return this
}

func (this GPUPipelineLayout) Ref() js.Ref {
	return this.ref
}

func (this GPUPipelineLayout) FromRef(ref js.Ref) GPUPipelineLayout {
	this.ref = ref
	return this
}

func (this GPUPipelineLayout) Free() {
	this.Ref().Free()
}

// Label returns the value of property "GPUPipelineLayout.label".
//
// The returned bool will be false if there is no such property.
func (this GPUPipelineLayout) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUPipelineLayoutLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPUPipelineLayout.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUPipelineLayout) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUPipelineLayoutLabel(
		this.Ref(),
		val.Ref(),
	)
}

type GPUPipelineLayoutDescriptor struct {
	// BindGroupLayouts is "GPUPipelineLayoutDescriptor.bindGroupLayouts"
	//
	// Required
	BindGroupLayouts js.Array[GPUBindGroupLayout]
	// Label is "GPUPipelineLayoutDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUPipelineLayoutDescriptor with all fields set.
func (p GPUPipelineLayoutDescriptor) FromRef(ref js.Ref) GPUPipelineLayoutDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUPipelineLayoutDescriptor GPUPipelineLayoutDescriptor [// GPUPipelineLayoutDescriptor] [0x140003d8280 0x140003d8320] 0x1400037dc08 {0 0}} in the application heap.
func (p GPUPipelineLayoutDescriptor) New() js.Ref {
	return bindings.GPUPipelineLayoutDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUPipelineLayoutDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUPipelineLayoutDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUPipelineLayoutDescriptor) Update(ref js.Ref) {
	bindings.GPUPipelineLayoutDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUBindGroup struct {
	ref js.Ref
}

func (this GPUBindGroup) Once() GPUBindGroup {
	this.Ref().Once()
	return this
}

func (this GPUBindGroup) Ref() js.Ref {
	return this.ref
}

func (this GPUBindGroup) FromRef(ref js.Ref) GPUBindGroup {
	this.ref = ref
	return this
}

func (this GPUBindGroup) Free() {
	this.Ref().Free()
}

// Label returns the value of property "GPUBindGroup.label".
//
// The returned bool will be false if there is no such property.
func (this GPUBindGroup) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUBindGroupLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPUBindGroup.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUBindGroup) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUBindGroupLabel(
		this.Ref(),
		val.Ref(),
	)
}

type GPUBufferBinding struct {
	// Buffer is "GPUBufferBinding.buffer"
	//
	// Required
	Buffer GPUBuffer
	// Offset is "GPUBufferBinding.offset"
	//
	// Optional, defaults to 0.
	Offset GPUSize64
	// Size is "GPUBufferBinding.size"
	//
	// Optional
	Size GPUSize64

	FFI_USE_Offset bool // for Offset.
	FFI_USE_Size   bool // for Size.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUBufferBinding with all fields set.
func (p GPUBufferBinding) FromRef(ref js.Ref) GPUBufferBinding {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUBufferBinding GPUBufferBinding [// GPUBufferBinding] [0x140003d8500 0x140003d85a0 0x140003d86e0 0x140003d8640 0x140003d8780] 0x1400037dc98 {0 0}} in the application heap.
func (p GPUBufferBinding) New() js.Ref {
	return bindings.GPUBufferBindingJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUBufferBinding) UpdateFrom(ref js.Ref) {
	bindings.GPUBufferBindingJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUBufferBinding) Update(ref js.Ref) {
	bindings.GPUBufferBindingJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_GPUSampler_GPUTextureView_GPUBufferBinding_GPUExternalTexture struct {
	ref js.Ref
}

func (x OneOf_GPUSampler_GPUTextureView_GPUBufferBinding_GPUExternalTexture) Ref() js.Ref {
	return x.ref
}

func (x OneOf_GPUSampler_GPUTextureView_GPUBufferBinding_GPUExternalTexture) Free() {
	x.ref.Free()
}

func (x OneOf_GPUSampler_GPUTextureView_GPUBufferBinding_GPUExternalTexture) FromRef(ref js.Ref) OneOf_GPUSampler_GPUTextureView_GPUBufferBinding_GPUExternalTexture {
	return OneOf_GPUSampler_GPUTextureView_GPUBufferBinding_GPUExternalTexture{
		ref: ref,
	}
}

func (x OneOf_GPUSampler_GPUTextureView_GPUBufferBinding_GPUExternalTexture) GPUSampler() GPUSampler {
	return GPUSampler{}.FromRef(x.ref)
}

func (x OneOf_GPUSampler_GPUTextureView_GPUBufferBinding_GPUExternalTexture) GPUTextureView() GPUTextureView {
	return GPUTextureView{}.FromRef(x.ref)
}

func (x OneOf_GPUSampler_GPUTextureView_GPUBufferBinding_GPUExternalTexture) GPUBufferBinding() GPUBufferBinding {
	var ret GPUBufferBinding
	ret.UpdateFrom(x.ref)
	return ret
}

func (x OneOf_GPUSampler_GPUTextureView_GPUBufferBinding_GPUExternalTexture) GPUExternalTexture() GPUExternalTexture {
	return GPUExternalTexture{}.FromRef(x.ref)
}

type GPUBindingResource = OneOf_GPUSampler_GPUTextureView_GPUBufferBinding_GPUExternalTexture

type GPUBindGroupEntry struct {
	// Binding is "GPUBindGroupEntry.binding"
	//
	// Required
	Binding GPUIndex32
	// Resource is "GPUBindGroupEntry.resource"
	//
	// Required
	Resource GPUBindingResource

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUBindGroupEntry with all fields set.
func (p GPUBindGroupEntry) FromRef(ref js.Ref) GPUBindGroupEntry {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUBindGroupEntry GPUBindGroupEntry [// GPUBindGroupEntry] [0x140003d8460 0x140003d8820] 0x1400037dc50 {0 0}} in the application heap.
func (p GPUBindGroupEntry) New() js.Ref {
	return bindings.GPUBindGroupEntryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUBindGroupEntry) UpdateFrom(ref js.Ref) {
	bindings.GPUBindGroupEntryJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUBindGroupEntry) Update(ref js.Ref) {
	bindings.GPUBindGroupEntryJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUBindGroupDescriptor struct {
	// Layout is "GPUBindGroupDescriptor.layout"
	//
	// Required
	Layout GPUBindGroupLayout
	// Entries is "GPUBindGroupDescriptor.entries"
	//
	// Required
	Entries js.Array[GPUBindGroupEntry]
	// Label is "GPUBindGroupDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUBindGroupDescriptor with all fields set.
func (p GPUBindGroupDescriptor) FromRef(ref js.Ref) GPUBindGroupDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUBindGroupDescriptor GPUBindGroupDescriptor [// GPUBindGroupDescriptor] [0x140003d83c0 0x140003d88c0 0x140003d8960] 0x1400037dc20 {0 0}} in the application heap.
func (p GPUBindGroupDescriptor) New() js.Ref {
	return bindings.GPUBindGroupDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUBindGroupDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUBindGroupDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUBindGroupDescriptor) Update(ref js.Ref) {
	bindings.GPUBindGroupDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUCompilationMessageType uint32

const (
	_ GPUCompilationMessageType = iota

	GPUCompilationMessageType_ERROR
	GPUCompilationMessageType_WARNING
	GPUCompilationMessageType_INFO
)

func (GPUCompilationMessageType) FromRef(str js.Ref) GPUCompilationMessageType {
	return GPUCompilationMessageType(bindings.ConstOfGPUCompilationMessageType(str))
}

func (x GPUCompilationMessageType) String() (string, bool) {
	switch x {
	case GPUCompilationMessageType_ERROR:
		return "error", true
	case GPUCompilationMessageType_WARNING:
		return "warning", true
	case GPUCompilationMessageType_INFO:
		return "info", true
	default:
		return "", false
	}
}

type GPUCompilationMessage struct {
	ref js.Ref
}

func (this GPUCompilationMessage) Once() GPUCompilationMessage {
	this.Ref().Once()
	return this
}

func (this GPUCompilationMessage) Ref() js.Ref {
	return this.ref
}

func (this GPUCompilationMessage) FromRef(ref js.Ref) GPUCompilationMessage {
	this.ref = ref
	return this
}

func (this GPUCompilationMessage) Free() {
	this.Ref().Free()
}

// Message returns the value of property "GPUCompilationMessage.message".
//
// The returned bool will be false if there is no such property.
func (this GPUCompilationMessage) Message() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUCompilationMessageMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type returns the value of property "GPUCompilationMessage.type".
//
// The returned bool will be false if there is no such property.
func (this GPUCompilationMessage) Type() (GPUCompilationMessageType, bool) {
	var _ok bool
	_ret := bindings.GetGPUCompilationMessageType(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUCompilationMessageType(_ret), _ok
}

// LineNum returns the value of property "GPUCompilationMessage.lineNum".
//
// The returned bool will be false if there is no such property.
func (this GPUCompilationMessage) LineNum() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetGPUCompilationMessageLineNum(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// LinePos returns the value of property "GPUCompilationMessage.linePos".
//
// The returned bool will be false if there is no such property.
func (this GPUCompilationMessage) LinePos() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetGPUCompilationMessageLinePos(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// Offset returns the value of property "GPUCompilationMessage.offset".
//
// The returned bool will be false if there is no such property.
func (this GPUCompilationMessage) Offset() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetGPUCompilationMessageOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// Length returns the value of property "GPUCompilationMessage.length".
//
// The returned bool will be false if there is no such property.
func (this GPUCompilationMessage) Length() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetGPUCompilationMessageLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

type GPUCompilationInfo struct {
	ref js.Ref
}

func (this GPUCompilationInfo) Once() GPUCompilationInfo {
	this.Ref().Once()
	return this
}

func (this GPUCompilationInfo) Ref() js.Ref {
	return this.ref
}

func (this GPUCompilationInfo) FromRef(ref js.Ref) GPUCompilationInfo {
	this.ref = ref
	return this
}

func (this GPUCompilationInfo) Free() {
	this.Ref().Free()
}

// Messages returns the value of property "GPUCompilationInfo.messages".
//
// The returned bool will be false if there is no such property.
func (this GPUCompilationInfo) Messages() (js.FrozenArray[GPUCompilationMessage], bool) {
	var _ok bool
	_ret := bindings.GetGPUCompilationInfoMessages(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[GPUCompilationMessage]{}.FromRef(_ret), _ok
}

type GPUShaderModule struct {
	ref js.Ref
}

func (this GPUShaderModule) Once() GPUShaderModule {
	this.Ref().Once()
	return this
}

func (this GPUShaderModule) Ref() js.Ref {
	return this.ref
}

func (this GPUShaderModule) FromRef(ref js.Ref) GPUShaderModule {
	this.ref = ref
	return this
}

func (this GPUShaderModule) Free() {
	this.Ref().Free()
}

// Label returns the value of property "GPUShaderModule.label".
//
// The returned bool will be false if there is no such property.
func (this GPUShaderModule) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUShaderModuleLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPUShaderModule.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUShaderModule) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUShaderModuleLabel(
		this.Ref(),
		val.Ref(),
	)
}

// GetCompilationInfo calls the method "GPUShaderModule.getCompilationInfo".
//
// The returned bool will be false if there is no such method.
func (this GPUShaderModule) GetCompilationInfo() (js.Promise[GPUCompilationInfo], bool) {
	var _ok bool
	_ret := bindings.CallGPUShaderModuleGetCompilationInfo(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[GPUCompilationInfo]{}.FromRef(_ret), _ok
}

// GetCompilationInfoFunc returns the method "GPUShaderModule.getCompilationInfo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUShaderModule) GetCompilationInfoFunc() (fn js.Func[func() js.Promise[GPUCompilationInfo]]) {
	return fn.FromRef(
		bindings.GPUShaderModuleGetCompilationInfoFunc(
			this.Ref(),
		),
	)
}

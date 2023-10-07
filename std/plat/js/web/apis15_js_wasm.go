// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

type GPUExtent3DDict struct {
	// Width is "GPUExtent3DDict.width"
	//
	// Required
	Width GPUIntegerCoordinate
	// Height is "GPUExtent3DDict.height"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height GPUIntegerCoordinate
	// DepthOrArrayLayers is "GPUExtent3DDict.depthOrArrayLayers"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_DepthOrArrayLayers MUST be set to true to make this field effective.
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

// New creates a new GPUExtent3DDict in the application heap.
func (p GPUExtent3DDict) New() js.Ref {
	return bindings.GPUExtent3DDictJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUExtent3DDict) UpdateFrom(ref js.Ref) {
	bindings.GPUExtent3DDictJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUExtent3DDict) Update(ref js.Ref) {
	bindings.GPUExtent3DDictJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUExtent3DDict) FreeMembers(recursive bool) {
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
	//
	// NOTE: FFI_USE_MipLevelCount MUST be set to true to make this field effective.
	MipLevelCount GPUIntegerCoordinate
	// SampleCount is "GPUTextureDescriptor.sampleCount"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_SampleCount MUST be set to true to make this field effective.
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

// New creates a new GPUTextureDescriptor in the application heap.
func (p GPUTextureDescriptor) New() js.Ref {
	return bindings.GPUTextureDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUTextureDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUTextureDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUTextureDescriptor) Update(ref js.Ref) {
	bindings.GPUTextureDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUTextureDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Size.Ref(),
		p.ViewFormats.Ref(),
		p.Label.Ref(),
	)
	p.Size = p.Size.FromRef(js.Undefined)
	p.ViewFormats = p.ViewFormats.FromRef(js.Undefined)
	p.Label = p.Label.FromRef(js.Undefined)
}

type GPUSampler struct {
	ref js.Ref
}

func (this GPUSampler) Once() GPUSampler {
	this.ref.Once()
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
	this.ref.Free()
}

// Label returns the value of property "GPUSampler.label".
//
// It returns ok=false if there is no such property.
func (this GPUSampler) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUSamplerLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUSampler.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUSampler) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUSamplerLabel(
		this.ref,
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
	//
	// NOTE: FFI_USE_LodMinClamp MUST be set to true to make this field effective.
	LodMinClamp float32
	// LodMaxClamp is "GPUSamplerDescriptor.lodMaxClamp"
	//
	// Optional, defaults to 32.
	//
	// NOTE: FFI_USE_LodMaxClamp MUST be set to true to make this field effective.
	LodMaxClamp float32
	// Compare is "GPUSamplerDescriptor.compare"
	//
	// Optional
	Compare GPUCompareFunction
	// MaxAnisotropy is "GPUSamplerDescriptor.maxAnisotropy"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_MaxAnisotropy MUST be set to true to make this field effective.
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

// New creates a new GPUSamplerDescriptor in the application heap.
func (p GPUSamplerDescriptor) New() js.Ref {
	return bindings.GPUSamplerDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUSamplerDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUSamplerDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUSamplerDescriptor) Update(ref js.Ref) {
	bindings.GPUSamplerDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUSamplerDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Label.Ref(),
	)
	p.Label = p.Label.FromRef(js.Undefined)
}

type GPUExternalTexture struct {
	ref js.Ref
}

func (this GPUExternalTexture) Once() GPUExternalTexture {
	this.ref.Once()
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
	this.ref.Free()
}

// Label returns the value of property "GPUExternalTexture.label".
//
// It returns ok=false if there is no such property.
func (this GPUExternalTexture) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUExternalTextureLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUExternalTexture.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUExternalTexture) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUExternalTextureLabel(
		this.ref,
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

// New creates a new GPUExternalTextureDescriptor in the application heap.
func (p GPUExternalTextureDescriptor) New() js.Ref {
	return bindings.GPUExternalTextureDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUExternalTextureDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUExternalTextureDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUExternalTextureDescriptor) Update(ref js.Ref) {
	bindings.GPUExternalTextureDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUExternalTextureDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Source.Ref(),
		p.Label.Ref(),
	)
	p.Source = p.Source.FromRef(js.Undefined)
	p.Label = p.Label.FromRef(js.Undefined)
}

type GPUBindGroupLayout struct {
	ref js.Ref
}

func (this GPUBindGroupLayout) Once() GPUBindGroupLayout {
	this.ref.Once()
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
	this.ref.Free()
}

// Label returns the value of property "GPUBindGroupLayout.label".
//
// It returns ok=false if there is no such property.
func (this GPUBindGroupLayout) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUBindGroupLayoutLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUBindGroupLayout.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUBindGroupLayout) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUBindGroupLayoutLabel(
		this.ref,
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
	//
	// NOTE: FFI_USE_HasDynamicOffset MUST be set to true to make this field effective.
	HasDynamicOffset bool
	// MinBindingSize is "GPUBufferBindingLayout.minBindingSize"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_MinBindingSize MUST be set to true to make this field effective.
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

// New creates a new GPUBufferBindingLayout in the application heap.
func (p GPUBufferBindingLayout) New() js.Ref {
	return bindings.GPUBufferBindingLayoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUBufferBindingLayout) UpdateFrom(ref js.Ref) {
	bindings.GPUBufferBindingLayoutJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUBufferBindingLayout) Update(ref js.Ref) {
	bindings.GPUBufferBindingLayoutJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUBufferBindingLayout) FreeMembers(recursive bool) {
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

// New creates a new GPUSamplerBindingLayout in the application heap.
func (p GPUSamplerBindingLayout) New() js.Ref {
	return bindings.GPUSamplerBindingLayoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUSamplerBindingLayout) UpdateFrom(ref js.Ref) {
	bindings.GPUSamplerBindingLayoutJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUSamplerBindingLayout) Update(ref js.Ref) {
	bindings.GPUSamplerBindingLayoutJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUSamplerBindingLayout) FreeMembers(recursive bool) {
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
	//
	// NOTE: FFI_USE_Multisampled MUST be set to true to make this field effective.
	Multisampled bool

	FFI_USE_Multisampled bool // for Multisampled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUTextureBindingLayout with all fields set.
func (p GPUTextureBindingLayout) FromRef(ref js.Ref) GPUTextureBindingLayout {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUTextureBindingLayout in the application heap.
func (p GPUTextureBindingLayout) New() js.Ref {
	return bindings.GPUTextureBindingLayoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUTextureBindingLayout) UpdateFrom(ref js.Ref) {
	bindings.GPUTextureBindingLayoutJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUTextureBindingLayout) Update(ref js.Ref) {
	bindings.GPUTextureBindingLayoutJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUTextureBindingLayout) FreeMembers(recursive bool) {
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

// New creates a new GPUStorageTextureBindingLayout in the application heap.
func (p GPUStorageTextureBindingLayout) New() js.Ref {
	return bindings.GPUStorageTextureBindingLayoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUStorageTextureBindingLayout) UpdateFrom(ref js.Ref) {
	bindings.GPUStorageTextureBindingLayoutJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUStorageTextureBindingLayout) Update(ref js.Ref) {
	bindings.GPUStorageTextureBindingLayoutJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUStorageTextureBindingLayout) FreeMembers(recursive bool) {
}

type GPUExternalTextureBindingLayout struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUExternalTextureBindingLayout with all fields set.
func (p GPUExternalTextureBindingLayout) FromRef(ref js.Ref) GPUExternalTextureBindingLayout {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUExternalTextureBindingLayout in the application heap.
func (p GPUExternalTextureBindingLayout) New() js.Ref {
	return bindings.GPUExternalTextureBindingLayoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUExternalTextureBindingLayout) UpdateFrom(ref js.Ref) {
	bindings.GPUExternalTextureBindingLayoutJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUExternalTextureBindingLayout) Update(ref js.Ref) {
	bindings.GPUExternalTextureBindingLayoutJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUExternalTextureBindingLayout) FreeMembers(recursive bool) {
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
	//
	// NOTE: Buffer.FFI_USE MUST be set to true to get Buffer used.
	Buffer GPUBufferBindingLayout
	// Sampler is "GPUBindGroupLayoutEntry.sampler"
	//
	// Optional
	//
	// NOTE: Sampler.FFI_USE MUST be set to true to get Sampler used.
	Sampler GPUSamplerBindingLayout
	// Texture is "GPUBindGroupLayoutEntry.texture"
	//
	// Optional
	//
	// NOTE: Texture.FFI_USE MUST be set to true to get Texture used.
	Texture GPUTextureBindingLayout
	// StorageTexture is "GPUBindGroupLayoutEntry.storageTexture"
	//
	// Optional
	//
	// NOTE: StorageTexture.FFI_USE MUST be set to true to get StorageTexture used.
	StorageTexture GPUStorageTextureBindingLayout
	// ExternalTexture is "GPUBindGroupLayoutEntry.externalTexture"
	//
	// Optional
	//
	// NOTE: ExternalTexture.FFI_USE MUST be set to true to get ExternalTexture used.
	ExternalTexture GPUExternalTextureBindingLayout

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUBindGroupLayoutEntry with all fields set.
func (p GPUBindGroupLayoutEntry) FromRef(ref js.Ref) GPUBindGroupLayoutEntry {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUBindGroupLayoutEntry in the application heap.
func (p GPUBindGroupLayoutEntry) New() js.Ref {
	return bindings.GPUBindGroupLayoutEntryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUBindGroupLayoutEntry) UpdateFrom(ref js.Ref) {
	bindings.GPUBindGroupLayoutEntryJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUBindGroupLayoutEntry) Update(ref js.Ref) {
	bindings.GPUBindGroupLayoutEntryJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUBindGroupLayoutEntry) FreeMembers(recursive bool) {
	if recursive {
		p.Buffer.FreeMembers(true)
		p.Sampler.FreeMembers(true)
		p.Texture.FreeMembers(true)
		p.StorageTexture.FreeMembers(true)
		p.ExternalTexture.FreeMembers(true)
	}
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

// New creates a new GPUBindGroupLayoutDescriptor in the application heap.
func (p GPUBindGroupLayoutDescriptor) New() js.Ref {
	return bindings.GPUBindGroupLayoutDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUBindGroupLayoutDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUBindGroupLayoutDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUBindGroupLayoutDescriptor) Update(ref js.Ref) {
	bindings.GPUBindGroupLayoutDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUBindGroupLayoutDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Entries.Ref(),
		p.Label.Ref(),
	)
	p.Entries = p.Entries.FromRef(js.Undefined)
	p.Label = p.Label.FromRef(js.Undefined)
}

type GPUPipelineLayout struct {
	ref js.Ref
}

func (this GPUPipelineLayout) Once() GPUPipelineLayout {
	this.ref.Once()
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
	this.ref.Free()
}

// Label returns the value of property "GPUPipelineLayout.label".
//
// It returns ok=false if there is no such property.
func (this GPUPipelineLayout) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUPipelineLayoutLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUPipelineLayout.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUPipelineLayout) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUPipelineLayoutLabel(
		this.ref,
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

// New creates a new GPUPipelineLayoutDescriptor in the application heap.
func (p GPUPipelineLayoutDescriptor) New() js.Ref {
	return bindings.GPUPipelineLayoutDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUPipelineLayoutDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUPipelineLayoutDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUPipelineLayoutDescriptor) Update(ref js.Ref) {
	bindings.GPUPipelineLayoutDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUPipelineLayoutDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.BindGroupLayouts.Ref(),
		p.Label.Ref(),
	)
	p.BindGroupLayouts = p.BindGroupLayouts.FromRef(js.Undefined)
	p.Label = p.Label.FromRef(js.Undefined)
}

type GPUBindGroup struct {
	ref js.Ref
}

func (this GPUBindGroup) Once() GPUBindGroup {
	this.ref.Once()
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
	this.ref.Free()
}

// Label returns the value of property "GPUBindGroup.label".
//
// It returns ok=false if there is no such property.
func (this GPUBindGroup) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUBindGroupLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUBindGroup.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUBindGroup) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUBindGroupLabel(
		this.ref,
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
	//
	// NOTE: FFI_USE_Offset MUST be set to true to make this field effective.
	Offset GPUSize64
	// Size is "GPUBufferBinding.size"
	//
	// Optional
	//
	// NOTE: FFI_USE_Size MUST be set to true to make this field effective.
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

// New creates a new GPUBufferBinding in the application heap.
func (p GPUBufferBinding) New() js.Ref {
	return bindings.GPUBufferBindingJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUBufferBinding) UpdateFrom(ref js.Ref) {
	bindings.GPUBufferBindingJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUBufferBinding) Update(ref js.Ref) {
	bindings.GPUBufferBindingJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUBufferBinding) FreeMembers(recursive bool) {
	js.Free(
		p.Buffer.Ref(),
	)
	p.Buffer = p.Buffer.FromRef(js.Undefined)
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

// New creates a new GPUBindGroupEntry in the application heap.
func (p GPUBindGroupEntry) New() js.Ref {
	return bindings.GPUBindGroupEntryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUBindGroupEntry) UpdateFrom(ref js.Ref) {
	bindings.GPUBindGroupEntryJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUBindGroupEntry) Update(ref js.Ref) {
	bindings.GPUBindGroupEntryJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUBindGroupEntry) FreeMembers(recursive bool) {
	js.Free(
		p.Resource.Ref(),
	)
	p.Resource = p.Resource.FromRef(js.Undefined)
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

// New creates a new GPUBindGroupDescriptor in the application heap.
func (p GPUBindGroupDescriptor) New() js.Ref {
	return bindings.GPUBindGroupDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUBindGroupDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUBindGroupDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUBindGroupDescriptor) Update(ref js.Ref) {
	bindings.GPUBindGroupDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUBindGroupDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Layout.Ref(),
		p.Entries.Ref(),
		p.Label.Ref(),
	)
	p.Layout = p.Layout.FromRef(js.Undefined)
	p.Entries = p.Entries.FromRef(js.Undefined)
	p.Label = p.Label.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Message returns the value of property "GPUCompilationMessage.message".
//
// It returns ok=false if there is no such property.
func (this GPUCompilationMessage) Message() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUCompilationMessageMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "GPUCompilationMessage.type".
//
// It returns ok=false if there is no such property.
func (this GPUCompilationMessage) Type() (ret GPUCompilationMessageType, ok bool) {
	ok = js.True == bindings.GetGPUCompilationMessageType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LineNum returns the value of property "GPUCompilationMessage.lineNum".
//
// It returns ok=false if there is no such property.
func (this GPUCompilationMessage) LineNum() (ret uint64, ok bool) {
	ok = js.True == bindings.GetGPUCompilationMessageLineNum(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LinePos returns the value of property "GPUCompilationMessage.linePos".
//
// It returns ok=false if there is no such property.
func (this GPUCompilationMessage) LinePos() (ret uint64, ok bool) {
	ok = js.True == bindings.GetGPUCompilationMessageLinePos(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Offset returns the value of property "GPUCompilationMessage.offset".
//
// It returns ok=false if there is no such property.
func (this GPUCompilationMessage) Offset() (ret uint64, ok bool) {
	ok = js.True == bindings.GetGPUCompilationMessageOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Length returns the value of property "GPUCompilationMessage.length".
//
// It returns ok=false if there is no such property.
func (this GPUCompilationMessage) Length() (ret uint64, ok bool) {
	ok = js.True == bindings.GetGPUCompilationMessageLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

type GPUCompilationInfo struct {
	ref js.Ref
}

func (this GPUCompilationInfo) Once() GPUCompilationInfo {
	this.ref.Once()
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
	this.ref.Free()
}

// Messages returns the value of property "GPUCompilationInfo.messages".
//
// It returns ok=false if there is no such property.
func (this GPUCompilationInfo) Messages() (ret js.FrozenArray[GPUCompilationMessage], ok bool) {
	ok = js.True == bindings.GetGPUCompilationInfoMessages(
		this.ref, js.Pointer(&ret),
	)
	return
}

type GPUShaderModule struct {
	ref js.Ref
}

func (this GPUShaderModule) Once() GPUShaderModule {
	this.ref.Once()
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
	this.ref.Free()
}

// Label returns the value of property "GPUShaderModule.label".
//
// It returns ok=false if there is no such property.
func (this GPUShaderModule) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUShaderModuleLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUShaderModule.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUShaderModule) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUShaderModuleLabel(
		this.ref,
		val.Ref(),
	)
}

// HasFuncGetCompilationInfo returns true if the method "GPUShaderModule.getCompilationInfo" exists.
func (this GPUShaderModule) HasFuncGetCompilationInfo() bool {
	return js.True == bindings.HasFuncGPUShaderModuleGetCompilationInfo(
		this.ref,
	)
}

// FuncGetCompilationInfo returns the method "GPUShaderModule.getCompilationInfo".
func (this GPUShaderModule) FuncGetCompilationInfo() (fn js.Func[func() js.Promise[GPUCompilationInfo]]) {
	bindings.FuncGPUShaderModuleGetCompilationInfo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetCompilationInfo calls the method "GPUShaderModule.getCompilationInfo".
func (this GPUShaderModule) GetCompilationInfo() (ret js.Promise[GPUCompilationInfo]) {
	bindings.CallGPUShaderModuleGetCompilationInfo(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetCompilationInfo calls the method "GPUShaderModule.getCompilationInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUShaderModule) TryGetCompilationInfo() (ret js.Promise[GPUCompilationInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUShaderModuleGetCompilationInfo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

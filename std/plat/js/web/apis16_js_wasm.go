// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

type GPUAutoLayoutMode uint32

const (
	_ GPUAutoLayoutMode = iota

	GPUAutoLayoutMode_AUTO
)

func (GPUAutoLayoutMode) FromRef(str js.Ref) GPUAutoLayoutMode {
	return GPUAutoLayoutMode(bindings.ConstOfGPUAutoLayoutMode(str))
}

func (x GPUAutoLayoutMode) String() (string, bool) {
	switch x {
	case GPUAutoLayoutMode_AUTO:
		return "auto", true
	default:
		return "", false
	}
}

type OneOf_GPUPipelineLayout_GPUAutoLayoutMode struct {
	ref js.Ref
}

func (x OneOf_GPUPipelineLayout_GPUAutoLayoutMode) Ref() js.Ref {
	return x.ref
}

func (x OneOf_GPUPipelineLayout_GPUAutoLayoutMode) Free() {
	x.ref.Free()
}

func (x OneOf_GPUPipelineLayout_GPUAutoLayoutMode) FromRef(ref js.Ref) OneOf_GPUPipelineLayout_GPUAutoLayoutMode {
	return OneOf_GPUPipelineLayout_GPUAutoLayoutMode{
		ref: ref,
	}
}

func (x OneOf_GPUPipelineLayout_GPUAutoLayoutMode) GPUPipelineLayout() GPUPipelineLayout {
	return GPUPipelineLayout{}.FromRef(x.ref)
}

func (x OneOf_GPUPipelineLayout_GPUAutoLayoutMode) GPUAutoLayoutMode() GPUAutoLayoutMode {
	return GPUAutoLayoutMode(0).FromRef(x.ref)
}

type GPUShaderModuleCompilationHint struct {
	// Layout is "GPUShaderModuleCompilationHint.layout"
	//
	// Optional
	Layout OneOf_GPUPipelineLayout_GPUAutoLayoutMode

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUShaderModuleCompilationHint with all fields set.
func (p GPUShaderModuleCompilationHint) FromRef(ref js.Ref) GPUShaderModuleCompilationHint {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUShaderModuleCompilationHint in the application heap.
func (p GPUShaderModuleCompilationHint) New() js.Ref {
	return bindings.GPUShaderModuleCompilationHintJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUShaderModuleCompilationHint) UpdateFrom(ref js.Ref) {
	bindings.GPUShaderModuleCompilationHintJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUShaderModuleCompilationHint) Update(ref js.Ref) {
	bindings.GPUShaderModuleCompilationHintJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUShaderModuleDescriptor struct {
	// Code is "GPUShaderModuleDescriptor.code"
	//
	// Required
	Code js.String
	// SourceMap is "GPUShaderModuleDescriptor.sourceMap"
	//
	// Optional
	SourceMap js.Object
	// Hints is "GPUShaderModuleDescriptor.hints"
	//
	// Optional
	Hints js.Record[GPUShaderModuleCompilationHint]
	// Label is "GPUShaderModuleDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUShaderModuleDescriptor with all fields set.
func (p GPUShaderModuleDescriptor) FromRef(ref js.Ref) GPUShaderModuleDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUShaderModuleDescriptor in the application heap.
func (p GPUShaderModuleDescriptor) New() js.Ref {
	return bindings.GPUShaderModuleDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUShaderModuleDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUShaderModuleDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUShaderModuleDescriptor) Update(ref js.Ref) {
	bindings.GPUShaderModuleDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUComputePipeline struct {
	ref js.Ref
}

func (this GPUComputePipeline) Once() GPUComputePipeline {
	this.Ref().Once()
	return this
}

func (this GPUComputePipeline) Ref() js.Ref {
	return this.ref
}

func (this GPUComputePipeline) FromRef(ref js.Ref) GPUComputePipeline {
	this.ref = ref
	return this
}

func (this GPUComputePipeline) Free() {
	this.Ref().Free()
}

// Label returns the value of property "GPUComputePipeline.label".
//
// It returns ok=false if there is no such property.
func (this GPUComputePipeline) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUComputePipelineLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUComputePipeline.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUComputePipeline) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUComputePipelineLabel(
		this.Ref(),
		val.Ref(),
	)
}

// HasGetBindGroupLayout returns true if the method "GPUComputePipeline.getBindGroupLayout" exists.
func (this GPUComputePipeline) HasGetBindGroupLayout() bool {
	return js.True == bindings.HasGPUComputePipelineGetBindGroupLayout(
		this.Ref(),
	)
}

// GetBindGroupLayoutFunc returns the method "GPUComputePipeline.getBindGroupLayout".
func (this GPUComputePipeline) GetBindGroupLayoutFunc() (fn js.Func[func(index uint32) GPUBindGroupLayout]) {
	return fn.FromRef(
		bindings.GPUComputePipelineGetBindGroupLayoutFunc(
			this.Ref(),
		),
	)
}

// GetBindGroupLayout calls the method "GPUComputePipeline.getBindGroupLayout".
func (this GPUComputePipeline) GetBindGroupLayout(index uint32) (ret GPUBindGroupLayout) {
	bindings.CallGPUComputePipelineGetBindGroupLayout(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGetBindGroupLayout calls the method "GPUComputePipeline.getBindGroupLayout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePipeline) TryGetBindGroupLayout(index uint32) (ret GPUBindGroupLayout, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePipelineGetBindGroupLayout(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

type GPUPipelineConstantValue float64

type GPUProgrammableStage struct {
	// Module is "GPUProgrammableStage.module"
	//
	// Required
	Module GPUShaderModule
	// EntryPoint is "GPUProgrammableStage.entryPoint"
	//
	// Required
	EntryPoint js.String
	// Constants is "GPUProgrammableStage.constants"
	//
	// Optional
	Constants js.Record[GPUPipelineConstantValue]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUProgrammableStage with all fields set.
func (p GPUProgrammableStage) FromRef(ref js.Ref) GPUProgrammableStage {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUProgrammableStage in the application heap.
func (p GPUProgrammableStage) New() js.Ref {
	return bindings.GPUProgrammableStageJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUProgrammableStage) UpdateFrom(ref js.Ref) {
	bindings.GPUProgrammableStageJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUProgrammableStage) Update(ref js.Ref) {
	bindings.GPUProgrammableStageJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUComputePipelineDescriptor struct {
	// Compute is "GPUComputePipelineDescriptor.compute"
	//
	// Required
	//
	// NOTE: Compute.FFI_USE MUST be set to true to get Compute used.
	Compute GPUProgrammableStage
	// Layout is "GPUComputePipelineDescriptor.layout"
	//
	// Required
	Layout OneOf_GPUPipelineLayout_GPUAutoLayoutMode
	// Label is "GPUComputePipelineDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUComputePipelineDescriptor with all fields set.
func (p GPUComputePipelineDescriptor) FromRef(ref js.Ref) GPUComputePipelineDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUComputePipelineDescriptor in the application heap.
func (p GPUComputePipelineDescriptor) New() js.Ref {
	return bindings.GPUComputePipelineDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUComputePipelineDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUComputePipelineDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUComputePipelineDescriptor) Update(ref js.Ref) {
	bindings.GPUComputePipelineDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPURenderPipeline struct {
	ref js.Ref
}

func (this GPURenderPipeline) Once() GPURenderPipeline {
	this.Ref().Once()
	return this
}

func (this GPURenderPipeline) Ref() js.Ref {
	return this.ref
}

func (this GPURenderPipeline) FromRef(ref js.Ref) GPURenderPipeline {
	this.ref = ref
	return this
}

func (this GPURenderPipeline) Free() {
	this.Ref().Free()
}

// Label returns the value of property "GPURenderPipeline.label".
//
// It returns ok=false if there is no such property.
func (this GPURenderPipeline) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPURenderPipelineLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPURenderPipeline.label" to val.
//
// It returns false if the property cannot be set.
func (this GPURenderPipeline) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPURenderPipelineLabel(
		this.Ref(),
		val.Ref(),
	)
}

// HasGetBindGroupLayout returns true if the method "GPURenderPipeline.getBindGroupLayout" exists.
func (this GPURenderPipeline) HasGetBindGroupLayout() bool {
	return js.True == bindings.HasGPURenderPipelineGetBindGroupLayout(
		this.Ref(),
	)
}

// GetBindGroupLayoutFunc returns the method "GPURenderPipeline.getBindGroupLayout".
func (this GPURenderPipeline) GetBindGroupLayoutFunc() (fn js.Func[func(index uint32) GPUBindGroupLayout]) {
	return fn.FromRef(
		bindings.GPURenderPipelineGetBindGroupLayoutFunc(
			this.Ref(),
		),
	)
}

// GetBindGroupLayout calls the method "GPURenderPipeline.getBindGroupLayout".
func (this GPURenderPipeline) GetBindGroupLayout(index uint32) (ret GPUBindGroupLayout) {
	bindings.CallGPURenderPipelineGetBindGroupLayout(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGetBindGroupLayout calls the method "GPURenderPipeline.getBindGroupLayout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPipeline) TryGetBindGroupLayout(index uint32) (ret GPUBindGroupLayout, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPipelineGetBindGroupLayout(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

type GPUVertexStepMode uint32

const (
	_ GPUVertexStepMode = iota

	GPUVertexStepMode_VERTEX
	GPUVertexStepMode_INSTANCE
)

func (GPUVertexStepMode) FromRef(str js.Ref) GPUVertexStepMode {
	return GPUVertexStepMode(bindings.ConstOfGPUVertexStepMode(str))
}

func (x GPUVertexStepMode) String() (string, bool) {
	switch x {
	case GPUVertexStepMode_VERTEX:
		return "vertex", true
	case GPUVertexStepMode_INSTANCE:
		return "instance", true
	default:
		return "", false
	}
}

type GPUVertexFormat uint32

const (
	_ GPUVertexFormat = iota

	GPUVertexFormat_UINT_8X2
	GPUVertexFormat_UINT_8X4
	GPUVertexFormat_SINT_8X2
	GPUVertexFormat_SINT_8X4
	GPUVertexFormat_UNORM_8X2
	GPUVertexFormat_UNORM_8X4
	GPUVertexFormat_SNORM_8X2
	GPUVertexFormat_SNORM_8X4
	GPUVertexFormat_UINT_16X2
	GPUVertexFormat_UINT_16X4
	GPUVertexFormat_SINT_16X2
	GPUVertexFormat_SINT_16X4
	GPUVertexFormat_UNORM_16X2
	GPUVertexFormat_UNORM_16X4
	GPUVertexFormat_SNORM_16X2
	GPUVertexFormat_SNORM_16X4
	GPUVertexFormat_FLOAT_16X2
	GPUVertexFormat_FLOAT_16X4
	GPUVertexFormat_FLOAT32
	GPUVertexFormat_FLOAT_32X2
	GPUVertexFormat_FLOAT_32X3
	GPUVertexFormat_FLOAT_32X4
	GPUVertexFormat_UINT32
	GPUVertexFormat_UINT_32X2
	GPUVertexFormat_UINT_32X3
	GPUVertexFormat_UINT_32X4
	GPUVertexFormat_SINT32
	GPUVertexFormat_SINT_32X2
	GPUVertexFormat_SINT_32X3
	GPUVertexFormat_SINT_32X4
)

func (GPUVertexFormat) FromRef(str js.Ref) GPUVertexFormat {
	return GPUVertexFormat(bindings.ConstOfGPUVertexFormat(str))
}

func (x GPUVertexFormat) String() (string, bool) {
	switch x {
	case GPUVertexFormat_UINT_8X2:
		return "uint8x2", true
	case GPUVertexFormat_UINT_8X4:
		return "uint8x4", true
	case GPUVertexFormat_SINT_8X2:
		return "sint8x2", true
	case GPUVertexFormat_SINT_8X4:
		return "sint8x4", true
	case GPUVertexFormat_UNORM_8X2:
		return "unorm8x2", true
	case GPUVertexFormat_UNORM_8X4:
		return "unorm8x4", true
	case GPUVertexFormat_SNORM_8X2:
		return "snorm8x2", true
	case GPUVertexFormat_SNORM_8X4:
		return "snorm8x4", true
	case GPUVertexFormat_UINT_16X2:
		return "uint16x2", true
	case GPUVertexFormat_UINT_16X4:
		return "uint16x4", true
	case GPUVertexFormat_SINT_16X2:
		return "sint16x2", true
	case GPUVertexFormat_SINT_16X4:
		return "sint16x4", true
	case GPUVertexFormat_UNORM_16X2:
		return "unorm16x2", true
	case GPUVertexFormat_UNORM_16X4:
		return "unorm16x4", true
	case GPUVertexFormat_SNORM_16X2:
		return "snorm16x2", true
	case GPUVertexFormat_SNORM_16X4:
		return "snorm16x4", true
	case GPUVertexFormat_FLOAT_16X2:
		return "float16x2", true
	case GPUVertexFormat_FLOAT_16X4:
		return "float16x4", true
	case GPUVertexFormat_FLOAT32:
		return "float32", true
	case GPUVertexFormat_FLOAT_32X2:
		return "float32x2", true
	case GPUVertexFormat_FLOAT_32X3:
		return "float32x3", true
	case GPUVertexFormat_FLOAT_32X4:
		return "float32x4", true
	case GPUVertexFormat_UINT32:
		return "uint32", true
	case GPUVertexFormat_UINT_32X2:
		return "uint32x2", true
	case GPUVertexFormat_UINT_32X3:
		return "uint32x3", true
	case GPUVertexFormat_UINT_32X4:
		return "uint32x4", true
	case GPUVertexFormat_SINT32:
		return "sint32", true
	case GPUVertexFormat_SINT_32X2:
		return "sint32x2", true
	case GPUVertexFormat_SINT_32X3:
		return "sint32x3", true
	case GPUVertexFormat_SINT_32X4:
		return "sint32x4", true
	default:
		return "", false
	}
}

type GPUVertexAttribute struct {
	// Format is "GPUVertexAttribute.format"
	//
	// Required
	Format GPUVertexFormat
	// Offset is "GPUVertexAttribute.offset"
	//
	// Required
	Offset GPUSize64
	// ShaderLocation is "GPUVertexAttribute.shaderLocation"
	//
	// Required
	ShaderLocation GPUIndex32

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUVertexAttribute with all fields set.
func (p GPUVertexAttribute) FromRef(ref js.Ref) GPUVertexAttribute {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUVertexAttribute in the application heap.
func (p GPUVertexAttribute) New() js.Ref {
	return bindings.GPUVertexAttributeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUVertexAttribute) UpdateFrom(ref js.Ref) {
	bindings.GPUVertexAttributeJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUVertexAttribute) Update(ref js.Ref) {
	bindings.GPUVertexAttributeJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUVertexBufferLayout struct {
	// ArrayStride is "GPUVertexBufferLayout.arrayStride"
	//
	// Required
	ArrayStride GPUSize64
	// StepMode is "GPUVertexBufferLayout.stepMode"
	//
	// Optional, defaults to "vertex".
	StepMode GPUVertexStepMode
	// Attributes is "GPUVertexBufferLayout.attributes"
	//
	// Required
	Attributes js.Array[GPUVertexAttribute]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUVertexBufferLayout with all fields set.
func (p GPUVertexBufferLayout) FromRef(ref js.Ref) GPUVertexBufferLayout {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUVertexBufferLayout in the application heap.
func (p GPUVertexBufferLayout) New() js.Ref {
	return bindings.GPUVertexBufferLayoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUVertexBufferLayout) UpdateFrom(ref js.Ref) {
	bindings.GPUVertexBufferLayoutJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUVertexBufferLayout) Update(ref js.Ref) {
	bindings.GPUVertexBufferLayoutJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUVertexState struct {
	// Buffers is "GPUVertexState.buffers"
	//
	// Optional, defaults to [].
	Buffers js.Array[GPUVertexBufferLayout]
	// Module is "GPUVertexState.module"
	//
	// Required
	Module GPUShaderModule
	// EntryPoint is "GPUVertexState.entryPoint"
	//
	// Required
	EntryPoint js.String
	// Constants is "GPUVertexState.constants"
	//
	// Optional
	Constants js.Record[GPUPipelineConstantValue]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUVertexState with all fields set.
func (p GPUVertexState) FromRef(ref js.Ref) GPUVertexState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUVertexState in the application heap.
func (p GPUVertexState) New() js.Ref {
	return bindings.GPUVertexStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUVertexState) UpdateFrom(ref js.Ref) {
	bindings.GPUVertexStateJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUVertexState) Update(ref js.Ref) {
	bindings.GPUVertexStateJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUPrimitiveTopology uint32

const (
	_ GPUPrimitiveTopology = iota

	GPUPrimitiveTopology_POINT_LIST
	GPUPrimitiveTopology_LINE_LIST
	GPUPrimitiveTopology_LINE_STRIP
	GPUPrimitiveTopology_TRIANGLE_LIST
	GPUPrimitiveTopology_TRIANGLE_STRIP
)

func (GPUPrimitiveTopology) FromRef(str js.Ref) GPUPrimitiveTopology {
	return GPUPrimitiveTopology(bindings.ConstOfGPUPrimitiveTopology(str))
}

func (x GPUPrimitiveTopology) String() (string, bool) {
	switch x {
	case GPUPrimitiveTopology_POINT_LIST:
		return "point-list", true
	case GPUPrimitiveTopology_LINE_LIST:
		return "line-list", true
	case GPUPrimitiveTopology_LINE_STRIP:
		return "line-strip", true
	case GPUPrimitiveTopology_TRIANGLE_LIST:
		return "triangle-list", true
	case GPUPrimitiveTopology_TRIANGLE_STRIP:
		return "triangle-strip", true
	default:
		return "", false
	}
}

type GPUIndexFormat uint32

const (
	_ GPUIndexFormat = iota

	GPUIndexFormat_UINT16
	GPUIndexFormat_UINT32
)

func (GPUIndexFormat) FromRef(str js.Ref) GPUIndexFormat {
	return GPUIndexFormat(bindings.ConstOfGPUIndexFormat(str))
}

func (x GPUIndexFormat) String() (string, bool) {
	switch x {
	case GPUIndexFormat_UINT16:
		return "uint16", true
	case GPUIndexFormat_UINT32:
		return "uint32", true
	default:
		return "", false
	}
}

type GPUFrontFace uint32

const (
	_ GPUFrontFace = iota

	GPUFrontFace_CCW
	GPUFrontFace_CW
)

func (GPUFrontFace) FromRef(str js.Ref) GPUFrontFace {
	return GPUFrontFace(bindings.ConstOfGPUFrontFace(str))
}

func (x GPUFrontFace) String() (string, bool) {
	switch x {
	case GPUFrontFace_CCW:
		return "ccw", true
	case GPUFrontFace_CW:
		return "cw", true
	default:
		return "", false
	}
}

type GPUCullMode uint32

const (
	_ GPUCullMode = iota

	GPUCullMode_NONE
	GPUCullMode_FRONT
	GPUCullMode_BACK
)

func (GPUCullMode) FromRef(str js.Ref) GPUCullMode {
	return GPUCullMode(bindings.ConstOfGPUCullMode(str))
}

func (x GPUCullMode) String() (string, bool) {
	switch x {
	case GPUCullMode_NONE:
		return "none", true
	case GPUCullMode_FRONT:
		return "front", true
	case GPUCullMode_BACK:
		return "back", true
	default:
		return "", false
	}
}

type GPUPrimitiveState struct {
	// Topology is "GPUPrimitiveState.topology"
	//
	// Optional, defaults to "triangle-list".
	Topology GPUPrimitiveTopology
	// StripIndexFormat is "GPUPrimitiveState.stripIndexFormat"
	//
	// Optional
	StripIndexFormat GPUIndexFormat
	// FrontFace is "GPUPrimitiveState.frontFace"
	//
	// Optional, defaults to "ccw".
	FrontFace GPUFrontFace
	// CullMode is "GPUPrimitiveState.cullMode"
	//
	// Optional, defaults to "none".
	CullMode GPUCullMode
	// UnclippedDepth is "GPUPrimitiveState.unclippedDepth"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_UnclippedDepth MUST be set to true to make this field effective.
	UnclippedDepth bool

	FFI_USE_UnclippedDepth bool // for UnclippedDepth.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUPrimitiveState with all fields set.
func (p GPUPrimitiveState) FromRef(ref js.Ref) GPUPrimitiveState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUPrimitiveState in the application heap.
func (p GPUPrimitiveState) New() js.Ref {
	return bindings.GPUPrimitiveStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUPrimitiveState) UpdateFrom(ref js.Ref) {
	bindings.GPUPrimitiveStateJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUPrimitiveState) Update(ref js.Ref) {
	bindings.GPUPrimitiveStateJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUStencilOperation uint32

const (
	_ GPUStencilOperation = iota

	GPUStencilOperation_KEEP
	GPUStencilOperation_ZERO
	GPUStencilOperation_REPLACE
	GPUStencilOperation_INVERT
	GPUStencilOperation_INCREMENT_CLAMP
	GPUStencilOperation_DECREMENT_CLAMP
	GPUStencilOperation_INCREMENT_WRAP
	GPUStencilOperation_DECREMENT_WRAP
)

func (GPUStencilOperation) FromRef(str js.Ref) GPUStencilOperation {
	return GPUStencilOperation(bindings.ConstOfGPUStencilOperation(str))
}

func (x GPUStencilOperation) String() (string, bool) {
	switch x {
	case GPUStencilOperation_KEEP:
		return "keep", true
	case GPUStencilOperation_ZERO:
		return "zero", true
	case GPUStencilOperation_REPLACE:
		return "replace", true
	case GPUStencilOperation_INVERT:
		return "invert", true
	case GPUStencilOperation_INCREMENT_CLAMP:
		return "increment-clamp", true
	case GPUStencilOperation_DECREMENT_CLAMP:
		return "decrement-clamp", true
	case GPUStencilOperation_INCREMENT_WRAP:
		return "increment-wrap", true
	case GPUStencilOperation_DECREMENT_WRAP:
		return "decrement-wrap", true
	default:
		return "", false
	}
}

type GPUStencilFaceState struct {
	// Compare is "GPUStencilFaceState.compare"
	//
	// Optional, defaults to "always".
	Compare GPUCompareFunction
	// FailOp is "GPUStencilFaceState.failOp"
	//
	// Optional, defaults to "keep".
	FailOp GPUStencilOperation
	// DepthFailOp is "GPUStencilFaceState.depthFailOp"
	//
	// Optional, defaults to "keep".
	DepthFailOp GPUStencilOperation
	// PassOp is "GPUStencilFaceState.passOp"
	//
	// Optional, defaults to "keep".
	PassOp GPUStencilOperation

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUStencilFaceState with all fields set.
func (p GPUStencilFaceState) FromRef(ref js.Ref) GPUStencilFaceState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUStencilFaceState in the application heap.
func (p GPUStencilFaceState) New() js.Ref {
	return bindings.GPUStencilFaceStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUStencilFaceState) UpdateFrom(ref js.Ref) {
	bindings.GPUStencilFaceStateJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUStencilFaceState) Update(ref js.Ref) {
	bindings.GPUStencilFaceStateJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUStencilValue uint32

type GPUDepthBias int32

type GPUDepthStencilState struct {
	// Format is "GPUDepthStencilState.format"
	//
	// Required
	Format GPUTextureFormat
	// DepthWriteEnabled is "GPUDepthStencilState.depthWriteEnabled"
	//
	// Required
	DepthWriteEnabled bool
	// DepthCompare is "GPUDepthStencilState.depthCompare"
	//
	// Required
	DepthCompare GPUCompareFunction
	// StencilFront is "GPUDepthStencilState.stencilFront"
	//
	// Optional, defaults to {}.
	//
	// NOTE: StencilFront.FFI_USE MUST be set to true to get StencilFront used.
	StencilFront GPUStencilFaceState
	// StencilBack is "GPUDepthStencilState.stencilBack"
	//
	// Optional, defaults to {}.
	//
	// NOTE: StencilBack.FFI_USE MUST be set to true to get StencilBack used.
	StencilBack GPUStencilFaceState
	// StencilReadMask is "GPUDepthStencilState.stencilReadMask"
	//
	// Optional, defaults to 0xFFFFFFFF.
	//
	// NOTE: FFI_USE_StencilReadMask MUST be set to true to make this field effective.
	StencilReadMask GPUStencilValue
	// StencilWriteMask is "GPUDepthStencilState.stencilWriteMask"
	//
	// Optional, defaults to 0xFFFFFFFF.
	//
	// NOTE: FFI_USE_StencilWriteMask MUST be set to true to make this field effective.
	StencilWriteMask GPUStencilValue
	// DepthBias is "GPUDepthStencilState.depthBias"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_DepthBias MUST be set to true to make this field effective.
	DepthBias GPUDepthBias
	// DepthBiasSlopeScale is "GPUDepthStencilState.depthBiasSlopeScale"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_DepthBiasSlopeScale MUST be set to true to make this field effective.
	DepthBiasSlopeScale float32
	// DepthBiasClamp is "GPUDepthStencilState.depthBiasClamp"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_DepthBiasClamp MUST be set to true to make this field effective.
	DepthBiasClamp float32

	FFI_USE_StencilReadMask     bool // for StencilReadMask.
	FFI_USE_StencilWriteMask    bool // for StencilWriteMask.
	FFI_USE_DepthBias           bool // for DepthBias.
	FFI_USE_DepthBiasSlopeScale bool // for DepthBiasSlopeScale.
	FFI_USE_DepthBiasClamp      bool // for DepthBiasClamp.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUDepthStencilState with all fields set.
func (p GPUDepthStencilState) FromRef(ref js.Ref) GPUDepthStencilState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUDepthStencilState in the application heap.
func (p GPUDepthStencilState) New() js.Ref {
	return bindings.GPUDepthStencilStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUDepthStencilState) UpdateFrom(ref js.Ref) {
	bindings.GPUDepthStencilStateJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUDepthStencilState) Update(ref js.Ref) {
	bindings.GPUDepthStencilStateJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUSampleMask uint32

type GPUMultisampleState struct {
	// Count is "GPUMultisampleState.count"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_Count MUST be set to true to make this field effective.
	Count GPUSize32
	// Mask is "GPUMultisampleState.mask"
	//
	// Optional, defaults to 0xFFFFFFFF.
	//
	// NOTE: FFI_USE_Mask MUST be set to true to make this field effective.
	Mask GPUSampleMask
	// AlphaToCoverageEnabled is "GPUMultisampleState.alphaToCoverageEnabled"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_AlphaToCoverageEnabled MUST be set to true to make this field effective.
	AlphaToCoverageEnabled bool

	FFI_USE_Count                  bool // for Count.
	FFI_USE_Mask                   bool // for Mask.
	FFI_USE_AlphaToCoverageEnabled bool // for AlphaToCoverageEnabled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUMultisampleState with all fields set.
func (p GPUMultisampleState) FromRef(ref js.Ref) GPUMultisampleState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUMultisampleState in the application heap.
func (p GPUMultisampleState) New() js.Ref {
	return bindings.GPUMultisampleStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUMultisampleState) UpdateFrom(ref js.Ref) {
	bindings.GPUMultisampleStateJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUMultisampleState) Update(ref js.Ref) {
	bindings.GPUMultisampleStateJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUBlendOperation uint32

const (
	_ GPUBlendOperation = iota

	GPUBlendOperation_ADD
	GPUBlendOperation_SUBTRACT
	GPUBlendOperation_REVERSE_SUBTRACT
	GPUBlendOperation_MIN
	GPUBlendOperation_MAX
)

func (GPUBlendOperation) FromRef(str js.Ref) GPUBlendOperation {
	return GPUBlendOperation(bindings.ConstOfGPUBlendOperation(str))
}

func (x GPUBlendOperation) String() (string, bool) {
	switch x {
	case GPUBlendOperation_ADD:
		return "add", true
	case GPUBlendOperation_SUBTRACT:
		return "subtract", true
	case GPUBlendOperation_REVERSE_SUBTRACT:
		return "reverse-subtract", true
	case GPUBlendOperation_MIN:
		return "min", true
	case GPUBlendOperation_MAX:
		return "max", true
	default:
		return "", false
	}
}

type GPUBlendFactor uint32

const (
	_ GPUBlendFactor = iota

	GPUBlendFactor_ZERO
	GPUBlendFactor_ONE
	GPUBlendFactor_SRC
	GPUBlendFactor_ONE_MINUS_SRC
	GPUBlendFactor_SRC_ALPHA
	GPUBlendFactor_ONE_MINUS_SRC_ALPHA
	GPUBlendFactor_DST
	GPUBlendFactor_ONE_MINUS_DST
	GPUBlendFactor_DST_ALPHA
	GPUBlendFactor_ONE_MINUS_DST_ALPHA
	GPUBlendFactor_SRC_ALPHA_SATURATED
	GPUBlendFactor_CONSTANT
	GPUBlendFactor_ONE_MINUS_CONSTANT
)

func (GPUBlendFactor) FromRef(str js.Ref) GPUBlendFactor {
	return GPUBlendFactor(bindings.ConstOfGPUBlendFactor(str))
}

func (x GPUBlendFactor) String() (string, bool) {
	switch x {
	case GPUBlendFactor_ZERO:
		return "zero", true
	case GPUBlendFactor_ONE:
		return "one", true
	case GPUBlendFactor_SRC:
		return "src", true
	case GPUBlendFactor_ONE_MINUS_SRC:
		return "one-minus-src", true
	case GPUBlendFactor_SRC_ALPHA:
		return "src-alpha", true
	case GPUBlendFactor_ONE_MINUS_SRC_ALPHA:
		return "one-minus-src-alpha", true
	case GPUBlendFactor_DST:
		return "dst", true
	case GPUBlendFactor_ONE_MINUS_DST:
		return "one-minus-dst", true
	case GPUBlendFactor_DST_ALPHA:
		return "dst-alpha", true
	case GPUBlendFactor_ONE_MINUS_DST_ALPHA:
		return "one-minus-dst-alpha", true
	case GPUBlendFactor_SRC_ALPHA_SATURATED:
		return "src-alpha-saturated", true
	case GPUBlendFactor_CONSTANT:
		return "constant", true
	case GPUBlendFactor_ONE_MINUS_CONSTANT:
		return "one-minus-constant", true
	default:
		return "", false
	}
}

type GPUBlendComponent struct {
	// Operation is "GPUBlendComponent.operation"
	//
	// Optional, defaults to "add".
	Operation GPUBlendOperation
	// SrcFactor is "GPUBlendComponent.srcFactor"
	//
	// Optional, defaults to "one".
	SrcFactor GPUBlendFactor
	// DstFactor is "GPUBlendComponent.dstFactor"
	//
	// Optional, defaults to "zero".
	DstFactor GPUBlendFactor

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUBlendComponent with all fields set.
func (p GPUBlendComponent) FromRef(ref js.Ref) GPUBlendComponent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUBlendComponent in the application heap.
func (p GPUBlendComponent) New() js.Ref {
	return bindings.GPUBlendComponentJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUBlendComponent) UpdateFrom(ref js.Ref) {
	bindings.GPUBlendComponentJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUBlendComponent) Update(ref js.Ref) {
	bindings.GPUBlendComponentJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUBlendState struct {
	// Color is "GPUBlendState.color"
	//
	// Required
	//
	// NOTE: Color.FFI_USE MUST be set to true to get Color used.
	Color GPUBlendComponent
	// Alpha is "GPUBlendState.alpha"
	//
	// Required
	//
	// NOTE: Alpha.FFI_USE MUST be set to true to get Alpha used.
	Alpha GPUBlendComponent

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUBlendState with all fields set.
func (p GPUBlendState) FromRef(ref js.Ref) GPUBlendState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUBlendState in the application heap.
func (p GPUBlendState) New() js.Ref {
	return bindings.GPUBlendStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUBlendState) UpdateFrom(ref js.Ref) {
	bindings.GPUBlendStateJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUBlendState) Update(ref js.Ref) {
	bindings.GPUBlendStateJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUColorWriteFlags uint32

type GPUColorTargetState struct {
	// Format is "GPUColorTargetState.format"
	//
	// Required
	Format GPUTextureFormat
	// Blend is "GPUColorTargetState.blend"
	//
	// Optional
	//
	// NOTE: Blend.FFI_USE MUST be set to true to get Blend used.
	Blend GPUBlendState
	// WriteMask is "GPUColorTargetState.writeMask"
	//
	// Optional, defaults to 0xF.
	//
	// NOTE: FFI_USE_WriteMask MUST be set to true to make this field effective.
	WriteMask GPUColorWriteFlags

	FFI_USE_WriteMask bool // for WriteMask.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUColorTargetState with all fields set.
func (p GPUColorTargetState) FromRef(ref js.Ref) GPUColorTargetState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUColorTargetState in the application heap.
func (p GPUColorTargetState) New() js.Ref {
	return bindings.GPUColorTargetStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUColorTargetState) UpdateFrom(ref js.Ref) {
	bindings.GPUColorTargetStateJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUColorTargetState) Update(ref js.Ref) {
	bindings.GPUColorTargetStateJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUFragmentState struct {
	// Targets is "GPUFragmentState.targets"
	//
	// Required
	Targets js.Array[GPUColorTargetState]
	// Module is "GPUFragmentState.module"
	//
	// Required
	Module GPUShaderModule
	// EntryPoint is "GPUFragmentState.entryPoint"
	//
	// Required
	EntryPoint js.String
	// Constants is "GPUFragmentState.constants"
	//
	// Optional
	Constants js.Record[GPUPipelineConstantValue]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUFragmentState with all fields set.
func (p GPUFragmentState) FromRef(ref js.Ref) GPUFragmentState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUFragmentState in the application heap.
func (p GPUFragmentState) New() js.Ref {
	return bindings.GPUFragmentStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUFragmentState) UpdateFrom(ref js.Ref) {
	bindings.GPUFragmentStateJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUFragmentState) Update(ref js.Ref) {
	bindings.GPUFragmentStateJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPURenderPipelineDescriptor struct {
	// Vertex is "GPURenderPipelineDescriptor.vertex"
	//
	// Required
	//
	// NOTE: Vertex.FFI_USE MUST be set to true to get Vertex used.
	Vertex GPUVertexState
	// Primitive is "GPURenderPipelineDescriptor.primitive"
	//
	// Optional, defaults to {}.
	//
	// NOTE: Primitive.FFI_USE MUST be set to true to get Primitive used.
	Primitive GPUPrimitiveState
	// DepthStencil is "GPURenderPipelineDescriptor.depthStencil"
	//
	// Optional
	//
	// NOTE: DepthStencil.FFI_USE MUST be set to true to get DepthStencil used.
	DepthStencil GPUDepthStencilState
	// Multisample is "GPURenderPipelineDescriptor.multisample"
	//
	// Optional, defaults to {}.
	//
	// NOTE: Multisample.FFI_USE MUST be set to true to get Multisample used.
	Multisample GPUMultisampleState
	// Fragment is "GPURenderPipelineDescriptor.fragment"
	//
	// Optional
	//
	// NOTE: Fragment.FFI_USE MUST be set to true to get Fragment used.
	Fragment GPUFragmentState
	// Layout is "GPURenderPipelineDescriptor.layout"
	//
	// Required
	Layout OneOf_GPUPipelineLayout_GPUAutoLayoutMode
	// Label is "GPURenderPipelineDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPURenderPipelineDescriptor with all fields set.
func (p GPURenderPipelineDescriptor) FromRef(ref js.Ref) GPURenderPipelineDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPURenderPipelineDescriptor in the application heap.
func (p GPURenderPipelineDescriptor) New() js.Ref {
	return bindings.GPURenderPipelineDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPURenderPipelineDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPURenderPipelineDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPURenderPipelineDescriptor) Update(ref js.Ref) {
	bindings.GPURenderPipelineDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUColorDict struct {
	// R is "GPUColorDict.r"
	//
	// Required
	R float64
	// G is "GPUColorDict.g"
	//
	// Required
	G float64
	// B is "GPUColorDict.b"
	//
	// Required
	B float64
	// A is "GPUColorDict.a"
	//
	// Required
	A float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUColorDict with all fields set.
func (p GPUColorDict) FromRef(ref js.Ref) GPUColorDict {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUColorDict in the application heap.
func (p GPUColorDict) New() js.Ref {
	return bindings.GPUColorDictJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUColorDict) UpdateFrom(ref js.Ref) {
	bindings.GPUColorDictJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUColorDict) Update(ref js.Ref) {
	bindings.GPUColorDictJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_ArrayFloat64_GPUColorDict struct {
	ref js.Ref
}

func (x OneOf_ArrayFloat64_GPUColorDict) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ArrayFloat64_GPUColorDict) Free() {
	x.ref.Free()
}

func (x OneOf_ArrayFloat64_GPUColorDict) FromRef(ref js.Ref) OneOf_ArrayFloat64_GPUColorDict {
	return OneOf_ArrayFloat64_GPUColorDict{
		ref: ref,
	}
}

func (x OneOf_ArrayFloat64_GPUColorDict) ArrayFloat64() js.Array[float64] {
	return js.Array[float64]{}.FromRef(x.ref)
}

func (x OneOf_ArrayFloat64_GPUColorDict) GPUColorDict() GPUColorDict {
	var ret GPUColorDict
	ret.UpdateFrom(x.ref)
	return ret
}

type GPUColor = OneOf_ArrayFloat64_GPUColorDict

type GPURenderBundle struct {
	ref js.Ref
}

func (this GPURenderBundle) Once() GPURenderBundle {
	this.Ref().Once()
	return this
}

func (this GPURenderBundle) Ref() js.Ref {
	return this.ref
}

func (this GPURenderBundle) FromRef(ref js.Ref) GPURenderBundle {
	this.ref = ref
	return this
}

func (this GPURenderBundle) Free() {
	this.Ref().Free()
}

// Label returns the value of property "GPURenderBundle.label".
//
// It returns ok=false if there is no such property.
func (this GPURenderBundle) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPURenderBundleLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPURenderBundle.label" to val.
//
// It returns false if the property cannot be set.
func (this GPURenderBundle) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPURenderBundleLabel(
		this.Ref(),
		val.Ref(),
	)
}

type GPUSignedOffset32 int32

type GPUBufferDynamicOffset uint32

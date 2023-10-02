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

type OneOf_GPUBuffer_GPUTexture struct {
	ref js.Ref
}

func (x OneOf_GPUBuffer_GPUTexture) Ref() js.Ref {
	return x.ref
}

func (x OneOf_GPUBuffer_GPUTexture) Free() {
	x.ref.Free()
}

func (x OneOf_GPUBuffer_GPUTexture) FromRef(ref js.Ref) OneOf_GPUBuffer_GPUTexture {
	return OneOf_GPUBuffer_GPUTexture{
		ref: ref,
	}
}

func (x OneOf_GPUBuffer_GPUTexture) GPUBuffer() GPUBuffer {
	return GPUBuffer{}.FromRef(x.ref)
}

func (x OneOf_GPUBuffer_GPUTexture) GPUTexture() GPUTexture {
	return GPUTexture{}.FromRef(x.ref)
}

type MLGPUResource = OneOf_GPUBuffer_GPUTexture

type MLNamedGPUResources = js.Record[MLGPUResource]

type MLCommandEncoder struct {
	ref js.Ref
}

func (this MLCommandEncoder) Once() MLCommandEncoder {
	this.Ref().Once()
	return this
}

func (this MLCommandEncoder) Ref() js.Ref {
	return this.ref
}

func (this MLCommandEncoder) FromRef(ref js.Ref) MLCommandEncoder {
	this.ref = ref
	return this
}

func (this MLCommandEncoder) Free() {
	this.Ref().Free()
}

// Dispatch calls the method "MLCommandEncoder.dispatch".
//
// The returned bool will be false if there is no such method.
func (this MLCommandEncoder) Dispatch(graph MLGraph, inputs MLNamedGPUResources, outputs MLNamedGPUResources) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMLCommandEncoderDispatch(
		this.Ref(), js.Pointer(&_ok),
		graph.Ref(),
		inputs.Ref(),
		outputs.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DispatchFunc returns the method "MLCommandEncoder.dispatch".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLCommandEncoder) DispatchFunc() (fn js.Func[func(graph MLGraph, inputs MLNamedGPUResources, outputs MLNamedGPUResources)]) {
	return fn.FromRef(
		bindings.MLCommandEncoderDispatchFunc(
			this.Ref(),
		),
	)
}

// Finish calls the method "MLCommandEncoder.finish".
//
// The returned bool will be false if there is no such method.
func (this MLCommandEncoder) Finish(descriptor GPUCommandBufferDescriptor) (GPUCommandBuffer, bool) {
	var _ok bool
	_ret := bindings.CallMLCommandEncoderFinish(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPUCommandBuffer{}.FromRef(_ret), _ok
}

// FinishFunc returns the method "MLCommandEncoder.finish".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLCommandEncoder) FinishFunc() (fn js.Func[func(descriptor GPUCommandBufferDescriptor) GPUCommandBuffer]) {
	return fn.FromRef(
		bindings.MLCommandEncoderFinishFunc(
			this.Ref(),
		),
	)
}

// Finish1 calls the method "MLCommandEncoder.finish".
//
// The returned bool will be false if there is no such method.
func (this MLCommandEncoder) Finish1() (GPUCommandBuffer, bool) {
	var _ok bool
	_ret := bindings.CallMLCommandEncoderFinish1(
		this.Ref(), js.Pointer(&_ok),
	)

	return GPUCommandBuffer{}.FromRef(_ret), _ok
}

// Finish1Func returns the method "MLCommandEncoder.finish".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLCommandEncoder) Finish1Func() (fn js.Func[func() GPUCommandBuffer]) {
	return fn.FromRef(
		bindings.MLCommandEncoderFinish1Func(
			this.Ref(),
		),
	)
}

// InitializeGraph calls the method "MLCommandEncoder.initializeGraph".
//
// The returned bool will be false if there is no such method.
func (this MLCommandEncoder) InitializeGraph(graph MLGraph) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMLCommandEncoderInitializeGraph(
		this.Ref(), js.Pointer(&_ok),
		graph.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitializeGraphFunc returns the method "MLCommandEncoder.initializeGraph".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLCommandEncoder) InitializeGraphFunc() (fn js.Func[func(graph MLGraph)]) {
	return fn.FromRef(
		bindings.MLCommandEncoderInitializeGraphFunc(
			this.Ref(),
		),
	)
}

type MLContext struct {
	ref js.Ref
}

func (this MLContext) Once() MLContext {
	this.Ref().Once()
	return this
}

func (this MLContext) Ref() js.Ref {
	return this.ref
}

func (this MLContext) FromRef(ref js.Ref) MLContext {
	this.ref = ref
	return this
}

func (this MLContext) Free() {
	this.Ref().Free()
}

// Compute calls the method "MLContext.compute".
//
// The returned bool will be false if there is no such method.
func (this MLContext) Compute(graph MLGraph, inputs MLNamedArrayBufferViews, outputs MLNamedArrayBufferViews) (js.Promise[MLComputeResult], bool) {
	var _ok bool
	_ret := bindings.CallMLContextCompute(
		this.Ref(), js.Pointer(&_ok),
		graph.Ref(),
		inputs.Ref(),
		outputs.Ref(),
	)

	return js.Promise[MLComputeResult]{}.FromRef(_ret), _ok
}

// ComputeFunc returns the method "MLContext.compute".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLContext) ComputeFunc() (fn js.Func[func(graph MLGraph, inputs MLNamedArrayBufferViews, outputs MLNamedArrayBufferViews) js.Promise[MLComputeResult]]) {
	return fn.FromRef(
		bindings.MLContextComputeFunc(
			this.Ref(),
		),
	)
}

// ComputeSync calls the method "MLContext.computeSync".
//
// The returned bool will be false if there is no such method.
func (this MLContext) ComputeSync(graph MLGraph, inputs MLNamedArrayBufferViews, outputs MLNamedArrayBufferViews) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMLContextComputeSync(
		this.Ref(), js.Pointer(&_ok),
		graph.Ref(),
		inputs.Ref(),
		outputs.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ComputeSyncFunc returns the method "MLContext.computeSync".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLContext) ComputeSyncFunc() (fn js.Func[func(graph MLGraph, inputs MLNamedArrayBufferViews, outputs MLNamedArrayBufferViews)]) {
	return fn.FromRef(
		bindings.MLContextComputeSyncFunc(
			this.Ref(),
		),
	)
}

// CreateCommandEncoder calls the method "MLContext.createCommandEncoder".
//
// The returned bool will be false if there is no such method.
func (this MLContext) CreateCommandEncoder() (MLCommandEncoder, bool) {
	var _ok bool
	_ret := bindings.CallMLContextCreateCommandEncoder(
		this.Ref(), js.Pointer(&_ok),
	)

	return MLCommandEncoder{}.FromRef(_ret), _ok
}

// CreateCommandEncoderFunc returns the method "MLContext.createCommandEncoder".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLContext) CreateCommandEncoderFunc() (fn js.Func[func() MLCommandEncoder]) {
	return fn.FromRef(
		bindings.MLContextCreateCommandEncoderFunc(
			this.Ref(),
		),
	)
}

type MLDeviceType uint32

const (
	_ MLDeviceType = iota

	MLDeviceType_CPU
	MLDeviceType_GPU
)

func (MLDeviceType) FromRef(str js.Ref) MLDeviceType {
	return MLDeviceType(bindings.ConstOfMLDeviceType(str))
}

func (x MLDeviceType) String() (string, bool) {
	switch x {
	case MLDeviceType_CPU:
		return "cpu", true
	case MLDeviceType_GPU:
		return "gpu", true
	default:
		return "", false
	}
}

type MLPowerPreference uint32

const (
	_ MLPowerPreference = iota

	MLPowerPreference_DEFAULT
	MLPowerPreference_HIGH_PERFORMANCE
	MLPowerPreference_LOW_POWER
)

func (MLPowerPreference) FromRef(str js.Ref) MLPowerPreference {
	return MLPowerPreference(bindings.ConstOfMLPowerPreference(str))
}

func (x MLPowerPreference) String() (string, bool) {
	switch x {
	case MLPowerPreference_DEFAULT:
		return "default", true
	case MLPowerPreference_HIGH_PERFORMANCE:
		return "high-performance", true
	case MLPowerPreference_LOW_POWER:
		return "low-power", true
	default:
		return "", false
	}
}

type MLContextOptions struct {
	// DeviceType is "MLContextOptions.deviceType"
	//
	// Optional, defaults to "cpu".
	DeviceType MLDeviceType
	// PowerPreference is "MLContextOptions.powerPreference"
	//
	// Optional, defaults to "default".
	PowerPreference MLPowerPreference

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLContextOptions with all fields set.
func (p MLContextOptions) FromRef(ref js.Ref) MLContextOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MLContextOptions MLContextOptions [// MLContextOptions] [0x140007d3180 0x140007d3220] 0x1400081f200 {0 0}} in the application heap.
func (p MLContextOptions) New() js.Ref {
	return bindings.MLContextOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MLContextOptions) UpdateFrom(ref js.Ref) {
	bindings.MLContextOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLContextOptions) Update(ref js.Ref) {
	bindings.MLContextOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ML struct {
	ref js.Ref
}

func (this ML) Once() ML {
	this.Ref().Once()
	return this
}

func (this ML) Ref() js.Ref {
	return this.ref
}

func (this ML) FromRef(ref js.Ref) ML {
	this.ref = ref
	return this
}

func (this ML) Free() {
	this.Ref().Free()
}

// CreateContext calls the method "ML.createContext".
//
// The returned bool will be false if there is no such method.
func (this ML) CreateContext(options MLContextOptions) (js.Promise[MLContext], bool) {
	var _ok bool
	_ret := bindings.CallMLCreateContext(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[MLContext]{}.FromRef(_ret), _ok
}

// CreateContextFunc returns the method "ML.createContext".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ML) CreateContextFunc() (fn js.Func[func(options MLContextOptions) js.Promise[MLContext]]) {
	return fn.FromRef(
		bindings.MLCreateContextFunc(
			this.Ref(),
		),
	)
}

// CreateContext1 calls the method "ML.createContext".
//
// The returned bool will be false if there is no such method.
func (this ML) CreateContext1() (js.Promise[MLContext], bool) {
	var _ok bool
	_ret := bindings.CallMLCreateContext1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[MLContext]{}.FromRef(_ret), _ok
}

// CreateContext1Func returns the method "ML.createContext".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ML) CreateContext1Func() (fn js.Func[func() js.Promise[MLContext]]) {
	return fn.FromRef(
		bindings.MLCreateContext1Func(
			this.Ref(),
		),
	)
}

// CreateContext2 calls the method "ML.createContext".
//
// The returned bool will be false if there is no such method.
func (this ML) CreateContext2(gpuDevice GPUDevice) (js.Promise[MLContext], bool) {
	var _ok bool
	_ret := bindings.CallMLCreateContext2(
		this.Ref(), js.Pointer(&_ok),
		gpuDevice.Ref(),
	)

	return js.Promise[MLContext]{}.FromRef(_ret), _ok
}

// CreateContext2Func returns the method "ML.createContext".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ML) CreateContext2Func() (fn js.Func[func(gpuDevice GPUDevice) js.Promise[MLContext]]) {
	return fn.FromRef(
		bindings.MLCreateContext2Func(
			this.Ref(),
		),
	)
}

// CreateContextSync calls the method "ML.createContextSync".
//
// The returned bool will be false if there is no such method.
func (this ML) CreateContextSync(options MLContextOptions) (MLContext, bool) {
	var _ok bool
	_ret := bindings.CallMLCreateContextSync(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return MLContext{}.FromRef(_ret), _ok
}

// CreateContextSyncFunc returns the method "ML.createContextSync".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ML) CreateContextSyncFunc() (fn js.Func[func(options MLContextOptions) MLContext]) {
	return fn.FromRef(
		bindings.MLCreateContextSyncFunc(
			this.Ref(),
		),
	)
}

// CreateContextSync1 calls the method "ML.createContextSync".
//
// The returned bool will be false if there is no such method.
func (this ML) CreateContextSync1() (MLContext, bool) {
	var _ok bool
	_ret := bindings.CallMLCreateContextSync1(
		this.Ref(), js.Pointer(&_ok),
	)

	return MLContext{}.FromRef(_ret), _ok
}

// CreateContextSync1Func returns the method "ML.createContextSync".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ML) CreateContextSync1Func() (fn js.Func[func() MLContext]) {
	return fn.FromRef(
		bindings.MLCreateContextSync1Func(
			this.Ref(),
		),
	)
}

// CreateContextSync2 calls the method "ML.createContextSync".
//
// The returned bool will be false if there is no such method.
func (this ML) CreateContextSync2(gpuDevice GPUDevice) (MLContext, bool) {
	var _ok bool
	_ret := bindings.CallMLCreateContextSync2(
		this.Ref(), js.Pointer(&_ok),
		gpuDevice.Ref(),
	)

	return MLContext{}.FromRef(_ret), _ok
}

// CreateContextSync2Func returns the method "ML.createContextSync".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ML) CreateContextSync2Func() (fn js.Func[func(gpuDevice GPUDevice) MLContext]) {
	return fn.FromRef(
		bindings.MLCreateContextSync2Func(
			this.Ref(),
		),
	)
}

type ConnectionType uint32

const (
	_ ConnectionType = iota

	ConnectionType_BLUETOOTH
	ConnectionType_CELLULAR
	ConnectionType_ETHERNET
	ConnectionType_MIXED
	ConnectionType_NONE
	ConnectionType_OTHER
	ConnectionType_UNKNOWN
	ConnectionType_WIFI
	ConnectionType_WIMAX
)

func (ConnectionType) FromRef(str js.Ref) ConnectionType {
	return ConnectionType(bindings.ConstOfConnectionType(str))
}

func (x ConnectionType) String() (string, bool) {
	switch x {
	case ConnectionType_BLUETOOTH:
		return "bluetooth", true
	case ConnectionType_CELLULAR:
		return "cellular", true
	case ConnectionType_ETHERNET:
		return "ethernet", true
	case ConnectionType_MIXED:
		return "mixed", true
	case ConnectionType_NONE:
		return "none", true
	case ConnectionType_OTHER:
		return "other", true
	case ConnectionType_UNKNOWN:
		return "unknown", true
	case ConnectionType_WIFI:
		return "wifi", true
	case ConnectionType_WIMAX:
		return "wimax", true
	default:
		return "", false
	}
}

type EffectiveConnectionType uint32

const (
	_ EffectiveConnectionType = iota

	EffectiveConnectionType_2G
	EffectiveConnectionType_3G
	EffectiveConnectionType_4G
	EffectiveConnectionType_SLOW_2G
)

func (EffectiveConnectionType) FromRef(str js.Ref) EffectiveConnectionType {
	return EffectiveConnectionType(bindings.ConstOfEffectiveConnectionType(str))
}

func (x EffectiveConnectionType) String() (string, bool) {
	switch x {
	case EffectiveConnectionType_2G:
		return "2g", true
	case EffectiveConnectionType_3G:
		return "3g", true
	case EffectiveConnectionType_4G:
		return "4g", true
	case EffectiveConnectionType_SLOW_2G:
		return "slow-2g", true
	default:
		return "", false
	}
}

type Megabit float64

type Millisecond uint64

type NetworkInformation struct {
	EventTarget
}

func (this NetworkInformation) Once() NetworkInformation {
	this.Ref().Once()
	return this
}

func (this NetworkInformation) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this NetworkInformation) FromRef(ref js.Ref) NetworkInformation {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this NetworkInformation) Free() {
	this.Ref().Free()
}

// Type returns the value of property "NetworkInformation.type".
//
// The returned bool will be false if there is no such property.
func (this NetworkInformation) Type() (ConnectionType, bool) {
	var _ok bool
	_ret := bindings.GetNetworkInformationType(
		this.Ref(), js.Pointer(&_ok),
	)
	return ConnectionType(_ret), _ok
}

// EffectiveType returns the value of property "NetworkInformation.effectiveType".
//
// The returned bool will be false if there is no such property.
func (this NetworkInformation) EffectiveType() (EffectiveConnectionType, bool) {
	var _ok bool
	_ret := bindings.GetNetworkInformationEffectiveType(
		this.Ref(), js.Pointer(&_ok),
	)
	return EffectiveConnectionType(_ret), _ok
}

// DownlinkMax returns the value of property "NetworkInformation.downlinkMax".
//
// The returned bool will be false if there is no such property.
func (this NetworkInformation) DownlinkMax() (Megabit, bool) {
	var _ok bool
	_ret := bindings.GetNetworkInformationDownlinkMax(
		this.Ref(), js.Pointer(&_ok),
	)
	return Megabit(_ret), _ok
}

// Downlink returns the value of property "NetworkInformation.downlink".
//
// The returned bool will be false if there is no such property.
func (this NetworkInformation) Downlink() (Megabit, bool) {
	var _ok bool
	_ret := bindings.GetNetworkInformationDownlink(
		this.Ref(), js.Pointer(&_ok),
	)
	return Megabit(_ret), _ok
}

// Rtt returns the value of property "NetworkInformation.rtt".
//
// The returned bool will be false if there is no such property.
func (this NetworkInformation) Rtt() (Millisecond, bool) {
	var _ok bool
	_ret := bindings.GetNetworkInformationRtt(
		this.Ref(), js.Pointer(&_ok),
	)
	return Millisecond(_ret), _ok
}

// SaveData returns the value of property "NetworkInformation.saveData".
//
// The returned bool will be false if there is no such property.
func (this NetworkInformation) SaveData() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNetworkInformationSaveData(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

type GPUFeatureName uint32

const (
	_ GPUFeatureName = iota

	GPUFeatureName_DEPTH_CLIP_CONTROL
	GPUFeatureName_DEPTH_32FLOAT_STENCIL8
	GPUFeatureName_TEXTURE_COMPRESSION_BC
	GPUFeatureName_TEXTURE_COMPRESSION_ETC2
	GPUFeatureName_TEXTURE_COMPRESSION_ASTC
	GPUFeatureName_TIMESTAMP_QUERY
	GPUFeatureName_INDIRECT_FIRST_INSTANCE
	GPUFeatureName_SHADER_F16
	GPUFeatureName_RG_11B10UFLOAT_RENDERABLE
	GPUFeatureName_BGRA_8UNORM_STORAGE
	GPUFeatureName_FLOAT32_FILTERABLE
)

func (GPUFeatureName) FromRef(str js.Ref) GPUFeatureName {
	return GPUFeatureName(bindings.ConstOfGPUFeatureName(str))
}

func (x GPUFeatureName) String() (string, bool) {
	switch x {
	case GPUFeatureName_DEPTH_CLIP_CONTROL:
		return "depth-clip-control", true
	case GPUFeatureName_DEPTH_32FLOAT_STENCIL8:
		return "depth32float-stencil8", true
	case GPUFeatureName_TEXTURE_COMPRESSION_BC:
		return "texture-compression-bc", true
	case GPUFeatureName_TEXTURE_COMPRESSION_ETC2:
		return "texture-compression-etc2", true
	case GPUFeatureName_TEXTURE_COMPRESSION_ASTC:
		return "texture-compression-astc", true
	case GPUFeatureName_TIMESTAMP_QUERY:
		return "timestamp-query", true
	case GPUFeatureName_INDIRECT_FIRST_INSTANCE:
		return "indirect-first-instance", true
	case GPUFeatureName_SHADER_F16:
		return "shader-f16", true
	case GPUFeatureName_RG_11B10UFLOAT_RENDERABLE:
		return "rg11b10ufloat-renderable", true
	case GPUFeatureName_BGRA_8UNORM_STORAGE:
		return "bgra8unorm-storage", true
	case GPUFeatureName_FLOAT32_FILTERABLE:
		return "float32-filterable", true
	default:
		return "", false
	}
}

type GPUQueueDescriptor struct {
	// Label is "GPUQueueDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUQueueDescriptor with all fields set.
func (p GPUQueueDescriptor) FromRef(ref js.Ref) GPUQueueDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUQueueDescriptor GPUQueueDescriptor [// GPUQueueDescriptor] [0x140007d34a0] 0x1400081f5a8 {0 0}} in the application heap.
func (p GPUQueueDescriptor) New() js.Ref {
	return bindings.GPUQueueDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUQueueDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUQueueDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUQueueDescriptor) Update(ref js.Ref) {
	bindings.GPUQueueDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUDeviceDescriptor struct {
	// RequiredFeatures is "GPUDeviceDescriptor.requiredFeatures"
	//
	// Optional, defaults to [].
	RequiredFeatures js.Array[GPUFeatureName]
	// RequiredLimits is "GPUDeviceDescriptor.requiredLimits"
	//
	// Optional, defaults to {}.
	RequiredLimits js.Record[GPUSize64]
	// DefaultQueue is "GPUDeviceDescriptor.defaultQueue"
	//
	// Optional, defaults to {}.
	DefaultQueue GPUQueueDescriptor
	// Label is "GPUDeviceDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUDeviceDescriptor with all fields set.
func (p GPUDeviceDescriptor) FromRef(ref js.Ref) GPUDeviceDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPUDeviceDescriptor GPUDeviceDescriptor [// GPUDeviceDescriptor] [0x140007d3360 0x140007d3400 0x140007d3540 0x140007d35e0] 0x1400081f2f0 {0 0}} in the application heap.
func (p GPUDeviceDescriptor) New() js.Ref {
	return bindings.GPUDeviceDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUDeviceDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUDeviceDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUDeviceDescriptor) Update(ref js.Ref) {
	bindings.GPUDeviceDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUAdapterInfo struct {
	ref js.Ref
}

func (this GPUAdapterInfo) Once() GPUAdapterInfo {
	this.Ref().Once()
	return this
}

func (this GPUAdapterInfo) Ref() js.Ref {
	return this.ref
}

func (this GPUAdapterInfo) FromRef(ref js.Ref) GPUAdapterInfo {
	this.ref = ref
	return this
}

func (this GPUAdapterInfo) Free() {
	this.Ref().Free()
}

// Vendor returns the value of property "GPUAdapterInfo.vendor".
//
// The returned bool will be false if there is no such property.
func (this GPUAdapterInfo) Vendor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUAdapterInfoVendor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Architecture returns the value of property "GPUAdapterInfo.architecture".
//
// The returned bool will be false if there is no such property.
func (this GPUAdapterInfo) Architecture() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUAdapterInfoArchitecture(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Device returns the value of property "GPUAdapterInfo.device".
//
// The returned bool will be false if there is no such property.
func (this GPUAdapterInfo) Device() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUAdapterInfoDevice(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Description returns the value of property "GPUAdapterInfo.description".
//
// The returned bool will be false if there is no such property.
func (this GPUAdapterInfo) Description() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUAdapterInfoDescription(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type GPUAdapter struct {
	ref js.Ref
}

func (this GPUAdapter) Once() GPUAdapter {
	this.Ref().Once()
	return this
}

func (this GPUAdapter) Ref() js.Ref {
	return this.ref
}

func (this GPUAdapter) FromRef(ref js.Ref) GPUAdapter {
	this.ref = ref
	return this
}

func (this GPUAdapter) Free() {
	this.Ref().Free()
}

// Features returns the value of property "GPUAdapter.features".
//
// The returned bool will be false if there is no such property.
func (this GPUAdapter) Features() (GPUSupportedFeatures, bool) {
	var _ok bool
	_ret := bindings.GetGPUAdapterFeatures(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUSupportedFeatures{}.FromRef(_ret), _ok
}

// Limits returns the value of property "GPUAdapter.limits".
//
// The returned bool will be false if there is no such property.
func (this GPUAdapter) Limits() (GPUSupportedLimits, bool) {
	var _ok bool
	_ret := bindings.GetGPUAdapterLimits(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUSupportedLimits{}.FromRef(_ret), _ok
}

// IsFallbackAdapter returns the value of property "GPUAdapter.isFallbackAdapter".
//
// The returned bool will be false if there is no such property.
func (this GPUAdapter) IsFallbackAdapter() (bool, bool) {
	var _ok bool
	_ret := bindings.GetGPUAdapterIsFallbackAdapter(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// RequestDevice calls the method "GPUAdapter.requestDevice".
//
// The returned bool will be false if there is no such method.
func (this GPUAdapter) RequestDevice(descriptor GPUDeviceDescriptor) (js.Promise[GPUDevice], bool) {
	var _ok bool
	_ret := bindings.CallGPUAdapterRequestDevice(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return js.Promise[GPUDevice]{}.FromRef(_ret), _ok
}

// RequestDeviceFunc returns the method "GPUAdapter.requestDevice".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUAdapter) RequestDeviceFunc() (fn js.Func[func(descriptor GPUDeviceDescriptor) js.Promise[GPUDevice]]) {
	return fn.FromRef(
		bindings.GPUAdapterRequestDeviceFunc(
			this.Ref(),
		),
	)
}

// RequestDevice1 calls the method "GPUAdapter.requestDevice".
//
// The returned bool will be false if there is no such method.
func (this GPUAdapter) RequestDevice1() (js.Promise[GPUDevice], bool) {
	var _ok bool
	_ret := bindings.CallGPUAdapterRequestDevice1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[GPUDevice]{}.FromRef(_ret), _ok
}

// RequestDevice1Func returns the method "GPUAdapter.requestDevice".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUAdapter) RequestDevice1Func() (fn js.Func[func() js.Promise[GPUDevice]]) {
	return fn.FromRef(
		bindings.GPUAdapterRequestDevice1Func(
			this.Ref(),
		),
	)
}

// RequestAdapterInfo calls the method "GPUAdapter.requestAdapterInfo".
//
// The returned bool will be false if there is no such method.
func (this GPUAdapter) RequestAdapterInfo(unmaskHints js.Array[js.String]) (js.Promise[GPUAdapterInfo], bool) {
	var _ok bool
	_ret := bindings.CallGPUAdapterRequestAdapterInfo(
		this.Ref(), js.Pointer(&_ok),
		unmaskHints.Ref(),
	)

	return js.Promise[GPUAdapterInfo]{}.FromRef(_ret), _ok
}

// RequestAdapterInfoFunc returns the method "GPUAdapter.requestAdapterInfo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUAdapter) RequestAdapterInfoFunc() (fn js.Func[func(unmaskHints js.Array[js.String]) js.Promise[GPUAdapterInfo]]) {
	return fn.FromRef(
		bindings.GPUAdapterRequestAdapterInfoFunc(
			this.Ref(),
		),
	)
}

// RequestAdapterInfo1 calls the method "GPUAdapter.requestAdapterInfo".
//
// The returned bool will be false if there is no such method.
func (this GPUAdapter) RequestAdapterInfo1() (js.Promise[GPUAdapterInfo], bool) {
	var _ok bool
	_ret := bindings.CallGPUAdapterRequestAdapterInfo1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[GPUAdapterInfo]{}.FromRef(_ret), _ok
}

// RequestAdapterInfo1Func returns the method "GPUAdapter.requestAdapterInfo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUAdapter) RequestAdapterInfo1Func() (fn js.Func[func() js.Promise[GPUAdapterInfo]]) {
	return fn.FromRef(
		bindings.GPUAdapterRequestAdapterInfo1Func(
			this.Ref(),
		),
	)
}

type GPUPowerPreference uint32

const (
	_ GPUPowerPreference = iota

	GPUPowerPreference_LOW_POWER
	GPUPowerPreference_HIGH_PERFORMANCE
)

func (GPUPowerPreference) FromRef(str js.Ref) GPUPowerPreference {
	return GPUPowerPreference(bindings.ConstOfGPUPowerPreference(str))
}

func (x GPUPowerPreference) String() (string, bool) {
	switch x {
	case GPUPowerPreference_LOW_POWER:
		return "low-power", true
	case GPUPowerPreference_HIGH_PERFORMANCE:
		return "high-performance", true
	default:
		return "", false
	}
}

type GPURequestAdapterOptions struct {
	// PowerPreference is "GPURequestAdapterOptions.powerPreference"
	//
	// Optional
	PowerPreference GPUPowerPreference
	// ForceFallbackAdapter is "GPURequestAdapterOptions.forceFallbackAdapter"
	//
	// Optional, defaults to false.
	ForceFallbackAdapter bool

	FFI_USE_ForceFallbackAdapter bool // for ForceFallbackAdapter.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPURequestAdapterOptions with all fields set.
func (p GPURequestAdapterOptions) FromRef(ref js.Ref) GPURequestAdapterOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GPURequestAdapterOptions GPURequestAdapterOptions [// GPURequestAdapterOptions] [0x140007d37c0 0x140007d3860 0x140007d3900] 0x1400081f608 {0 0}} in the application heap.
func (p GPURequestAdapterOptions) New() js.Ref {
	return bindings.GPURequestAdapterOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPURequestAdapterOptions) UpdateFrom(ref js.Ref) {
	bindings.GPURequestAdapterOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPURequestAdapterOptions) Update(ref js.Ref) {
	bindings.GPURequestAdapterOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type WGSLLanguageFeatures struct {
	ref js.Ref
}

func (this WGSLLanguageFeatures) Once() WGSLLanguageFeatures {
	this.Ref().Once()
	return this
}

func (this WGSLLanguageFeatures) Ref() js.Ref {
	return this.ref
}

func (this WGSLLanguageFeatures) FromRef(ref js.Ref) WGSLLanguageFeatures {
	this.ref = ref
	return this
}

func (this WGSLLanguageFeatures) Free() {
	this.Ref().Free()
}

type GPU struct {
	ref js.Ref
}

func (this GPU) Once() GPU {
	this.Ref().Once()
	return this
}

func (this GPU) Ref() js.Ref {
	return this.ref
}

func (this GPU) FromRef(ref js.Ref) GPU {
	this.ref = ref
	return this
}

func (this GPU) Free() {
	this.Ref().Free()
}

// WgslLanguageFeatures returns the value of property "GPU.wgslLanguageFeatures".
//
// The returned bool will be false if there is no such property.
func (this GPU) WgslLanguageFeatures() (WGSLLanguageFeatures, bool) {
	var _ok bool
	_ret := bindings.GetGPUWgslLanguageFeatures(
		this.Ref(), js.Pointer(&_ok),
	)
	return WGSLLanguageFeatures{}.FromRef(_ret), _ok
}

// RequestAdapter calls the method "GPU.requestAdapter".
//
// The returned bool will be false if there is no such method.
func (this GPU) RequestAdapter(options GPURequestAdapterOptions) (js.Promise[GPUAdapter], bool) {
	var _ok bool
	_ret := bindings.CallGPURequestAdapter(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[GPUAdapter]{}.FromRef(_ret), _ok
}

// RequestAdapterFunc returns the method "GPU.requestAdapter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPU) RequestAdapterFunc() (fn js.Func[func(options GPURequestAdapterOptions) js.Promise[GPUAdapter]]) {
	return fn.FromRef(
		bindings.GPURequestAdapterFunc(
			this.Ref(),
		),
	)
}

// RequestAdapter1 calls the method "GPU.requestAdapter".
//
// The returned bool will be false if there is no such method.
func (this GPU) RequestAdapter1() (js.Promise[GPUAdapter], bool) {
	var _ok bool
	_ret := bindings.CallGPURequestAdapter1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[GPUAdapter]{}.FromRef(_ret), _ok
}

// RequestAdapter1Func returns the method "GPU.requestAdapter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPU) RequestAdapter1Func() (fn js.Func[func() js.Promise[GPUAdapter]]) {
	return fn.FromRef(
		bindings.GPURequestAdapter1Func(
			this.Ref(),
		),
	)
}

// GetPreferredCanvasFormat calls the method "GPU.getPreferredCanvasFormat".
//
// The returned bool will be false if there is no such method.
func (this GPU) GetPreferredCanvasFormat() (GPUTextureFormat, bool) {
	var _ok bool
	_ret := bindings.CallGPUGetPreferredCanvasFormat(
		this.Ref(), js.Pointer(&_ok),
	)

	return GPUTextureFormat(_ret), _ok
}

// GetPreferredCanvasFormatFunc returns the method "GPU.getPreferredCanvasFormat".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPU) GetPreferredCanvasFormatFunc() (fn js.Func[func() GPUTextureFormat]) {
	return fn.FromRef(
		bindings.GPUGetPreferredCanvasFormatFunc(
			this.Ref(),
		),
	)
}

type Navigator struct {
	ref js.Ref
}

func (this Navigator) Once() Navigator {
	this.Ref().Once()
	return this
}

func (this Navigator) Ref() js.Ref {
	return this.ref
}

func (this Navigator) FromRef(ref js.Ref) Navigator {
	this.ref = ref
	return this
}

func (this Navigator) Free() {
	this.Ref().Free()
}

// Hid returns the value of property "Navigator.hid".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Hid() (HID, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorHid(
		this.Ref(), js.Pointer(&_ok),
	)
	return HID{}.FromRef(_ret), _ok
}

// WindowControlsOverlay returns the value of property "Navigator.windowControlsOverlay".
//
// The returned bool will be false if there is no such property.
func (this Navigator) WindowControlsOverlay() (WindowControlsOverlay, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorWindowControlsOverlay(
		this.Ref(), js.Pointer(&_ok),
	)
	return WindowControlsOverlay{}.FromRef(_ret), _ok
}

// Credentials returns the value of property "Navigator.credentials".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Credentials() (CredentialsContainer, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorCredentials(
		this.Ref(), js.Pointer(&_ok),
	)
	return CredentialsContainer{}.FromRef(_ret), _ok
}

// Clipboard returns the value of property "Navigator.clipboard".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Clipboard() (Clipboard, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorClipboard(
		this.Ref(), js.Pointer(&_ok),
	)
	return Clipboard{}.FromRef(_ret), _ok
}

// Geolocation returns the value of property "Navigator.geolocation".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Geolocation() (Geolocation, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorGeolocation(
		this.Ref(), js.Pointer(&_ok),
	)
	return Geolocation{}.FromRef(_ret), _ok
}

// Usb returns the value of property "Navigator.usb".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Usb() (USB, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorUsb(
		this.Ref(), js.Pointer(&_ok),
	)
	return USB{}.FromRef(_ret), _ok
}

// EpubReadingSystem returns the value of property "Navigator.epubReadingSystem".
//
// The returned bool will be false if there is no such property.
func (this Navigator) EpubReadingSystem() (EpubReadingSystem, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorEpubReadingSystem(
		this.Ref(), js.Pointer(&_ok),
	)
	return EpubReadingSystem{}.FromRef(_ret), _ok
}

// Xr returns the value of property "Navigator.xr".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Xr() (XRSystem, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorXr(
		this.Ref(), js.Pointer(&_ok),
	)
	return XRSystem{}.FromRef(_ret), _ok
}

// ServiceWorker returns the value of property "Navigator.serviceWorker".
//
// The returned bool will be false if there is no such property.
func (this Navigator) ServiceWorker() (ServiceWorkerContainer, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorServiceWorker(
		this.Ref(), js.Pointer(&_ok),
	)
	return ServiceWorkerContainer{}.FromRef(_ret), _ok
}

// MediaDevices returns the value of property "Navigator.mediaDevices".
//
// The returned bool will be false if there is no such property.
func (this Navigator) MediaDevices() (MediaDevices, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorMediaDevices(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaDevices{}.FromRef(_ret), _ok
}

// Bluetooth returns the value of property "Navigator.bluetooth".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Bluetooth() (Bluetooth, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorBluetooth(
		this.Ref(), js.Pointer(&_ok),
	)
	return Bluetooth{}.FromRef(_ret), _ok
}

// Serial returns the value of property "Navigator.serial".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Serial() (Serial, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorSerial(
		this.Ref(), js.Pointer(&_ok),
	)
	return Serial{}.FromRef(_ret), _ok
}

// MediaCapabilities returns the value of property "Navigator.mediaCapabilities".
//
// The returned bool will be false if there is no such property.
func (this Navigator) MediaCapabilities() (MediaCapabilities, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorMediaCapabilities(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaCapabilities{}.FromRef(_ret), _ok
}

// UserActivation returns the value of property "Navigator.userActivation".
//
// The returned bool will be false if there is no such property.
func (this Navigator) UserActivation() (UserActivation, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorUserActivation(
		this.Ref(), js.Pointer(&_ok),
	)
	return UserActivation{}.FromRef(_ret), _ok
}

// Permissions returns the value of property "Navigator.permissions".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Permissions() (Permissions, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorPermissions(
		this.Ref(), js.Pointer(&_ok),
	)
	return Permissions{}.FromRef(_ret), _ok
}

// Contacts returns the value of property "Navigator.contacts".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Contacts() (ContactsManager, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorContacts(
		this.Ref(), js.Pointer(&_ok),
	)
	return ContactsManager{}.FromRef(_ret), _ok
}

// Keyboard returns the value of property "Navigator.keyboard".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Keyboard() (Keyboard, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorKeyboard(
		this.Ref(), js.Pointer(&_ok),
	)
	return Keyboard{}.FromRef(_ret), _ok
}

// MediaSession returns the value of property "Navigator.mediaSession".
//
// The returned bool will be false if there is no such property.
func (this Navigator) MediaSession() (MediaSession, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorMediaSession(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaSession{}.FromRef(_ret), _ok
}

// DevicePosture returns the value of property "Navigator.devicePosture".
//
// The returned bool will be false if there is no such property.
func (this Navigator) DevicePosture() (DevicePosture, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorDevicePosture(
		this.Ref(), js.Pointer(&_ok),
	)
	return DevicePosture{}.FromRef(_ret), _ok
}

// MaxTouchPoints returns the value of property "Navigator.maxTouchPoints".
//
// The returned bool will be false if there is no such property.
func (this Navigator) MaxTouchPoints() (int32, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorMaxTouchPoints(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Scheduling returns the value of property "Navigator.scheduling".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Scheduling() (Scheduling, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorScheduling(
		this.Ref(), js.Pointer(&_ok),
	)
	return Scheduling{}.FromRef(_ret), _ok
}

// WakeLock returns the value of property "Navigator.wakeLock".
//
// The returned bool will be false if there is no such property.
func (this Navigator) WakeLock() (WakeLock, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorWakeLock(
		this.Ref(), js.Pointer(&_ok),
	)
	return WakeLock{}.FromRef(_ret), _ok
}

// Ink returns the value of property "Navigator.ink".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Ink() (Ink, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorInk(
		this.Ref(), js.Pointer(&_ok),
	)
	return Ink{}.FromRef(_ret), _ok
}

// Presentation returns the value of property "Navigator.presentation".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Presentation() (Presentation, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorPresentation(
		this.Ref(), js.Pointer(&_ok),
	)
	return Presentation{}.FromRef(_ret), _ok
}

// VirtualKeyboard returns the value of property "Navigator.virtualKeyboard".
//
// The returned bool will be false if there is no such property.
func (this Navigator) VirtualKeyboard() (VirtualKeyboard, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorVirtualKeyboard(
		this.Ref(), js.Pointer(&_ok),
	)
	return VirtualKeyboard{}.FromRef(_ret), _ok
}

// UserAgentData returns the value of property "Navigator.userAgentData".
//
// The returned bool will be false if there is no such property.
func (this Navigator) UserAgentData() (NavigatorUAData, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorUserAgentData(
		this.Ref(), js.Pointer(&_ok),
	)
	return NavigatorUAData{}.FromRef(_ret), _ok
}

// HardwareConcurrency returns the value of property "Navigator.hardwareConcurrency".
//
// The returned bool will be false if there is no such property.
func (this Navigator) HardwareConcurrency() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorHardwareConcurrency(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// DeviceMemory returns the value of property "Navigator.deviceMemory".
//
// The returned bool will be false if there is no such property.
func (this Navigator) DeviceMemory() (float64, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorDeviceMemory(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Storage returns the value of property "Navigator.storage".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Storage() (StorageManager, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorStorage(
		this.Ref(), js.Pointer(&_ok),
	)
	return StorageManager{}.FromRef(_ret), _ok
}

// StorageBuckets returns the value of property "Navigator.storageBuckets".
//
// The returned bool will be false if there is no such property.
func (this Navigator) StorageBuckets() (StorageBucketManager, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorStorageBuckets(
		this.Ref(), js.Pointer(&_ok),
	)
	return StorageBucketManager{}.FromRef(_ret), _ok
}

// Locks returns the value of property "Navigator.locks".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Locks() (LockManager, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorLocks(
		this.Ref(), js.Pointer(&_ok),
	)
	return LockManager{}.FromRef(_ret), _ok
}

// AppCodeName returns the value of property "Navigator.appCodeName".
//
// The returned bool will be false if there is no such property.
func (this Navigator) AppCodeName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorAppCodeName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AppName returns the value of property "Navigator.appName".
//
// The returned bool will be false if there is no such property.
func (this Navigator) AppName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorAppName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AppVersion returns the value of property "Navigator.appVersion".
//
// The returned bool will be false if there is no such property.
func (this Navigator) AppVersion() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorAppVersion(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Platform returns the value of property "Navigator.platform".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Platform() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorPlatform(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Product returns the value of property "Navigator.product".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Product() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorProduct(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ProductSub returns the value of property "Navigator.productSub".
//
// The returned bool will be false if there is no such property.
func (this Navigator) ProductSub() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorProductSub(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// UserAgent returns the value of property "Navigator.userAgent".
//
// The returned bool will be false if there is no such property.
func (this Navigator) UserAgent() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorUserAgent(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Vendor returns the value of property "Navigator.vendor".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Vendor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorVendor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// VendorSub returns the value of property "Navigator.vendorSub".
//
// The returned bool will be false if there is no such property.
func (this Navigator) VendorSub() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorVendorSub(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Oscpu returns the value of property "Navigator.oscpu".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Oscpu() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorOscpu(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Language returns the value of property "Navigator.language".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Language() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorLanguage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Languages returns the value of property "Navigator.languages".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Languages() (js.FrozenArray[js.String], bool) {
	var _ok bool
	_ret := bindings.GetNavigatorLanguages(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[js.String]{}.FromRef(_ret), _ok
}

// OnLine returns the value of property "Navigator.onLine".
//
// The returned bool will be false if there is no such property.
func (this Navigator) OnLine() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorOnLine(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// CookieEnabled returns the value of property "Navigator.cookieEnabled".
//
// The returned bool will be false if there is no such property.
func (this Navigator) CookieEnabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorCookieEnabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Plugins returns the value of property "Navigator.plugins".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Plugins() (PluginArray, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorPlugins(
		this.Ref(), js.Pointer(&_ok),
	)
	return PluginArray{}.FromRef(_ret), _ok
}

// MimeTypes returns the value of property "Navigator.mimeTypes".
//
// The returned bool will be false if there is no such property.
func (this Navigator) MimeTypes() (MimeTypeArray, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorMimeTypes(
		this.Ref(), js.Pointer(&_ok),
	)
	return MimeTypeArray{}.FromRef(_ret), _ok
}

// PdfViewerEnabled returns the value of property "Navigator.pdfViewerEnabled".
//
// The returned bool will be false if there is no such property.
func (this Navigator) PdfViewerEnabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorPdfViewerEnabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Webdriver returns the value of property "Navigator.webdriver".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Webdriver() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorWebdriver(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Ml returns the value of property "Navigator.ml".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Ml() (ML, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorMl(
		this.Ref(), js.Pointer(&_ok),
	)
	return ML{}.FromRef(_ret), _ok
}

// Connection returns the value of property "Navigator.connection".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Connection() (NetworkInformation, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorConnection(
		this.Ref(), js.Pointer(&_ok),
	)
	return NetworkInformation{}.FromRef(_ret), _ok
}

// Gpu returns the value of property "Navigator.gpu".
//
// The returned bool will be false if there is no such property.
func (this Navigator) Gpu() (GPU, bool) {
	var _ok bool
	_ret := bindings.GetNavigatorGpu(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPU{}.FromRef(_ret), _ok
}

// UpdateAdInterestGroups calls the method "Navigator.updateAdInterestGroups".
//
// The returned bool will be false if there is no such method.
func (this Navigator) UpdateAdInterestGroups() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallNavigatorUpdateAdInterestGroups(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UpdateAdInterestGroupsFunc returns the method "Navigator.updateAdInterestGroups".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) UpdateAdInterestGroupsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.NavigatorUpdateAdInterestGroupsFunc(
			this.Ref(),
		),
	)
}

// SendBeacon calls the method "Navigator.sendBeacon".
//
// The returned bool will be false if there is no such method.
func (this Navigator) SendBeacon(url js.String, data BodyInit) (bool, bool) {
	var _ok bool
	_ret := bindings.CallNavigatorSendBeacon(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
		data.Ref(),
	)

	return _ret == js.True, _ok
}

// SendBeaconFunc returns the method "Navigator.sendBeacon".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) SendBeaconFunc() (fn js.Func[func(url js.String, data BodyInit) bool]) {
	return fn.FromRef(
		bindings.NavigatorSendBeaconFunc(
			this.Ref(),
		),
	)
}

// SendBeacon1 calls the method "Navigator.sendBeacon".
//
// The returned bool will be false if there is no such method.
func (this Navigator) SendBeacon1(url js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallNavigatorSendBeacon1(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
	)

	return _ret == js.True, _ok
}

// SendBeacon1Func returns the method "Navigator.sendBeacon".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) SendBeacon1Func() (fn js.Func[func(url js.String) bool]) {
	return fn.FromRef(
		bindings.NavigatorSendBeacon1Func(
			this.Ref(),
		),
	)
}

// GetBattery calls the method "Navigator.getBattery".
//
// The returned bool will be false if there is no such method.
func (this Navigator) GetBattery() (js.Promise[BatteryManager], bool) {
	var _ok bool
	_ret := bindings.CallNavigatorGetBattery(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[BatteryManager]{}.FromRef(_ret), _ok
}

// GetBatteryFunc returns the method "Navigator.getBattery".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) GetBatteryFunc() (fn js.Func[func() js.Promise[BatteryManager]]) {
	return fn.FromRef(
		bindings.NavigatorGetBatteryFunc(
			this.Ref(),
		),
	)
}

// Vibrate calls the method "Navigator.vibrate".
//
// The returned bool will be false if there is no such method.
func (this Navigator) Vibrate(pattern VibratePattern) (bool, bool) {
	var _ok bool
	_ret := bindings.CallNavigatorVibrate(
		this.Ref(), js.Pointer(&_ok),
		pattern.Ref(),
	)

	return _ret == js.True, _ok
}

// VibrateFunc returns the method "Navigator.vibrate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) VibrateFunc() (fn js.Func[func(pattern VibratePattern) bool]) {
	return fn.FromRef(
		bindings.NavigatorVibrateFunc(
			this.Ref(),
		),
	)
}

// GetGamepads calls the method "Navigator.getGamepads".
//
// The returned bool will be false if there is no such method.
func (this Navigator) GetGamepads() (js.Array[Gamepad], bool) {
	var _ok bool
	_ret := bindings.CallNavigatorGetGamepads(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[Gamepad]{}.FromRef(_ret), _ok
}

// GetGamepadsFunc returns the method "Navigator.getGamepads".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) GetGamepadsFunc() (fn js.Func[func() js.Array[Gamepad]]) {
	return fn.FromRef(
		bindings.NavigatorGetGamepadsFunc(
			this.Ref(),
		),
	)
}

// GetInstalledRelatedApps calls the method "Navigator.getInstalledRelatedApps".
//
// The returned bool will be false if there is no such method.
func (this Navigator) GetInstalledRelatedApps() (js.Promise[js.Array[RelatedApplication]], bool) {
	var _ok bool
	_ret := bindings.CallNavigatorGetInstalledRelatedApps(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[RelatedApplication]]{}.FromRef(_ret), _ok
}

// GetInstalledRelatedAppsFunc returns the method "Navigator.getInstalledRelatedApps".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) GetInstalledRelatedAppsFunc() (fn js.Func[func() js.Promise[js.Array[RelatedApplication]]]) {
	return fn.FromRef(
		bindings.NavigatorGetInstalledRelatedAppsFunc(
			this.Ref(),
		),
	)
}

// RequestMediaKeySystemAccess calls the method "Navigator.requestMediaKeySystemAccess".
//
// The returned bool will be false if there is no such method.
func (this Navigator) RequestMediaKeySystemAccess(keySystem js.String, supportedConfigurations js.Array[MediaKeySystemConfiguration]) (js.Promise[MediaKeySystemAccess], bool) {
	var _ok bool
	_ret := bindings.CallNavigatorRequestMediaKeySystemAccess(
		this.Ref(), js.Pointer(&_ok),
		keySystem.Ref(),
		supportedConfigurations.Ref(),
	)

	return js.Promise[MediaKeySystemAccess]{}.FromRef(_ret), _ok
}

// RequestMediaKeySystemAccessFunc returns the method "Navigator.requestMediaKeySystemAccess".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) RequestMediaKeySystemAccessFunc() (fn js.Func[func(keySystem js.String, supportedConfigurations js.Array[MediaKeySystemConfiguration]) js.Promise[MediaKeySystemAccess]]) {
	return fn.FromRef(
		bindings.NavigatorRequestMediaKeySystemAccessFunc(
			this.Ref(),
		),
	)
}

// JoinAdInterestGroup calls the method "Navigator.joinAdInterestGroup".
//
// The returned bool will be false if there is no such method.
func (this Navigator) JoinAdInterestGroup(group AuctionAdInterestGroup) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNavigatorJoinAdInterestGroup(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&group),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// JoinAdInterestGroupFunc returns the method "Navigator.joinAdInterestGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) JoinAdInterestGroupFunc() (fn js.Func[func(group AuctionAdInterestGroup) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigatorJoinAdInterestGroupFunc(
			this.Ref(),
		),
	)
}

// GetUserMedia calls the method "Navigator.getUserMedia".
//
// The returned bool will be false if there is no such method.
func (this Navigator) GetUserMedia(constraints MediaStreamConstraints, successCallback js.Func[func(stream MediaStream)], errorCallback js.Func[func(err DOMException)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallNavigatorGetUserMedia(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&constraints),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetUserMediaFunc returns the method "Navigator.getUserMedia".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) GetUserMediaFunc() (fn js.Func[func(constraints MediaStreamConstraints, successCallback js.Func[func(stream MediaStream)], errorCallback js.Func[func(err DOMException)])]) {
	return fn.FromRef(
		bindings.NavigatorGetUserMediaFunc(
			this.Ref(),
		),
	)
}

// LeaveAdInterestGroup calls the method "Navigator.leaveAdInterestGroup".
//
// The returned bool will be false if there is no such method.
func (this Navigator) LeaveAdInterestGroup(group AuctionAdInterestGroupKey) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNavigatorLeaveAdInterestGroup(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&group),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// LeaveAdInterestGroupFunc returns the method "Navigator.leaveAdInterestGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) LeaveAdInterestGroupFunc() (fn js.Func[func(group AuctionAdInterestGroupKey) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigatorLeaveAdInterestGroupFunc(
			this.Ref(),
		),
	)
}

// LeaveAdInterestGroup1 calls the method "Navigator.leaveAdInterestGroup".
//
// The returned bool will be false if there is no such method.
func (this Navigator) LeaveAdInterestGroup1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNavigatorLeaveAdInterestGroup1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// LeaveAdInterestGroup1Func returns the method "Navigator.leaveAdInterestGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) LeaveAdInterestGroup1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigatorLeaveAdInterestGroup1Func(
			this.Ref(),
		),
	)
}

// RunAdAuction calls the method "Navigator.runAdAuction".
//
// The returned bool will be false if there is no such method.
func (this Navigator) RunAdAuction(config AuctionAdConfig) (js.Promise[OneOf_String_FencedFrameConfig], bool) {
	var _ok bool
	_ret := bindings.CallNavigatorRunAdAuction(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&config),
	)

	return js.Promise[OneOf_String_FencedFrameConfig]{}.FromRef(_ret), _ok
}

// RunAdAuctionFunc returns the method "Navigator.runAdAuction".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) RunAdAuctionFunc() (fn js.Func[func(config AuctionAdConfig) js.Promise[OneOf_String_FencedFrameConfig]]) {
	return fn.FromRef(
		bindings.NavigatorRunAdAuctionFunc(
			this.Ref(),
		),
	)
}

// Share calls the method "Navigator.share".
//
// The returned bool will be false if there is no such method.
func (this Navigator) Share(data ShareData) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNavigatorShare(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&data),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ShareFunc returns the method "Navigator.share".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) ShareFunc() (fn js.Func[func(data ShareData) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigatorShareFunc(
			this.Ref(),
		),
	)
}

// Share1 calls the method "Navigator.share".
//
// The returned bool will be false if there is no such method.
func (this Navigator) Share1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNavigatorShare1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Share1Func returns the method "Navigator.share".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) Share1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigatorShare1Func(
			this.Ref(),
		),
	)
}

// CanShare calls the method "Navigator.canShare".
//
// The returned bool will be false if there is no such method.
func (this Navigator) CanShare(data ShareData) (bool, bool) {
	var _ok bool
	_ret := bindings.CallNavigatorCanShare(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&data),
	)

	return _ret == js.True, _ok
}

// CanShareFunc returns the method "Navigator.canShare".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) CanShareFunc() (fn js.Func[func(data ShareData) bool]) {
	return fn.FromRef(
		bindings.NavigatorCanShareFunc(
			this.Ref(),
		),
	)
}

// CanShare1 calls the method "Navigator.canShare".
//
// The returned bool will be false if there is no such method.
func (this Navigator) CanShare1() (bool, bool) {
	var _ok bool
	_ret := bindings.CallNavigatorCanShare1(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// CanShare1Func returns the method "Navigator.canShare".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) CanShare1Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.NavigatorCanShare1Func(
			this.Ref(),
		),
	)
}

// RequestMIDIAccess calls the method "Navigator.requestMIDIAccess".
//
// The returned bool will be false if there is no such method.
func (this Navigator) RequestMIDIAccess(options MIDIOptions) (js.Promise[MIDIAccess], bool) {
	var _ok bool
	_ret := bindings.CallNavigatorRequestMIDIAccess(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[MIDIAccess]{}.FromRef(_ret), _ok
}

// RequestMIDIAccessFunc returns the method "Navigator.requestMIDIAccess".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) RequestMIDIAccessFunc() (fn js.Func[func(options MIDIOptions) js.Promise[MIDIAccess]]) {
	return fn.FromRef(
		bindings.NavigatorRequestMIDIAccessFunc(
			this.Ref(),
		),
	)
}

// RequestMIDIAccess1 calls the method "Navigator.requestMIDIAccess".
//
// The returned bool will be false if there is no such method.
func (this Navigator) RequestMIDIAccess1() (js.Promise[MIDIAccess], bool) {
	var _ok bool
	_ret := bindings.CallNavigatorRequestMIDIAccess1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[MIDIAccess]{}.FromRef(_ret), _ok
}

// RequestMIDIAccess1Func returns the method "Navigator.requestMIDIAccess".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) RequestMIDIAccess1Func() (fn js.Func[func() js.Promise[MIDIAccess]]) {
	return fn.FromRef(
		bindings.NavigatorRequestMIDIAccess1Func(
			this.Ref(),
		),
	)
}

// GetAutoplayPolicy calls the method "Navigator.getAutoplayPolicy".
//
// The returned bool will be false if there is no such method.
func (this Navigator) GetAutoplayPolicy(typ AutoplayPolicyMediaType) (AutoplayPolicy, bool) {
	var _ok bool
	_ret := bindings.CallNavigatorGetAutoplayPolicy(
		this.Ref(), js.Pointer(&_ok),
		uint32(typ),
	)

	return AutoplayPolicy(_ret), _ok
}

// GetAutoplayPolicyFunc returns the method "Navigator.getAutoplayPolicy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) GetAutoplayPolicyFunc() (fn js.Func[func(typ AutoplayPolicyMediaType) AutoplayPolicy]) {
	return fn.FromRef(
		bindings.NavigatorGetAutoplayPolicyFunc(
			this.Ref(),
		),
	)
}

// GetAutoplayPolicy1 calls the method "Navigator.getAutoplayPolicy".
//
// The returned bool will be false if there is no such method.
func (this Navigator) GetAutoplayPolicy1(element HTMLMediaElement) (AutoplayPolicy, bool) {
	var _ok bool
	_ret := bindings.CallNavigatorGetAutoplayPolicy1(
		this.Ref(), js.Pointer(&_ok),
		element.Ref(),
	)

	return AutoplayPolicy(_ret), _ok
}

// GetAutoplayPolicy1Func returns the method "Navigator.getAutoplayPolicy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) GetAutoplayPolicy1Func() (fn js.Func[func(element HTMLMediaElement) AutoplayPolicy]) {
	return fn.FromRef(
		bindings.NavigatorGetAutoplayPolicy1Func(
			this.Ref(),
		),
	)
}

// GetAutoplayPolicy2 calls the method "Navigator.getAutoplayPolicy".
//
// The returned bool will be false if there is no such method.
func (this Navigator) GetAutoplayPolicy2(context AudioContext) (AutoplayPolicy, bool) {
	var _ok bool
	_ret := bindings.CallNavigatorGetAutoplayPolicy2(
		this.Ref(), js.Pointer(&_ok),
		context.Ref(),
	)

	return AutoplayPolicy(_ret), _ok
}

// GetAutoplayPolicy2Func returns the method "Navigator.getAutoplayPolicy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) GetAutoplayPolicy2Func() (fn js.Func[func(context AudioContext) AutoplayPolicy]) {
	return fn.FromRef(
		bindings.NavigatorGetAutoplayPolicy2Func(
			this.Ref(),
		),
	)
}

// TaintEnabled calls the method "Navigator.taintEnabled".
//
// The returned bool will be false if there is no such method.
func (this Navigator) TaintEnabled() (bool, bool) {
	var _ok bool
	_ret := bindings.CallNavigatorTaintEnabled(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// TaintEnabledFunc returns the method "Navigator.taintEnabled".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) TaintEnabledFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.NavigatorTaintEnabledFunc(
			this.Ref(),
		),
	)
}

// RegisterProtocolHandler calls the method "Navigator.registerProtocolHandler".
//
// The returned bool will be false if there is no such method.
func (this Navigator) RegisterProtocolHandler(scheme js.String, url js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallNavigatorRegisterProtocolHandler(
		this.Ref(), js.Pointer(&_ok),
		scheme.Ref(),
		url.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RegisterProtocolHandlerFunc returns the method "Navigator.registerProtocolHandler".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) RegisterProtocolHandlerFunc() (fn js.Func[func(scheme js.String, url js.String)]) {
	return fn.FromRef(
		bindings.NavigatorRegisterProtocolHandlerFunc(
			this.Ref(),
		),
	)
}

// UnregisterProtocolHandler calls the method "Navigator.unregisterProtocolHandler".
//
// The returned bool will be false if there is no such method.
func (this Navigator) UnregisterProtocolHandler(scheme js.String, url js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallNavigatorUnregisterProtocolHandler(
		this.Ref(), js.Pointer(&_ok),
		scheme.Ref(),
		url.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UnregisterProtocolHandlerFunc returns the method "Navigator.unregisterProtocolHandler".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) UnregisterProtocolHandlerFunc() (fn js.Func[func(scheme js.String, url js.String)]) {
	return fn.FromRef(
		bindings.NavigatorUnregisterProtocolHandlerFunc(
			this.Ref(),
		),
	)
}

// JavaEnabled calls the method "Navigator.javaEnabled".
//
// The returned bool will be false if there is no such method.
func (this Navigator) JavaEnabled() (bool, bool) {
	var _ok bool
	_ret := bindings.CallNavigatorJavaEnabled(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// JavaEnabledFunc returns the method "Navigator.javaEnabled".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) JavaEnabledFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.NavigatorJavaEnabledFunc(
			this.Ref(),
		),
	)
}

// SetAppBadge calls the method "Navigator.setAppBadge".
//
// The returned bool will be false if there is no such method.
func (this Navigator) SetAppBadge(contents uint64) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNavigatorSetAppBadge(
		this.Ref(), js.Pointer(&_ok),
		float64(contents),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetAppBadgeFunc returns the method "Navigator.setAppBadge".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) SetAppBadgeFunc() (fn js.Func[func(contents uint64) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigatorSetAppBadgeFunc(
			this.Ref(),
		),
	)
}

// SetAppBadge1 calls the method "Navigator.setAppBadge".
//
// The returned bool will be false if there is no such method.
func (this Navigator) SetAppBadge1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNavigatorSetAppBadge1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetAppBadge1Func returns the method "Navigator.setAppBadge".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) SetAppBadge1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigatorSetAppBadge1Func(
			this.Ref(),
		),
	)
}

// ClearAppBadge calls the method "Navigator.clearAppBadge".
//
// The returned bool will be false if there is no such method.
func (this Navigator) ClearAppBadge() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNavigatorClearAppBadge(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ClearAppBadgeFunc returns the method "Navigator.clearAppBadge".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigator) ClearAppBadgeFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigatorClearAppBadgeFunc(
			this.Ref(),
		),
	)
}

type LaunchConsumerFunc func(this js.Ref, params LaunchParams) js.Ref

func (fn LaunchConsumerFunc) Register() js.Func[func(params LaunchParams) js.Any] {
	return js.RegisterCallback[func(params LaunchParams) js.Any](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn LaunchConsumerFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		LaunchParams{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LaunchConsumer[T any] struct {
	Fn  func(arg T, this js.Ref, params LaunchParams) js.Ref
	Arg T
}

func (cb *LaunchConsumer[T]) Register() js.Func[func(params LaunchParams) js.Any] {
	return js.RegisterCallback[func(params LaunchParams) js.Any](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *LaunchConsumer[T]) DispatchCallback(
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

		LaunchParams{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LaunchParams struct {
	ref js.Ref
}

func (this LaunchParams) Once() LaunchParams {
	this.Ref().Once()
	return this
}

func (this LaunchParams) Ref() js.Ref {
	return this.ref
}

func (this LaunchParams) FromRef(ref js.Ref) LaunchParams {
	this.ref = ref
	return this
}

func (this LaunchParams) Free() {
	this.Ref().Free()
}

// TargetURL returns the value of property "LaunchParams.targetURL".
//
// The returned bool will be false if there is no such property.
func (this LaunchParams) TargetURL() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetLaunchParamsTargetURL(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Files returns the value of property "LaunchParams.files".
//
// The returned bool will be false if there is no such property.
func (this LaunchParams) Files() (js.FrozenArray[FileSystemHandle], bool) {
	var _ok bool
	_ret := bindings.GetLaunchParamsFiles(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[FileSystemHandle]{}.FromRef(_ret), _ok
}

type LaunchQueue struct {
	ref js.Ref
}

func (this LaunchQueue) Once() LaunchQueue {
	this.Ref().Once()
	return this
}

func (this LaunchQueue) Ref() js.Ref {
	return this.ref
}

func (this LaunchQueue) FromRef(ref js.Ref) LaunchQueue {
	this.ref = ref
	return this
}

func (this LaunchQueue) Free() {
	this.Ref().Free()
}

// SetConsumer calls the method "LaunchQueue.setConsumer".
//
// The returned bool will be false if there is no such method.
func (this LaunchQueue) SetConsumer(consumer js.Func[func(params LaunchParams) js.Any]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallLaunchQueueSetConsumer(
		this.Ref(), js.Pointer(&_ok),
		consumer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetConsumerFunc returns the method "LaunchQueue.setConsumer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this LaunchQueue) SetConsumerFunc() (fn js.Func[func(consumer js.Func[func(params LaunchParams) js.Any])]) {
	return fn.FromRef(
		bindings.LaunchQueueSetConsumerFunc(
			this.Ref(),
		),
	)
}

type FenceReportingDestination uint32

const (
	_ FenceReportingDestination = iota

	FenceReportingDestination_BUYER
	FenceReportingDestination_SELLER
	FenceReportingDestination_COMPONENT_SELLER
	FenceReportingDestination_DIRECT_SELLER
	FenceReportingDestination_SHARED_STORAGE_SELECT_URL
)

func (FenceReportingDestination) FromRef(str js.Ref) FenceReportingDestination {
	return FenceReportingDestination(bindings.ConstOfFenceReportingDestination(str))
}

func (x FenceReportingDestination) String() (string, bool) {
	switch x {
	case FenceReportingDestination_BUYER:
		return "buyer", true
	case FenceReportingDestination_SELLER:
		return "seller", true
	case FenceReportingDestination_COMPONENT_SELLER:
		return "component-seller", true
	case FenceReportingDestination_DIRECT_SELLER:
		return "direct-seller", true
	case FenceReportingDestination_SHARED_STORAGE_SELECT_URL:
		return "shared-storage-select-url", true
	default:
		return "", false
	}
}

type FenceEvent struct {
	// EventType is "FenceEvent.eventType"
	//
	// Required
	EventType js.String
	// EventData is "FenceEvent.eventData"
	//
	// Required
	EventData js.String
	// Destination is "FenceEvent.destination"
	//
	// Required
	Destination js.Array[FenceReportingDestination]
	// Once is "FenceEvent.once"
	//
	// Optional, defaults to false.
	Once bool

	FFI_USE_Once bool // for Once.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FenceEvent with all fields set.
func (p FenceEvent) FromRef(ref js.Ref) FenceEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 FenceEvent FenceEvent [// FenceEvent] [0x140007d3a40 0x140007d3ae0 0x140007d3b80 0x140007d3c20 0x140007d3cc0] 0x1400081f680 {0 0}} in the application heap.
func (p FenceEvent) New() js.Ref {
	return bindings.FenceEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FenceEvent) UpdateFrom(ref js.Ref) {
	bindings.FenceEventJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FenceEvent) Update(ref js.Ref) {
	bindings.FenceEventJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_FenceEvent_String struct {
	ref js.Ref
}

func (x OneOf_FenceEvent_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_FenceEvent_String) Free() {
	x.ref.Free()
}

func (x OneOf_FenceEvent_String) FromRef(ref js.Ref) OneOf_FenceEvent_String {
	return OneOf_FenceEvent_String{
		ref: ref,
	}
}

func (x OneOf_FenceEvent_String) FenceEvent() FenceEvent {
	var ret FenceEvent
	ret.UpdateFrom(x.ref)
	return ret
}

func (x OneOf_FenceEvent_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type ReportEventType = OneOf_FenceEvent_String

type Fence struct {
	ref js.Ref
}

func (this Fence) Once() Fence {
	this.Ref().Once()
	return this
}

func (this Fence) Ref() js.Ref {
	return this.ref
}

func (this Fence) FromRef(ref js.Ref) Fence {
	this.ref = ref
	return this
}

func (this Fence) Free() {
	this.Ref().Free()
}

// ReportEvent calls the method "Fence.reportEvent".
//
// The returned bool will be false if there is no such method.
func (this Fence) ReportEvent(event ReportEventType) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFenceReportEvent(
		this.Ref(), js.Pointer(&_ok),
		event.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReportEventFunc returns the method "Fence.reportEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Fence) ReportEventFunc() (fn js.Func[func(event ReportEventType)]) {
	return fn.FromRef(
		bindings.FenceReportEventFunc(
			this.Ref(),
		),
	)
}

// SetReportEventDataForAutomaticBeacons calls the method "Fence.setReportEventDataForAutomaticBeacons".
//
// The returned bool will be false if there is no such method.
func (this Fence) SetReportEventDataForAutomaticBeacons(event FenceEvent) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFenceSetReportEventDataForAutomaticBeacons(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&event),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetReportEventDataForAutomaticBeaconsFunc returns the method "Fence.setReportEventDataForAutomaticBeacons".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Fence) SetReportEventDataForAutomaticBeaconsFunc() (fn js.Func[func(event FenceEvent)]) {
	return fn.FromRef(
		bindings.FenceSetReportEventDataForAutomaticBeaconsFunc(
			this.Ref(),
		),
	)
}

// GetNestedConfigs calls the method "Fence.getNestedConfigs".
//
// The returned bool will be false if there is no such method.
func (this Fence) GetNestedConfigs() (js.Array[FencedFrameConfig], bool) {
	var _ok bool
	_ret := bindings.CallFenceGetNestedConfigs(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[FencedFrameConfig]{}.FromRef(_ret), _ok
}

// GetNestedConfigsFunc returns the method "Fence.getNestedConfigs".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Fence) GetNestedConfigsFunc() (fn js.Func[func() js.Array[FencedFrameConfig]]) {
	return fn.FromRef(
		bindings.FenceGetNestedConfigsFunc(
			this.Ref(),
		),
	)
}

type PortalHost struct {
	EventTarget
}

func (this PortalHost) Once() PortalHost {
	this.Ref().Once()
	return this
}

func (this PortalHost) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this PortalHost) FromRef(ref js.Ref) PortalHost {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this PortalHost) Free() {
	this.Ref().Free()
}

// PostMessage calls the method "PortalHost.postMessage".
//
// The returned bool will be false if there is no such method.
func (this PortalHost) PostMessage(message js.Any, options StructuredSerializeOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPortalHostPostMessage(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessageFunc returns the method "PortalHost.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PortalHost) PostMessageFunc() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	return fn.FromRef(
		bindings.PortalHostPostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage1 calls the method "PortalHost.postMessage".
//
// The returned bool will be false if there is no such method.
func (this PortalHost) PostMessage1(message js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPortalHostPostMessage1(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessage1Func returns the method "PortalHost.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PortalHost) PostMessage1Func() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.PortalHostPostMessage1Func(
			this.Ref(),
		),
	)
}

type OrientationLockType uint32

const (
	_ OrientationLockType = iota

	OrientationLockType_ANY
	OrientationLockType_NATURAL
	OrientationLockType_LANDSCAPE
	OrientationLockType_PORTRAIT
	OrientationLockType_PORTRAIT_PRIMARY
	OrientationLockType_PORTRAIT_SECONDARY
	OrientationLockType_LANDSCAPE_PRIMARY
	OrientationLockType_LANDSCAPE_SECONDARY
)

func (OrientationLockType) FromRef(str js.Ref) OrientationLockType {
	return OrientationLockType(bindings.ConstOfOrientationLockType(str))
}

func (x OrientationLockType) String() (string, bool) {
	switch x {
	case OrientationLockType_ANY:
		return "any", true
	case OrientationLockType_NATURAL:
		return "natural", true
	case OrientationLockType_LANDSCAPE:
		return "landscape", true
	case OrientationLockType_PORTRAIT:
		return "portrait", true
	case OrientationLockType_PORTRAIT_PRIMARY:
		return "portrait-primary", true
	case OrientationLockType_PORTRAIT_SECONDARY:
		return "portrait-secondary", true
	case OrientationLockType_LANDSCAPE_PRIMARY:
		return "landscape-primary", true
	case OrientationLockType_LANDSCAPE_SECONDARY:
		return "landscape-secondary", true
	default:
		return "", false
	}
}

type OrientationType uint32

const (
	_ OrientationType = iota

	OrientationType_PORTRAIT_PRIMARY
	OrientationType_PORTRAIT_SECONDARY
	OrientationType_LANDSCAPE_PRIMARY
	OrientationType_LANDSCAPE_SECONDARY
)

func (OrientationType) FromRef(str js.Ref) OrientationType {
	return OrientationType(bindings.ConstOfOrientationType(str))
}

func (x OrientationType) String() (string, bool) {
	switch x {
	case OrientationType_PORTRAIT_PRIMARY:
		return "portrait-primary", true
	case OrientationType_PORTRAIT_SECONDARY:
		return "portrait-secondary", true
	case OrientationType_LANDSCAPE_PRIMARY:
		return "landscape-primary", true
	case OrientationType_LANDSCAPE_SECONDARY:
		return "landscape-secondary", true
	default:
		return "", false
	}
}

type ScreenOrientation struct {
	EventTarget
}

func (this ScreenOrientation) Once() ScreenOrientation {
	this.Ref().Once()
	return this
}

func (this ScreenOrientation) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this ScreenOrientation) FromRef(ref js.Ref) ScreenOrientation {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this ScreenOrientation) Free() {
	this.Ref().Free()
}

// Type returns the value of property "ScreenOrientation.type".
//
// The returned bool will be false if there is no such property.
func (this ScreenOrientation) Type() (OrientationType, bool) {
	var _ok bool
	_ret := bindings.GetScreenOrientationType(
		this.Ref(), js.Pointer(&_ok),
	)
	return OrientationType(_ret), _ok
}

// Angle returns the value of property "ScreenOrientation.angle".
//
// The returned bool will be false if there is no such property.
func (this ScreenOrientation) Angle() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetScreenOrientationAngle(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Lock calls the method "ScreenOrientation.lock".
//
// The returned bool will be false if there is no such method.
func (this ScreenOrientation) Lock(orientation OrientationLockType) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallScreenOrientationLock(
		this.Ref(), js.Pointer(&_ok),
		uint32(orientation),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// LockFunc returns the method "ScreenOrientation.lock".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ScreenOrientation) LockFunc() (fn js.Func[func(orientation OrientationLockType) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ScreenOrientationLockFunc(
			this.Ref(),
		),
	)
}

// Unlock calls the method "ScreenOrientation.unlock".
//
// The returned bool will be false if there is no such method.
func (this ScreenOrientation) Unlock() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallScreenOrientationUnlock(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UnlockFunc returns the method "ScreenOrientation.unlock".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ScreenOrientation) UnlockFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ScreenOrientationUnlockFunc(
			this.Ref(),
		),
	)
}

type Screen struct {
	ref js.Ref
}

func (this Screen) Once() Screen {
	this.Ref().Once()
	return this
}

func (this Screen) Ref() js.Ref {
	return this.ref
}

func (this Screen) FromRef(ref js.Ref) Screen {
	this.ref = ref
	return this
}

func (this Screen) Free() {
	this.Ref().Free()
}

// AvailWidth returns the value of property "Screen.availWidth".
//
// The returned bool will be false if there is no such property.
func (this Screen) AvailWidth() (int32, bool) {
	var _ok bool
	_ret := bindings.GetScreenAvailWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// AvailHeight returns the value of property "Screen.availHeight".
//
// The returned bool will be false if there is no such property.
func (this Screen) AvailHeight() (int32, bool) {
	var _ok bool
	_ret := bindings.GetScreenAvailHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Width returns the value of property "Screen.width".
//
// The returned bool will be false if there is no such property.
func (this Screen) Width() (int32, bool) {
	var _ok bool
	_ret := bindings.GetScreenWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Height returns the value of property "Screen.height".
//
// The returned bool will be false if there is no such property.
func (this Screen) Height() (int32, bool) {
	var _ok bool
	_ret := bindings.GetScreenHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// ColorDepth returns the value of property "Screen.colorDepth".
//
// The returned bool will be false if there is no such property.
func (this Screen) ColorDepth() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetScreenColorDepth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// PixelDepth returns the value of property "Screen.pixelDepth".
//
// The returned bool will be false if there is no such property.
func (this Screen) PixelDepth() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetScreenPixelDepth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// IsExtended returns the value of property "Screen.isExtended".
//
// The returned bool will be false if there is no such property.
func (this Screen) IsExtended() (bool, bool) {
	var _ok bool
	_ret := bindings.GetScreenIsExtended(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Orientation returns the value of property "Screen.orientation".
//
// The returned bool will be false if there is no such property.
func (this Screen) Orientation() (ScreenOrientation, bool) {
	var _ok bool
	_ret := bindings.GetScreenOrientation(
		this.Ref(), js.Pointer(&_ok),
	)
	return ScreenOrientation{}.FromRef(_ret), _ok
}

type VisualViewport struct {
	EventTarget
}

func (this VisualViewport) Once() VisualViewport {
	this.Ref().Once()
	return this
}

func (this VisualViewport) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this VisualViewport) FromRef(ref js.Ref) VisualViewport {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this VisualViewport) Free() {
	this.Ref().Free()
}

// OffsetLeft returns the value of property "VisualViewport.offsetLeft".
//
// The returned bool will be false if there is no such property.
func (this VisualViewport) OffsetLeft() (float64, bool) {
	var _ok bool
	_ret := bindings.GetVisualViewportOffsetLeft(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// OffsetTop returns the value of property "VisualViewport.offsetTop".
//
// The returned bool will be false if there is no such property.
func (this VisualViewport) OffsetTop() (float64, bool) {
	var _ok bool
	_ret := bindings.GetVisualViewportOffsetTop(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// PageLeft returns the value of property "VisualViewport.pageLeft".
//
// The returned bool will be false if there is no such property.
func (this VisualViewport) PageLeft() (float64, bool) {
	var _ok bool
	_ret := bindings.GetVisualViewportPageLeft(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// PageTop returns the value of property "VisualViewport.pageTop".
//
// The returned bool will be false if there is no such property.
func (this VisualViewport) PageTop() (float64, bool) {
	var _ok bool
	_ret := bindings.GetVisualViewportPageTop(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Width returns the value of property "VisualViewport.width".
//
// The returned bool will be false if there is no such property.
func (this VisualViewport) Width() (float64, bool) {
	var _ok bool
	_ret := bindings.GetVisualViewportWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Height returns the value of property "VisualViewport.height".
//
// The returned bool will be false if there is no such property.
func (this VisualViewport) Height() (float64, bool) {
	var _ok bool
	_ret := bindings.GetVisualViewportHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Scale returns the value of property "VisualViewport.scale".
//
// The returned bool will be false if there is no such property.
func (this VisualViewport) Scale() (float64, bool) {
	var _ok bool
	_ret := bindings.GetVisualViewportScale(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

type SharedStorageRunOperationMethodOptions struct {
	// Data is "SharedStorageRunOperationMethodOptions.data"
	//
	// Optional
	Data js.Object
	// ResolveToConfig is "SharedStorageRunOperationMethodOptions.resolveToConfig"
	//
	// Optional, defaults to false.
	ResolveToConfig bool
	// KeepAlive is "SharedStorageRunOperationMethodOptions.keepAlive"
	//
	// Optional, defaults to false.
	KeepAlive bool

	FFI_USE_ResolveToConfig bool // for ResolveToConfig.
	FFI_USE_KeepAlive       bool // for KeepAlive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SharedStorageRunOperationMethodOptions with all fields set.
func (p SharedStorageRunOperationMethodOptions) FromRef(ref js.Ref) SharedStorageRunOperationMethodOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 SharedStorageRunOperationMethodOptions SharedStorageRunOperationMethodOptions [// SharedStorageRunOperationMethodOptions] [0x140007d3f40 0x140007da000 0x140007da140 0x140007da0a0 0x140007da1e0] 0x1400081f9e0 {0 0}} in the application heap.
func (p SharedStorageRunOperationMethodOptions) New() js.Ref {
	return bindings.SharedStorageRunOperationMethodOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SharedStorageRunOperationMethodOptions) UpdateFrom(ref js.Ref) {
	bindings.SharedStorageRunOperationMethodOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SharedStorageRunOperationMethodOptions) Update(ref js.Ref) {
	bindings.SharedStorageRunOperationMethodOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type SharedStorageResponse = OneOf_String_FencedFrameConfig

type SharedStorageUrlWithMetadata struct {
	// Url is "SharedStorageUrlWithMetadata.url"
	//
	// Required
	Url js.String
	// ReportingMetadata is "SharedStorageUrlWithMetadata.reportingMetadata"
	//
	// Optional
	ReportingMetadata js.Object

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SharedStorageUrlWithMetadata with all fields set.
func (p SharedStorageUrlWithMetadata) FromRef(ref js.Ref) SharedStorageUrlWithMetadata {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 SharedStorageUrlWithMetadata SharedStorageUrlWithMetadata [// SharedStorageUrlWithMetadata] [0x140007da280 0x140007da320] 0x1400081fa58 {0 0}} in the application heap.
func (p SharedStorageUrlWithMetadata) New() js.Ref {
	return bindings.SharedStorageUrlWithMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SharedStorageUrlWithMetadata) UpdateFrom(ref js.Ref) {
	bindings.SharedStorageUrlWithMetadataJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SharedStorageUrlWithMetadata) Update(ref js.Ref) {
	bindings.SharedStorageUrlWithMetadataJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

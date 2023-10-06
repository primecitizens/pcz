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

// HasDispatch returns true if the method "MLCommandEncoder.dispatch" exists.
func (this MLCommandEncoder) HasDispatch() bool {
	return js.True == bindings.HasMLCommandEncoderDispatch(
		this.Ref(),
	)
}

// DispatchFunc returns the method "MLCommandEncoder.dispatch".
func (this MLCommandEncoder) DispatchFunc() (fn js.Func[func(graph MLGraph, inputs MLNamedGPUResources, outputs MLNamedGPUResources)]) {
	return fn.FromRef(
		bindings.MLCommandEncoderDispatchFunc(
			this.Ref(),
		),
	)
}

// Dispatch calls the method "MLCommandEncoder.dispatch".
func (this MLCommandEncoder) Dispatch(graph MLGraph, inputs MLNamedGPUResources, outputs MLNamedGPUResources) (ret js.Void) {
	bindings.CallMLCommandEncoderDispatch(
		this.Ref(), js.Pointer(&ret),
		graph.Ref(),
		inputs.Ref(),
		outputs.Ref(),
	)

	return
}

// TryDispatch calls the method "MLCommandEncoder.dispatch"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MLCommandEncoder) TryDispatch(graph MLGraph, inputs MLNamedGPUResources, outputs MLNamedGPUResources) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCommandEncoderDispatch(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		graph.Ref(),
		inputs.Ref(),
		outputs.Ref(),
	)

	return
}

// HasFinish returns true if the method "MLCommandEncoder.finish" exists.
func (this MLCommandEncoder) HasFinish() bool {
	return js.True == bindings.HasMLCommandEncoderFinish(
		this.Ref(),
	)
}

// FinishFunc returns the method "MLCommandEncoder.finish".
func (this MLCommandEncoder) FinishFunc() (fn js.Func[func(descriptor GPUCommandBufferDescriptor) GPUCommandBuffer]) {
	return fn.FromRef(
		bindings.MLCommandEncoderFinishFunc(
			this.Ref(),
		),
	)
}

// Finish calls the method "MLCommandEncoder.finish".
func (this MLCommandEncoder) Finish(descriptor GPUCommandBufferDescriptor) (ret GPUCommandBuffer) {
	bindings.CallMLCommandEncoderFinish(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryFinish calls the method "MLCommandEncoder.finish"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MLCommandEncoder) TryFinish(descriptor GPUCommandBufferDescriptor) (ret GPUCommandBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCommandEncoderFinish(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFinish1 returns true if the method "MLCommandEncoder.finish" exists.
func (this MLCommandEncoder) HasFinish1() bool {
	return js.True == bindings.HasMLCommandEncoderFinish1(
		this.Ref(),
	)
}

// Finish1Func returns the method "MLCommandEncoder.finish".
func (this MLCommandEncoder) Finish1Func() (fn js.Func[func() GPUCommandBuffer]) {
	return fn.FromRef(
		bindings.MLCommandEncoderFinish1Func(
			this.Ref(),
		),
	)
}

// Finish1 calls the method "MLCommandEncoder.finish".
func (this MLCommandEncoder) Finish1() (ret GPUCommandBuffer) {
	bindings.CallMLCommandEncoderFinish1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFinish1 calls the method "MLCommandEncoder.finish"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MLCommandEncoder) TryFinish1() (ret GPUCommandBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCommandEncoderFinish1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasInitializeGraph returns true if the method "MLCommandEncoder.initializeGraph" exists.
func (this MLCommandEncoder) HasInitializeGraph() bool {
	return js.True == bindings.HasMLCommandEncoderInitializeGraph(
		this.Ref(),
	)
}

// InitializeGraphFunc returns the method "MLCommandEncoder.initializeGraph".
func (this MLCommandEncoder) InitializeGraphFunc() (fn js.Func[func(graph MLGraph)]) {
	return fn.FromRef(
		bindings.MLCommandEncoderInitializeGraphFunc(
			this.Ref(),
		),
	)
}

// InitializeGraph calls the method "MLCommandEncoder.initializeGraph".
func (this MLCommandEncoder) InitializeGraph(graph MLGraph) (ret js.Void) {
	bindings.CallMLCommandEncoderInitializeGraph(
		this.Ref(), js.Pointer(&ret),
		graph.Ref(),
	)

	return
}

// TryInitializeGraph calls the method "MLCommandEncoder.initializeGraph"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MLCommandEncoder) TryInitializeGraph(graph MLGraph) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCommandEncoderInitializeGraph(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		graph.Ref(),
	)

	return
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

// HasCompute returns true if the method "MLContext.compute" exists.
func (this MLContext) HasCompute() bool {
	return js.True == bindings.HasMLContextCompute(
		this.Ref(),
	)
}

// ComputeFunc returns the method "MLContext.compute".
func (this MLContext) ComputeFunc() (fn js.Func[func(graph MLGraph, inputs MLNamedArrayBufferViews, outputs MLNamedArrayBufferViews) js.Promise[MLComputeResult]]) {
	return fn.FromRef(
		bindings.MLContextComputeFunc(
			this.Ref(),
		),
	)
}

// Compute calls the method "MLContext.compute".
func (this MLContext) Compute(graph MLGraph, inputs MLNamedArrayBufferViews, outputs MLNamedArrayBufferViews) (ret js.Promise[MLComputeResult]) {
	bindings.CallMLContextCompute(
		this.Ref(), js.Pointer(&ret),
		graph.Ref(),
		inputs.Ref(),
		outputs.Ref(),
	)

	return
}

// TryCompute calls the method "MLContext.compute"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MLContext) TryCompute(graph MLGraph, inputs MLNamedArrayBufferViews, outputs MLNamedArrayBufferViews) (ret js.Promise[MLComputeResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLContextCompute(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		graph.Ref(),
		inputs.Ref(),
		outputs.Ref(),
	)

	return
}

// HasComputeSync returns true if the method "MLContext.computeSync" exists.
func (this MLContext) HasComputeSync() bool {
	return js.True == bindings.HasMLContextComputeSync(
		this.Ref(),
	)
}

// ComputeSyncFunc returns the method "MLContext.computeSync".
func (this MLContext) ComputeSyncFunc() (fn js.Func[func(graph MLGraph, inputs MLNamedArrayBufferViews, outputs MLNamedArrayBufferViews)]) {
	return fn.FromRef(
		bindings.MLContextComputeSyncFunc(
			this.Ref(),
		),
	)
}

// ComputeSync calls the method "MLContext.computeSync".
func (this MLContext) ComputeSync(graph MLGraph, inputs MLNamedArrayBufferViews, outputs MLNamedArrayBufferViews) (ret js.Void) {
	bindings.CallMLContextComputeSync(
		this.Ref(), js.Pointer(&ret),
		graph.Ref(),
		inputs.Ref(),
		outputs.Ref(),
	)

	return
}

// TryComputeSync calls the method "MLContext.computeSync"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MLContext) TryComputeSync(graph MLGraph, inputs MLNamedArrayBufferViews, outputs MLNamedArrayBufferViews) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLContextComputeSync(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		graph.Ref(),
		inputs.Ref(),
		outputs.Ref(),
	)

	return
}

// HasCreateCommandEncoder returns true if the method "MLContext.createCommandEncoder" exists.
func (this MLContext) HasCreateCommandEncoder() bool {
	return js.True == bindings.HasMLContextCreateCommandEncoder(
		this.Ref(),
	)
}

// CreateCommandEncoderFunc returns the method "MLContext.createCommandEncoder".
func (this MLContext) CreateCommandEncoderFunc() (fn js.Func[func() MLCommandEncoder]) {
	return fn.FromRef(
		bindings.MLContextCreateCommandEncoderFunc(
			this.Ref(),
		),
	)
}

// CreateCommandEncoder calls the method "MLContext.createCommandEncoder".
func (this MLContext) CreateCommandEncoder() (ret MLCommandEncoder) {
	bindings.CallMLContextCreateCommandEncoder(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateCommandEncoder calls the method "MLContext.createCommandEncoder"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MLContext) TryCreateCommandEncoder() (ret MLCommandEncoder, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLContextCreateCommandEncoder(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// New creates a new MLContextOptions in the application heap.
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

// HasCreateContext returns true if the method "ML.createContext" exists.
func (this ML) HasCreateContext() bool {
	return js.True == bindings.HasMLCreateContext(
		this.Ref(),
	)
}

// CreateContextFunc returns the method "ML.createContext".
func (this ML) CreateContextFunc() (fn js.Func[func(options MLContextOptions) js.Promise[MLContext]]) {
	return fn.FromRef(
		bindings.MLCreateContextFunc(
			this.Ref(),
		),
	)
}

// CreateContext calls the method "ML.createContext".
func (this ML) CreateContext(options MLContextOptions) (ret js.Promise[MLContext]) {
	bindings.CallMLCreateContext(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryCreateContext calls the method "ML.createContext"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this ML) TryCreateContext(options MLContextOptions) (ret js.Promise[MLContext], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCreateContext(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasCreateContext1 returns true if the method "ML.createContext" exists.
func (this ML) HasCreateContext1() bool {
	return js.True == bindings.HasMLCreateContext1(
		this.Ref(),
	)
}

// CreateContext1Func returns the method "ML.createContext".
func (this ML) CreateContext1Func() (fn js.Func[func() js.Promise[MLContext]]) {
	return fn.FromRef(
		bindings.MLCreateContext1Func(
			this.Ref(),
		),
	)
}

// CreateContext1 calls the method "ML.createContext".
func (this ML) CreateContext1() (ret js.Promise[MLContext]) {
	bindings.CallMLCreateContext1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateContext1 calls the method "ML.createContext"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this ML) TryCreateContext1() (ret js.Promise[MLContext], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCreateContext1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateContext2 returns true if the method "ML.createContext" exists.
func (this ML) HasCreateContext2() bool {
	return js.True == bindings.HasMLCreateContext2(
		this.Ref(),
	)
}

// CreateContext2Func returns the method "ML.createContext".
func (this ML) CreateContext2Func() (fn js.Func[func(gpuDevice GPUDevice) js.Promise[MLContext]]) {
	return fn.FromRef(
		bindings.MLCreateContext2Func(
			this.Ref(),
		),
	)
}

// CreateContext2 calls the method "ML.createContext".
func (this ML) CreateContext2(gpuDevice GPUDevice) (ret js.Promise[MLContext]) {
	bindings.CallMLCreateContext2(
		this.Ref(), js.Pointer(&ret),
		gpuDevice.Ref(),
	)

	return
}

// TryCreateContext2 calls the method "ML.createContext"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this ML) TryCreateContext2(gpuDevice GPUDevice) (ret js.Promise[MLContext], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCreateContext2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		gpuDevice.Ref(),
	)

	return
}

// HasCreateContextSync returns true if the method "ML.createContextSync" exists.
func (this ML) HasCreateContextSync() bool {
	return js.True == bindings.HasMLCreateContextSync(
		this.Ref(),
	)
}

// CreateContextSyncFunc returns the method "ML.createContextSync".
func (this ML) CreateContextSyncFunc() (fn js.Func[func(options MLContextOptions) MLContext]) {
	return fn.FromRef(
		bindings.MLCreateContextSyncFunc(
			this.Ref(),
		),
	)
}

// CreateContextSync calls the method "ML.createContextSync".
func (this ML) CreateContextSync(options MLContextOptions) (ret MLContext) {
	bindings.CallMLCreateContextSync(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryCreateContextSync calls the method "ML.createContextSync"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this ML) TryCreateContextSync(options MLContextOptions) (ret MLContext, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCreateContextSync(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasCreateContextSync1 returns true if the method "ML.createContextSync" exists.
func (this ML) HasCreateContextSync1() bool {
	return js.True == bindings.HasMLCreateContextSync1(
		this.Ref(),
	)
}

// CreateContextSync1Func returns the method "ML.createContextSync".
func (this ML) CreateContextSync1Func() (fn js.Func[func() MLContext]) {
	return fn.FromRef(
		bindings.MLCreateContextSync1Func(
			this.Ref(),
		),
	)
}

// CreateContextSync1 calls the method "ML.createContextSync".
func (this ML) CreateContextSync1() (ret MLContext) {
	bindings.CallMLCreateContextSync1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateContextSync1 calls the method "ML.createContextSync"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this ML) TryCreateContextSync1() (ret MLContext, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCreateContextSync1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateContextSync2 returns true if the method "ML.createContextSync" exists.
func (this ML) HasCreateContextSync2() bool {
	return js.True == bindings.HasMLCreateContextSync2(
		this.Ref(),
	)
}

// CreateContextSync2Func returns the method "ML.createContextSync".
func (this ML) CreateContextSync2Func() (fn js.Func[func(gpuDevice GPUDevice) MLContext]) {
	return fn.FromRef(
		bindings.MLCreateContextSync2Func(
			this.Ref(),
		),
	)
}

// CreateContextSync2 calls the method "ML.createContextSync".
func (this ML) CreateContextSync2(gpuDevice GPUDevice) (ret MLContext) {
	bindings.CallMLCreateContextSync2(
		this.Ref(), js.Pointer(&ret),
		gpuDevice.Ref(),
	)

	return
}

// TryCreateContextSync2 calls the method "ML.createContextSync"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this ML) TryCreateContextSync2(gpuDevice GPUDevice) (ret MLContext, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCreateContextSync2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		gpuDevice.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this NetworkInformation) Type() (ret ConnectionType, ok bool) {
	ok = js.True == bindings.GetNetworkInformationType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EffectiveType returns the value of property "NetworkInformation.effectiveType".
//
// It returns ok=false if there is no such property.
func (this NetworkInformation) EffectiveType() (ret EffectiveConnectionType, ok bool) {
	ok = js.True == bindings.GetNetworkInformationEffectiveType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DownlinkMax returns the value of property "NetworkInformation.downlinkMax".
//
// It returns ok=false if there is no such property.
func (this NetworkInformation) DownlinkMax() (ret Megabit, ok bool) {
	ok = js.True == bindings.GetNetworkInformationDownlinkMax(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Downlink returns the value of property "NetworkInformation.downlink".
//
// It returns ok=false if there is no such property.
func (this NetworkInformation) Downlink() (ret Megabit, ok bool) {
	ok = js.True == bindings.GetNetworkInformationDownlink(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Rtt returns the value of property "NetworkInformation.rtt".
//
// It returns ok=false if there is no such property.
func (this NetworkInformation) Rtt() (ret Millisecond, ok bool) {
	ok = js.True == bindings.GetNetworkInformationRtt(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SaveData returns the value of property "NetworkInformation.saveData".
//
// It returns ok=false if there is no such property.
func (this NetworkInformation) SaveData() (ret bool, ok bool) {
	ok = js.True == bindings.GetNetworkInformationSaveData(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// New creates a new GPUQueueDescriptor in the application heap.
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
	//
	// NOTE: DefaultQueue.FFI_USE MUST be set to true to get DefaultQueue used.
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

// New creates a new GPUDeviceDescriptor in the application heap.
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
// It returns ok=false if there is no such property.
func (this GPUAdapterInfo) Vendor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUAdapterInfoVendor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Architecture returns the value of property "GPUAdapterInfo.architecture".
//
// It returns ok=false if there is no such property.
func (this GPUAdapterInfo) Architecture() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUAdapterInfoArchitecture(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Device returns the value of property "GPUAdapterInfo.device".
//
// It returns ok=false if there is no such property.
func (this GPUAdapterInfo) Device() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUAdapterInfoDevice(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Description returns the value of property "GPUAdapterInfo.description".
//
// It returns ok=false if there is no such property.
func (this GPUAdapterInfo) Description() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUAdapterInfoDescription(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this GPUAdapter) Features() (ret GPUSupportedFeatures, ok bool) {
	ok = js.True == bindings.GetGPUAdapterFeatures(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Limits returns the value of property "GPUAdapter.limits".
//
// It returns ok=false if there is no such property.
func (this GPUAdapter) Limits() (ret GPUSupportedLimits, ok bool) {
	ok = js.True == bindings.GetGPUAdapterLimits(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsFallbackAdapter returns the value of property "GPUAdapter.isFallbackAdapter".
//
// It returns ok=false if there is no such property.
func (this GPUAdapter) IsFallbackAdapter() (ret bool, ok bool) {
	ok = js.True == bindings.GetGPUAdapterIsFallbackAdapter(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasRequestDevice returns true if the method "GPUAdapter.requestDevice" exists.
func (this GPUAdapter) HasRequestDevice() bool {
	return js.True == bindings.HasGPUAdapterRequestDevice(
		this.Ref(),
	)
}

// RequestDeviceFunc returns the method "GPUAdapter.requestDevice".
func (this GPUAdapter) RequestDeviceFunc() (fn js.Func[func(descriptor GPUDeviceDescriptor) js.Promise[GPUDevice]]) {
	return fn.FromRef(
		bindings.GPUAdapterRequestDeviceFunc(
			this.Ref(),
		),
	)
}

// RequestDevice calls the method "GPUAdapter.requestDevice".
func (this GPUAdapter) RequestDevice(descriptor GPUDeviceDescriptor) (ret js.Promise[GPUDevice]) {
	bindings.CallGPUAdapterRequestDevice(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryRequestDevice calls the method "GPUAdapter.requestDevice"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this GPUAdapter) TryRequestDevice(descriptor GPUDeviceDescriptor) (ret js.Promise[GPUDevice], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUAdapterRequestDevice(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasRequestDevice1 returns true if the method "GPUAdapter.requestDevice" exists.
func (this GPUAdapter) HasRequestDevice1() bool {
	return js.True == bindings.HasGPUAdapterRequestDevice1(
		this.Ref(),
	)
}

// RequestDevice1Func returns the method "GPUAdapter.requestDevice".
func (this GPUAdapter) RequestDevice1Func() (fn js.Func[func() js.Promise[GPUDevice]]) {
	return fn.FromRef(
		bindings.GPUAdapterRequestDevice1Func(
			this.Ref(),
		),
	)
}

// RequestDevice1 calls the method "GPUAdapter.requestDevice".
func (this GPUAdapter) RequestDevice1() (ret js.Promise[GPUDevice]) {
	bindings.CallGPUAdapterRequestDevice1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestDevice1 calls the method "GPUAdapter.requestDevice"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this GPUAdapter) TryRequestDevice1() (ret js.Promise[GPUDevice], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUAdapterRequestDevice1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRequestAdapterInfo returns true if the method "GPUAdapter.requestAdapterInfo" exists.
func (this GPUAdapter) HasRequestAdapterInfo() bool {
	return js.True == bindings.HasGPUAdapterRequestAdapterInfo(
		this.Ref(),
	)
}

// RequestAdapterInfoFunc returns the method "GPUAdapter.requestAdapterInfo".
func (this GPUAdapter) RequestAdapterInfoFunc() (fn js.Func[func(unmaskHints js.Array[js.String]) js.Promise[GPUAdapterInfo]]) {
	return fn.FromRef(
		bindings.GPUAdapterRequestAdapterInfoFunc(
			this.Ref(),
		),
	)
}

// RequestAdapterInfo calls the method "GPUAdapter.requestAdapterInfo".
func (this GPUAdapter) RequestAdapterInfo(unmaskHints js.Array[js.String]) (ret js.Promise[GPUAdapterInfo]) {
	bindings.CallGPUAdapterRequestAdapterInfo(
		this.Ref(), js.Pointer(&ret),
		unmaskHints.Ref(),
	)

	return
}

// TryRequestAdapterInfo calls the method "GPUAdapter.requestAdapterInfo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this GPUAdapter) TryRequestAdapterInfo(unmaskHints js.Array[js.String]) (ret js.Promise[GPUAdapterInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUAdapterRequestAdapterInfo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		unmaskHints.Ref(),
	)

	return
}

// HasRequestAdapterInfo1 returns true if the method "GPUAdapter.requestAdapterInfo" exists.
func (this GPUAdapter) HasRequestAdapterInfo1() bool {
	return js.True == bindings.HasGPUAdapterRequestAdapterInfo1(
		this.Ref(),
	)
}

// RequestAdapterInfo1Func returns the method "GPUAdapter.requestAdapterInfo".
func (this GPUAdapter) RequestAdapterInfo1Func() (fn js.Func[func() js.Promise[GPUAdapterInfo]]) {
	return fn.FromRef(
		bindings.GPUAdapterRequestAdapterInfo1Func(
			this.Ref(),
		),
	)
}

// RequestAdapterInfo1 calls the method "GPUAdapter.requestAdapterInfo".
func (this GPUAdapter) RequestAdapterInfo1() (ret js.Promise[GPUAdapterInfo]) {
	bindings.CallGPUAdapterRequestAdapterInfo1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestAdapterInfo1 calls the method "GPUAdapter.requestAdapterInfo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this GPUAdapter) TryRequestAdapterInfo1() (ret js.Promise[GPUAdapterInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUAdapterRequestAdapterInfo1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
	//
	// NOTE: FFI_USE_ForceFallbackAdapter MUST be set to true to make this field effective.
	ForceFallbackAdapter bool

	FFI_USE_ForceFallbackAdapter bool // for ForceFallbackAdapter.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPURequestAdapterOptions with all fields set.
func (p GPURequestAdapterOptions) FromRef(ref js.Ref) GPURequestAdapterOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPURequestAdapterOptions in the application heap.
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
// It returns ok=false if there is no such property.
func (this GPU) WgslLanguageFeatures() (ret WGSLLanguageFeatures, ok bool) {
	ok = js.True == bindings.GetGPUWgslLanguageFeatures(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasRequestAdapter returns true if the method "GPU.requestAdapter" exists.
func (this GPU) HasRequestAdapter() bool {
	return js.True == bindings.HasGPURequestAdapter(
		this.Ref(),
	)
}

// RequestAdapterFunc returns the method "GPU.requestAdapter".
func (this GPU) RequestAdapterFunc() (fn js.Func[func(options GPURequestAdapterOptions) js.Promise[GPUAdapter]]) {
	return fn.FromRef(
		bindings.GPURequestAdapterFunc(
			this.Ref(),
		),
	)
}

// RequestAdapter calls the method "GPU.requestAdapter".
func (this GPU) RequestAdapter(options GPURequestAdapterOptions) (ret js.Promise[GPUAdapter]) {
	bindings.CallGPURequestAdapter(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRequestAdapter calls the method "GPU.requestAdapter"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this GPU) TryRequestAdapter(options GPURequestAdapterOptions) (ret js.Promise[GPUAdapter], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURequestAdapter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasRequestAdapter1 returns true if the method "GPU.requestAdapter" exists.
func (this GPU) HasRequestAdapter1() bool {
	return js.True == bindings.HasGPURequestAdapter1(
		this.Ref(),
	)
}

// RequestAdapter1Func returns the method "GPU.requestAdapter".
func (this GPU) RequestAdapter1Func() (fn js.Func[func() js.Promise[GPUAdapter]]) {
	return fn.FromRef(
		bindings.GPURequestAdapter1Func(
			this.Ref(),
		),
	)
}

// RequestAdapter1 calls the method "GPU.requestAdapter".
func (this GPU) RequestAdapter1() (ret js.Promise[GPUAdapter]) {
	bindings.CallGPURequestAdapter1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestAdapter1 calls the method "GPU.requestAdapter"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this GPU) TryRequestAdapter1() (ret js.Promise[GPUAdapter], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURequestAdapter1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetPreferredCanvasFormat returns true if the method "GPU.getPreferredCanvasFormat" exists.
func (this GPU) HasGetPreferredCanvasFormat() bool {
	return js.True == bindings.HasGPUGetPreferredCanvasFormat(
		this.Ref(),
	)
}

// GetPreferredCanvasFormatFunc returns the method "GPU.getPreferredCanvasFormat".
func (this GPU) GetPreferredCanvasFormatFunc() (fn js.Func[func() GPUTextureFormat]) {
	return fn.FromRef(
		bindings.GPUGetPreferredCanvasFormatFunc(
			this.Ref(),
		),
	)
}

// GetPreferredCanvasFormat calls the method "GPU.getPreferredCanvasFormat".
func (this GPU) GetPreferredCanvasFormat() (ret GPUTextureFormat) {
	bindings.CallGPUGetPreferredCanvasFormat(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetPreferredCanvasFormat calls the method "GPU.getPreferredCanvasFormat"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this GPU) TryGetPreferredCanvasFormat() (ret GPUTextureFormat, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUGetPreferredCanvasFormat(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this Navigator) Hid() (ret HID, ok bool) {
	ok = js.True == bindings.GetNavigatorHid(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// WindowControlsOverlay returns the value of property "Navigator.windowControlsOverlay".
//
// It returns ok=false if there is no such property.
func (this Navigator) WindowControlsOverlay() (ret WindowControlsOverlay, ok bool) {
	ok = js.True == bindings.GetNavigatorWindowControlsOverlay(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Credentials returns the value of property "Navigator.credentials".
//
// It returns ok=false if there is no such property.
func (this Navigator) Credentials() (ret CredentialsContainer, ok bool) {
	ok = js.True == bindings.GetNavigatorCredentials(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Clipboard returns the value of property "Navigator.clipboard".
//
// It returns ok=false if there is no such property.
func (this Navigator) Clipboard() (ret Clipboard, ok bool) {
	ok = js.True == bindings.GetNavigatorClipboard(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Geolocation returns the value of property "Navigator.geolocation".
//
// It returns ok=false if there is no such property.
func (this Navigator) Geolocation() (ret Geolocation, ok bool) {
	ok = js.True == bindings.GetNavigatorGeolocation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Usb returns the value of property "Navigator.usb".
//
// It returns ok=false if there is no such property.
func (this Navigator) Usb() (ret USB, ok bool) {
	ok = js.True == bindings.GetNavigatorUsb(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EpubReadingSystem returns the value of property "Navigator.epubReadingSystem".
//
// It returns ok=false if there is no such property.
func (this Navigator) EpubReadingSystem() (ret EpubReadingSystem, ok bool) {
	ok = js.True == bindings.GetNavigatorEpubReadingSystem(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Xr returns the value of property "Navigator.xr".
//
// It returns ok=false if there is no such property.
func (this Navigator) Xr() (ret XRSystem, ok bool) {
	ok = js.True == bindings.GetNavigatorXr(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ServiceWorker returns the value of property "Navigator.serviceWorker".
//
// It returns ok=false if there is no such property.
func (this Navigator) ServiceWorker() (ret ServiceWorkerContainer, ok bool) {
	ok = js.True == bindings.GetNavigatorServiceWorker(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MediaDevices returns the value of property "Navigator.mediaDevices".
//
// It returns ok=false if there is no such property.
func (this Navigator) MediaDevices() (ret MediaDevices, ok bool) {
	ok = js.True == bindings.GetNavigatorMediaDevices(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Bluetooth returns the value of property "Navigator.bluetooth".
//
// It returns ok=false if there is no such property.
func (this Navigator) Bluetooth() (ret Bluetooth, ok bool) {
	ok = js.True == bindings.GetNavigatorBluetooth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Serial returns the value of property "Navigator.serial".
//
// It returns ok=false if there is no such property.
func (this Navigator) Serial() (ret Serial, ok bool) {
	ok = js.True == bindings.GetNavigatorSerial(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MediaCapabilities returns the value of property "Navigator.mediaCapabilities".
//
// It returns ok=false if there is no such property.
func (this Navigator) MediaCapabilities() (ret MediaCapabilities, ok bool) {
	ok = js.True == bindings.GetNavigatorMediaCapabilities(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UserActivation returns the value of property "Navigator.userActivation".
//
// It returns ok=false if there is no such property.
func (this Navigator) UserActivation() (ret UserActivation, ok bool) {
	ok = js.True == bindings.GetNavigatorUserActivation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Permissions returns the value of property "Navigator.permissions".
//
// It returns ok=false if there is no such property.
func (this Navigator) Permissions() (ret Permissions, ok bool) {
	ok = js.True == bindings.GetNavigatorPermissions(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Contacts returns the value of property "Navigator.contacts".
//
// It returns ok=false if there is no such property.
func (this Navigator) Contacts() (ret ContactsManager, ok bool) {
	ok = js.True == bindings.GetNavigatorContacts(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Keyboard returns the value of property "Navigator.keyboard".
//
// It returns ok=false if there is no such property.
func (this Navigator) Keyboard() (ret Keyboard, ok bool) {
	ok = js.True == bindings.GetNavigatorKeyboard(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MediaSession returns the value of property "Navigator.mediaSession".
//
// It returns ok=false if there is no such property.
func (this Navigator) MediaSession() (ret MediaSession, ok bool) {
	ok = js.True == bindings.GetNavigatorMediaSession(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DevicePosture returns the value of property "Navigator.devicePosture".
//
// It returns ok=false if there is no such property.
func (this Navigator) DevicePosture() (ret DevicePosture, ok bool) {
	ok = js.True == bindings.GetNavigatorDevicePosture(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxTouchPoints returns the value of property "Navigator.maxTouchPoints".
//
// It returns ok=false if there is no such property.
func (this Navigator) MaxTouchPoints() (ret int32, ok bool) {
	ok = js.True == bindings.GetNavigatorMaxTouchPoints(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Scheduling returns the value of property "Navigator.scheduling".
//
// It returns ok=false if there is no such property.
func (this Navigator) Scheduling() (ret Scheduling, ok bool) {
	ok = js.True == bindings.GetNavigatorScheduling(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// WakeLock returns the value of property "Navigator.wakeLock".
//
// It returns ok=false if there is no such property.
func (this Navigator) WakeLock() (ret WakeLock, ok bool) {
	ok = js.True == bindings.GetNavigatorWakeLock(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Ink returns the value of property "Navigator.ink".
//
// It returns ok=false if there is no such property.
func (this Navigator) Ink() (ret Ink, ok bool) {
	ok = js.True == bindings.GetNavigatorInk(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Presentation returns the value of property "Navigator.presentation".
//
// It returns ok=false if there is no such property.
func (this Navigator) Presentation() (ret Presentation, ok bool) {
	ok = js.True == bindings.GetNavigatorPresentation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// VirtualKeyboard returns the value of property "Navigator.virtualKeyboard".
//
// It returns ok=false if there is no such property.
func (this Navigator) VirtualKeyboard() (ret VirtualKeyboard, ok bool) {
	ok = js.True == bindings.GetNavigatorVirtualKeyboard(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UserAgentData returns the value of property "Navigator.userAgentData".
//
// It returns ok=false if there is no such property.
func (this Navigator) UserAgentData() (ret NavigatorUAData, ok bool) {
	ok = js.True == bindings.GetNavigatorUserAgentData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HardwareConcurrency returns the value of property "Navigator.hardwareConcurrency".
//
// It returns ok=false if there is no such property.
func (this Navigator) HardwareConcurrency() (ret uint64, ok bool) {
	ok = js.True == bindings.GetNavigatorHardwareConcurrency(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DeviceMemory returns the value of property "Navigator.deviceMemory".
//
// It returns ok=false if there is no such property.
func (this Navigator) DeviceMemory() (ret float64, ok bool) {
	ok = js.True == bindings.GetNavigatorDeviceMemory(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Storage returns the value of property "Navigator.storage".
//
// It returns ok=false if there is no such property.
func (this Navigator) Storage() (ret StorageManager, ok bool) {
	ok = js.True == bindings.GetNavigatorStorage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// StorageBuckets returns the value of property "Navigator.storageBuckets".
//
// It returns ok=false if there is no such property.
func (this Navigator) StorageBuckets() (ret StorageBucketManager, ok bool) {
	ok = js.True == bindings.GetNavigatorStorageBuckets(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Locks returns the value of property "Navigator.locks".
//
// It returns ok=false if there is no such property.
func (this Navigator) Locks() (ret LockManager, ok bool) {
	ok = js.True == bindings.GetNavigatorLocks(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AppCodeName returns the value of property "Navigator.appCodeName".
//
// It returns ok=false if there is no such property.
func (this Navigator) AppCodeName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorAppCodeName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AppName returns the value of property "Navigator.appName".
//
// It returns ok=false if there is no such property.
func (this Navigator) AppName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorAppName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AppVersion returns the value of property "Navigator.appVersion".
//
// It returns ok=false if there is no such property.
func (this Navigator) AppVersion() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorAppVersion(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Platform returns the value of property "Navigator.platform".
//
// It returns ok=false if there is no such property.
func (this Navigator) Platform() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorPlatform(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Product returns the value of property "Navigator.product".
//
// It returns ok=false if there is no such property.
func (this Navigator) Product() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorProduct(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ProductSub returns the value of property "Navigator.productSub".
//
// It returns ok=false if there is no such property.
func (this Navigator) ProductSub() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorProductSub(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UserAgent returns the value of property "Navigator.userAgent".
//
// It returns ok=false if there is no such property.
func (this Navigator) UserAgent() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorUserAgent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Vendor returns the value of property "Navigator.vendor".
//
// It returns ok=false if there is no such property.
func (this Navigator) Vendor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorVendor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// VendorSub returns the value of property "Navigator.vendorSub".
//
// It returns ok=false if there is no such property.
func (this Navigator) VendorSub() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorVendorSub(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Oscpu returns the value of property "Navigator.oscpu".
//
// It returns ok=false if there is no such property.
func (this Navigator) Oscpu() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorOscpu(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Language returns the value of property "Navigator.language".
//
// It returns ok=false if there is no such property.
func (this Navigator) Language() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorLanguage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Languages returns the value of property "Navigator.languages".
//
// It returns ok=false if there is no such property.
func (this Navigator) Languages() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetNavigatorLanguages(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OnLine returns the value of property "Navigator.onLine".
//
// It returns ok=false if there is no such property.
func (this Navigator) OnLine() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigatorOnLine(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CookieEnabled returns the value of property "Navigator.cookieEnabled".
//
// It returns ok=false if there is no such property.
func (this Navigator) CookieEnabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigatorCookieEnabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Plugins returns the value of property "Navigator.plugins".
//
// It returns ok=false if there is no such property.
func (this Navigator) Plugins() (ret PluginArray, ok bool) {
	ok = js.True == bindings.GetNavigatorPlugins(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MimeTypes returns the value of property "Navigator.mimeTypes".
//
// It returns ok=false if there is no such property.
func (this Navigator) MimeTypes() (ret MimeTypeArray, ok bool) {
	ok = js.True == bindings.GetNavigatorMimeTypes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PdfViewerEnabled returns the value of property "Navigator.pdfViewerEnabled".
//
// It returns ok=false if there is no such property.
func (this Navigator) PdfViewerEnabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigatorPdfViewerEnabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Webdriver returns the value of property "Navigator.webdriver".
//
// It returns ok=false if there is no such property.
func (this Navigator) Webdriver() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigatorWebdriver(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Ml returns the value of property "Navigator.ml".
//
// It returns ok=false if there is no such property.
func (this Navigator) Ml() (ret ML, ok bool) {
	ok = js.True == bindings.GetNavigatorMl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Connection returns the value of property "Navigator.connection".
//
// It returns ok=false if there is no such property.
func (this Navigator) Connection() (ret NetworkInformation, ok bool) {
	ok = js.True == bindings.GetNavigatorConnection(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Gpu returns the value of property "Navigator.gpu".
//
// It returns ok=false if there is no such property.
func (this Navigator) Gpu() (ret GPU, ok bool) {
	ok = js.True == bindings.GetNavigatorGpu(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasUpdateAdInterestGroups returns true if the method "Navigator.updateAdInterestGroups" exists.
func (this Navigator) HasUpdateAdInterestGroups() bool {
	return js.True == bindings.HasNavigatorUpdateAdInterestGroups(
		this.Ref(),
	)
}

// UpdateAdInterestGroupsFunc returns the method "Navigator.updateAdInterestGroups".
func (this Navigator) UpdateAdInterestGroupsFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.NavigatorUpdateAdInterestGroupsFunc(
			this.Ref(),
		),
	)
}

// UpdateAdInterestGroups calls the method "Navigator.updateAdInterestGroups".
func (this Navigator) UpdateAdInterestGroups() (ret js.Void) {
	bindings.CallNavigatorUpdateAdInterestGroups(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryUpdateAdInterestGroups calls the method "Navigator.updateAdInterestGroups"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryUpdateAdInterestGroups() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorUpdateAdInterestGroups(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSendBeacon returns true if the method "Navigator.sendBeacon" exists.
func (this Navigator) HasSendBeacon() bool {
	return js.True == bindings.HasNavigatorSendBeacon(
		this.Ref(),
	)
}

// SendBeaconFunc returns the method "Navigator.sendBeacon".
func (this Navigator) SendBeaconFunc() (fn js.Func[func(url js.String, data BodyInit) bool]) {
	return fn.FromRef(
		bindings.NavigatorSendBeaconFunc(
			this.Ref(),
		),
	)
}

// SendBeacon calls the method "Navigator.sendBeacon".
func (this Navigator) SendBeacon(url js.String, data BodyInit) (ret bool) {
	bindings.CallNavigatorSendBeacon(
		this.Ref(), js.Pointer(&ret),
		url.Ref(),
		data.Ref(),
	)

	return
}

// TrySendBeacon calls the method "Navigator.sendBeacon"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TrySendBeacon(url js.String, data BodyInit) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorSendBeacon(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		data.Ref(),
	)

	return
}

// HasSendBeacon1 returns true if the method "Navigator.sendBeacon" exists.
func (this Navigator) HasSendBeacon1() bool {
	return js.True == bindings.HasNavigatorSendBeacon1(
		this.Ref(),
	)
}

// SendBeacon1Func returns the method "Navigator.sendBeacon".
func (this Navigator) SendBeacon1Func() (fn js.Func[func(url js.String) bool]) {
	return fn.FromRef(
		bindings.NavigatorSendBeacon1Func(
			this.Ref(),
		),
	)
}

// SendBeacon1 calls the method "Navigator.sendBeacon".
func (this Navigator) SendBeacon1(url js.String) (ret bool) {
	bindings.CallNavigatorSendBeacon1(
		this.Ref(), js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TrySendBeacon1 calls the method "Navigator.sendBeacon"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TrySendBeacon1(url js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorSendBeacon1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasGetBattery returns true if the method "Navigator.getBattery" exists.
func (this Navigator) HasGetBattery() bool {
	return js.True == bindings.HasNavigatorGetBattery(
		this.Ref(),
	)
}

// GetBatteryFunc returns the method "Navigator.getBattery".
func (this Navigator) GetBatteryFunc() (fn js.Func[func() js.Promise[BatteryManager]]) {
	return fn.FromRef(
		bindings.NavigatorGetBatteryFunc(
			this.Ref(),
		),
	)
}

// GetBattery calls the method "Navigator.getBattery".
func (this Navigator) GetBattery() (ret js.Promise[BatteryManager]) {
	bindings.CallNavigatorGetBattery(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetBattery calls the method "Navigator.getBattery"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryGetBattery() (ret js.Promise[BatteryManager], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorGetBattery(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasVibrate returns true if the method "Navigator.vibrate" exists.
func (this Navigator) HasVibrate() bool {
	return js.True == bindings.HasNavigatorVibrate(
		this.Ref(),
	)
}

// VibrateFunc returns the method "Navigator.vibrate".
func (this Navigator) VibrateFunc() (fn js.Func[func(pattern VibratePattern) bool]) {
	return fn.FromRef(
		bindings.NavigatorVibrateFunc(
			this.Ref(),
		),
	)
}

// Vibrate calls the method "Navigator.vibrate".
func (this Navigator) Vibrate(pattern VibratePattern) (ret bool) {
	bindings.CallNavigatorVibrate(
		this.Ref(), js.Pointer(&ret),
		pattern.Ref(),
	)

	return
}

// TryVibrate calls the method "Navigator.vibrate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryVibrate(pattern VibratePattern) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorVibrate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		pattern.Ref(),
	)

	return
}

// HasGetGamepads returns true if the method "Navigator.getGamepads" exists.
func (this Navigator) HasGetGamepads() bool {
	return js.True == bindings.HasNavigatorGetGamepads(
		this.Ref(),
	)
}

// GetGamepadsFunc returns the method "Navigator.getGamepads".
func (this Navigator) GetGamepadsFunc() (fn js.Func[func() js.Array[Gamepad]]) {
	return fn.FromRef(
		bindings.NavigatorGetGamepadsFunc(
			this.Ref(),
		),
	)
}

// GetGamepads calls the method "Navigator.getGamepads".
func (this Navigator) GetGamepads() (ret js.Array[Gamepad]) {
	bindings.CallNavigatorGetGamepads(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetGamepads calls the method "Navigator.getGamepads"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryGetGamepads() (ret js.Array[Gamepad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorGetGamepads(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetInstalledRelatedApps returns true if the method "Navigator.getInstalledRelatedApps" exists.
func (this Navigator) HasGetInstalledRelatedApps() bool {
	return js.True == bindings.HasNavigatorGetInstalledRelatedApps(
		this.Ref(),
	)
}

// GetInstalledRelatedAppsFunc returns the method "Navigator.getInstalledRelatedApps".
func (this Navigator) GetInstalledRelatedAppsFunc() (fn js.Func[func() js.Promise[js.Array[RelatedApplication]]]) {
	return fn.FromRef(
		bindings.NavigatorGetInstalledRelatedAppsFunc(
			this.Ref(),
		),
	)
}

// GetInstalledRelatedApps calls the method "Navigator.getInstalledRelatedApps".
func (this Navigator) GetInstalledRelatedApps() (ret js.Promise[js.Array[RelatedApplication]]) {
	bindings.CallNavigatorGetInstalledRelatedApps(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetInstalledRelatedApps calls the method "Navigator.getInstalledRelatedApps"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryGetInstalledRelatedApps() (ret js.Promise[js.Array[RelatedApplication]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorGetInstalledRelatedApps(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRequestMediaKeySystemAccess returns true if the method "Navigator.requestMediaKeySystemAccess" exists.
func (this Navigator) HasRequestMediaKeySystemAccess() bool {
	return js.True == bindings.HasNavigatorRequestMediaKeySystemAccess(
		this.Ref(),
	)
}

// RequestMediaKeySystemAccessFunc returns the method "Navigator.requestMediaKeySystemAccess".
func (this Navigator) RequestMediaKeySystemAccessFunc() (fn js.Func[func(keySystem js.String, supportedConfigurations js.Array[MediaKeySystemConfiguration]) js.Promise[MediaKeySystemAccess]]) {
	return fn.FromRef(
		bindings.NavigatorRequestMediaKeySystemAccessFunc(
			this.Ref(),
		),
	)
}

// RequestMediaKeySystemAccess calls the method "Navigator.requestMediaKeySystemAccess".
func (this Navigator) RequestMediaKeySystemAccess(keySystem js.String, supportedConfigurations js.Array[MediaKeySystemConfiguration]) (ret js.Promise[MediaKeySystemAccess]) {
	bindings.CallNavigatorRequestMediaKeySystemAccess(
		this.Ref(), js.Pointer(&ret),
		keySystem.Ref(),
		supportedConfigurations.Ref(),
	)

	return
}

// TryRequestMediaKeySystemAccess calls the method "Navigator.requestMediaKeySystemAccess"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryRequestMediaKeySystemAccess(keySystem js.String, supportedConfigurations js.Array[MediaKeySystemConfiguration]) (ret js.Promise[MediaKeySystemAccess], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorRequestMediaKeySystemAccess(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		keySystem.Ref(),
		supportedConfigurations.Ref(),
	)

	return
}

// HasJoinAdInterestGroup returns true if the method "Navigator.joinAdInterestGroup" exists.
func (this Navigator) HasJoinAdInterestGroup() bool {
	return js.True == bindings.HasNavigatorJoinAdInterestGroup(
		this.Ref(),
	)
}

// JoinAdInterestGroupFunc returns the method "Navigator.joinAdInterestGroup".
func (this Navigator) JoinAdInterestGroupFunc() (fn js.Func[func(group AuctionAdInterestGroup) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigatorJoinAdInterestGroupFunc(
			this.Ref(),
		),
	)
}

// JoinAdInterestGroup calls the method "Navigator.joinAdInterestGroup".
func (this Navigator) JoinAdInterestGroup(group AuctionAdInterestGroup) (ret js.Promise[js.Void]) {
	bindings.CallNavigatorJoinAdInterestGroup(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&group),
	)

	return
}

// TryJoinAdInterestGroup calls the method "Navigator.joinAdInterestGroup"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryJoinAdInterestGroup(group AuctionAdInterestGroup) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorJoinAdInterestGroup(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&group),
	)

	return
}

// HasGetUserMedia returns true if the method "Navigator.getUserMedia" exists.
func (this Navigator) HasGetUserMedia() bool {
	return js.True == bindings.HasNavigatorGetUserMedia(
		this.Ref(),
	)
}

// GetUserMediaFunc returns the method "Navigator.getUserMedia".
func (this Navigator) GetUserMediaFunc() (fn js.Func[func(constraints MediaStreamConstraints, successCallback js.Func[func(stream MediaStream)], errorCallback js.Func[func(err DOMException)])]) {
	return fn.FromRef(
		bindings.NavigatorGetUserMediaFunc(
			this.Ref(),
		),
	)
}

// GetUserMedia calls the method "Navigator.getUserMedia".
func (this Navigator) GetUserMedia(constraints MediaStreamConstraints, successCallback js.Func[func(stream MediaStream)], errorCallback js.Func[func(err DOMException)]) (ret js.Void) {
	bindings.CallNavigatorGetUserMedia(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&constraints),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// TryGetUserMedia calls the method "Navigator.getUserMedia"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryGetUserMedia(constraints MediaStreamConstraints, successCallback js.Func[func(stream MediaStream)], errorCallback js.Func[func(err DOMException)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorGetUserMedia(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&constraints),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// HasLeaveAdInterestGroup returns true if the method "Navigator.leaveAdInterestGroup" exists.
func (this Navigator) HasLeaveAdInterestGroup() bool {
	return js.True == bindings.HasNavigatorLeaveAdInterestGroup(
		this.Ref(),
	)
}

// LeaveAdInterestGroupFunc returns the method "Navigator.leaveAdInterestGroup".
func (this Navigator) LeaveAdInterestGroupFunc() (fn js.Func[func(group AuctionAdInterestGroupKey) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigatorLeaveAdInterestGroupFunc(
			this.Ref(),
		),
	)
}

// LeaveAdInterestGroup calls the method "Navigator.leaveAdInterestGroup".
func (this Navigator) LeaveAdInterestGroup(group AuctionAdInterestGroupKey) (ret js.Promise[js.Void]) {
	bindings.CallNavigatorLeaveAdInterestGroup(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&group),
	)

	return
}

// TryLeaveAdInterestGroup calls the method "Navigator.leaveAdInterestGroup"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryLeaveAdInterestGroup(group AuctionAdInterestGroupKey) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorLeaveAdInterestGroup(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&group),
	)

	return
}

// HasLeaveAdInterestGroup1 returns true if the method "Navigator.leaveAdInterestGroup" exists.
func (this Navigator) HasLeaveAdInterestGroup1() bool {
	return js.True == bindings.HasNavigatorLeaveAdInterestGroup1(
		this.Ref(),
	)
}

// LeaveAdInterestGroup1Func returns the method "Navigator.leaveAdInterestGroup".
func (this Navigator) LeaveAdInterestGroup1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigatorLeaveAdInterestGroup1Func(
			this.Ref(),
		),
	)
}

// LeaveAdInterestGroup1 calls the method "Navigator.leaveAdInterestGroup".
func (this Navigator) LeaveAdInterestGroup1() (ret js.Promise[js.Void]) {
	bindings.CallNavigatorLeaveAdInterestGroup1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryLeaveAdInterestGroup1 calls the method "Navigator.leaveAdInterestGroup"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryLeaveAdInterestGroup1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorLeaveAdInterestGroup1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRunAdAuction returns true if the method "Navigator.runAdAuction" exists.
func (this Navigator) HasRunAdAuction() bool {
	return js.True == bindings.HasNavigatorRunAdAuction(
		this.Ref(),
	)
}

// RunAdAuctionFunc returns the method "Navigator.runAdAuction".
func (this Navigator) RunAdAuctionFunc() (fn js.Func[func(config AuctionAdConfig) js.Promise[OneOf_String_FencedFrameConfig]]) {
	return fn.FromRef(
		bindings.NavigatorRunAdAuctionFunc(
			this.Ref(),
		),
	)
}

// RunAdAuction calls the method "Navigator.runAdAuction".
func (this Navigator) RunAdAuction(config AuctionAdConfig) (ret js.Promise[OneOf_String_FencedFrameConfig]) {
	bindings.CallNavigatorRunAdAuction(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&config),
	)

	return
}

// TryRunAdAuction calls the method "Navigator.runAdAuction"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryRunAdAuction(config AuctionAdConfig) (ret js.Promise[OneOf_String_FencedFrameConfig], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorRunAdAuction(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&config),
	)

	return
}

// HasShare returns true if the method "Navigator.share" exists.
func (this Navigator) HasShare() bool {
	return js.True == bindings.HasNavigatorShare(
		this.Ref(),
	)
}

// ShareFunc returns the method "Navigator.share".
func (this Navigator) ShareFunc() (fn js.Func[func(data ShareData) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigatorShareFunc(
			this.Ref(),
		),
	)
}

// Share calls the method "Navigator.share".
func (this Navigator) Share(data ShareData) (ret js.Promise[js.Void]) {
	bindings.CallNavigatorShare(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&data),
	)

	return
}

// TryShare calls the method "Navigator.share"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryShare(data ShareData) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorShare(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&data),
	)

	return
}

// HasShare1 returns true if the method "Navigator.share" exists.
func (this Navigator) HasShare1() bool {
	return js.True == bindings.HasNavigatorShare1(
		this.Ref(),
	)
}

// Share1Func returns the method "Navigator.share".
func (this Navigator) Share1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigatorShare1Func(
			this.Ref(),
		),
	)
}

// Share1 calls the method "Navigator.share".
func (this Navigator) Share1() (ret js.Promise[js.Void]) {
	bindings.CallNavigatorShare1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryShare1 calls the method "Navigator.share"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryShare1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorShare1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCanShare returns true if the method "Navigator.canShare" exists.
func (this Navigator) HasCanShare() bool {
	return js.True == bindings.HasNavigatorCanShare(
		this.Ref(),
	)
}

// CanShareFunc returns the method "Navigator.canShare".
func (this Navigator) CanShareFunc() (fn js.Func[func(data ShareData) bool]) {
	return fn.FromRef(
		bindings.NavigatorCanShareFunc(
			this.Ref(),
		),
	)
}

// CanShare calls the method "Navigator.canShare".
func (this Navigator) CanShare(data ShareData) (ret bool) {
	bindings.CallNavigatorCanShare(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&data),
	)

	return
}

// TryCanShare calls the method "Navigator.canShare"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryCanShare(data ShareData) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorCanShare(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&data),
	)

	return
}

// HasCanShare1 returns true if the method "Navigator.canShare" exists.
func (this Navigator) HasCanShare1() bool {
	return js.True == bindings.HasNavigatorCanShare1(
		this.Ref(),
	)
}

// CanShare1Func returns the method "Navigator.canShare".
func (this Navigator) CanShare1Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.NavigatorCanShare1Func(
			this.Ref(),
		),
	)
}

// CanShare1 calls the method "Navigator.canShare".
func (this Navigator) CanShare1() (ret bool) {
	bindings.CallNavigatorCanShare1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCanShare1 calls the method "Navigator.canShare"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryCanShare1() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorCanShare1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRequestMIDIAccess returns true if the method "Navigator.requestMIDIAccess" exists.
func (this Navigator) HasRequestMIDIAccess() bool {
	return js.True == bindings.HasNavigatorRequestMIDIAccess(
		this.Ref(),
	)
}

// RequestMIDIAccessFunc returns the method "Navigator.requestMIDIAccess".
func (this Navigator) RequestMIDIAccessFunc() (fn js.Func[func(options MIDIOptions) js.Promise[MIDIAccess]]) {
	return fn.FromRef(
		bindings.NavigatorRequestMIDIAccessFunc(
			this.Ref(),
		),
	)
}

// RequestMIDIAccess calls the method "Navigator.requestMIDIAccess".
func (this Navigator) RequestMIDIAccess(options MIDIOptions) (ret js.Promise[MIDIAccess]) {
	bindings.CallNavigatorRequestMIDIAccess(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRequestMIDIAccess calls the method "Navigator.requestMIDIAccess"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryRequestMIDIAccess(options MIDIOptions) (ret js.Promise[MIDIAccess], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorRequestMIDIAccess(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasRequestMIDIAccess1 returns true if the method "Navigator.requestMIDIAccess" exists.
func (this Navigator) HasRequestMIDIAccess1() bool {
	return js.True == bindings.HasNavigatorRequestMIDIAccess1(
		this.Ref(),
	)
}

// RequestMIDIAccess1Func returns the method "Navigator.requestMIDIAccess".
func (this Navigator) RequestMIDIAccess1Func() (fn js.Func[func() js.Promise[MIDIAccess]]) {
	return fn.FromRef(
		bindings.NavigatorRequestMIDIAccess1Func(
			this.Ref(),
		),
	)
}

// RequestMIDIAccess1 calls the method "Navigator.requestMIDIAccess".
func (this Navigator) RequestMIDIAccess1() (ret js.Promise[MIDIAccess]) {
	bindings.CallNavigatorRequestMIDIAccess1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestMIDIAccess1 calls the method "Navigator.requestMIDIAccess"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryRequestMIDIAccess1() (ret js.Promise[MIDIAccess], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorRequestMIDIAccess1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetAutoplayPolicy returns true if the method "Navigator.getAutoplayPolicy" exists.
func (this Navigator) HasGetAutoplayPolicy() bool {
	return js.True == bindings.HasNavigatorGetAutoplayPolicy(
		this.Ref(),
	)
}

// GetAutoplayPolicyFunc returns the method "Navigator.getAutoplayPolicy".
func (this Navigator) GetAutoplayPolicyFunc() (fn js.Func[func(typ AutoplayPolicyMediaType) AutoplayPolicy]) {
	return fn.FromRef(
		bindings.NavigatorGetAutoplayPolicyFunc(
			this.Ref(),
		),
	)
}

// GetAutoplayPolicy calls the method "Navigator.getAutoplayPolicy".
func (this Navigator) GetAutoplayPolicy(typ AutoplayPolicyMediaType) (ret AutoplayPolicy) {
	bindings.CallNavigatorGetAutoplayPolicy(
		this.Ref(), js.Pointer(&ret),
		uint32(typ),
	)

	return
}

// TryGetAutoplayPolicy calls the method "Navigator.getAutoplayPolicy"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryGetAutoplayPolicy(typ AutoplayPolicyMediaType) (ret AutoplayPolicy, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorGetAutoplayPolicy(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
	)

	return
}

// HasGetAutoplayPolicy1 returns true if the method "Navigator.getAutoplayPolicy" exists.
func (this Navigator) HasGetAutoplayPolicy1() bool {
	return js.True == bindings.HasNavigatorGetAutoplayPolicy1(
		this.Ref(),
	)
}

// GetAutoplayPolicy1Func returns the method "Navigator.getAutoplayPolicy".
func (this Navigator) GetAutoplayPolicy1Func() (fn js.Func[func(element HTMLMediaElement) AutoplayPolicy]) {
	return fn.FromRef(
		bindings.NavigatorGetAutoplayPolicy1Func(
			this.Ref(),
		),
	)
}

// GetAutoplayPolicy1 calls the method "Navigator.getAutoplayPolicy".
func (this Navigator) GetAutoplayPolicy1(element HTMLMediaElement) (ret AutoplayPolicy) {
	bindings.CallNavigatorGetAutoplayPolicy1(
		this.Ref(), js.Pointer(&ret),
		element.Ref(),
	)

	return
}

// TryGetAutoplayPolicy1 calls the method "Navigator.getAutoplayPolicy"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryGetAutoplayPolicy1(element HTMLMediaElement) (ret AutoplayPolicy, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorGetAutoplayPolicy1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
	)

	return
}

// HasGetAutoplayPolicy2 returns true if the method "Navigator.getAutoplayPolicy" exists.
func (this Navigator) HasGetAutoplayPolicy2() bool {
	return js.True == bindings.HasNavigatorGetAutoplayPolicy2(
		this.Ref(),
	)
}

// GetAutoplayPolicy2Func returns the method "Navigator.getAutoplayPolicy".
func (this Navigator) GetAutoplayPolicy2Func() (fn js.Func[func(context AudioContext) AutoplayPolicy]) {
	return fn.FromRef(
		bindings.NavigatorGetAutoplayPolicy2Func(
			this.Ref(),
		),
	)
}

// GetAutoplayPolicy2 calls the method "Navigator.getAutoplayPolicy".
func (this Navigator) GetAutoplayPolicy2(context AudioContext) (ret AutoplayPolicy) {
	bindings.CallNavigatorGetAutoplayPolicy2(
		this.Ref(), js.Pointer(&ret),
		context.Ref(),
	)

	return
}

// TryGetAutoplayPolicy2 calls the method "Navigator.getAutoplayPolicy"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryGetAutoplayPolicy2(context AudioContext) (ret AutoplayPolicy, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorGetAutoplayPolicy2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		context.Ref(),
	)

	return
}

// HasTaintEnabled returns true if the method "Navigator.taintEnabled" exists.
func (this Navigator) HasTaintEnabled() bool {
	return js.True == bindings.HasNavigatorTaintEnabled(
		this.Ref(),
	)
}

// TaintEnabledFunc returns the method "Navigator.taintEnabled".
func (this Navigator) TaintEnabledFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.NavigatorTaintEnabledFunc(
			this.Ref(),
		),
	)
}

// TaintEnabled calls the method "Navigator.taintEnabled".
func (this Navigator) TaintEnabled() (ret bool) {
	bindings.CallNavigatorTaintEnabled(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTaintEnabled calls the method "Navigator.taintEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryTaintEnabled() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorTaintEnabled(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRegisterProtocolHandler returns true if the method "Navigator.registerProtocolHandler" exists.
func (this Navigator) HasRegisterProtocolHandler() bool {
	return js.True == bindings.HasNavigatorRegisterProtocolHandler(
		this.Ref(),
	)
}

// RegisterProtocolHandlerFunc returns the method "Navigator.registerProtocolHandler".
func (this Navigator) RegisterProtocolHandlerFunc() (fn js.Func[func(scheme js.String, url js.String)]) {
	return fn.FromRef(
		bindings.NavigatorRegisterProtocolHandlerFunc(
			this.Ref(),
		),
	)
}

// RegisterProtocolHandler calls the method "Navigator.registerProtocolHandler".
func (this Navigator) RegisterProtocolHandler(scheme js.String, url js.String) (ret js.Void) {
	bindings.CallNavigatorRegisterProtocolHandler(
		this.Ref(), js.Pointer(&ret),
		scheme.Ref(),
		url.Ref(),
	)

	return
}

// TryRegisterProtocolHandler calls the method "Navigator.registerProtocolHandler"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryRegisterProtocolHandler(scheme js.String, url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorRegisterProtocolHandler(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		scheme.Ref(),
		url.Ref(),
	)

	return
}

// HasUnregisterProtocolHandler returns true if the method "Navigator.unregisterProtocolHandler" exists.
func (this Navigator) HasUnregisterProtocolHandler() bool {
	return js.True == bindings.HasNavigatorUnregisterProtocolHandler(
		this.Ref(),
	)
}

// UnregisterProtocolHandlerFunc returns the method "Navigator.unregisterProtocolHandler".
func (this Navigator) UnregisterProtocolHandlerFunc() (fn js.Func[func(scheme js.String, url js.String)]) {
	return fn.FromRef(
		bindings.NavigatorUnregisterProtocolHandlerFunc(
			this.Ref(),
		),
	)
}

// UnregisterProtocolHandler calls the method "Navigator.unregisterProtocolHandler".
func (this Navigator) UnregisterProtocolHandler(scheme js.String, url js.String) (ret js.Void) {
	bindings.CallNavigatorUnregisterProtocolHandler(
		this.Ref(), js.Pointer(&ret),
		scheme.Ref(),
		url.Ref(),
	)

	return
}

// TryUnregisterProtocolHandler calls the method "Navigator.unregisterProtocolHandler"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryUnregisterProtocolHandler(scheme js.String, url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorUnregisterProtocolHandler(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		scheme.Ref(),
		url.Ref(),
	)

	return
}

// HasJavaEnabled returns true if the method "Navigator.javaEnabled" exists.
func (this Navigator) HasJavaEnabled() bool {
	return js.True == bindings.HasNavigatorJavaEnabled(
		this.Ref(),
	)
}

// JavaEnabledFunc returns the method "Navigator.javaEnabled".
func (this Navigator) JavaEnabledFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.NavigatorJavaEnabledFunc(
			this.Ref(),
		),
	)
}

// JavaEnabled calls the method "Navigator.javaEnabled".
func (this Navigator) JavaEnabled() (ret bool) {
	bindings.CallNavigatorJavaEnabled(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryJavaEnabled calls the method "Navigator.javaEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryJavaEnabled() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorJavaEnabled(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetAppBadge returns true if the method "Navigator.setAppBadge" exists.
func (this Navigator) HasSetAppBadge() bool {
	return js.True == bindings.HasNavigatorSetAppBadge(
		this.Ref(),
	)
}

// SetAppBadgeFunc returns the method "Navigator.setAppBadge".
func (this Navigator) SetAppBadgeFunc() (fn js.Func[func(contents uint64) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigatorSetAppBadgeFunc(
			this.Ref(),
		),
	)
}

// SetAppBadge calls the method "Navigator.setAppBadge".
func (this Navigator) SetAppBadge(contents uint64) (ret js.Promise[js.Void]) {
	bindings.CallNavigatorSetAppBadge(
		this.Ref(), js.Pointer(&ret),
		float64(contents),
	)

	return
}

// TrySetAppBadge calls the method "Navigator.setAppBadge"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TrySetAppBadge(contents uint64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorSetAppBadge(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(contents),
	)

	return
}

// HasSetAppBadge1 returns true if the method "Navigator.setAppBadge" exists.
func (this Navigator) HasSetAppBadge1() bool {
	return js.True == bindings.HasNavigatorSetAppBadge1(
		this.Ref(),
	)
}

// SetAppBadge1Func returns the method "Navigator.setAppBadge".
func (this Navigator) SetAppBadge1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigatorSetAppBadge1Func(
			this.Ref(),
		),
	)
}

// SetAppBadge1 calls the method "Navigator.setAppBadge".
func (this Navigator) SetAppBadge1() (ret js.Promise[js.Void]) {
	bindings.CallNavigatorSetAppBadge1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySetAppBadge1 calls the method "Navigator.setAppBadge"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TrySetAppBadge1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorSetAppBadge1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClearAppBadge returns true if the method "Navigator.clearAppBadge" exists.
func (this Navigator) HasClearAppBadge() bool {
	return js.True == bindings.HasNavigatorClearAppBadge(
		this.Ref(),
	)
}

// ClearAppBadgeFunc returns the method "Navigator.clearAppBadge".
func (this Navigator) ClearAppBadgeFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigatorClearAppBadgeFunc(
			this.Ref(),
		),
	)
}

// ClearAppBadge calls the method "Navigator.clearAppBadge".
func (this Navigator) ClearAppBadge() (ret js.Promise[js.Void]) {
	bindings.CallNavigatorClearAppBadge(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClearAppBadge calls the method "Navigator.clearAppBadge"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Navigator) TryClearAppBadge() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorClearAppBadge(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this LaunchParams) TargetURL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLaunchParamsTargetURL(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Files returns the value of property "LaunchParams.files".
//
// It returns ok=false if there is no such property.
func (this LaunchParams) Files() (ret js.FrozenArray[FileSystemHandle], ok bool) {
	ok = js.True == bindings.GetLaunchParamsFiles(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// HasSetConsumer returns true if the method "LaunchQueue.setConsumer" exists.
func (this LaunchQueue) HasSetConsumer() bool {
	return js.True == bindings.HasLaunchQueueSetConsumer(
		this.Ref(),
	)
}

// SetConsumerFunc returns the method "LaunchQueue.setConsumer".
func (this LaunchQueue) SetConsumerFunc() (fn js.Func[func(consumer js.Func[func(params LaunchParams) js.Any])]) {
	return fn.FromRef(
		bindings.LaunchQueueSetConsumerFunc(
			this.Ref(),
		),
	)
}

// SetConsumer calls the method "LaunchQueue.setConsumer".
func (this LaunchQueue) SetConsumer(consumer js.Func[func(params LaunchParams) js.Any]) (ret js.Void) {
	bindings.CallLaunchQueueSetConsumer(
		this.Ref(), js.Pointer(&ret),
		consumer.Ref(),
	)

	return
}

// TrySetConsumer calls the method "LaunchQueue.setConsumer"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this LaunchQueue) TrySetConsumer(consumer js.Func[func(params LaunchParams) js.Any]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLaunchQueueSetConsumer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		consumer.Ref(),
	)

	return
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
	//
	// NOTE: FFI_USE_Once MUST be set to true to make this field effective.
	Once bool

	FFI_USE_Once bool // for Once.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FenceEvent with all fields set.
func (p FenceEvent) FromRef(ref js.Ref) FenceEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FenceEvent in the application heap.
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

// HasReportEvent returns true if the method "Fence.reportEvent" exists.
func (this Fence) HasReportEvent() bool {
	return js.True == bindings.HasFenceReportEvent(
		this.Ref(),
	)
}

// ReportEventFunc returns the method "Fence.reportEvent".
func (this Fence) ReportEventFunc() (fn js.Func[func(event ReportEventType)]) {
	return fn.FromRef(
		bindings.FenceReportEventFunc(
			this.Ref(),
		),
	)
}

// ReportEvent calls the method "Fence.reportEvent".
func (this Fence) ReportEvent(event ReportEventType) (ret js.Void) {
	bindings.CallFenceReportEvent(
		this.Ref(), js.Pointer(&ret),
		event.Ref(),
	)

	return
}

// TryReportEvent calls the method "Fence.reportEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Fence) TryReportEvent(event ReportEventType) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFenceReportEvent(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		event.Ref(),
	)

	return
}

// HasSetReportEventDataForAutomaticBeacons returns true if the method "Fence.setReportEventDataForAutomaticBeacons" exists.
func (this Fence) HasSetReportEventDataForAutomaticBeacons() bool {
	return js.True == bindings.HasFenceSetReportEventDataForAutomaticBeacons(
		this.Ref(),
	)
}

// SetReportEventDataForAutomaticBeaconsFunc returns the method "Fence.setReportEventDataForAutomaticBeacons".
func (this Fence) SetReportEventDataForAutomaticBeaconsFunc() (fn js.Func[func(event FenceEvent)]) {
	return fn.FromRef(
		bindings.FenceSetReportEventDataForAutomaticBeaconsFunc(
			this.Ref(),
		),
	)
}

// SetReportEventDataForAutomaticBeacons calls the method "Fence.setReportEventDataForAutomaticBeacons".
func (this Fence) SetReportEventDataForAutomaticBeacons(event FenceEvent) (ret js.Void) {
	bindings.CallFenceSetReportEventDataForAutomaticBeacons(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&event),
	)

	return
}

// TrySetReportEventDataForAutomaticBeacons calls the method "Fence.setReportEventDataForAutomaticBeacons"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Fence) TrySetReportEventDataForAutomaticBeacons(event FenceEvent) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFenceSetReportEventDataForAutomaticBeacons(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&event),
	)

	return
}

// HasGetNestedConfigs returns true if the method "Fence.getNestedConfigs" exists.
func (this Fence) HasGetNestedConfigs() bool {
	return js.True == bindings.HasFenceGetNestedConfigs(
		this.Ref(),
	)
}

// GetNestedConfigsFunc returns the method "Fence.getNestedConfigs".
func (this Fence) GetNestedConfigsFunc() (fn js.Func[func() js.Array[FencedFrameConfig]]) {
	return fn.FromRef(
		bindings.FenceGetNestedConfigsFunc(
			this.Ref(),
		),
	)
}

// GetNestedConfigs calls the method "Fence.getNestedConfigs".
func (this Fence) GetNestedConfigs() (ret js.Array[FencedFrameConfig]) {
	bindings.CallFenceGetNestedConfigs(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetNestedConfigs calls the method "Fence.getNestedConfigs"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Fence) TryGetNestedConfigs() (ret js.Array[FencedFrameConfig], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFenceGetNestedConfigs(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// HasPostMessage returns true if the method "PortalHost.postMessage" exists.
func (this PortalHost) HasPostMessage() bool {
	return js.True == bindings.HasPortalHostPostMessage(
		this.Ref(),
	)
}

// PostMessageFunc returns the method "PortalHost.postMessage".
func (this PortalHost) PostMessageFunc() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	return fn.FromRef(
		bindings.PortalHostPostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage calls the method "PortalHost.postMessage".
func (this PortalHost) PostMessage(message js.Any, options StructuredSerializeOptions) (ret js.Void) {
	bindings.CallPortalHostPostMessage(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryPostMessage calls the method "PortalHost.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PortalHost) TryPostMessage(message js.Any, options StructuredSerializeOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPortalHostPostMessage(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasPostMessage1 returns true if the method "PortalHost.postMessage" exists.
func (this PortalHost) HasPostMessage1() bool {
	return js.True == bindings.HasPortalHostPostMessage1(
		this.Ref(),
	)
}

// PostMessage1Func returns the method "PortalHost.postMessage".
func (this PortalHost) PostMessage1Func() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.PortalHostPostMessage1Func(
			this.Ref(),
		),
	)
}

// PostMessage1 calls the method "PortalHost.postMessage".
func (this PortalHost) PostMessage1(message js.Any) (ret js.Void) {
	bindings.CallPortalHostPostMessage1(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPostMessage1 calls the method "PortalHost.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PortalHost) TryPostMessage1(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPortalHostPostMessage1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this ScreenOrientation) Type() (ret OrientationType, ok bool) {
	ok = js.True == bindings.GetScreenOrientationType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Angle returns the value of property "ScreenOrientation.angle".
//
// It returns ok=false if there is no such property.
func (this ScreenOrientation) Angle() (ret uint16, ok bool) {
	ok = js.True == bindings.GetScreenOrientationAngle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasLock returns true if the method "ScreenOrientation.lock" exists.
func (this ScreenOrientation) HasLock() bool {
	return js.True == bindings.HasScreenOrientationLock(
		this.Ref(),
	)
}

// LockFunc returns the method "ScreenOrientation.lock".
func (this ScreenOrientation) LockFunc() (fn js.Func[func(orientation OrientationLockType) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ScreenOrientationLockFunc(
			this.Ref(),
		),
	)
}

// Lock calls the method "ScreenOrientation.lock".
func (this ScreenOrientation) Lock(orientation OrientationLockType) (ret js.Promise[js.Void]) {
	bindings.CallScreenOrientationLock(
		this.Ref(), js.Pointer(&ret),
		uint32(orientation),
	)

	return
}

// TryLock calls the method "ScreenOrientation.lock"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this ScreenOrientation) TryLock(orientation OrientationLockType) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryScreenOrientationLock(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(orientation),
	)

	return
}

// HasUnlock returns true if the method "ScreenOrientation.unlock" exists.
func (this ScreenOrientation) HasUnlock() bool {
	return js.True == bindings.HasScreenOrientationUnlock(
		this.Ref(),
	)
}

// UnlockFunc returns the method "ScreenOrientation.unlock".
func (this ScreenOrientation) UnlockFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ScreenOrientationUnlockFunc(
			this.Ref(),
		),
	)
}

// Unlock calls the method "ScreenOrientation.unlock".
func (this ScreenOrientation) Unlock() (ret js.Void) {
	bindings.CallScreenOrientationUnlock(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryUnlock calls the method "ScreenOrientation.unlock"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this ScreenOrientation) TryUnlock() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryScreenOrientationUnlock(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this Screen) AvailWidth() (ret int32, ok bool) {
	ok = js.True == bindings.GetScreenAvailWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AvailHeight returns the value of property "Screen.availHeight".
//
// It returns ok=false if there is no such property.
func (this Screen) AvailHeight() (ret int32, ok bool) {
	ok = js.True == bindings.GetScreenAvailHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "Screen.width".
//
// It returns ok=false if there is no such property.
func (this Screen) Width() (ret int32, ok bool) {
	ok = js.True == bindings.GetScreenWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "Screen.height".
//
// It returns ok=false if there is no such property.
func (this Screen) Height() (ret int32, ok bool) {
	ok = js.True == bindings.GetScreenHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ColorDepth returns the value of property "Screen.colorDepth".
//
// It returns ok=false if there is no such property.
func (this Screen) ColorDepth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetScreenColorDepth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PixelDepth returns the value of property "Screen.pixelDepth".
//
// It returns ok=false if there is no such property.
func (this Screen) PixelDepth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetScreenPixelDepth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsExtended returns the value of property "Screen.isExtended".
//
// It returns ok=false if there is no such property.
func (this Screen) IsExtended() (ret bool, ok bool) {
	ok = js.True == bindings.GetScreenIsExtended(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Orientation returns the value of property "Screen.orientation".
//
// It returns ok=false if there is no such property.
func (this Screen) Orientation() (ret ScreenOrientation, ok bool) {
	ok = js.True == bindings.GetScreenOrientation(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this VisualViewport) OffsetLeft() (ret float64, ok bool) {
	ok = js.True == bindings.GetVisualViewportOffsetLeft(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OffsetTop returns the value of property "VisualViewport.offsetTop".
//
// It returns ok=false if there is no such property.
func (this VisualViewport) OffsetTop() (ret float64, ok bool) {
	ok = js.True == bindings.GetVisualViewportOffsetTop(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PageLeft returns the value of property "VisualViewport.pageLeft".
//
// It returns ok=false if there is no such property.
func (this VisualViewport) PageLeft() (ret float64, ok bool) {
	ok = js.True == bindings.GetVisualViewportPageLeft(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PageTop returns the value of property "VisualViewport.pageTop".
//
// It returns ok=false if there is no such property.
func (this VisualViewport) PageTop() (ret float64, ok bool) {
	ok = js.True == bindings.GetVisualViewportPageTop(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "VisualViewport.width".
//
// It returns ok=false if there is no such property.
func (this VisualViewport) Width() (ret float64, ok bool) {
	ok = js.True == bindings.GetVisualViewportWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "VisualViewport.height".
//
// It returns ok=false if there is no such property.
func (this VisualViewport) Height() (ret float64, ok bool) {
	ok = js.True == bindings.GetVisualViewportHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Scale returns the value of property "VisualViewport.scale".
//
// It returns ok=false if there is no such property.
func (this VisualViewport) Scale() (ret float64, ok bool) {
	ok = js.True == bindings.GetVisualViewportScale(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type SharedStorageRunOperationMethodOptions struct {
	// Data is "SharedStorageRunOperationMethodOptions.data"
	//
	// Optional
	Data js.Object
	// ResolveToConfig is "SharedStorageRunOperationMethodOptions.resolveToConfig"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ResolveToConfig MUST be set to true to make this field effective.
	ResolveToConfig bool
	// KeepAlive is "SharedStorageRunOperationMethodOptions.keepAlive"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_KeepAlive MUST be set to true to make this field effective.
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

// New creates a new SharedStorageRunOperationMethodOptions in the application heap.
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

// New creates a new SharedStorageUrlWithMetadata in the application heap.
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

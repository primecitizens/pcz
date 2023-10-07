// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncDispatch returns true if the method "MLCommandEncoder.dispatch" exists.
func (this MLCommandEncoder) HasFuncDispatch() bool {
	return js.True == bindings.HasFuncMLCommandEncoderDispatch(
		this.ref,
	)
}

// FuncDispatch returns the method "MLCommandEncoder.dispatch".
func (this MLCommandEncoder) FuncDispatch() (fn js.Func[func(graph MLGraph, inputs MLNamedGPUResources, outputs MLNamedGPUResources)]) {
	bindings.FuncMLCommandEncoderDispatch(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Dispatch calls the method "MLCommandEncoder.dispatch".
func (this MLCommandEncoder) Dispatch(graph MLGraph, inputs MLNamedGPUResources, outputs MLNamedGPUResources) (ret js.Void) {
	bindings.CallMLCommandEncoderDispatch(
		this.ref, js.Pointer(&ret),
		graph.Ref(),
		inputs.Ref(),
		outputs.Ref(),
	)

	return
}

// TryDispatch calls the method "MLCommandEncoder.dispatch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLCommandEncoder) TryDispatch(graph MLGraph, inputs MLNamedGPUResources, outputs MLNamedGPUResources) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCommandEncoderDispatch(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		graph.Ref(),
		inputs.Ref(),
		outputs.Ref(),
	)

	return
}

// HasFuncFinish returns true if the method "MLCommandEncoder.finish" exists.
func (this MLCommandEncoder) HasFuncFinish() bool {
	return js.True == bindings.HasFuncMLCommandEncoderFinish(
		this.ref,
	)
}

// FuncFinish returns the method "MLCommandEncoder.finish".
func (this MLCommandEncoder) FuncFinish() (fn js.Func[func(descriptor GPUCommandBufferDescriptor) GPUCommandBuffer]) {
	bindings.FuncMLCommandEncoderFinish(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Finish calls the method "MLCommandEncoder.finish".
func (this MLCommandEncoder) Finish(descriptor GPUCommandBufferDescriptor) (ret GPUCommandBuffer) {
	bindings.CallMLCommandEncoderFinish(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryFinish calls the method "MLCommandEncoder.finish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLCommandEncoder) TryFinish(descriptor GPUCommandBufferDescriptor) (ret GPUCommandBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCommandEncoderFinish(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncFinish1 returns true if the method "MLCommandEncoder.finish" exists.
func (this MLCommandEncoder) HasFuncFinish1() bool {
	return js.True == bindings.HasFuncMLCommandEncoderFinish1(
		this.ref,
	)
}

// FuncFinish1 returns the method "MLCommandEncoder.finish".
func (this MLCommandEncoder) FuncFinish1() (fn js.Func[func() GPUCommandBuffer]) {
	bindings.FuncMLCommandEncoderFinish1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Finish1 calls the method "MLCommandEncoder.finish".
func (this MLCommandEncoder) Finish1() (ret GPUCommandBuffer) {
	bindings.CallMLCommandEncoderFinish1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFinish1 calls the method "MLCommandEncoder.finish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLCommandEncoder) TryFinish1() (ret GPUCommandBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCommandEncoderFinish1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInitializeGraph returns true if the method "MLCommandEncoder.initializeGraph" exists.
func (this MLCommandEncoder) HasFuncInitializeGraph() bool {
	return js.True == bindings.HasFuncMLCommandEncoderInitializeGraph(
		this.ref,
	)
}

// FuncInitializeGraph returns the method "MLCommandEncoder.initializeGraph".
func (this MLCommandEncoder) FuncInitializeGraph() (fn js.Func[func(graph MLGraph)]) {
	bindings.FuncMLCommandEncoderInitializeGraph(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitializeGraph calls the method "MLCommandEncoder.initializeGraph".
func (this MLCommandEncoder) InitializeGraph(graph MLGraph) (ret js.Void) {
	bindings.CallMLCommandEncoderInitializeGraph(
		this.ref, js.Pointer(&ret),
		graph.Ref(),
	)

	return
}

// TryInitializeGraph calls the method "MLCommandEncoder.initializeGraph"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLCommandEncoder) TryInitializeGraph(graph MLGraph) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCommandEncoderInitializeGraph(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		graph.Ref(),
	)

	return
}

type MLContext struct {
	ref js.Ref
}

func (this MLContext) Once() MLContext {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncCompute returns true if the method "MLContext.compute" exists.
func (this MLContext) HasFuncCompute() bool {
	return js.True == bindings.HasFuncMLContextCompute(
		this.ref,
	)
}

// FuncCompute returns the method "MLContext.compute".
func (this MLContext) FuncCompute() (fn js.Func[func(graph MLGraph, inputs MLNamedArrayBufferViews, outputs MLNamedArrayBufferViews) js.Promise[MLComputeResult]]) {
	bindings.FuncMLContextCompute(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Compute calls the method "MLContext.compute".
func (this MLContext) Compute(graph MLGraph, inputs MLNamedArrayBufferViews, outputs MLNamedArrayBufferViews) (ret js.Promise[MLComputeResult]) {
	bindings.CallMLContextCompute(
		this.ref, js.Pointer(&ret),
		graph.Ref(),
		inputs.Ref(),
		outputs.Ref(),
	)

	return
}

// TryCompute calls the method "MLContext.compute"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLContext) TryCompute(graph MLGraph, inputs MLNamedArrayBufferViews, outputs MLNamedArrayBufferViews) (ret js.Promise[MLComputeResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLContextCompute(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		graph.Ref(),
		inputs.Ref(),
		outputs.Ref(),
	)

	return
}

// HasFuncComputeSync returns true if the method "MLContext.computeSync" exists.
func (this MLContext) HasFuncComputeSync() bool {
	return js.True == bindings.HasFuncMLContextComputeSync(
		this.ref,
	)
}

// FuncComputeSync returns the method "MLContext.computeSync".
func (this MLContext) FuncComputeSync() (fn js.Func[func(graph MLGraph, inputs MLNamedArrayBufferViews, outputs MLNamedArrayBufferViews)]) {
	bindings.FuncMLContextComputeSync(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ComputeSync calls the method "MLContext.computeSync".
func (this MLContext) ComputeSync(graph MLGraph, inputs MLNamedArrayBufferViews, outputs MLNamedArrayBufferViews) (ret js.Void) {
	bindings.CallMLContextComputeSync(
		this.ref, js.Pointer(&ret),
		graph.Ref(),
		inputs.Ref(),
		outputs.Ref(),
	)

	return
}

// TryComputeSync calls the method "MLContext.computeSync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLContext) TryComputeSync(graph MLGraph, inputs MLNamedArrayBufferViews, outputs MLNamedArrayBufferViews) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLContextComputeSync(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		graph.Ref(),
		inputs.Ref(),
		outputs.Ref(),
	)

	return
}

// HasFuncCreateCommandEncoder returns true if the method "MLContext.createCommandEncoder" exists.
func (this MLContext) HasFuncCreateCommandEncoder() bool {
	return js.True == bindings.HasFuncMLContextCreateCommandEncoder(
		this.ref,
	)
}

// FuncCreateCommandEncoder returns the method "MLContext.createCommandEncoder".
func (this MLContext) FuncCreateCommandEncoder() (fn js.Func[func() MLCommandEncoder]) {
	bindings.FuncMLContextCreateCommandEncoder(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateCommandEncoder calls the method "MLContext.createCommandEncoder".
func (this MLContext) CreateCommandEncoder() (ret MLCommandEncoder) {
	bindings.CallMLContextCreateCommandEncoder(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateCommandEncoder calls the method "MLContext.createCommandEncoder"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLContext) TryCreateCommandEncoder() (ret MLCommandEncoder, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLContextCreateCommandEncoder(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *MLContextOptions) UpdateFrom(ref js.Ref) {
	bindings.MLContextOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLContextOptions) Update(ref js.Ref) {
	bindings.MLContextOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLContextOptions) FreeMembers(recursive bool) {
}

type ML struct {
	ref js.Ref
}

func (this ML) Once() ML {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncCreateContext returns true if the method "ML.createContext" exists.
func (this ML) HasFuncCreateContext() bool {
	return js.True == bindings.HasFuncMLCreateContext(
		this.ref,
	)
}

// FuncCreateContext returns the method "ML.createContext".
func (this ML) FuncCreateContext() (fn js.Func[func(options MLContextOptions) js.Promise[MLContext]]) {
	bindings.FuncMLCreateContext(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateContext calls the method "ML.createContext".
func (this ML) CreateContext(options MLContextOptions) (ret js.Promise[MLContext]) {
	bindings.CallMLCreateContext(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryCreateContext calls the method "ML.createContext"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ML) TryCreateContext(options MLContextOptions) (ret js.Promise[MLContext], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCreateContext(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncCreateContext1 returns true if the method "ML.createContext" exists.
func (this ML) HasFuncCreateContext1() bool {
	return js.True == bindings.HasFuncMLCreateContext1(
		this.ref,
	)
}

// FuncCreateContext1 returns the method "ML.createContext".
func (this ML) FuncCreateContext1() (fn js.Func[func() js.Promise[MLContext]]) {
	bindings.FuncMLCreateContext1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateContext1 calls the method "ML.createContext".
func (this ML) CreateContext1() (ret js.Promise[MLContext]) {
	bindings.CallMLCreateContext1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateContext1 calls the method "ML.createContext"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ML) TryCreateContext1() (ret js.Promise[MLContext], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCreateContext1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateContext2 returns true if the method "ML.createContext" exists.
func (this ML) HasFuncCreateContext2() bool {
	return js.True == bindings.HasFuncMLCreateContext2(
		this.ref,
	)
}

// FuncCreateContext2 returns the method "ML.createContext".
func (this ML) FuncCreateContext2() (fn js.Func[func(gpuDevice GPUDevice) js.Promise[MLContext]]) {
	bindings.FuncMLCreateContext2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateContext2 calls the method "ML.createContext".
func (this ML) CreateContext2(gpuDevice GPUDevice) (ret js.Promise[MLContext]) {
	bindings.CallMLCreateContext2(
		this.ref, js.Pointer(&ret),
		gpuDevice.Ref(),
	)

	return
}

// TryCreateContext2 calls the method "ML.createContext"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ML) TryCreateContext2(gpuDevice GPUDevice) (ret js.Promise[MLContext], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCreateContext2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		gpuDevice.Ref(),
	)

	return
}

// HasFuncCreateContextSync returns true if the method "ML.createContextSync" exists.
func (this ML) HasFuncCreateContextSync() bool {
	return js.True == bindings.HasFuncMLCreateContextSync(
		this.ref,
	)
}

// FuncCreateContextSync returns the method "ML.createContextSync".
func (this ML) FuncCreateContextSync() (fn js.Func[func(options MLContextOptions) MLContext]) {
	bindings.FuncMLCreateContextSync(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateContextSync calls the method "ML.createContextSync".
func (this ML) CreateContextSync(options MLContextOptions) (ret MLContext) {
	bindings.CallMLCreateContextSync(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryCreateContextSync calls the method "ML.createContextSync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ML) TryCreateContextSync(options MLContextOptions) (ret MLContext, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCreateContextSync(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncCreateContextSync1 returns true if the method "ML.createContextSync" exists.
func (this ML) HasFuncCreateContextSync1() bool {
	return js.True == bindings.HasFuncMLCreateContextSync1(
		this.ref,
	)
}

// FuncCreateContextSync1 returns the method "ML.createContextSync".
func (this ML) FuncCreateContextSync1() (fn js.Func[func() MLContext]) {
	bindings.FuncMLCreateContextSync1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateContextSync1 calls the method "ML.createContextSync".
func (this ML) CreateContextSync1() (ret MLContext) {
	bindings.CallMLCreateContextSync1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateContextSync1 calls the method "ML.createContextSync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ML) TryCreateContextSync1() (ret MLContext, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCreateContextSync1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateContextSync2 returns true if the method "ML.createContextSync" exists.
func (this ML) HasFuncCreateContextSync2() bool {
	return js.True == bindings.HasFuncMLCreateContextSync2(
		this.ref,
	)
}

// FuncCreateContextSync2 returns the method "ML.createContextSync".
func (this ML) FuncCreateContextSync2() (fn js.Func[func(gpuDevice GPUDevice) MLContext]) {
	bindings.FuncMLCreateContextSync2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateContextSync2 calls the method "ML.createContextSync".
func (this ML) CreateContextSync2(gpuDevice GPUDevice) (ret MLContext) {
	bindings.CallMLCreateContextSync2(
		this.ref, js.Pointer(&ret),
		gpuDevice.Ref(),
	)

	return
}

// TryCreateContextSync2 calls the method "ML.createContextSync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ML) TryCreateContextSync2(gpuDevice GPUDevice) (ret MLContext, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLCreateContextSync2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Type returns the value of property "NetworkInformation.type".
//
// It returns ok=false if there is no such property.
func (this NetworkInformation) Type() (ret ConnectionType, ok bool) {
	ok = js.True == bindings.GetNetworkInformationType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EffectiveType returns the value of property "NetworkInformation.effectiveType".
//
// It returns ok=false if there is no such property.
func (this NetworkInformation) EffectiveType() (ret EffectiveConnectionType, ok bool) {
	ok = js.True == bindings.GetNetworkInformationEffectiveType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DownlinkMax returns the value of property "NetworkInformation.downlinkMax".
//
// It returns ok=false if there is no such property.
func (this NetworkInformation) DownlinkMax() (ret Megabit, ok bool) {
	ok = js.True == bindings.GetNetworkInformationDownlinkMax(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Downlink returns the value of property "NetworkInformation.downlink".
//
// It returns ok=false if there is no such property.
func (this NetworkInformation) Downlink() (ret Megabit, ok bool) {
	ok = js.True == bindings.GetNetworkInformationDownlink(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Rtt returns the value of property "NetworkInformation.rtt".
//
// It returns ok=false if there is no such property.
func (this NetworkInformation) Rtt() (ret Millisecond, ok bool) {
	ok = js.True == bindings.GetNetworkInformationRtt(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SaveData returns the value of property "NetworkInformation.saveData".
//
// It returns ok=false if there is no such property.
func (this NetworkInformation) SaveData() (ret bool, ok bool) {
	ok = js.True == bindings.GetNetworkInformationSaveData(
		this.ref, js.Pointer(&ret),
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
func (p *GPUQueueDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUQueueDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUQueueDescriptor) Update(ref js.Ref) {
	bindings.GPUQueueDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUQueueDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Label.Ref(),
	)
	p.Label = p.Label.FromRef(js.Undefined)
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
func (p *GPUDeviceDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUDeviceDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUDeviceDescriptor) Update(ref js.Ref) {
	bindings.GPUDeviceDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUDeviceDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.RequiredFeatures.Ref(),
		p.RequiredLimits.Ref(),
		p.Label.Ref(),
	)
	p.RequiredFeatures = p.RequiredFeatures.FromRef(js.Undefined)
	p.RequiredLimits = p.RequiredLimits.FromRef(js.Undefined)
	p.Label = p.Label.FromRef(js.Undefined)
	if recursive {
		p.DefaultQueue.FreeMembers(true)
	}
}

type GPUAdapterInfo struct {
	ref js.Ref
}

func (this GPUAdapterInfo) Once() GPUAdapterInfo {
	this.ref.Once()
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
	this.ref.Free()
}

// Vendor returns the value of property "GPUAdapterInfo.vendor".
//
// It returns ok=false if there is no such property.
func (this GPUAdapterInfo) Vendor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUAdapterInfoVendor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Architecture returns the value of property "GPUAdapterInfo.architecture".
//
// It returns ok=false if there is no such property.
func (this GPUAdapterInfo) Architecture() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUAdapterInfoArchitecture(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Device returns the value of property "GPUAdapterInfo.device".
//
// It returns ok=false if there is no such property.
func (this GPUAdapterInfo) Device() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUAdapterInfoDevice(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Description returns the value of property "GPUAdapterInfo.description".
//
// It returns ok=false if there is no such property.
func (this GPUAdapterInfo) Description() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUAdapterInfoDescription(
		this.ref, js.Pointer(&ret),
	)
	return
}

type GPUAdapter struct {
	ref js.Ref
}

func (this GPUAdapter) Once() GPUAdapter {
	this.ref.Once()
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
	this.ref.Free()
}

// Features returns the value of property "GPUAdapter.features".
//
// It returns ok=false if there is no such property.
func (this GPUAdapter) Features() (ret GPUSupportedFeatures, ok bool) {
	ok = js.True == bindings.GetGPUAdapterFeatures(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Limits returns the value of property "GPUAdapter.limits".
//
// It returns ok=false if there is no such property.
func (this GPUAdapter) Limits() (ret GPUSupportedLimits, ok bool) {
	ok = js.True == bindings.GetGPUAdapterLimits(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsFallbackAdapter returns the value of property "GPUAdapter.isFallbackAdapter".
//
// It returns ok=false if there is no such property.
func (this GPUAdapter) IsFallbackAdapter() (ret bool, ok bool) {
	ok = js.True == bindings.GetGPUAdapterIsFallbackAdapter(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncRequestDevice returns true if the method "GPUAdapter.requestDevice" exists.
func (this GPUAdapter) HasFuncRequestDevice() bool {
	return js.True == bindings.HasFuncGPUAdapterRequestDevice(
		this.ref,
	)
}

// FuncRequestDevice returns the method "GPUAdapter.requestDevice".
func (this GPUAdapter) FuncRequestDevice() (fn js.Func[func(descriptor GPUDeviceDescriptor) js.Promise[GPUDevice]]) {
	bindings.FuncGPUAdapterRequestDevice(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestDevice calls the method "GPUAdapter.requestDevice".
func (this GPUAdapter) RequestDevice(descriptor GPUDeviceDescriptor) (ret js.Promise[GPUDevice]) {
	bindings.CallGPUAdapterRequestDevice(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryRequestDevice calls the method "GPUAdapter.requestDevice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUAdapter) TryRequestDevice(descriptor GPUDeviceDescriptor) (ret js.Promise[GPUDevice], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUAdapterRequestDevice(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncRequestDevice1 returns true if the method "GPUAdapter.requestDevice" exists.
func (this GPUAdapter) HasFuncRequestDevice1() bool {
	return js.True == bindings.HasFuncGPUAdapterRequestDevice1(
		this.ref,
	)
}

// FuncRequestDevice1 returns the method "GPUAdapter.requestDevice".
func (this GPUAdapter) FuncRequestDevice1() (fn js.Func[func() js.Promise[GPUDevice]]) {
	bindings.FuncGPUAdapterRequestDevice1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestDevice1 calls the method "GPUAdapter.requestDevice".
func (this GPUAdapter) RequestDevice1() (ret js.Promise[GPUDevice]) {
	bindings.CallGPUAdapterRequestDevice1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestDevice1 calls the method "GPUAdapter.requestDevice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUAdapter) TryRequestDevice1() (ret js.Promise[GPUDevice], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUAdapterRequestDevice1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestAdapterInfo returns true if the method "GPUAdapter.requestAdapterInfo" exists.
func (this GPUAdapter) HasFuncRequestAdapterInfo() bool {
	return js.True == bindings.HasFuncGPUAdapterRequestAdapterInfo(
		this.ref,
	)
}

// FuncRequestAdapterInfo returns the method "GPUAdapter.requestAdapterInfo".
func (this GPUAdapter) FuncRequestAdapterInfo() (fn js.Func[func(unmaskHints js.Array[js.String]) js.Promise[GPUAdapterInfo]]) {
	bindings.FuncGPUAdapterRequestAdapterInfo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestAdapterInfo calls the method "GPUAdapter.requestAdapterInfo".
func (this GPUAdapter) RequestAdapterInfo(unmaskHints js.Array[js.String]) (ret js.Promise[GPUAdapterInfo]) {
	bindings.CallGPUAdapterRequestAdapterInfo(
		this.ref, js.Pointer(&ret),
		unmaskHints.Ref(),
	)

	return
}

// TryRequestAdapterInfo calls the method "GPUAdapter.requestAdapterInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUAdapter) TryRequestAdapterInfo(unmaskHints js.Array[js.String]) (ret js.Promise[GPUAdapterInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUAdapterRequestAdapterInfo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		unmaskHints.Ref(),
	)

	return
}

// HasFuncRequestAdapterInfo1 returns true if the method "GPUAdapter.requestAdapterInfo" exists.
func (this GPUAdapter) HasFuncRequestAdapterInfo1() bool {
	return js.True == bindings.HasFuncGPUAdapterRequestAdapterInfo1(
		this.ref,
	)
}

// FuncRequestAdapterInfo1 returns the method "GPUAdapter.requestAdapterInfo".
func (this GPUAdapter) FuncRequestAdapterInfo1() (fn js.Func[func() js.Promise[GPUAdapterInfo]]) {
	bindings.FuncGPUAdapterRequestAdapterInfo1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestAdapterInfo1 calls the method "GPUAdapter.requestAdapterInfo".
func (this GPUAdapter) RequestAdapterInfo1() (ret js.Promise[GPUAdapterInfo]) {
	bindings.CallGPUAdapterRequestAdapterInfo1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestAdapterInfo1 calls the method "GPUAdapter.requestAdapterInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUAdapter) TryRequestAdapterInfo1() (ret js.Promise[GPUAdapterInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUAdapterRequestAdapterInfo1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *GPURequestAdapterOptions) UpdateFrom(ref js.Ref) {
	bindings.GPURequestAdapterOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPURequestAdapterOptions) Update(ref js.Ref) {
	bindings.GPURequestAdapterOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPURequestAdapterOptions) FreeMembers(recursive bool) {
}

type WGSLLanguageFeatures struct {
	ref js.Ref
}

func (this WGSLLanguageFeatures) Once() WGSLLanguageFeatures {
	this.ref.Once()
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
	this.ref.Free()
}

type GPU struct {
	ref js.Ref
}

func (this GPU) Once() GPU {
	this.ref.Once()
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
	this.ref.Free()
}

// WgslLanguageFeatures returns the value of property "GPU.wgslLanguageFeatures".
//
// It returns ok=false if there is no such property.
func (this GPU) WgslLanguageFeatures() (ret WGSLLanguageFeatures, ok bool) {
	ok = js.True == bindings.GetGPUWgslLanguageFeatures(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncRequestAdapter returns true if the method "GPU.requestAdapter" exists.
func (this GPU) HasFuncRequestAdapter() bool {
	return js.True == bindings.HasFuncGPURequestAdapter(
		this.ref,
	)
}

// FuncRequestAdapter returns the method "GPU.requestAdapter".
func (this GPU) FuncRequestAdapter() (fn js.Func[func(options GPURequestAdapterOptions) js.Promise[GPUAdapter]]) {
	bindings.FuncGPURequestAdapter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestAdapter calls the method "GPU.requestAdapter".
func (this GPU) RequestAdapter(options GPURequestAdapterOptions) (ret js.Promise[GPUAdapter]) {
	bindings.CallGPURequestAdapter(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRequestAdapter calls the method "GPU.requestAdapter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPU) TryRequestAdapter(options GPURequestAdapterOptions) (ret js.Promise[GPUAdapter], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURequestAdapter(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRequestAdapter1 returns true if the method "GPU.requestAdapter" exists.
func (this GPU) HasFuncRequestAdapter1() bool {
	return js.True == bindings.HasFuncGPURequestAdapter1(
		this.ref,
	)
}

// FuncRequestAdapter1 returns the method "GPU.requestAdapter".
func (this GPU) FuncRequestAdapter1() (fn js.Func[func() js.Promise[GPUAdapter]]) {
	bindings.FuncGPURequestAdapter1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestAdapter1 calls the method "GPU.requestAdapter".
func (this GPU) RequestAdapter1() (ret js.Promise[GPUAdapter]) {
	bindings.CallGPURequestAdapter1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestAdapter1 calls the method "GPU.requestAdapter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPU) TryRequestAdapter1() (ret js.Promise[GPUAdapter], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURequestAdapter1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetPreferredCanvasFormat returns true if the method "GPU.getPreferredCanvasFormat" exists.
func (this GPU) HasFuncGetPreferredCanvasFormat() bool {
	return js.True == bindings.HasFuncGPUGetPreferredCanvasFormat(
		this.ref,
	)
}

// FuncGetPreferredCanvasFormat returns the method "GPU.getPreferredCanvasFormat".
func (this GPU) FuncGetPreferredCanvasFormat() (fn js.Func[func() GPUTextureFormat]) {
	bindings.FuncGPUGetPreferredCanvasFormat(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetPreferredCanvasFormat calls the method "GPU.getPreferredCanvasFormat".
func (this GPU) GetPreferredCanvasFormat() (ret GPUTextureFormat) {
	bindings.CallGPUGetPreferredCanvasFormat(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetPreferredCanvasFormat calls the method "GPU.getPreferredCanvasFormat"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPU) TryGetPreferredCanvasFormat() (ret GPUTextureFormat, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUGetPreferredCanvasFormat(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type Navigator struct {
	ref js.Ref
}

func (this Navigator) Once() Navigator {
	this.ref.Once()
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
	this.ref.Free()
}

// Hid returns the value of property "Navigator.hid".
//
// It returns ok=false if there is no such property.
func (this Navigator) Hid() (ret HID, ok bool) {
	ok = js.True == bindings.GetNavigatorHid(
		this.ref, js.Pointer(&ret),
	)
	return
}

// WindowControlsOverlay returns the value of property "Navigator.windowControlsOverlay".
//
// It returns ok=false if there is no such property.
func (this Navigator) WindowControlsOverlay() (ret WindowControlsOverlay, ok bool) {
	ok = js.True == bindings.GetNavigatorWindowControlsOverlay(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Credentials returns the value of property "Navigator.credentials".
//
// It returns ok=false if there is no such property.
func (this Navigator) Credentials() (ret CredentialsContainer, ok bool) {
	ok = js.True == bindings.GetNavigatorCredentials(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Clipboard returns the value of property "Navigator.clipboard".
//
// It returns ok=false if there is no such property.
func (this Navigator) Clipboard() (ret Clipboard, ok bool) {
	ok = js.True == bindings.GetNavigatorClipboard(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Geolocation returns the value of property "Navigator.geolocation".
//
// It returns ok=false if there is no such property.
func (this Navigator) Geolocation() (ret Geolocation, ok bool) {
	ok = js.True == bindings.GetNavigatorGeolocation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Usb returns the value of property "Navigator.usb".
//
// It returns ok=false if there is no such property.
func (this Navigator) Usb() (ret USB, ok bool) {
	ok = js.True == bindings.GetNavigatorUsb(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EpubReadingSystem returns the value of property "Navigator.epubReadingSystem".
//
// It returns ok=false if there is no such property.
func (this Navigator) EpubReadingSystem() (ret EpubReadingSystem, ok bool) {
	ok = js.True == bindings.GetNavigatorEpubReadingSystem(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Xr returns the value of property "Navigator.xr".
//
// It returns ok=false if there is no such property.
func (this Navigator) Xr() (ret XRSystem, ok bool) {
	ok = js.True == bindings.GetNavigatorXr(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ServiceWorker returns the value of property "Navigator.serviceWorker".
//
// It returns ok=false if there is no such property.
func (this Navigator) ServiceWorker() (ret ServiceWorkerContainer, ok bool) {
	ok = js.True == bindings.GetNavigatorServiceWorker(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MediaDevices returns the value of property "Navigator.mediaDevices".
//
// It returns ok=false if there is no such property.
func (this Navigator) MediaDevices() (ret MediaDevices, ok bool) {
	ok = js.True == bindings.GetNavigatorMediaDevices(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Bluetooth returns the value of property "Navigator.bluetooth".
//
// It returns ok=false if there is no such property.
func (this Navigator) Bluetooth() (ret Bluetooth, ok bool) {
	ok = js.True == bindings.GetNavigatorBluetooth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Serial returns the value of property "Navigator.serial".
//
// It returns ok=false if there is no such property.
func (this Navigator) Serial() (ret Serial, ok bool) {
	ok = js.True == bindings.GetNavigatorSerial(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MediaCapabilities returns the value of property "Navigator.mediaCapabilities".
//
// It returns ok=false if there is no such property.
func (this Navigator) MediaCapabilities() (ret MediaCapabilities, ok bool) {
	ok = js.True == bindings.GetNavigatorMediaCapabilities(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UserActivation returns the value of property "Navigator.userActivation".
//
// It returns ok=false if there is no such property.
func (this Navigator) UserActivation() (ret UserActivation, ok bool) {
	ok = js.True == bindings.GetNavigatorUserActivation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Permissions returns the value of property "Navigator.permissions".
//
// It returns ok=false if there is no such property.
func (this Navigator) Permissions() (ret Permissions, ok bool) {
	ok = js.True == bindings.GetNavigatorPermissions(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Contacts returns the value of property "Navigator.contacts".
//
// It returns ok=false if there is no such property.
func (this Navigator) Contacts() (ret ContactsManager, ok bool) {
	ok = js.True == bindings.GetNavigatorContacts(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Keyboard returns the value of property "Navigator.keyboard".
//
// It returns ok=false if there is no such property.
func (this Navigator) Keyboard() (ret Keyboard, ok bool) {
	ok = js.True == bindings.GetNavigatorKeyboard(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MediaSession returns the value of property "Navigator.mediaSession".
//
// It returns ok=false if there is no such property.
func (this Navigator) MediaSession() (ret MediaSession, ok bool) {
	ok = js.True == bindings.GetNavigatorMediaSession(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DevicePosture returns the value of property "Navigator.devicePosture".
//
// It returns ok=false if there is no such property.
func (this Navigator) DevicePosture() (ret DevicePosture, ok bool) {
	ok = js.True == bindings.GetNavigatorDevicePosture(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxTouchPoints returns the value of property "Navigator.maxTouchPoints".
//
// It returns ok=false if there is no such property.
func (this Navigator) MaxTouchPoints() (ret int32, ok bool) {
	ok = js.True == bindings.GetNavigatorMaxTouchPoints(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Scheduling returns the value of property "Navigator.scheduling".
//
// It returns ok=false if there is no such property.
func (this Navigator) Scheduling() (ret Scheduling, ok bool) {
	ok = js.True == bindings.GetNavigatorScheduling(
		this.ref, js.Pointer(&ret),
	)
	return
}

// WakeLock returns the value of property "Navigator.wakeLock".
//
// It returns ok=false if there is no such property.
func (this Navigator) WakeLock() (ret WakeLock, ok bool) {
	ok = js.True == bindings.GetNavigatorWakeLock(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Ink returns the value of property "Navigator.ink".
//
// It returns ok=false if there is no such property.
func (this Navigator) Ink() (ret Ink, ok bool) {
	ok = js.True == bindings.GetNavigatorInk(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Presentation returns the value of property "Navigator.presentation".
//
// It returns ok=false if there is no such property.
func (this Navigator) Presentation() (ret Presentation, ok bool) {
	ok = js.True == bindings.GetNavigatorPresentation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// VirtualKeyboard returns the value of property "Navigator.virtualKeyboard".
//
// It returns ok=false if there is no such property.
func (this Navigator) VirtualKeyboard() (ret VirtualKeyboard, ok bool) {
	ok = js.True == bindings.GetNavigatorVirtualKeyboard(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UserAgentData returns the value of property "Navigator.userAgentData".
//
// It returns ok=false if there is no such property.
func (this Navigator) UserAgentData() (ret NavigatorUAData, ok bool) {
	ok = js.True == bindings.GetNavigatorUserAgentData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HardwareConcurrency returns the value of property "Navigator.hardwareConcurrency".
//
// It returns ok=false if there is no such property.
func (this Navigator) HardwareConcurrency() (ret uint64, ok bool) {
	ok = js.True == bindings.GetNavigatorHardwareConcurrency(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DeviceMemory returns the value of property "Navigator.deviceMemory".
//
// It returns ok=false if there is no such property.
func (this Navigator) DeviceMemory() (ret float64, ok bool) {
	ok = js.True == bindings.GetNavigatorDeviceMemory(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Storage returns the value of property "Navigator.storage".
//
// It returns ok=false if there is no such property.
func (this Navigator) Storage() (ret StorageManager, ok bool) {
	ok = js.True == bindings.GetNavigatorStorage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StorageBuckets returns the value of property "Navigator.storageBuckets".
//
// It returns ok=false if there is no such property.
func (this Navigator) StorageBuckets() (ret StorageBucketManager, ok bool) {
	ok = js.True == bindings.GetNavigatorStorageBuckets(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Locks returns the value of property "Navigator.locks".
//
// It returns ok=false if there is no such property.
func (this Navigator) Locks() (ret LockManager, ok bool) {
	ok = js.True == bindings.GetNavigatorLocks(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AppCodeName returns the value of property "Navigator.appCodeName".
//
// It returns ok=false if there is no such property.
func (this Navigator) AppCodeName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorAppCodeName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AppName returns the value of property "Navigator.appName".
//
// It returns ok=false if there is no such property.
func (this Navigator) AppName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorAppName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AppVersion returns the value of property "Navigator.appVersion".
//
// It returns ok=false if there is no such property.
func (this Navigator) AppVersion() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorAppVersion(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Platform returns the value of property "Navigator.platform".
//
// It returns ok=false if there is no such property.
func (this Navigator) Platform() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorPlatform(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Product returns the value of property "Navigator.product".
//
// It returns ok=false if there is no such property.
func (this Navigator) Product() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorProduct(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ProductSub returns the value of property "Navigator.productSub".
//
// It returns ok=false if there is no such property.
func (this Navigator) ProductSub() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorProductSub(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UserAgent returns the value of property "Navigator.userAgent".
//
// It returns ok=false if there is no such property.
func (this Navigator) UserAgent() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorUserAgent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Vendor returns the value of property "Navigator.vendor".
//
// It returns ok=false if there is no such property.
func (this Navigator) Vendor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorVendor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// VendorSub returns the value of property "Navigator.vendorSub".
//
// It returns ok=false if there is no such property.
func (this Navigator) VendorSub() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorVendorSub(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Oscpu returns the value of property "Navigator.oscpu".
//
// It returns ok=false if there is no such property.
func (this Navigator) Oscpu() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorOscpu(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Language returns the value of property "Navigator.language".
//
// It returns ok=false if there is no such property.
func (this Navigator) Language() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigatorLanguage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Languages returns the value of property "Navigator.languages".
//
// It returns ok=false if there is no such property.
func (this Navigator) Languages() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetNavigatorLanguages(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OnLine returns the value of property "Navigator.onLine".
//
// It returns ok=false if there is no such property.
func (this Navigator) OnLine() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigatorOnLine(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CookieEnabled returns the value of property "Navigator.cookieEnabled".
//
// It returns ok=false if there is no such property.
func (this Navigator) CookieEnabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigatorCookieEnabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Plugins returns the value of property "Navigator.plugins".
//
// It returns ok=false if there is no such property.
func (this Navigator) Plugins() (ret PluginArray, ok bool) {
	ok = js.True == bindings.GetNavigatorPlugins(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MimeTypes returns the value of property "Navigator.mimeTypes".
//
// It returns ok=false if there is no such property.
func (this Navigator) MimeTypes() (ret MimeTypeArray, ok bool) {
	ok = js.True == bindings.GetNavigatorMimeTypes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PdfViewerEnabled returns the value of property "Navigator.pdfViewerEnabled".
//
// It returns ok=false if there is no such property.
func (this Navigator) PdfViewerEnabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigatorPdfViewerEnabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Webdriver returns the value of property "Navigator.webdriver".
//
// It returns ok=false if there is no such property.
func (this Navigator) Webdriver() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigatorWebdriver(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Ml returns the value of property "Navigator.ml".
//
// It returns ok=false if there is no such property.
func (this Navigator) Ml() (ret ML, ok bool) {
	ok = js.True == bindings.GetNavigatorMl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Connection returns the value of property "Navigator.connection".
//
// It returns ok=false if there is no such property.
func (this Navigator) Connection() (ret NetworkInformation, ok bool) {
	ok = js.True == bindings.GetNavigatorConnection(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Gpu returns the value of property "Navigator.gpu".
//
// It returns ok=false if there is no such property.
func (this Navigator) Gpu() (ret GPU, ok bool) {
	ok = js.True == bindings.GetNavigatorGpu(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncUpdateAdInterestGroups returns true if the method "Navigator.updateAdInterestGroups" exists.
func (this Navigator) HasFuncUpdateAdInterestGroups() bool {
	return js.True == bindings.HasFuncNavigatorUpdateAdInterestGroups(
		this.ref,
	)
}

// FuncUpdateAdInterestGroups returns the method "Navigator.updateAdInterestGroups".
func (this Navigator) FuncUpdateAdInterestGroups() (fn js.Func[func()]) {
	bindings.FuncNavigatorUpdateAdInterestGroups(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpdateAdInterestGroups calls the method "Navigator.updateAdInterestGroups".
func (this Navigator) UpdateAdInterestGroups() (ret js.Void) {
	bindings.CallNavigatorUpdateAdInterestGroups(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryUpdateAdInterestGroups calls the method "Navigator.updateAdInterestGroups"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryUpdateAdInterestGroups() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorUpdateAdInterestGroups(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSendBeacon returns true if the method "Navigator.sendBeacon" exists.
func (this Navigator) HasFuncSendBeacon() bool {
	return js.True == bindings.HasFuncNavigatorSendBeacon(
		this.ref,
	)
}

// FuncSendBeacon returns the method "Navigator.sendBeacon".
func (this Navigator) FuncSendBeacon() (fn js.Func[func(url js.String, data BodyInit) bool]) {
	bindings.FuncNavigatorSendBeacon(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SendBeacon calls the method "Navigator.sendBeacon".
func (this Navigator) SendBeacon(url js.String, data BodyInit) (ret bool) {
	bindings.CallNavigatorSendBeacon(
		this.ref, js.Pointer(&ret),
		url.Ref(),
		data.Ref(),
	)

	return
}

// TrySendBeacon calls the method "Navigator.sendBeacon"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TrySendBeacon(url js.String, data BodyInit) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorSendBeacon(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		data.Ref(),
	)

	return
}

// HasFuncSendBeacon1 returns true if the method "Navigator.sendBeacon" exists.
func (this Navigator) HasFuncSendBeacon1() bool {
	return js.True == bindings.HasFuncNavigatorSendBeacon1(
		this.ref,
	)
}

// FuncSendBeacon1 returns the method "Navigator.sendBeacon".
func (this Navigator) FuncSendBeacon1() (fn js.Func[func(url js.String) bool]) {
	bindings.FuncNavigatorSendBeacon1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SendBeacon1 calls the method "Navigator.sendBeacon".
func (this Navigator) SendBeacon1(url js.String) (ret bool) {
	bindings.CallNavigatorSendBeacon1(
		this.ref, js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TrySendBeacon1 calls the method "Navigator.sendBeacon"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TrySendBeacon1(url js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorSendBeacon1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncGetBattery returns true if the method "Navigator.getBattery" exists.
func (this Navigator) HasFuncGetBattery() bool {
	return js.True == bindings.HasFuncNavigatorGetBattery(
		this.ref,
	)
}

// FuncGetBattery returns the method "Navigator.getBattery".
func (this Navigator) FuncGetBattery() (fn js.Func[func() js.Promise[BatteryManager]]) {
	bindings.FuncNavigatorGetBattery(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetBattery calls the method "Navigator.getBattery".
func (this Navigator) GetBattery() (ret js.Promise[BatteryManager]) {
	bindings.CallNavigatorGetBattery(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetBattery calls the method "Navigator.getBattery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryGetBattery() (ret js.Promise[BatteryManager], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorGetBattery(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncVibrate returns true if the method "Navigator.vibrate" exists.
func (this Navigator) HasFuncVibrate() bool {
	return js.True == bindings.HasFuncNavigatorVibrate(
		this.ref,
	)
}

// FuncVibrate returns the method "Navigator.vibrate".
func (this Navigator) FuncVibrate() (fn js.Func[func(pattern VibratePattern) bool]) {
	bindings.FuncNavigatorVibrate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Vibrate calls the method "Navigator.vibrate".
func (this Navigator) Vibrate(pattern VibratePattern) (ret bool) {
	bindings.CallNavigatorVibrate(
		this.ref, js.Pointer(&ret),
		pattern.Ref(),
	)

	return
}

// TryVibrate calls the method "Navigator.vibrate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryVibrate(pattern VibratePattern) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorVibrate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		pattern.Ref(),
	)

	return
}

// HasFuncGetGamepads returns true if the method "Navigator.getGamepads" exists.
func (this Navigator) HasFuncGetGamepads() bool {
	return js.True == bindings.HasFuncNavigatorGetGamepads(
		this.ref,
	)
}

// FuncGetGamepads returns the method "Navigator.getGamepads".
func (this Navigator) FuncGetGamepads() (fn js.Func[func() js.Array[Gamepad]]) {
	bindings.FuncNavigatorGetGamepads(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetGamepads calls the method "Navigator.getGamepads".
func (this Navigator) GetGamepads() (ret js.Array[Gamepad]) {
	bindings.CallNavigatorGetGamepads(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetGamepads calls the method "Navigator.getGamepads"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryGetGamepads() (ret js.Array[Gamepad], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorGetGamepads(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetInstalledRelatedApps returns true if the method "Navigator.getInstalledRelatedApps" exists.
func (this Navigator) HasFuncGetInstalledRelatedApps() bool {
	return js.True == bindings.HasFuncNavigatorGetInstalledRelatedApps(
		this.ref,
	)
}

// FuncGetInstalledRelatedApps returns the method "Navigator.getInstalledRelatedApps".
func (this Navigator) FuncGetInstalledRelatedApps() (fn js.Func[func() js.Promise[js.Array[RelatedApplication]]]) {
	bindings.FuncNavigatorGetInstalledRelatedApps(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetInstalledRelatedApps calls the method "Navigator.getInstalledRelatedApps".
func (this Navigator) GetInstalledRelatedApps() (ret js.Promise[js.Array[RelatedApplication]]) {
	bindings.CallNavigatorGetInstalledRelatedApps(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetInstalledRelatedApps calls the method "Navigator.getInstalledRelatedApps"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryGetInstalledRelatedApps() (ret js.Promise[js.Array[RelatedApplication]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorGetInstalledRelatedApps(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestMediaKeySystemAccess returns true if the method "Navigator.requestMediaKeySystemAccess" exists.
func (this Navigator) HasFuncRequestMediaKeySystemAccess() bool {
	return js.True == bindings.HasFuncNavigatorRequestMediaKeySystemAccess(
		this.ref,
	)
}

// FuncRequestMediaKeySystemAccess returns the method "Navigator.requestMediaKeySystemAccess".
func (this Navigator) FuncRequestMediaKeySystemAccess() (fn js.Func[func(keySystem js.String, supportedConfigurations js.Array[MediaKeySystemConfiguration]) js.Promise[MediaKeySystemAccess]]) {
	bindings.FuncNavigatorRequestMediaKeySystemAccess(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestMediaKeySystemAccess calls the method "Navigator.requestMediaKeySystemAccess".
func (this Navigator) RequestMediaKeySystemAccess(keySystem js.String, supportedConfigurations js.Array[MediaKeySystemConfiguration]) (ret js.Promise[MediaKeySystemAccess]) {
	bindings.CallNavigatorRequestMediaKeySystemAccess(
		this.ref, js.Pointer(&ret),
		keySystem.Ref(),
		supportedConfigurations.Ref(),
	)

	return
}

// TryRequestMediaKeySystemAccess calls the method "Navigator.requestMediaKeySystemAccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryRequestMediaKeySystemAccess(keySystem js.String, supportedConfigurations js.Array[MediaKeySystemConfiguration]) (ret js.Promise[MediaKeySystemAccess], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorRequestMediaKeySystemAccess(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		keySystem.Ref(),
		supportedConfigurations.Ref(),
	)

	return
}

// HasFuncJoinAdInterestGroup returns true if the method "Navigator.joinAdInterestGroup" exists.
func (this Navigator) HasFuncJoinAdInterestGroup() bool {
	return js.True == bindings.HasFuncNavigatorJoinAdInterestGroup(
		this.ref,
	)
}

// FuncJoinAdInterestGroup returns the method "Navigator.joinAdInterestGroup".
func (this Navigator) FuncJoinAdInterestGroup() (fn js.Func[func(group AuctionAdInterestGroup) js.Promise[js.Void]]) {
	bindings.FuncNavigatorJoinAdInterestGroup(
		this.ref, js.Pointer(&fn),
	)
	return
}

// JoinAdInterestGroup calls the method "Navigator.joinAdInterestGroup".
func (this Navigator) JoinAdInterestGroup(group AuctionAdInterestGroup) (ret js.Promise[js.Void]) {
	bindings.CallNavigatorJoinAdInterestGroup(
		this.ref, js.Pointer(&ret),
		js.Pointer(&group),
	)

	return
}

// TryJoinAdInterestGroup calls the method "Navigator.joinAdInterestGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryJoinAdInterestGroup(group AuctionAdInterestGroup) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorJoinAdInterestGroup(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&group),
	)

	return
}

// HasFuncGetUserMedia returns true if the method "Navigator.getUserMedia" exists.
func (this Navigator) HasFuncGetUserMedia() bool {
	return js.True == bindings.HasFuncNavigatorGetUserMedia(
		this.ref,
	)
}

// FuncGetUserMedia returns the method "Navigator.getUserMedia".
func (this Navigator) FuncGetUserMedia() (fn js.Func[func(constraints MediaStreamConstraints, successCallback js.Func[func(stream MediaStream)], errorCallback js.Func[func(err DOMException)])]) {
	bindings.FuncNavigatorGetUserMedia(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetUserMedia calls the method "Navigator.getUserMedia".
func (this Navigator) GetUserMedia(constraints MediaStreamConstraints, successCallback js.Func[func(stream MediaStream)], errorCallback js.Func[func(err DOMException)]) (ret js.Void) {
	bindings.CallNavigatorGetUserMedia(
		this.ref, js.Pointer(&ret),
		js.Pointer(&constraints),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// TryGetUserMedia calls the method "Navigator.getUserMedia"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryGetUserMedia(constraints MediaStreamConstraints, successCallback js.Func[func(stream MediaStream)], errorCallback js.Func[func(err DOMException)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorGetUserMedia(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&constraints),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// HasFuncLeaveAdInterestGroup returns true if the method "Navigator.leaveAdInterestGroup" exists.
func (this Navigator) HasFuncLeaveAdInterestGroup() bool {
	return js.True == bindings.HasFuncNavigatorLeaveAdInterestGroup(
		this.ref,
	)
}

// FuncLeaveAdInterestGroup returns the method "Navigator.leaveAdInterestGroup".
func (this Navigator) FuncLeaveAdInterestGroup() (fn js.Func[func(group AuctionAdInterestGroupKey) js.Promise[js.Void]]) {
	bindings.FuncNavigatorLeaveAdInterestGroup(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LeaveAdInterestGroup calls the method "Navigator.leaveAdInterestGroup".
func (this Navigator) LeaveAdInterestGroup(group AuctionAdInterestGroupKey) (ret js.Promise[js.Void]) {
	bindings.CallNavigatorLeaveAdInterestGroup(
		this.ref, js.Pointer(&ret),
		js.Pointer(&group),
	)

	return
}

// TryLeaveAdInterestGroup calls the method "Navigator.leaveAdInterestGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryLeaveAdInterestGroup(group AuctionAdInterestGroupKey) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorLeaveAdInterestGroup(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&group),
	)

	return
}

// HasFuncLeaveAdInterestGroup1 returns true if the method "Navigator.leaveAdInterestGroup" exists.
func (this Navigator) HasFuncLeaveAdInterestGroup1() bool {
	return js.True == bindings.HasFuncNavigatorLeaveAdInterestGroup1(
		this.ref,
	)
}

// FuncLeaveAdInterestGroup1 returns the method "Navigator.leaveAdInterestGroup".
func (this Navigator) FuncLeaveAdInterestGroup1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncNavigatorLeaveAdInterestGroup1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LeaveAdInterestGroup1 calls the method "Navigator.leaveAdInterestGroup".
func (this Navigator) LeaveAdInterestGroup1() (ret js.Promise[js.Void]) {
	bindings.CallNavigatorLeaveAdInterestGroup1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryLeaveAdInterestGroup1 calls the method "Navigator.leaveAdInterestGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryLeaveAdInterestGroup1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorLeaveAdInterestGroup1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRunAdAuction returns true if the method "Navigator.runAdAuction" exists.
func (this Navigator) HasFuncRunAdAuction() bool {
	return js.True == bindings.HasFuncNavigatorRunAdAuction(
		this.ref,
	)
}

// FuncRunAdAuction returns the method "Navigator.runAdAuction".
func (this Navigator) FuncRunAdAuction() (fn js.Func[func(config AuctionAdConfig) js.Promise[OneOf_String_FencedFrameConfig]]) {
	bindings.FuncNavigatorRunAdAuction(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RunAdAuction calls the method "Navigator.runAdAuction".
func (this Navigator) RunAdAuction(config AuctionAdConfig) (ret js.Promise[OneOf_String_FencedFrameConfig]) {
	bindings.CallNavigatorRunAdAuction(
		this.ref, js.Pointer(&ret),
		js.Pointer(&config),
	)

	return
}

// TryRunAdAuction calls the method "Navigator.runAdAuction"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryRunAdAuction(config AuctionAdConfig) (ret js.Promise[OneOf_String_FencedFrameConfig], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorRunAdAuction(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&config),
	)

	return
}

// HasFuncShare returns true if the method "Navigator.share" exists.
func (this Navigator) HasFuncShare() bool {
	return js.True == bindings.HasFuncNavigatorShare(
		this.ref,
	)
}

// FuncShare returns the method "Navigator.share".
func (this Navigator) FuncShare() (fn js.Func[func(data ShareData) js.Promise[js.Void]]) {
	bindings.FuncNavigatorShare(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Share calls the method "Navigator.share".
func (this Navigator) Share(data ShareData) (ret js.Promise[js.Void]) {
	bindings.CallNavigatorShare(
		this.ref, js.Pointer(&ret),
		js.Pointer(&data),
	)

	return
}

// TryShare calls the method "Navigator.share"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryShare(data ShareData) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorShare(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&data),
	)

	return
}

// HasFuncShare1 returns true if the method "Navigator.share" exists.
func (this Navigator) HasFuncShare1() bool {
	return js.True == bindings.HasFuncNavigatorShare1(
		this.ref,
	)
}

// FuncShare1 returns the method "Navigator.share".
func (this Navigator) FuncShare1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncNavigatorShare1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Share1 calls the method "Navigator.share".
func (this Navigator) Share1() (ret js.Promise[js.Void]) {
	bindings.CallNavigatorShare1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryShare1 calls the method "Navigator.share"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryShare1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorShare1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCanShare returns true if the method "Navigator.canShare" exists.
func (this Navigator) HasFuncCanShare() bool {
	return js.True == bindings.HasFuncNavigatorCanShare(
		this.ref,
	)
}

// FuncCanShare returns the method "Navigator.canShare".
func (this Navigator) FuncCanShare() (fn js.Func[func(data ShareData) bool]) {
	bindings.FuncNavigatorCanShare(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CanShare calls the method "Navigator.canShare".
func (this Navigator) CanShare(data ShareData) (ret bool) {
	bindings.CallNavigatorCanShare(
		this.ref, js.Pointer(&ret),
		js.Pointer(&data),
	)

	return
}

// TryCanShare calls the method "Navigator.canShare"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryCanShare(data ShareData) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorCanShare(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&data),
	)

	return
}

// HasFuncCanShare1 returns true if the method "Navigator.canShare" exists.
func (this Navigator) HasFuncCanShare1() bool {
	return js.True == bindings.HasFuncNavigatorCanShare1(
		this.ref,
	)
}

// FuncCanShare1 returns the method "Navigator.canShare".
func (this Navigator) FuncCanShare1() (fn js.Func[func() bool]) {
	bindings.FuncNavigatorCanShare1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CanShare1 calls the method "Navigator.canShare".
func (this Navigator) CanShare1() (ret bool) {
	bindings.CallNavigatorCanShare1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCanShare1 calls the method "Navigator.canShare"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryCanShare1() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorCanShare1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestMIDIAccess returns true if the method "Navigator.requestMIDIAccess" exists.
func (this Navigator) HasFuncRequestMIDIAccess() bool {
	return js.True == bindings.HasFuncNavigatorRequestMIDIAccess(
		this.ref,
	)
}

// FuncRequestMIDIAccess returns the method "Navigator.requestMIDIAccess".
func (this Navigator) FuncRequestMIDIAccess() (fn js.Func[func(options MIDIOptions) js.Promise[MIDIAccess]]) {
	bindings.FuncNavigatorRequestMIDIAccess(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestMIDIAccess calls the method "Navigator.requestMIDIAccess".
func (this Navigator) RequestMIDIAccess(options MIDIOptions) (ret js.Promise[MIDIAccess]) {
	bindings.CallNavigatorRequestMIDIAccess(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryRequestMIDIAccess calls the method "Navigator.requestMIDIAccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryRequestMIDIAccess(options MIDIOptions) (ret js.Promise[MIDIAccess], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorRequestMIDIAccess(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncRequestMIDIAccess1 returns true if the method "Navigator.requestMIDIAccess" exists.
func (this Navigator) HasFuncRequestMIDIAccess1() bool {
	return js.True == bindings.HasFuncNavigatorRequestMIDIAccess1(
		this.ref,
	)
}

// FuncRequestMIDIAccess1 returns the method "Navigator.requestMIDIAccess".
func (this Navigator) FuncRequestMIDIAccess1() (fn js.Func[func() js.Promise[MIDIAccess]]) {
	bindings.FuncNavigatorRequestMIDIAccess1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestMIDIAccess1 calls the method "Navigator.requestMIDIAccess".
func (this Navigator) RequestMIDIAccess1() (ret js.Promise[MIDIAccess]) {
	bindings.CallNavigatorRequestMIDIAccess1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestMIDIAccess1 calls the method "Navigator.requestMIDIAccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryRequestMIDIAccess1() (ret js.Promise[MIDIAccess], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorRequestMIDIAccess1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetAutoplayPolicy returns true if the method "Navigator.getAutoplayPolicy" exists.
func (this Navigator) HasFuncGetAutoplayPolicy() bool {
	return js.True == bindings.HasFuncNavigatorGetAutoplayPolicy(
		this.ref,
	)
}

// FuncGetAutoplayPolicy returns the method "Navigator.getAutoplayPolicy".
func (this Navigator) FuncGetAutoplayPolicy() (fn js.Func[func(typ AutoplayPolicyMediaType) AutoplayPolicy]) {
	bindings.FuncNavigatorGetAutoplayPolicy(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAutoplayPolicy calls the method "Navigator.getAutoplayPolicy".
func (this Navigator) GetAutoplayPolicy(typ AutoplayPolicyMediaType) (ret AutoplayPolicy) {
	bindings.CallNavigatorGetAutoplayPolicy(
		this.ref, js.Pointer(&ret),
		uint32(typ),
	)

	return
}

// TryGetAutoplayPolicy calls the method "Navigator.getAutoplayPolicy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryGetAutoplayPolicy(typ AutoplayPolicyMediaType) (ret AutoplayPolicy, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorGetAutoplayPolicy(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
	)

	return
}

// HasFuncGetAutoplayPolicy1 returns true if the method "Navigator.getAutoplayPolicy" exists.
func (this Navigator) HasFuncGetAutoplayPolicy1() bool {
	return js.True == bindings.HasFuncNavigatorGetAutoplayPolicy1(
		this.ref,
	)
}

// FuncGetAutoplayPolicy1 returns the method "Navigator.getAutoplayPolicy".
func (this Navigator) FuncGetAutoplayPolicy1() (fn js.Func[func(element HTMLMediaElement) AutoplayPolicy]) {
	bindings.FuncNavigatorGetAutoplayPolicy1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAutoplayPolicy1 calls the method "Navigator.getAutoplayPolicy".
func (this Navigator) GetAutoplayPolicy1(element HTMLMediaElement) (ret AutoplayPolicy) {
	bindings.CallNavigatorGetAutoplayPolicy1(
		this.ref, js.Pointer(&ret),
		element.Ref(),
	)

	return
}

// TryGetAutoplayPolicy1 calls the method "Navigator.getAutoplayPolicy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryGetAutoplayPolicy1(element HTMLMediaElement) (ret AutoplayPolicy, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorGetAutoplayPolicy1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
	)

	return
}

// HasFuncGetAutoplayPolicy2 returns true if the method "Navigator.getAutoplayPolicy" exists.
func (this Navigator) HasFuncGetAutoplayPolicy2() bool {
	return js.True == bindings.HasFuncNavigatorGetAutoplayPolicy2(
		this.ref,
	)
}

// FuncGetAutoplayPolicy2 returns the method "Navigator.getAutoplayPolicy".
func (this Navigator) FuncGetAutoplayPolicy2() (fn js.Func[func(context AudioContext) AutoplayPolicy]) {
	bindings.FuncNavigatorGetAutoplayPolicy2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAutoplayPolicy2 calls the method "Navigator.getAutoplayPolicy".
func (this Navigator) GetAutoplayPolicy2(context AudioContext) (ret AutoplayPolicy) {
	bindings.CallNavigatorGetAutoplayPolicy2(
		this.ref, js.Pointer(&ret),
		context.Ref(),
	)

	return
}

// TryGetAutoplayPolicy2 calls the method "Navigator.getAutoplayPolicy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryGetAutoplayPolicy2(context AudioContext) (ret AutoplayPolicy, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorGetAutoplayPolicy2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		context.Ref(),
	)

	return
}

// HasFuncTaintEnabled returns true if the method "Navigator.taintEnabled" exists.
func (this Navigator) HasFuncTaintEnabled() bool {
	return js.True == bindings.HasFuncNavigatorTaintEnabled(
		this.ref,
	)
}

// FuncTaintEnabled returns the method "Navigator.taintEnabled".
func (this Navigator) FuncTaintEnabled() (fn js.Func[func() bool]) {
	bindings.FuncNavigatorTaintEnabled(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TaintEnabled calls the method "Navigator.taintEnabled".
func (this Navigator) TaintEnabled() (ret bool) {
	bindings.CallNavigatorTaintEnabled(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryTaintEnabled calls the method "Navigator.taintEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryTaintEnabled() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorTaintEnabled(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRegisterProtocolHandler returns true if the method "Navigator.registerProtocolHandler" exists.
func (this Navigator) HasFuncRegisterProtocolHandler() bool {
	return js.True == bindings.HasFuncNavigatorRegisterProtocolHandler(
		this.ref,
	)
}

// FuncRegisterProtocolHandler returns the method "Navigator.registerProtocolHandler".
func (this Navigator) FuncRegisterProtocolHandler() (fn js.Func[func(scheme js.String, url js.String)]) {
	bindings.FuncNavigatorRegisterProtocolHandler(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RegisterProtocolHandler calls the method "Navigator.registerProtocolHandler".
func (this Navigator) RegisterProtocolHandler(scheme js.String, url js.String) (ret js.Void) {
	bindings.CallNavigatorRegisterProtocolHandler(
		this.ref, js.Pointer(&ret),
		scheme.Ref(),
		url.Ref(),
	)

	return
}

// TryRegisterProtocolHandler calls the method "Navigator.registerProtocolHandler"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryRegisterProtocolHandler(scheme js.String, url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorRegisterProtocolHandler(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		scheme.Ref(),
		url.Ref(),
	)

	return
}

// HasFuncUnregisterProtocolHandler returns true if the method "Navigator.unregisterProtocolHandler" exists.
func (this Navigator) HasFuncUnregisterProtocolHandler() bool {
	return js.True == bindings.HasFuncNavigatorUnregisterProtocolHandler(
		this.ref,
	)
}

// FuncUnregisterProtocolHandler returns the method "Navigator.unregisterProtocolHandler".
func (this Navigator) FuncUnregisterProtocolHandler() (fn js.Func[func(scheme js.String, url js.String)]) {
	bindings.FuncNavigatorUnregisterProtocolHandler(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UnregisterProtocolHandler calls the method "Navigator.unregisterProtocolHandler".
func (this Navigator) UnregisterProtocolHandler(scheme js.String, url js.String) (ret js.Void) {
	bindings.CallNavigatorUnregisterProtocolHandler(
		this.ref, js.Pointer(&ret),
		scheme.Ref(),
		url.Ref(),
	)

	return
}

// TryUnregisterProtocolHandler calls the method "Navigator.unregisterProtocolHandler"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryUnregisterProtocolHandler(scheme js.String, url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorUnregisterProtocolHandler(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		scheme.Ref(),
		url.Ref(),
	)

	return
}

// HasFuncJavaEnabled returns true if the method "Navigator.javaEnabled" exists.
func (this Navigator) HasFuncJavaEnabled() bool {
	return js.True == bindings.HasFuncNavigatorJavaEnabled(
		this.ref,
	)
}

// FuncJavaEnabled returns the method "Navigator.javaEnabled".
func (this Navigator) FuncJavaEnabled() (fn js.Func[func() bool]) {
	bindings.FuncNavigatorJavaEnabled(
		this.ref, js.Pointer(&fn),
	)
	return
}

// JavaEnabled calls the method "Navigator.javaEnabled".
func (this Navigator) JavaEnabled() (ret bool) {
	bindings.CallNavigatorJavaEnabled(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryJavaEnabled calls the method "Navigator.javaEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryJavaEnabled() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorJavaEnabled(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetAppBadge returns true if the method "Navigator.setAppBadge" exists.
func (this Navigator) HasFuncSetAppBadge() bool {
	return js.True == bindings.HasFuncNavigatorSetAppBadge(
		this.ref,
	)
}

// FuncSetAppBadge returns the method "Navigator.setAppBadge".
func (this Navigator) FuncSetAppBadge() (fn js.Func[func(contents uint64) js.Promise[js.Void]]) {
	bindings.FuncNavigatorSetAppBadge(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetAppBadge calls the method "Navigator.setAppBadge".
func (this Navigator) SetAppBadge(contents uint64) (ret js.Promise[js.Void]) {
	bindings.CallNavigatorSetAppBadge(
		this.ref, js.Pointer(&ret),
		float64(contents),
	)

	return
}

// TrySetAppBadge calls the method "Navigator.setAppBadge"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TrySetAppBadge(contents uint64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorSetAppBadge(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(contents),
	)

	return
}

// HasFuncSetAppBadge1 returns true if the method "Navigator.setAppBadge" exists.
func (this Navigator) HasFuncSetAppBadge1() bool {
	return js.True == bindings.HasFuncNavigatorSetAppBadge1(
		this.ref,
	)
}

// FuncSetAppBadge1 returns the method "Navigator.setAppBadge".
func (this Navigator) FuncSetAppBadge1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncNavigatorSetAppBadge1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetAppBadge1 calls the method "Navigator.setAppBadge".
func (this Navigator) SetAppBadge1() (ret js.Promise[js.Void]) {
	bindings.CallNavigatorSetAppBadge1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySetAppBadge1 calls the method "Navigator.setAppBadge"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TrySetAppBadge1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorSetAppBadge1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClearAppBadge returns true if the method "Navigator.clearAppBadge" exists.
func (this Navigator) HasFuncClearAppBadge() bool {
	return js.True == bindings.HasFuncNavigatorClearAppBadge(
		this.ref,
	)
}

// FuncClearAppBadge returns the method "Navigator.clearAppBadge".
func (this Navigator) FuncClearAppBadge() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncNavigatorClearAppBadge(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearAppBadge calls the method "Navigator.clearAppBadge".
func (this Navigator) ClearAppBadge() (ret js.Promise[js.Void]) {
	bindings.CallNavigatorClearAppBadge(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClearAppBadge calls the method "Navigator.clearAppBadge"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigator) TryClearAppBadge() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigatorClearAppBadge(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
		js.ThrowInvalidCallbackInvocation()
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
	this.ref.Once()
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
	this.ref.Free()
}

// TargetURL returns the value of property "LaunchParams.targetURL".
//
// It returns ok=false if there is no such property.
func (this LaunchParams) TargetURL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLaunchParamsTargetURL(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Files returns the value of property "LaunchParams.files".
//
// It returns ok=false if there is no such property.
func (this LaunchParams) Files() (ret js.FrozenArray[FileSystemHandle], ok bool) {
	ok = js.True == bindings.GetLaunchParamsFiles(
		this.ref, js.Pointer(&ret),
	)
	return
}

type LaunchQueue struct {
	ref js.Ref
}

func (this LaunchQueue) Once() LaunchQueue {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncSetConsumer returns true if the method "LaunchQueue.setConsumer" exists.
func (this LaunchQueue) HasFuncSetConsumer() bool {
	return js.True == bindings.HasFuncLaunchQueueSetConsumer(
		this.ref,
	)
}

// FuncSetConsumer returns the method "LaunchQueue.setConsumer".
func (this LaunchQueue) FuncSetConsumer() (fn js.Func[func(consumer js.Func[func(params LaunchParams) js.Any])]) {
	bindings.FuncLaunchQueueSetConsumer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetConsumer calls the method "LaunchQueue.setConsumer".
func (this LaunchQueue) SetConsumer(consumer js.Func[func(params LaunchParams) js.Any]) (ret js.Void) {
	bindings.CallLaunchQueueSetConsumer(
		this.ref, js.Pointer(&ret),
		consumer.Ref(),
	)

	return
}

// TrySetConsumer calls the method "LaunchQueue.setConsumer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this LaunchQueue) TrySetConsumer(consumer js.Func[func(params LaunchParams) js.Any]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLaunchQueueSetConsumer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *FenceEvent) UpdateFrom(ref js.Ref) {
	bindings.FenceEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FenceEvent) Update(ref js.Ref) {
	bindings.FenceEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FenceEvent) FreeMembers(recursive bool) {
	js.Free(
		p.EventType.Ref(),
		p.EventData.Ref(),
		p.Destination.Ref(),
	)
	p.EventType = p.EventType.FromRef(js.Undefined)
	p.EventData = p.EventData.FromRef(js.Undefined)
	p.Destination = p.Destination.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncReportEvent returns true if the method "Fence.reportEvent" exists.
func (this Fence) HasFuncReportEvent() bool {
	return js.True == bindings.HasFuncFenceReportEvent(
		this.ref,
	)
}

// FuncReportEvent returns the method "Fence.reportEvent".
func (this Fence) FuncReportEvent() (fn js.Func[func(event ReportEventType)]) {
	bindings.FuncFenceReportEvent(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReportEvent calls the method "Fence.reportEvent".
func (this Fence) ReportEvent(event ReportEventType) (ret js.Void) {
	bindings.CallFenceReportEvent(
		this.ref, js.Pointer(&ret),
		event.Ref(),
	)

	return
}

// TryReportEvent calls the method "Fence.reportEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Fence) TryReportEvent(event ReportEventType) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFenceReportEvent(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		event.Ref(),
	)

	return
}

// HasFuncSetReportEventDataForAutomaticBeacons returns true if the method "Fence.setReportEventDataForAutomaticBeacons" exists.
func (this Fence) HasFuncSetReportEventDataForAutomaticBeacons() bool {
	return js.True == bindings.HasFuncFenceSetReportEventDataForAutomaticBeacons(
		this.ref,
	)
}

// FuncSetReportEventDataForAutomaticBeacons returns the method "Fence.setReportEventDataForAutomaticBeacons".
func (this Fence) FuncSetReportEventDataForAutomaticBeacons() (fn js.Func[func(event FenceEvent)]) {
	bindings.FuncFenceSetReportEventDataForAutomaticBeacons(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetReportEventDataForAutomaticBeacons calls the method "Fence.setReportEventDataForAutomaticBeacons".
func (this Fence) SetReportEventDataForAutomaticBeacons(event FenceEvent) (ret js.Void) {
	bindings.CallFenceSetReportEventDataForAutomaticBeacons(
		this.ref, js.Pointer(&ret),
		js.Pointer(&event),
	)

	return
}

// TrySetReportEventDataForAutomaticBeacons calls the method "Fence.setReportEventDataForAutomaticBeacons"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Fence) TrySetReportEventDataForAutomaticBeacons(event FenceEvent) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFenceSetReportEventDataForAutomaticBeacons(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&event),
	)

	return
}

// HasFuncGetNestedConfigs returns true if the method "Fence.getNestedConfigs" exists.
func (this Fence) HasFuncGetNestedConfigs() bool {
	return js.True == bindings.HasFuncFenceGetNestedConfigs(
		this.ref,
	)
}

// FuncGetNestedConfigs returns the method "Fence.getNestedConfigs".
func (this Fence) FuncGetNestedConfigs() (fn js.Func[func() js.Array[FencedFrameConfig]]) {
	bindings.FuncFenceGetNestedConfigs(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetNestedConfigs calls the method "Fence.getNestedConfigs".
func (this Fence) GetNestedConfigs() (ret js.Array[FencedFrameConfig]) {
	bindings.CallFenceGetNestedConfigs(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetNestedConfigs calls the method "Fence.getNestedConfigs"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Fence) TryGetNestedConfigs() (ret js.Array[FencedFrameConfig], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFenceGetNestedConfigs(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PortalHost struct {
	EventTarget
}

func (this PortalHost) Once() PortalHost {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncPostMessage returns true if the method "PortalHost.postMessage" exists.
func (this PortalHost) HasFuncPostMessage() bool {
	return js.True == bindings.HasFuncPortalHostPostMessage(
		this.ref,
	)
}

// FuncPostMessage returns the method "PortalHost.postMessage".
func (this PortalHost) FuncPostMessage() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	bindings.FuncPortalHostPostMessage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PostMessage calls the method "PortalHost.postMessage".
func (this PortalHost) PostMessage(message js.Any, options StructuredSerializeOptions) (ret js.Void) {
	bindings.CallPortalHostPostMessage(
		this.ref, js.Pointer(&ret),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryPostMessage calls the method "PortalHost.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PortalHost) TryPostMessage(message js.Any, options StructuredSerializeOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPortalHostPostMessage(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncPostMessage1 returns true if the method "PortalHost.postMessage" exists.
func (this PortalHost) HasFuncPostMessage1() bool {
	return js.True == bindings.HasFuncPortalHostPostMessage1(
		this.ref,
	)
}

// FuncPostMessage1 returns the method "PortalHost.postMessage".
func (this PortalHost) FuncPostMessage1() (fn js.Func[func(message js.Any)]) {
	bindings.FuncPortalHostPostMessage1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PostMessage1 calls the method "PortalHost.postMessage".
func (this PortalHost) PostMessage1(message js.Any) (ret js.Void) {
	bindings.CallPortalHostPostMessage1(
		this.ref, js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPostMessage1 calls the method "PortalHost.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PortalHost) TryPostMessage1(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPortalHostPostMessage1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Type returns the value of property "ScreenOrientation.type".
//
// It returns ok=false if there is no such property.
func (this ScreenOrientation) Type() (ret OrientationType, ok bool) {
	ok = js.True == bindings.GetScreenOrientationType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Angle returns the value of property "ScreenOrientation.angle".
//
// It returns ok=false if there is no such property.
func (this ScreenOrientation) Angle() (ret uint16, ok bool) {
	ok = js.True == bindings.GetScreenOrientationAngle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncLock returns true if the method "ScreenOrientation.lock" exists.
func (this ScreenOrientation) HasFuncLock() bool {
	return js.True == bindings.HasFuncScreenOrientationLock(
		this.ref,
	)
}

// FuncLock returns the method "ScreenOrientation.lock".
func (this ScreenOrientation) FuncLock() (fn js.Func[func(orientation OrientationLockType) js.Promise[js.Void]]) {
	bindings.FuncScreenOrientationLock(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Lock calls the method "ScreenOrientation.lock".
func (this ScreenOrientation) Lock(orientation OrientationLockType) (ret js.Promise[js.Void]) {
	bindings.CallScreenOrientationLock(
		this.ref, js.Pointer(&ret),
		uint32(orientation),
	)

	return
}

// TryLock calls the method "ScreenOrientation.lock"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ScreenOrientation) TryLock(orientation OrientationLockType) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryScreenOrientationLock(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(orientation),
	)

	return
}

// HasFuncUnlock returns true if the method "ScreenOrientation.unlock" exists.
func (this ScreenOrientation) HasFuncUnlock() bool {
	return js.True == bindings.HasFuncScreenOrientationUnlock(
		this.ref,
	)
}

// FuncUnlock returns the method "ScreenOrientation.unlock".
func (this ScreenOrientation) FuncUnlock() (fn js.Func[func()]) {
	bindings.FuncScreenOrientationUnlock(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Unlock calls the method "ScreenOrientation.unlock".
func (this ScreenOrientation) Unlock() (ret js.Void) {
	bindings.CallScreenOrientationUnlock(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryUnlock calls the method "ScreenOrientation.unlock"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ScreenOrientation) TryUnlock() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryScreenOrientationUnlock(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type Screen struct {
	ref js.Ref
}

func (this Screen) Once() Screen {
	this.ref.Once()
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
	this.ref.Free()
}

// AvailWidth returns the value of property "Screen.availWidth".
//
// It returns ok=false if there is no such property.
func (this Screen) AvailWidth() (ret int32, ok bool) {
	ok = js.True == bindings.GetScreenAvailWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AvailHeight returns the value of property "Screen.availHeight".
//
// It returns ok=false if there is no such property.
func (this Screen) AvailHeight() (ret int32, ok bool) {
	ok = js.True == bindings.GetScreenAvailHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "Screen.width".
//
// It returns ok=false if there is no such property.
func (this Screen) Width() (ret int32, ok bool) {
	ok = js.True == bindings.GetScreenWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "Screen.height".
//
// It returns ok=false if there is no such property.
func (this Screen) Height() (ret int32, ok bool) {
	ok = js.True == bindings.GetScreenHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ColorDepth returns the value of property "Screen.colorDepth".
//
// It returns ok=false if there is no such property.
func (this Screen) ColorDepth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetScreenColorDepth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PixelDepth returns the value of property "Screen.pixelDepth".
//
// It returns ok=false if there is no such property.
func (this Screen) PixelDepth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetScreenPixelDepth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsExtended returns the value of property "Screen.isExtended".
//
// It returns ok=false if there is no such property.
func (this Screen) IsExtended() (ret bool, ok bool) {
	ok = js.True == bindings.GetScreenIsExtended(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Orientation returns the value of property "Screen.orientation".
//
// It returns ok=false if there is no such property.
func (this Screen) Orientation() (ret ScreenOrientation, ok bool) {
	ok = js.True == bindings.GetScreenOrientation(
		this.ref, js.Pointer(&ret),
	)
	return
}

type VisualViewport struct {
	EventTarget
}

func (this VisualViewport) Once() VisualViewport {
	this.ref.Once()
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
	this.ref.Free()
}

// OffsetLeft returns the value of property "VisualViewport.offsetLeft".
//
// It returns ok=false if there is no such property.
func (this VisualViewport) OffsetLeft() (ret float64, ok bool) {
	ok = js.True == bindings.GetVisualViewportOffsetLeft(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OffsetTop returns the value of property "VisualViewport.offsetTop".
//
// It returns ok=false if there is no such property.
func (this VisualViewport) OffsetTop() (ret float64, ok bool) {
	ok = js.True == bindings.GetVisualViewportOffsetTop(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PageLeft returns the value of property "VisualViewport.pageLeft".
//
// It returns ok=false if there is no such property.
func (this VisualViewport) PageLeft() (ret float64, ok bool) {
	ok = js.True == bindings.GetVisualViewportPageLeft(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PageTop returns the value of property "VisualViewport.pageTop".
//
// It returns ok=false if there is no such property.
func (this VisualViewport) PageTop() (ret float64, ok bool) {
	ok = js.True == bindings.GetVisualViewportPageTop(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "VisualViewport.width".
//
// It returns ok=false if there is no such property.
func (this VisualViewport) Width() (ret float64, ok bool) {
	ok = js.True == bindings.GetVisualViewportWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "VisualViewport.height".
//
// It returns ok=false if there is no such property.
func (this VisualViewport) Height() (ret float64, ok bool) {
	ok = js.True == bindings.GetVisualViewportHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Scale returns the value of property "VisualViewport.scale".
//
// It returns ok=false if there is no such property.
func (this VisualViewport) Scale() (ret float64, ok bool) {
	ok = js.True == bindings.GetVisualViewportScale(
		this.ref, js.Pointer(&ret),
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
func (p *SharedStorageRunOperationMethodOptions) UpdateFrom(ref js.Ref) {
	bindings.SharedStorageRunOperationMethodOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SharedStorageRunOperationMethodOptions) Update(ref js.Ref) {
	bindings.SharedStorageRunOperationMethodOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SharedStorageRunOperationMethodOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
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
func (p *SharedStorageUrlWithMetadata) UpdateFrom(ref js.Ref) {
	bindings.SharedStorageUrlWithMetadataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SharedStorageUrlWithMetadata) Update(ref js.Ref) {
	bindings.SharedStorageUrlWithMetadataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SharedStorageUrlWithMetadata) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.ReportingMetadata.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.ReportingMetadata = p.ReportingMetadata.FromRef(js.Undefined)
}

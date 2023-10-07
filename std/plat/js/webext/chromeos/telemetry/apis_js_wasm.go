// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package telemetry

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/chromeos/telemetry/bindings"
)

type AudioOutputNodeInfo struct {
	// Id is "AudioOutputNodeInfo.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id float64
	// Name is "AudioOutputNodeInfo.name"
	//
	// Optional
	Name js.String
	// DeviceName is "AudioOutputNodeInfo.deviceName"
	//
	// Optional
	DeviceName js.String
	// Active is "AudioOutputNodeInfo.active"
	//
	// Optional
	//
	// NOTE: FFI_USE_Active MUST be set to true to make this field effective.
	Active bool
	// NodeVolume is "AudioOutputNodeInfo.nodeVolume"
	//
	// Optional
	//
	// NOTE: FFI_USE_NodeVolume MUST be set to true to make this field effective.
	NodeVolume int32

	FFI_USE_Id         bool // for Id.
	FFI_USE_Active     bool // for Active.
	FFI_USE_NodeVolume bool // for NodeVolume.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioOutputNodeInfo with all fields set.
func (p AudioOutputNodeInfo) FromRef(ref js.Ref) AudioOutputNodeInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioOutputNodeInfo in the application heap.
func (p AudioOutputNodeInfo) New() js.Ref {
	return bindings.AudioOutputNodeInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioOutputNodeInfo) UpdateFrom(ref js.Ref) {
	bindings.AudioOutputNodeInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioOutputNodeInfo) Update(ref js.Ref) {
	bindings.AudioOutputNodeInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioOutputNodeInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.DeviceName.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.DeviceName = p.DeviceName.FromRef(js.Undefined)
}

type AudioInputNodeInfo struct {
	// Id is "AudioInputNodeInfo.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id float64
	// Name is "AudioInputNodeInfo.name"
	//
	// Optional
	Name js.String
	// DeviceName is "AudioInputNodeInfo.deviceName"
	//
	// Optional
	DeviceName js.String
	// Active is "AudioInputNodeInfo.active"
	//
	// Optional
	//
	// NOTE: FFI_USE_Active MUST be set to true to make this field effective.
	Active bool
	// NodeGain is "AudioInputNodeInfo.nodeGain"
	//
	// Optional
	//
	// NOTE: FFI_USE_NodeGain MUST be set to true to make this field effective.
	NodeGain int32

	FFI_USE_Id       bool // for Id.
	FFI_USE_Active   bool // for Active.
	FFI_USE_NodeGain bool // for NodeGain.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioInputNodeInfo with all fields set.
func (p AudioInputNodeInfo) FromRef(ref js.Ref) AudioInputNodeInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioInputNodeInfo in the application heap.
func (p AudioInputNodeInfo) New() js.Ref {
	return bindings.AudioInputNodeInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioInputNodeInfo) UpdateFrom(ref js.Ref) {
	bindings.AudioInputNodeInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioInputNodeInfo) Update(ref js.Ref) {
	bindings.AudioInputNodeInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioInputNodeInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.DeviceName.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.DeviceName = p.DeviceName.FromRef(js.Undefined)
}

type AudioInfo struct {
	// OutputMute is "AudioInfo.outputMute"
	//
	// Optional
	//
	// NOTE: FFI_USE_OutputMute MUST be set to true to make this field effective.
	OutputMute bool
	// InputMute is "AudioInfo.inputMute"
	//
	// Optional
	//
	// NOTE: FFI_USE_InputMute MUST be set to true to make this field effective.
	InputMute bool
	// Underruns is "AudioInfo.underruns"
	//
	// Optional
	//
	// NOTE: FFI_USE_Underruns MUST be set to true to make this field effective.
	Underruns int32
	// SevereUnderruns is "AudioInfo.severeUnderruns"
	//
	// Optional
	//
	// NOTE: FFI_USE_SevereUnderruns MUST be set to true to make this field effective.
	SevereUnderruns int32
	// OutputNodes is "AudioInfo.outputNodes"
	//
	// Optional
	OutputNodes js.Array[AudioOutputNodeInfo]
	// InputNodes is "AudioInfo.inputNodes"
	//
	// Optional
	InputNodes js.Array[AudioInputNodeInfo]

	FFI_USE_OutputMute      bool // for OutputMute.
	FFI_USE_InputMute       bool // for InputMute.
	FFI_USE_Underruns       bool // for Underruns.
	FFI_USE_SevereUnderruns bool // for SevereUnderruns.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioInfo with all fields set.
func (p AudioInfo) FromRef(ref js.Ref) AudioInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioInfo in the application heap.
func (p AudioInfo) New() js.Ref {
	return bindings.AudioInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioInfo) UpdateFrom(ref js.Ref) {
	bindings.AudioInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioInfo) Update(ref js.Ref) {
	bindings.AudioInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioInfo) FreeMembers(recursive bool) {
	js.Free(
		p.OutputNodes.Ref(),
		p.InputNodes.Ref(),
	)
	p.OutputNodes = p.OutputNodes.FromRef(js.Undefined)
	p.InputNodes = p.InputNodes.FromRef(js.Undefined)
}

type AudioInfoCallbackFunc func(this js.Ref, audioInfo *AudioInfo) js.Ref

func (fn AudioInfoCallbackFunc) Register() js.Func[func(audioInfo *AudioInfo)] {
	return js.RegisterCallback[func(audioInfo *AudioInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AudioInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AudioInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AudioInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, audioInfo *AudioInfo) js.Ref
	Arg T
}

func (cb *AudioInfoCallback[T]) Register() js.Func[func(audioInfo *AudioInfo)] {
	return js.RegisterCallback[func(audioInfo *AudioInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AudioInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AudioInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type BatteryInfo struct {
	// CycleCount is "BatteryInfo.cycleCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_CycleCount MUST be set to true to make this field effective.
	CycleCount float64
	// VoltageNow is "BatteryInfo.voltageNow"
	//
	// Optional
	//
	// NOTE: FFI_USE_VoltageNow MUST be set to true to make this field effective.
	VoltageNow float64
	// Vendor is "BatteryInfo.vendor"
	//
	// Optional
	Vendor js.String
	// SerialNumber is "BatteryInfo.serialNumber"
	//
	// Optional
	SerialNumber js.String
	// ChargeFullDesign is "BatteryInfo.chargeFullDesign"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChargeFullDesign MUST be set to true to make this field effective.
	ChargeFullDesign float64
	// ChargeFull is "BatteryInfo.chargeFull"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChargeFull MUST be set to true to make this field effective.
	ChargeFull float64
	// VoltageMinDesign is "BatteryInfo.voltageMinDesign"
	//
	// Optional
	//
	// NOTE: FFI_USE_VoltageMinDesign MUST be set to true to make this field effective.
	VoltageMinDesign float64
	// ModelName is "BatteryInfo.modelName"
	//
	// Optional
	ModelName js.String
	// ChargeNow is "BatteryInfo.chargeNow"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChargeNow MUST be set to true to make this field effective.
	ChargeNow float64
	// CurrentNow is "BatteryInfo.currentNow"
	//
	// Optional
	//
	// NOTE: FFI_USE_CurrentNow MUST be set to true to make this field effective.
	CurrentNow float64
	// Technology is "BatteryInfo.technology"
	//
	// Optional
	Technology js.String
	// Status is "BatteryInfo.status"
	//
	// Optional
	Status js.String
	// ManufactureDate is "BatteryInfo.manufactureDate"
	//
	// Optional
	ManufactureDate js.String
	// Temperature is "BatteryInfo.temperature"
	//
	// Optional
	//
	// NOTE: FFI_USE_Temperature MUST be set to true to make this field effective.
	Temperature float64

	FFI_USE_CycleCount       bool // for CycleCount.
	FFI_USE_VoltageNow       bool // for VoltageNow.
	FFI_USE_ChargeFullDesign bool // for ChargeFullDesign.
	FFI_USE_ChargeFull       bool // for ChargeFull.
	FFI_USE_VoltageMinDesign bool // for VoltageMinDesign.
	FFI_USE_ChargeNow        bool // for ChargeNow.
	FFI_USE_CurrentNow       bool // for CurrentNow.
	FFI_USE_Temperature      bool // for Temperature.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BatteryInfo with all fields set.
func (p BatteryInfo) FromRef(ref js.Ref) BatteryInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BatteryInfo in the application heap.
func (p BatteryInfo) New() js.Ref {
	return bindings.BatteryInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BatteryInfo) UpdateFrom(ref js.Ref) {
	bindings.BatteryInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BatteryInfo) Update(ref js.Ref) {
	bindings.BatteryInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BatteryInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Vendor.Ref(),
		p.SerialNumber.Ref(),
		p.ModelName.Ref(),
		p.Technology.Ref(),
		p.Status.Ref(),
		p.ManufactureDate.Ref(),
	)
	p.Vendor = p.Vendor.FromRef(js.Undefined)
	p.SerialNumber = p.SerialNumber.FromRef(js.Undefined)
	p.ModelName = p.ModelName.FromRef(js.Undefined)
	p.Technology = p.Technology.FromRef(js.Undefined)
	p.Status = p.Status.FromRef(js.Undefined)
	p.ManufactureDate = p.ManufactureDate.FromRef(js.Undefined)
}

type BatteryInfoCallbackFunc func(this js.Ref, batteryInfo *BatteryInfo) js.Ref

func (fn BatteryInfoCallbackFunc) Register() js.Func[func(batteryInfo *BatteryInfo)] {
	return js.RegisterCallback[func(batteryInfo *BatteryInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn BatteryInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 BatteryInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type BatteryInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, batteryInfo *BatteryInfo) js.Ref
	Arg T
}

func (cb *BatteryInfoCallback[T]) Register() js.Func[func(batteryInfo *BatteryInfo)] {
	return js.RegisterCallback[func(batteryInfo *BatteryInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *BatteryInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 BatteryInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CpuArchitectureEnum uint32

const (
	_ CpuArchitectureEnum = iota

	CpuArchitectureEnum_UNKNOWN
	CpuArchitectureEnum_X86_64
	CpuArchitectureEnum_AARCH64
	CpuArchitectureEnum_ARMV_7L
)

func (CpuArchitectureEnum) FromRef(str js.Ref) CpuArchitectureEnum {
	return CpuArchitectureEnum(bindings.ConstOfCpuArchitectureEnum(str))
}

func (x CpuArchitectureEnum) String() (string, bool) {
	switch x {
	case CpuArchitectureEnum_UNKNOWN:
		return "unknown", true
	case CpuArchitectureEnum_X86_64:
		return "x86_64", true
	case CpuArchitectureEnum_AARCH64:
		return "aarch64", true
	case CpuArchitectureEnum_ARMV_7L:
		return "armv7l", true
	default:
		return "", false
	}
}

type CpuCStateInfo struct {
	// Name is "CpuCStateInfo.name"
	//
	// Optional
	Name js.String
	// TimeInStateSinceLastBootUs is "CpuCStateInfo.timeInStateSinceLastBootUs"
	//
	// Optional
	//
	// NOTE: FFI_USE_TimeInStateSinceLastBootUs MUST be set to true to make this field effective.
	TimeInStateSinceLastBootUs float64

	FFI_USE_TimeInStateSinceLastBootUs bool // for TimeInStateSinceLastBootUs.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CpuCStateInfo with all fields set.
func (p CpuCStateInfo) FromRef(ref js.Ref) CpuCStateInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CpuCStateInfo in the application heap.
func (p CpuCStateInfo) New() js.Ref {
	return bindings.CpuCStateInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CpuCStateInfo) UpdateFrom(ref js.Ref) {
	bindings.CpuCStateInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CpuCStateInfo) Update(ref js.Ref) {
	bindings.CpuCStateInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CpuCStateInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
}

type LogicalCpuInfo struct {
	// MaxClockSpeedKhz is "LogicalCpuInfo.maxClockSpeedKhz"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxClockSpeedKhz MUST be set to true to make this field effective.
	MaxClockSpeedKhz int32
	// ScalingMaxFrequencyKhz is "LogicalCpuInfo.scalingMaxFrequencyKhz"
	//
	// Optional
	//
	// NOTE: FFI_USE_ScalingMaxFrequencyKhz MUST be set to true to make this field effective.
	ScalingMaxFrequencyKhz int32
	// ScalingCurrentFrequencyKhz is "LogicalCpuInfo.scalingCurrentFrequencyKhz"
	//
	// Optional
	//
	// NOTE: FFI_USE_ScalingCurrentFrequencyKhz MUST be set to true to make this field effective.
	ScalingCurrentFrequencyKhz int32
	// IdleTimeMs is "LogicalCpuInfo.idleTimeMs"
	//
	// Optional
	//
	// NOTE: FFI_USE_IdleTimeMs MUST be set to true to make this field effective.
	IdleTimeMs float64
	// CStates is "LogicalCpuInfo.cStates"
	//
	// Optional
	CStates js.Array[CpuCStateInfo]
	// CoreId is "LogicalCpuInfo.coreId"
	//
	// Optional
	//
	// NOTE: FFI_USE_CoreId MUST be set to true to make this field effective.
	CoreId int32

	FFI_USE_MaxClockSpeedKhz           bool // for MaxClockSpeedKhz.
	FFI_USE_ScalingMaxFrequencyKhz     bool // for ScalingMaxFrequencyKhz.
	FFI_USE_ScalingCurrentFrequencyKhz bool // for ScalingCurrentFrequencyKhz.
	FFI_USE_IdleTimeMs                 bool // for IdleTimeMs.
	FFI_USE_CoreId                     bool // for CoreId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LogicalCpuInfo with all fields set.
func (p LogicalCpuInfo) FromRef(ref js.Ref) LogicalCpuInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LogicalCpuInfo in the application heap.
func (p LogicalCpuInfo) New() js.Ref {
	return bindings.LogicalCpuInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LogicalCpuInfo) UpdateFrom(ref js.Ref) {
	bindings.LogicalCpuInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LogicalCpuInfo) Update(ref js.Ref) {
	bindings.LogicalCpuInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LogicalCpuInfo) FreeMembers(recursive bool) {
	js.Free(
		p.CStates.Ref(),
	)
	p.CStates = p.CStates.FromRef(js.Undefined)
}

type PhysicalCpuInfo struct {
	// ModelName is "PhysicalCpuInfo.modelName"
	//
	// Optional
	ModelName js.String
	// LogicalCpus is "PhysicalCpuInfo.logicalCpus"
	//
	// Optional
	LogicalCpus js.Array[LogicalCpuInfo]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PhysicalCpuInfo with all fields set.
func (p PhysicalCpuInfo) FromRef(ref js.Ref) PhysicalCpuInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PhysicalCpuInfo in the application heap.
func (p PhysicalCpuInfo) New() js.Ref {
	return bindings.PhysicalCpuInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PhysicalCpuInfo) UpdateFrom(ref js.Ref) {
	bindings.PhysicalCpuInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PhysicalCpuInfo) Update(ref js.Ref) {
	bindings.PhysicalCpuInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PhysicalCpuInfo) FreeMembers(recursive bool) {
	js.Free(
		p.ModelName.Ref(),
		p.LogicalCpus.Ref(),
	)
	p.ModelName = p.ModelName.FromRef(js.Undefined)
	p.LogicalCpus = p.LogicalCpus.FromRef(js.Undefined)
}

type CpuInfo struct {
	// NumTotalThreads is "CpuInfo.numTotalThreads"
	//
	// Optional
	//
	// NOTE: FFI_USE_NumTotalThreads MUST be set to true to make this field effective.
	NumTotalThreads int32
	// Architecture is "CpuInfo.architecture"
	//
	// Optional
	Architecture CpuArchitectureEnum
	// PhysicalCpus is "CpuInfo.physicalCpus"
	//
	// Optional
	PhysicalCpus js.Array[PhysicalCpuInfo]

	FFI_USE_NumTotalThreads bool // for NumTotalThreads.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CpuInfo with all fields set.
func (p CpuInfo) FromRef(ref js.Ref) CpuInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CpuInfo in the application heap.
func (p CpuInfo) New() js.Ref {
	return bindings.CpuInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CpuInfo) UpdateFrom(ref js.Ref) {
	bindings.CpuInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CpuInfo) Update(ref js.Ref) {
	bindings.CpuInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CpuInfo) FreeMembers(recursive bool) {
	js.Free(
		p.PhysicalCpus.Ref(),
	)
	p.PhysicalCpus = p.PhysicalCpus.FromRef(js.Undefined)
}

type CpuInfoCallbackFunc func(this js.Ref, cpuInfo *CpuInfo) js.Ref

func (fn CpuInfoCallbackFunc) Register() js.Func[func(cpuInfo *CpuInfo)] {
	return js.RegisterCallback[func(cpuInfo *CpuInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CpuInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CpuInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CpuInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, cpuInfo *CpuInfo) js.Ref
	Arg T
}

func (cb *CpuInfoCallback[T]) Register() js.Func[func(cpuInfo *CpuInfo)] {
	return js.RegisterCallback[func(cpuInfo *CpuInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CpuInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CpuInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DisplayInputType uint32

const (
	_ DisplayInputType = iota

	DisplayInputType_UNKNOWN
	DisplayInputType_DIGITAL
	DisplayInputType_ANALOG
)

func (DisplayInputType) FromRef(str js.Ref) DisplayInputType {
	return DisplayInputType(bindings.ConstOfDisplayInputType(str))
}

func (x DisplayInputType) String() (string, bool) {
	switch x {
	case DisplayInputType_UNKNOWN:
		return "unknown", true
	case DisplayInputType_DIGITAL:
		return "digital", true
	case DisplayInputType_ANALOG:
		return "analog", true
	default:
		return "", false
	}
}

type EmbeddedDisplayInfo struct {
	// PrivacyScreenSupported is "EmbeddedDisplayInfo.privacyScreenSupported"
	//
	// Optional
	//
	// NOTE: FFI_USE_PrivacyScreenSupported MUST be set to true to make this field effective.
	PrivacyScreenSupported bool
	// PrivacyScreenEnabled is "EmbeddedDisplayInfo.privacyScreenEnabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_PrivacyScreenEnabled MUST be set to true to make this field effective.
	PrivacyScreenEnabled bool
	// DisplayWidth is "EmbeddedDisplayInfo.displayWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_DisplayWidth MUST be set to true to make this field effective.
	DisplayWidth int32
	// DisplayHeight is "EmbeddedDisplayInfo.displayHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_DisplayHeight MUST be set to true to make this field effective.
	DisplayHeight int32
	// ResolutionHorizontal is "EmbeddedDisplayInfo.resolutionHorizontal"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResolutionHorizontal MUST be set to true to make this field effective.
	ResolutionHorizontal int32
	// ResolutionVertical is "EmbeddedDisplayInfo.resolutionVertical"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResolutionVertical MUST be set to true to make this field effective.
	ResolutionVertical int32
	// RefreshRate is "EmbeddedDisplayInfo.refreshRate"
	//
	// Optional
	//
	// NOTE: FFI_USE_RefreshRate MUST be set to true to make this field effective.
	RefreshRate float64
	// Manufacturer is "EmbeddedDisplayInfo.manufacturer"
	//
	// Optional
	Manufacturer js.String
	// ModelId is "EmbeddedDisplayInfo.modelId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ModelId MUST be set to true to make this field effective.
	ModelId int32
	// SerialNumber is "EmbeddedDisplayInfo.serialNumber"
	//
	// Optional
	//
	// NOTE: FFI_USE_SerialNumber MUST be set to true to make this field effective.
	SerialNumber int32
	// ManufactureWeek is "EmbeddedDisplayInfo.manufactureWeek"
	//
	// Optional
	//
	// NOTE: FFI_USE_ManufactureWeek MUST be set to true to make this field effective.
	ManufactureWeek int32
	// ManufactureYear is "EmbeddedDisplayInfo.manufactureYear"
	//
	// Optional
	//
	// NOTE: FFI_USE_ManufactureYear MUST be set to true to make this field effective.
	ManufactureYear int32
	// EdidVersion is "EmbeddedDisplayInfo.edidVersion"
	//
	// Optional
	EdidVersion js.String
	// InputType is "EmbeddedDisplayInfo.inputType"
	//
	// Optional
	InputType DisplayInputType
	// DisplayName is "EmbeddedDisplayInfo.displayName"
	//
	// Optional
	DisplayName js.String

	FFI_USE_PrivacyScreenSupported bool // for PrivacyScreenSupported.
	FFI_USE_PrivacyScreenEnabled   bool // for PrivacyScreenEnabled.
	FFI_USE_DisplayWidth           bool // for DisplayWidth.
	FFI_USE_DisplayHeight          bool // for DisplayHeight.
	FFI_USE_ResolutionHorizontal   bool // for ResolutionHorizontal.
	FFI_USE_ResolutionVertical     bool // for ResolutionVertical.
	FFI_USE_RefreshRate            bool // for RefreshRate.
	FFI_USE_ModelId                bool // for ModelId.
	FFI_USE_SerialNumber           bool // for SerialNumber.
	FFI_USE_ManufactureWeek        bool // for ManufactureWeek.
	FFI_USE_ManufactureYear        bool // for ManufactureYear.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EmbeddedDisplayInfo with all fields set.
func (p EmbeddedDisplayInfo) FromRef(ref js.Ref) EmbeddedDisplayInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EmbeddedDisplayInfo in the application heap.
func (p EmbeddedDisplayInfo) New() js.Ref {
	return bindings.EmbeddedDisplayInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *EmbeddedDisplayInfo) UpdateFrom(ref js.Ref) {
	bindings.EmbeddedDisplayInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EmbeddedDisplayInfo) Update(ref js.Ref) {
	bindings.EmbeddedDisplayInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EmbeddedDisplayInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Manufacturer.Ref(),
		p.EdidVersion.Ref(),
		p.DisplayName.Ref(),
	)
	p.Manufacturer = p.Manufacturer.FromRef(js.Undefined)
	p.EdidVersion = p.EdidVersion.FromRef(js.Undefined)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
}

type ExternalDisplayInfo struct {
	// DisplayWidth is "ExternalDisplayInfo.displayWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_DisplayWidth MUST be set to true to make this field effective.
	DisplayWidth int32
	// DisplayHeight is "ExternalDisplayInfo.displayHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_DisplayHeight MUST be set to true to make this field effective.
	DisplayHeight int32
	// ResolutionHorizontal is "ExternalDisplayInfo.resolutionHorizontal"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResolutionHorizontal MUST be set to true to make this field effective.
	ResolutionHorizontal int32
	// ResolutionVertical is "ExternalDisplayInfo.resolutionVertical"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResolutionVertical MUST be set to true to make this field effective.
	ResolutionVertical int32
	// RefreshRate is "ExternalDisplayInfo.refreshRate"
	//
	// Optional
	//
	// NOTE: FFI_USE_RefreshRate MUST be set to true to make this field effective.
	RefreshRate float64
	// Manufacturer is "ExternalDisplayInfo.manufacturer"
	//
	// Optional
	Manufacturer js.String
	// ModelId is "ExternalDisplayInfo.modelId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ModelId MUST be set to true to make this field effective.
	ModelId int32
	// SerialNumber is "ExternalDisplayInfo.serialNumber"
	//
	// Optional
	//
	// NOTE: FFI_USE_SerialNumber MUST be set to true to make this field effective.
	SerialNumber int32
	// ManufactureWeek is "ExternalDisplayInfo.manufactureWeek"
	//
	// Optional
	//
	// NOTE: FFI_USE_ManufactureWeek MUST be set to true to make this field effective.
	ManufactureWeek int32
	// ManufactureYear is "ExternalDisplayInfo.manufactureYear"
	//
	// Optional
	//
	// NOTE: FFI_USE_ManufactureYear MUST be set to true to make this field effective.
	ManufactureYear int32
	// EdidVersion is "ExternalDisplayInfo.edidVersion"
	//
	// Optional
	EdidVersion js.String
	// InputType is "ExternalDisplayInfo.inputType"
	//
	// Optional
	InputType DisplayInputType
	// DisplayName is "ExternalDisplayInfo.displayName"
	//
	// Optional
	DisplayName js.String

	FFI_USE_DisplayWidth         bool // for DisplayWidth.
	FFI_USE_DisplayHeight        bool // for DisplayHeight.
	FFI_USE_ResolutionHorizontal bool // for ResolutionHorizontal.
	FFI_USE_ResolutionVertical   bool // for ResolutionVertical.
	FFI_USE_RefreshRate          bool // for RefreshRate.
	FFI_USE_ModelId              bool // for ModelId.
	FFI_USE_SerialNumber         bool // for SerialNumber.
	FFI_USE_ManufactureWeek      bool // for ManufactureWeek.
	FFI_USE_ManufactureYear      bool // for ManufactureYear.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExternalDisplayInfo with all fields set.
func (p ExternalDisplayInfo) FromRef(ref js.Ref) ExternalDisplayInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExternalDisplayInfo in the application heap.
func (p ExternalDisplayInfo) New() js.Ref {
	return bindings.ExternalDisplayInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExternalDisplayInfo) UpdateFrom(ref js.Ref) {
	bindings.ExternalDisplayInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExternalDisplayInfo) Update(ref js.Ref) {
	bindings.ExternalDisplayInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExternalDisplayInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Manufacturer.Ref(),
		p.EdidVersion.Ref(),
		p.DisplayName.Ref(),
	)
	p.Manufacturer = p.Manufacturer.FromRef(js.Undefined)
	p.EdidVersion = p.EdidVersion.FromRef(js.Undefined)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
}

type DisplayInfo struct {
	// EmbeddedDisplay is "DisplayInfo.embeddedDisplay"
	//
	// Optional
	//
	// NOTE: EmbeddedDisplay.FFI_USE MUST be set to true to get EmbeddedDisplay used.
	EmbeddedDisplay EmbeddedDisplayInfo
	// ExternalDisplays is "DisplayInfo.externalDisplays"
	//
	// Optional
	ExternalDisplays js.Array[ExternalDisplayInfo]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DisplayInfo with all fields set.
func (p DisplayInfo) FromRef(ref js.Ref) DisplayInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DisplayInfo in the application heap.
func (p DisplayInfo) New() js.Ref {
	return bindings.DisplayInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DisplayInfo) UpdateFrom(ref js.Ref) {
	bindings.DisplayInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DisplayInfo) Update(ref js.Ref) {
	bindings.DisplayInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DisplayInfo) FreeMembers(recursive bool) {
	js.Free(
		p.ExternalDisplays.Ref(),
	)
	p.ExternalDisplays = p.ExternalDisplays.FromRef(js.Undefined)
	if recursive {
		p.EmbeddedDisplay.FreeMembers(true)
	}
}

type DisplayInfoCallbackFunc func(this js.Ref, displayInfo *DisplayInfo) js.Ref

func (fn DisplayInfoCallbackFunc) Register() js.Func[func(displayInfo *DisplayInfo)] {
	return js.RegisterCallback[func(displayInfo *DisplayInfo)](
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
	var arg0 DisplayInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DisplayInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, displayInfo *DisplayInfo) js.Ref
	Arg T
}

func (cb *DisplayInfoCallback[T]) Register() js.Func[func(displayInfo *DisplayInfo)] {
	return js.RegisterCallback[func(displayInfo *DisplayInfo)](
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
	var arg0 DisplayInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EOFFunc func(this js.Ref) js.Ref

func (fn EOFFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EOFFunc) DispatchCallback(
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

type EOF[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *EOF[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EOF[T]) DispatchCallback(
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

type FwupdVersionFormat uint32

const (
	_ FwupdVersionFormat = iota

	FwupdVersionFormat_PLAIN
	FwupdVersionFormat_NUMBER
	FwupdVersionFormat_PAIR
	FwupdVersionFormat_TRIPLET
	FwupdVersionFormat_QUAD
	FwupdVersionFormat_BCD
	FwupdVersionFormat_INTEL_ME
	FwupdVersionFormat_INTEL_ME2
	FwupdVersionFormat_SURFACE_LEGACY
	FwupdVersionFormat_SURFACE
	FwupdVersionFormat_DELL_BIOS
	FwupdVersionFormat_HEX
)

func (FwupdVersionFormat) FromRef(str js.Ref) FwupdVersionFormat {
	return FwupdVersionFormat(bindings.ConstOfFwupdVersionFormat(str))
}

func (x FwupdVersionFormat) String() (string, bool) {
	switch x {
	case FwupdVersionFormat_PLAIN:
		return "plain", true
	case FwupdVersionFormat_NUMBER:
		return "number", true
	case FwupdVersionFormat_PAIR:
		return "pair", true
	case FwupdVersionFormat_TRIPLET:
		return "triplet", true
	case FwupdVersionFormat_QUAD:
		return "quad", true
	case FwupdVersionFormat_BCD:
		return "bcd", true
	case FwupdVersionFormat_INTEL_ME:
		return "intelMe", true
	case FwupdVersionFormat_INTEL_ME2:
		return "intelMe2", true
	case FwupdVersionFormat_SURFACE_LEGACY:
		return "surfaceLegacy", true
	case FwupdVersionFormat_SURFACE:
		return "surface", true
	case FwupdVersionFormat_DELL_BIOS:
		return "dellBios", true
	case FwupdVersionFormat_HEX:
		return "hex", true
	default:
		return "", false
	}
}

type FwupdFirmwareVersionInfo struct {
	// Version is "FwupdFirmwareVersionInfo.version"
	//
	// Optional
	Version js.String
	// VersionFormat is "FwupdFirmwareVersionInfo.version_format"
	//
	// Optional
	VersionFormat FwupdVersionFormat

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FwupdFirmwareVersionInfo with all fields set.
func (p FwupdFirmwareVersionInfo) FromRef(ref js.Ref) FwupdFirmwareVersionInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FwupdFirmwareVersionInfo in the application heap.
func (p FwupdFirmwareVersionInfo) New() js.Ref {
	return bindings.FwupdFirmwareVersionInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FwupdFirmwareVersionInfo) UpdateFrom(ref js.Ref) {
	bindings.FwupdFirmwareVersionInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FwupdFirmwareVersionInfo) Update(ref js.Ref) {
	bindings.FwupdFirmwareVersionInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FwupdFirmwareVersionInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Version.Ref(),
	)
	p.Version = p.Version.FromRef(js.Undefined)
}

type NetworkType uint32

const (
	_ NetworkType = iota

	NetworkType_CELLULAR
	NetworkType_ETHERNET
	NetworkType_TETHER
	NetworkType_VPN
	NetworkType_WIFI
)

func (NetworkType) FromRef(str js.Ref) NetworkType {
	return NetworkType(bindings.ConstOfNetworkType(str))
}

func (x NetworkType) String() (string, bool) {
	switch x {
	case NetworkType_CELLULAR:
		return "cellular", true
	case NetworkType_ETHERNET:
		return "ethernet", true
	case NetworkType_TETHER:
		return "tether", true
	case NetworkType_VPN:
		return "vpn", true
	case NetworkType_WIFI:
		return "wifi", true
	default:
		return "", false
	}
}

type NetworkState uint32

const (
	_ NetworkState = iota

	NetworkState_UNINITIALIZED
	NetworkState_DISABLED
	NetworkState_PROHIBITED
	NetworkState_NOT_CONNECTED
	NetworkState_CONNECTING
	NetworkState_PORTAL
	NetworkState_CONNECTED
	NetworkState_ONLINE
)

func (NetworkState) FromRef(str js.Ref) NetworkState {
	return NetworkState(bindings.ConstOfNetworkState(str))
}

func (x NetworkState) String() (string, bool) {
	switch x {
	case NetworkState_UNINITIALIZED:
		return "uninitialized", true
	case NetworkState_DISABLED:
		return "disabled", true
	case NetworkState_PROHIBITED:
		return "prohibited", true
	case NetworkState_NOT_CONNECTED:
		return "not_connected", true
	case NetworkState_CONNECTING:
		return "connecting", true
	case NetworkState_PORTAL:
		return "portal", true
	case NetworkState_CONNECTED:
		return "connected", true
	case NetworkState_ONLINE:
		return "online", true
	default:
		return "", false
	}
}

type NetworkInfo struct {
	// Type is "NetworkInfo.type"
	//
	// Optional
	Type NetworkType
	// State is "NetworkInfo.state"
	//
	// Optional
	State NetworkState
	// MacAddress is "NetworkInfo.macAddress"
	//
	// Optional
	MacAddress js.String
	// Ipv4Address is "NetworkInfo.ipv4Address"
	//
	// Optional
	Ipv4Address js.String
	// Ipv6Addresses is "NetworkInfo.ipv6Addresses"
	//
	// Optional
	Ipv6Addresses js.Array[js.String]
	// SignalStrength is "NetworkInfo.signalStrength"
	//
	// Optional
	//
	// NOTE: FFI_USE_SignalStrength MUST be set to true to make this field effective.
	SignalStrength float64

	FFI_USE_SignalStrength bool // for SignalStrength.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NetworkInfo with all fields set.
func (p NetworkInfo) FromRef(ref js.Ref) NetworkInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NetworkInfo in the application heap.
func (p NetworkInfo) New() js.Ref {
	return bindings.NetworkInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NetworkInfo) UpdateFrom(ref js.Ref) {
	bindings.NetworkInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NetworkInfo) Update(ref js.Ref) {
	bindings.NetworkInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NetworkInfo) FreeMembers(recursive bool) {
	js.Free(
		p.MacAddress.Ref(),
		p.Ipv4Address.Ref(),
		p.Ipv6Addresses.Ref(),
	)
	p.MacAddress = p.MacAddress.FromRef(js.Undefined)
	p.Ipv4Address = p.Ipv4Address.FromRef(js.Undefined)
	p.Ipv6Addresses = p.Ipv6Addresses.FromRef(js.Undefined)
}

type InternetConnectivityInfo struct {
	// Networks is "InternetConnectivityInfo.networks"
	//
	// Optional
	Networks js.Array[NetworkInfo]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InternetConnectivityInfo with all fields set.
func (p InternetConnectivityInfo) FromRef(ref js.Ref) InternetConnectivityInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InternetConnectivityInfo in the application heap.
func (p InternetConnectivityInfo) New() js.Ref {
	return bindings.InternetConnectivityInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InternetConnectivityInfo) UpdateFrom(ref js.Ref) {
	bindings.InternetConnectivityInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InternetConnectivityInfo) Update(ref js.Ref) {
	bindings.InternetConnectivityInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InternetConnectivityInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Networks.Ref(),
	)
	p.Networks = p.Networks.FromRef(js.Undefined)
}

type InternetConnectivityInfoCallbackFunc func(this js.Ref, networkInfo *InternetConnectivityInfo) js.Ref

func (fn InternetConnectivityInfoCallbackFunc) Register() js.Func[func(networkInfo *InternetConnectivityInfo)] {
	return js.RegisterCallback[func(networkInfo *InternetConnectivityInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn InternetConnectivityInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 InternetConnectivityInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type InternetConnectivityInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, networkInfo *InternetConnectivityInfo) js.Ref
	Arg T
}

func (cb *InternetConnectivityInfoCallback[T]) Register() js.Func[func(networkInfo *InternetConnectivityInfo)] {
	return js.RegisterCallback[func(networkInfo *InternetConnectivityInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *InternetConnectivityInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 InternetConnectivityInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type MarketingInfo struct {
	// MarketingName is "MarketingInfo.marketingName"
	//
	// Optional
	MarketingName js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MarketingInfo with all fields set.
func (p MarketingInfo) FromRef(ref js.Ref) MarketingInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MarketingInfo in the application heap.
func (p MarketingInfo) New() js.Ref {
	return bindings.MarketingInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MarketingInfo) UpdateFrom(ref js.Ref) {
	bindings.MarketingInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MarketingInfo) Update(ref js.Ref) {
	bindings.MarketingInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MarketingInfo) FreeMembers(recursive bool) {
	js.Free(
		p.MarketingName.Ref(),
	)
	p.MarketingName = p.MarketingName.FromRef(js.Undefined)
}

type MarketingInfoCallbackFunc func(this js.Ref, marketingInfo *MarketingInfo) js.Ref

func (fn MarketingInfoCallbackFunc) Register() js.Func[func(marketingInfo *MarketingInfo)] {
	return js.RegisterCallback[func(marketingInfo *MarketingInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn MarketingInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MarketingInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type MarketingInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, marketingInfo *MarketingInfo) js.Ref
	Arg T
}

func (cb *MarketingInfoCallback[T]) Register() js.Func[func(marketingInfo *MarketingInfo)] {
	return js.RegisterCallback[func(marketingInfo *MarketingInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *MarketingInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MarketingInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type MemoryInfo struct {
	// TotalMemoryKiB is "MemoryInfo.totalMemoryKiB"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalMemoryKiB MUST be set to true to make this field effective.
	TotalMemoryKiB int32
	// FreeMemoryKiB is "MemoryInfo.freeMemoryKiB"
	//
	// Optional
	//
	// NOTE: FFI_USE_FreeMemoryKiB MUST be set to true to make this field effective.
	FreeMemoryKiB int32
	// AvailableMemoryKiB is "MemoryInfo.availableMemoryKiB"
	//
	// Optional
	//
	// NOTE: FFI_USE_AvailableMemoryKiB MUST be set to true to make this field effective.
	AvailableMemoryKiB int32
	// PageFaultsSinceLastBoot is "MemoryInfo.pageFaultsSinceLastBoot"
	//
	// Optional
	//
	// NOTE: FFI_USE_PageFaultsSinceLastBoot MUST be set to true to make this field effective.
	PageFaultsSinceLastBoot float64

	FFI_USE_TotalMemoryKiB          bool // for TotalMemoryKiB.
	FFI_USE_FreeMemoryKiB           bool // for FreeMemoryKiB.
	FFI_USE_AvailableMemoryKiB      bool // for AvailableMemoryKiB.
	FFI_USE_PageFaultsSinceLastBoot bool // for PageFaultsSinceLastBoot.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MemoryInfo with all fields set.
func (p MemoryInfo) FromRef(ref js.Ref) MemoryInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MemoryInfo in the application heap.
func (p MemoryInfo) New() js.Ref {
	return bindings.MemoryInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MemoryInfo) UpdateFrom(ref js.Ref) {
	bindings.MemoryInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MemoryInfo) Update(ref js.Ref) {
	bindings.MemoryInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MemoryInfo) FreeMembers(recursive bool) {
}

type MemoryInfoCallbackFunc func(this js.Ref, cpuInfo *MemoryInfo) js.Ref

func (fn MemoryInfoCallbackFunc) Register() js.Func[func(cpuInfo *MemoryInfo)] {
	return js.RegisterCallback[func(cpuInfo *MemoryInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn MemoryInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MemoryInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type MemoryInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, cpuInfo *MemoryInfo) js.Ref
	Arg T
}

func (cb *MemoryInfoCallback[T]) Register() js.Func[func(cpuInfo *MemoryInfo)] {
	return js.RegisterCallback[func(cpuInfo *MemoryInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *MemoryInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 MemoryInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type NonRemovableBlockDeviceInfo struct {
	// Name is "NonRemovableBlockDeviceInfo.name"
	//
	// Optional
	Name js.String
	// Type is "NonRemovableBlockDeviceInfo.type"
	//
	// Optional
	Type js.String
	// Size is "NonRemovableBlockDeviceInfo.size"
	//
	// Optional
	//
	// NOTE: FFI_USE_Size MUST be set to true to make this field effective.
	Size float64

	FFI_USE_Size bool // for Size.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NonRemovableBlockDeviceInfo with all fields set.
func (p NonRemovableBlockDeviceInfo) FromRef(ref js.Ref) NonRemovableBlockDeviceInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NonRemovableBlockDeviceInfo in the application heap.
func (p NonRemovableBlockDeviceInfo) New() js.Ref {
	return bindings.NonRemovableBlockDeviceInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NonRemovableBlockDeviceInfo) UpdateFrom(ref js.Ref) {
	bindings.NonRemovableBlockDeviceInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NonRemovableBlockDeviceInfo) Update(ref js.Ref) {
	bindings.NonRemovableBlockDeviceInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NonRemovableBlockDeviceInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.Type.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Type = p.Type.FromRef(js.Undefined)
}

type NonRemovableBlockDeviceInfoResponse struct {
	// DeviceInfos is "NonRemovableBlockDeviceInfoResponse.deviceInfos"
	//
	// Optional
	DeviceInfos js.Array[NonRemovableBlockDeviceInfo]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NonRemovableBlockDeviceInfoResponse with all fields set.
func (p NonRemovableBlockDeviceInfoResponse) FromRef(ref js.Ref) NonRemovableBlockDeviceInfoResponse {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NonRemovableBlockDeviceInfoResponse in the application heap.
func (p NonRemovableBlockDeviceInfoResponse) New() js.Ref {
	return bindings.NonRemovableBlockDeviceInfoResponseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NonRemovableBlockDeviceInfoResponse) UpdateFrom(ref js.Ref) {
	bindings.NonRemovableBlockDeviceInfoResponseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NonRemovableBlockDeviceInfoResponse) Update(ref js.Ref) {
	bindings.NonRemovableBlockDeviceInfoResponseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NonRemovableBlockDeviceInfoResponse) FreeMembers(recursive bool) {
	js.Free(
		p.DeviceInfos.Ref(),
	)
	p.DeviceInfos = p.DeviceInfos.FromRef(js.Undefined)
}

type NonRemovableBlockDevicesInfoCallbackFunc func(this js.Ref, deviceInfoResponse *NonRemovableBlockDeviceInfoResponse) js.Ref

func (fn NonRemovableBlockDevicesInfoCallbackFunc) Register() js.Func[func(deviceInfoResponse *NonRemovableBlockDeviceInfoResponse)] {
	return js.RegisterCallback[func(deviceInfoResponse *NonRemovableBlockDeviceInfoResponse)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn NonRemovableBlockDevicesInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 NonRemovableBlockDeviceInfoResponse
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type NonRemovableBlockDevicesInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, deviceInfoResponse *NonRemovableBlockDeviceInfoResponse) js.Ref
	Arg T
}

func (cb *NonRemovableBlockDevicesInfoCallback[T]) Register() js.Func[func(deviceInfoResponse *NonRemovableBlockDeviceInfoResponse)] {
	return js.RegisterCallback[func(deviceInfoResponse *NonRemovableBlockDeviceInfoResponse)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *NonRemovableBlockDevicesInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 NonRemovableBlockDeviceInfoResponse
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OemData struct {
	// OemData is "OemData.oemData"
	//
	// Optional
	OemData js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OemData with all fields set.
func (p OemData) FromRef(ref js.Ref) OemData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OemData in the application heap.
func (p OemData) New() js.Ref {
	return bindings.OemDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OemData) UpdateFrom(ref js.Ref) {
	bindings.OemDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OemData) Update(ref js.Ref) {
	bindings.OemDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OemData) FreeMembers(recursive bool) {
	js.Free(
		p.OemData.Ref(),
	)
	p.OemData = p.OemData.FromRef(js.Undefined)
}

type OemDataCallbackFunc func(this js.Ref, oemData *OemData) js.Ref

func (fn OemDataCallbackFunc) Register() js.Func[func(oemData *OemData)] {
	return js.RegisterCallback[func(oemData *OemData)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OemDataCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OemData
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OemDataCallback[T any] struct {
	Fn  func(arg T, this js.Ref, oemData *OemData) js.Ref
	Arg T
}

func (cb *OemDataCallback[T]) Register() js.Func[func(oemData *OemData)] {
	return js.RegisterCallback[func(oemData *OemData)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OemDataCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OemData
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OsVersionInfo struct {
	// ReleaseMilestone is "OsVersionInfo.releaseMilestone"
	//
	// Optional
	ReleaseMilestone js.String
	// BuildNumber is "OsVersionInfo.buildNumber"
	//
	// Optional
	BuildNumber js.String
	// PatchNumber is "OsVersionInfo.patchNumber"
	//
	// Optional
	PatchNumber js.String
	// ReleaseChannel is "OsVersionInfo.releaseChannel"
	//
	// Optional
	ReleaseChannel js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OsVersionInfo with all fields set.
func (p OsVersionInfo) FromRef(ref js.Ref) OsVersionInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OsVersionInfo in the application heap.
func (p OsVersionInfo) New() js.Ref {
	return bindings.OsVersionInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OsVersionInfo) UpdateFrom(ref js.Ref) {
	bindings.OsVersionInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OsVersionInfo) Update(ref js.Ref) {
	bindings.OsVersionInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OsVersionInfo) FreeMembers(recursive bool) {
	js.Free(
		p.ReleaseMilestone.Ref(),
		p.BuildNumber.Ref(),
		p.PatchNumber.Ref(),
		p.ReleaseChannel.Ref(),
	)
	p.ReleaseMilestone = p.ReleaseMilestone.FromRef(js.Undefined)
	p.BuildNumber = p.BuildNumber.FromRef(js.Undefined)
	p.PatchNumber = p.PatchNumber.FromRef(js.Undefined)
	p.ReleaseChannel = p.ReleaseChannel.FromRef(js.Undefined)
}

type OsVersionInfoCallbackFunc func(this js.Ref, osVersionInfo *OsVersionInfo) js.Ref

func (fn OsVersionInfoCallbackFunc) Register() js.Func[func(osVersionInfo *OsVersionInfo)] {
	return js.RegisterCallback[func(osVersionInfo *OsVersionInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OsVersionInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OsVersionInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OsVersionInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, osVersionInfo *OsVersionInfo) js.Ref
	Arg T
}

func (cb *OsVersionInfoCallback[T]) Register() js.Func[func(osVersionInfo *OsVersionInfo)] {
	return js.RegisterCallback[func(osVersionInfo *OsVersionInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OsVersionInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OsVersionInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type StatefulPartitionInfo struct {
	// AvailableSpace is "StatefulPartitionInfo.availableSpace"
	//
	// Optional
	//
	// NOTE: FFI_USE_AvailableSpace MUST be set to true to make this field effective.
	AvailableSpace float64
	// TotalSpace is "StatefulPartitionInfo.totalSpace"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalSpace MUST be set to true to make this field effective.
	TotalSpace float64

	FFI_USE_AvailableSpace bool // for AvailableSpace.
	FFI_USE_TotalSpace     bool // for TotalSpace.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StatefulPartitionInfo with all fields set.
func (p StatefulPartitionInfo) FromRef(ref js.Ref) StatefulPartitionInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StatefulPartitionInfo in the application heap.
func (p StatefulPartitionInfo) New() js.Ref {
	return bindings.StatefulPartitionInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StatefulPartitionInfo) UpdateFrom(ref js.Ref) {
	bindings.StatefulPartitionInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StatefulPartitionInfo) Update(ref js.Ref) {
	bindings.StatefulPartitionInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StatefulPartitionInfo) FreeMembers(recursive bool) {
}

type StatefulPartitionInfoCallbackFunc func(this js.Ref, statefulPartitionInfo *StatefulPartitionInfo) js.Ref

func (fn StatefulPartitionInfoCallbackFunc) Register() js.Func[func(statefulPartitionInfo *StatefulPartitionInfo)] {
	return js.RegisterCallback[func(statefulPartitionInfo *StatefulPartitionInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn StatefulPartitionInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StatefulPartitionInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type StatefulPartitionInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, statefulPartitionInfo *StatefulPartitionInfo) js.Ref
	Arg T
}

func (cb *StatefulPartitionInfoCallback[T]) Register() js.Func[func(statefulPartitionInfo *StatefulPartitionInfo)] {
	return js.RegisterCallback[func(statefulPartitionInfo *StatefulPartitionInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *StatefulPartitionInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StatefulPartitionInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TpmDictionaryAttack struct {
	// Counter is "TpmDictionaryAttack.counter"
	//
	// Optional
	//
	// NOTE: FFI_USE_Counter MUST be set to true to make this field effective.
	Counter int32
	// Threshold is "TpmDictionaryAttack.threshold"
	//
	// Optional
	//
	// NOTE: FFI_USE_Threshold MUST be set to true to make this field effective.
	Threshold int32
	// LockoutInEffect is "TpmDictionaryAttack.lockoutInEffect"
	//
	// Optional
	//
	// NOTE: FFI_USE_LockoutInEffect MUST be set to true to make this field effective.
	LockoutInEffect bool
	// LockoutSecondsRemaining is "TpmDictionaryAttack.lockoutSecondsRemaining"
	//
	// Optional
	//
	// NOTE: FFI_USE_LockoutSecondsRemaining MUST be set to true to make this field effective.
	LockoutSecondsRemaining int32

	FFI_USE_Counter                 bool // for Counter.
	FFI_USE_Threshold               bool // for Threshold.
	FFI_USE_LockoutInEffect         bool // for LockoutInEffect.
	FFI_USE_LockoutSecondsRemaining bool // for LockoutSecondsRemaining.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TpmDictionaryAttack with all fields set.
func (p TpmDictionaryAttack) FromRef(ref js.Ref) TpmDictionaryAttack {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TpmDictionaryAttack in the application heap.
func (p TpmDictionaryAttack) New() js.Ref {
	return bindings.TpmDictionaryAttackJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TpmDictionaryAttack) UpdateFrom(ref js.Ref) {
	bindings.TpmDictionaryAttackJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TpmDictionaryAttack) Update(ref js.Ref) {
	bindings.TpmDictionaryAttackJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TpmDictionaryAttack) FreeMembers(recursive bool) {
}

type TpmGSCVersion uint32

const (
	_ TpmGSCVersion = iota

	TpmGSCVersion_NOT_GSC
	TpmGSCVersion_CR50
	TpmGSCVersion_TI50
)

func (TpmGSCVersion) FromRef(str js.Ref) TpmGSCVersion {
	return TpmGSCVersion(bindings.ConstOfTpmGSCVersion(str))
}

func (x TpmGSCVersion) String() (string, bool) {
	switch x {
	case TpmGSCVersion_NOT_GSC:
		return "not_gsc", true
	case TpmGSCVersion_CR50:
		return "cr50", true
	case TpmGSCVersion_TI50:
		return "ti50", true
	default:
		return "", false
	}
}

type TpmVersion struct {
	// GscVersion is "TpmVersion.gscVersion"
	//
	// Optional
	GscVersion TpmGSCVersion
	// Family is "TpmVersion.family"
	//
	// Optional
	//
	// NOTE: FFI_USE_Family MUST be set to true to make this field effective.
	Family int32
	// SpecLevel is "TpmVersion.specLevel"
	//
	// Optional
	//
	// NOTE: FFI_USE_SpecLevel MUST be set to true to make this field effective.
	SpecLevel float64
	// Manufacturer is "TpmVersion.manufacturer"
	//
	// Optional
	//
	// NOTE: FFI_USE_Manufacturer MUST be set to true to make this field effective.
	Manufacturer int32
	// TpmModel is "TpmVersion.tpmModel"
	//
	// Optional
	//
	// NOTE: FFI_USE_TpmModel MUST be set to true to make this field effective.
	TpmModel int32
	// FirmwareVersion is "TpmVersion.firmwareVersion"
	//
	// Optional
	//
	// NOTE: FFI_USE_FirmwareVersion MUST be set to true to make this field effective.
	FirmwareVersion float64
	// VendorSpecific is "TpmVersion.vendorSpecific"
	//
	// Optional
	VendorSpecific js.String

	FFI_USE_Family          bool // for Family.
	FFI_USE_SpecLevel       bool // for SpecLevel.
	FFI_USE_Manufacturer    bool // for Manufacturer.
	FFI_USE_TpmModel        bool // for TpmModel.
	FFI_USE_FirmwareVersion bool // for FirmwareVersion.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TpmVersion with all fields set.
func (p TpmVersion) FromRef(ref js.Ref) TpmVersion {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TpmVersion in the application heap.
func (p TpmVersion) New() js.Ref {
	return bindings.TpmVersionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TpmVersion) UpdateFrom(ref js.Ref) {
	bindings.TpmVersionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TpmVersion) Update(ref js.Ref) {
	bindings.TpmVersionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TpmVersion) FreeMembers(recursive bool) {
	js.Free(
		p.VendorSpecific.Ref(),
	)
	p.VendorSpecific = p.VendorSpecific.FromRef(js.Undefined)
}

type TpmStatus struct {
	// Enabled is "TpmStatus.enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_Enabled MUST be set to true to make this field effective.
	Enabled bool
	// Owned is "TpmStatus.owned"
	//
	// Optional
	//
	// NOTE: FFI_USE_Owned MUST be set to true to make this field effective.
	Owned bool
	// OwnerPasswordIsPresent is "TpmStatus.ownerPasswordIsPresent"
	//
	// Optional
	//
	// NOTE: FFI_USE_OwnerPasswordIsPresent MUST be set to true to make this field effective.
	OwnerPasswordIsPresent bool

	FFI_USE_Enabled                bool // for Enabled.
	FFI_USE_Owned                  bool // for Owned.
	FFI_USE_OwnerPasswordIsPresent bool // for OwnerPasswordIsPresent.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TpmStatus with all fields set.
func (p TpmStatus) FromRef(ref js.Ref) TpmStatus {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TpmStatus in the application heap.
func (p TpmStatus) New() js.Ref {
	return bindings.TpmStatusJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TpmStatus) UpdateFrom(ref js.Ref) {
	bindings.TpmStatusJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TpmStatus) Update(ref js.Ref) {
	bindings.TpmStatusJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TpmStatus) FreeMembers(recursive bool) {
}

type TpmInfo struct {
	// Version is "TpmInfo.version"
	//
	// Optional
	//
	// NOTE: Version.FFI_USE MUST be set to true to get Version used.
	Version TpmVersion
	// Status is "TpmInfo.status"
	//
	// Optional
	//
	// NOTE: Status.FFI_USE MUST be set to true to get Status used.
	Status TpmStatus
	// DictionaryAttack is "TpmInfo.dictionaryAttack"
	//
	// Optional
	//
	// NOTE: DictionaryAttack.FFI_USE MUST be set to true to get DictionaryAttack used.
	DictionaryAttack TpmDictionaryAttack

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TpmInfo with all fields set.
func (p TpmInfo) FromRef(ref js.Ref) TpmInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TpmInfo in the application heap.
func (p TpmInfo) New() js.Ref {
	return bindings.TpmInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TpmInfo) UpdateFrom(ref js.Ref) {
	bindings.TpmInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TpmInfo) Update(ref js.Ref) {
	bindings.TpmInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TpmInfo) FreeMembers(recursive bool) {
	if recursive {
		p.Version.FreeMembers(true)
		p.Status.FreeMembers(true)
		p.DictionaryAttack.FreeMembers(true)
	}
}

type TpmInfoCallbackFunc func(this js.Ref, tpmInfo *TpmInfo) js.Ref

func (fn TpmInfoCallbackFunc) Register() js.Func[func(tpmInfo *TpmInfo)] {
	return js.RegisterCallback[func(tpmInfo *TpmInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn TpmInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TpmInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type TpmInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, tpmInfo *TpmInfo) js.Ref
	Arg T
}

func (cb *TpmInfoCallback[T]) Register() js.Func[func(tpmInfo *TpmInfo)] {
	return js.RegisterCallback[func(tpmInfo *TpmInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *TpmInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TpmInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UsbBusInterfaceInfo struct {
	// InterfaceNumber is "UsbBusInterfaceInfo.interfaceNumber"
	//
	// Optional
	//
	// NOTE: FFI_USE_InterfaceNumber MUST be set to true to make this field effective.
	InterfaceNumber float64
	// ClassId is "UsbBusInterfaceInfo.classId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ClassId MUST be set to true to make this field effective.
	ClassId float64
	// SubclassId is "UsbBusInterfaceInfo.subclassId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SubclassId MUST be set to true to make this field effective.
	SubclassId float64
	// ProtocolId is "UsbBusInterfaceInfo.protocolId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProtocolId MUST be set to true to make this field effective.
	ProtocolId float64
	// Driver is "UsbBusInterfaceInfo.driver"
	//
	// Optional
	Driver js.String

	FFI_USE_InterfaceNumber bool // for InterfaceNumber.
	FFI_USE_ClassId         bool // for ClassId.
	FFI_USE_SubclassId      bool // for SubclassId.
	FFI_USE_ProtocolId      bool // for ProtocolId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UsbBusInterfaceInfo with all fields set.
func (p UsbBusInterfaceInfo) FromRef(ref js.Ref) UsbBusInterfaceInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UsbBusInterfaceInfo in the application heap.
func (p UsbBusInterfaceInfo) New() js.Ref {
	return bindings.UsbBusInterfaceInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UsbBusInterfaceInfo) UpdateFrom(ref js.Ref) {
	bindings.UsbBusInterfaceInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UsbBusInterfaceInfo) Update(ref js.Ref) {
	bindings.UsbBusInterfaceInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UsbBusInterfaceInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Driver.Ref(),
	)
	p.Driver = p.Driver.FromRef(js.Undefined)
}

type UsbVersion uint32

const (
	_ UsbVersion = iota

	UsbVersion_UNKNOWN
	UsbVersion_USB1
	UsbVersion_USB2
	UsbVersion_USB3
)

func (UsbVersion) FromRef(str js.Ref) UsbVersion {
	return UsbVersion(bindings.ConstOfUsbVersion(str))
}

func (x UsbVersion) String() (string, bool) {
	switch x {
	case UsbVersion_UNKNOWN:
		return "unknown", true
	case UsbVersion_USB1:
		return "usb1", true
	case UsbVersion_USB2:
		return "usb2", true
	case UsbVersion_USB3:
		return "usb3", true
	default:
		return "", false
	}
}

type UsbSpecSpeed uint32

const (
	_ UsbSpecSpeed = iota

	UsbSpecSpeed_UNKNOWN
	UsbSpecSpeed_N1_5_MBPS
	UsbSpecSpeed_N12_MBPS
	UsbSpecSpeed_N480_MBPS
	UsbSpecSpeed_N5_GBPS
	UsbSpecSpeed_N10_GBPS
	UsbSpecSpeed_N20_GBPS
)

func (UsbSpecSpeed) FromRef(str js.Ref) UsbSpecSpeed {
	return UsbSpecSpeed(bindings.ConstOfUsbSpecSpeed(str))
}

func (x UsbSpecSpeed) String() (string, bool) {
	switch x {
	case UsbSpecSpeed_UNKNOWN:
		return "unknown", true
	case UsbSpecSpeed_N1_5_MBPS:
		return "n1_5Mbps", true
	case UsbSpecSpeed_N12_MBPS:
		return "n12Mbps", true
	case UsbSpecSpeed_N480_MBPS:
		return "n480Mbps", true
	case UsbSpecSpeed_N5_GBPS:
		return "n5Gbps", true
	case UsbSpecSpeed_N10_GBPS:
		return "n10Gbps", true
	case UsbSpecSpeed_N20_GBPS:
		return "n20Gbps", true
	default:
		return "", false
	}
}

type UsbBusInfo struct {
	// ClassId is "UsbBusInfo.classId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ClassId MUST be set to true to make this field effective.
	ClassId float64
	// SubclassId is "UsbBusInfo.subclassId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SubclassId MUST be set to true to make this field effective.
	SubclassId float64
	// ProtocolId is "UsbBusInfo.protocolId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProtocolId MUST be set to true to make this field effective.
	ProtocolId float64
	// VendorId is "UsbBusInfo.vendorId"
	//
	// Optional
	//
	// NOTE: FFI_USE_VendorId MUST be set to true to make this field effective.
	VendorId float64
	// ProductId is "UsbBusInfo.productId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProductId MUST be set to true to make this field effective.
	ProductId float64
	// Interfaces is "UsbBusInfo.interfaces"
	//
	// Optional
	Interfaces js.Array[UsbBusInterfaceInfo]
	// FwupdFirmwareVersionInfo is "UsbBusInfo.fwupdFirmwareVersionInfo"
	//
	// Optional
	//
	// NOTE: FwupdFirmwareVersionInfo.FFI_USE MUST be set to true to get FwupdFirmwareVersionInfo used.
	FwupdFirmwareVersionInfo FwupdFirmwareVersionInfo
	// Version is "UsbBusInfo.version"
	//
	// Optional
	Version UsbVersion
	// SpecSpeed is "UsbBusInfo.spec_speed"
	//
	// Optional
	SpecSpeed UsbSpecSpeed

	FFI_USE_ClassId    bool // for ClassId.
	FFI_USE_SubclassId bool // for SubclassId.
	FFI_USE_ProtocolId bool // for ProtocolId.
	FFI_USE_VendorId   bool // for VendorId.
	FFI_USE_ProductId  bool // for ProductId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UsbBusInfo with all fields set.
func (p UsbBusInfo) FromRef(ref js.Ref) UsbBusInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UsbBusInfo in the application heap.
func (p UsbBusInfo) New() js.Ref {
	return bindings.UsbBusInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UsbBusInfo) UpdateFrom(ref js.Ref) {
	bindings.UsbBusInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UsbBusInfo) Update(ref js.Ref) {
	bindings.UsbBusInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UsbBusInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Interfaces.Ref(),
	)
	p.Interfaces = p.Interfaces.FromRef(js.Undefined)
	if recursive {
		p.FwupdFirmwareVersionInfo.FreeMembers(true)
	}
}

type UsbBusDevices struct {
	// Devices is "UsbBusDevices.devices"
	//
	// Optional
	Devices js.Array[UsbBusInfo]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UsbBusDevices with all fields set.
func (p UsbBusDevices) FromRef(ref js.Ref) UsbBusDevices {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UsbBusDevices in the application heap.
func (p UsbBusDevices) New() js.Ref {
	return bindings.UsbBusDevicesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UsbBusDevices) UpdateFrom(ref js.Ref) {
	bindings.UsbBusDevicesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UsbBusDevices) Update(ref js.Ref) {
	bindings.UsbBusDevicesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UsbBusDevices) FreeMembers(recursive bool) {
	js.Free(
		p.Devices.Ref(),
	)
	p.Devices = p.Devices.FromRef(js.Undefined)
}

type UsbBusDevicesCallbackFunc func(this js.Ref, UsbBusDevices *UsbBusDevices) js.Ref

func (fn UsbBusDevicesCallbackFunc) Register() js.Func[func(UsbBusDevices *UsbBusDevices)] {
	return js.RegisterCallback[func(UsbBusDevices *UsbBusDevices)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UsbBusDevicesCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 UsbBusDevices
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UsbBusDevicesCallback[T any] struct {
	Fn  func(arg T, this js.Ref, UsbBusDevices *UsbBusDevices) js.Ref
	Arg T
}

func (cb *UsbBusDevicesCallback[T]) Register() js.Func[func(UsbBusDevices *UsbBusDevices)] {
	return js.RegisterCallback[func(UsbBusDevices *UsbBusDevices)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UsbBusDevicesCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 UsbBusDevices
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type VpdInfo struct {
	// ActivateDate is "VpdInfo.activateDate"
	//
	// Optional
	ActivateDate js.String
	// ModelName is "VpdInfo.modelName"
	//
	// Optional
	ModelName js.String
	// SerialNumber is "VpdInfo.serialNumber"
	//
	// Optional
	SerialNumber js.String
	// SkuNumber is "VpdInfo.skuNumber"
	//
	// Optional
	SkuNumber js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VpdInfo with all fields set.
func (p VpdInfo) FromRef(ref js.Ref) VpdInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VpdInfo in the application heap.
func (p VpdInfo) New() js.Ref {
	return bindings.VpdInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *VpdInfo) UpdateFrom(ref js.Ref) {
	bindings.VpdInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VpdInfo) Update(ref js.Ref) {
	bindings.VpdInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VpdInfo) FreeMembers(recursive bool) {
	js.Free(
		p.ActivateDate.Ref(),
		p.ModelName.Ref(),
		p.SerialNumber.Ref(),
		p.SkuNumber.Ref(),
	)
	p.ActivateDate = p.ActivateDate.FromRef(js.Undefined)
	p.ModelName = p.ModelName.FromRef(js.Undefined)
	p.SerialNumber = p.SerialNumber.FromRef(js.Undefined)
	p.SkuNumber = p.SkuNumber.FromRef(js.Undefined)
}

type VpdInfoCallbackFunc func(this js.Ref, vpdInfo *VpdInfo) js.Ref

func (fn VpdInfoCallbackFunc) Register() js.Func[func(vpdInfo *VpdInfo)] {
	return js.RegisterCallback[func(vpdInfo *VpdInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn VpdInfoCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 VpdInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type VpdInfoCallback[T any] struct {
	Fn  func(arg T, this js.Ref, vpdInfo *VpdInfo) js.Ref
	Arg T
}

func (cb *VpdInfoCallback[T]) Register() js.Func[func(vpdInfo *VpdInfo)] {
	return js.RegisterCallback[func(vpdInfo *VpdInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VpdInfoCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 VpdInfo
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncGetAudioInfo returns true if the function "WEBEXT.os.telemetry.getAudioInfo" exists.
func HasFuncGetAudioInfo() bool {
	return js.True == bindings.HasFuncGetAudioInfo()
}

// FuncGetAudioInfo returns the function "WEBEXT.os.telemetry.getAudioInfo".
func FuncGetAudioInfo() (fn js.Func[func() js.Promise[AudioInfo]]) {
	bindings.FuncGetAudioInfo(
		js.Pointer(&fn),
	)
	return
}

// GetAudioInfo calls the function "WEBEXT.os.telemetry.getAudioInfo" directly.
func GetAudioInfo() (ret js.Promise[AudioInfo]) {
	bindings.CallGetAudioInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetAudioInfo calls the function "WEBEXT.os.telemetry.getAudioInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAudioInfo() (ret js.Promise[AudioInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAudioInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetBatteryInfo returns true if the function "WEBEXT.os.telemetry.getBatteryInfo" exists.
func HasFuncGetBatteryInfo() bool {
	return js.True == bindings.HasFuncGetBatteryInfo()
}

// FuncGetBatteryInfo returns the function "WEBEXT.os.telemetry.getBatteryInfo".
func FuncGetBatteryInfo() (fn js.Func[func() js.Promise[BatteryInfo]]) {
	bindings.FuncGetBatteryInfo(
		js.Pointer(&fn),
	)
	return
}

// GetBatteryInfo calls the function "WEBEXT.os.telemetry.getBatteryInfo" directly.
func GetBatteryInfo() (ret js.Promise[BatteryInfo]) {
	bindings.CallGetBatteryInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetBatteryInfo calls the function "WEBEXT.os.telemetry.getBatteryInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetBatteryInfo() (ret js.Promise[BatteryInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetBatteryInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetCpuInfo returns true if the function "WEBEXT.os.telemetry.getCpuInfo" exists.
func HasFuncGetCpuInfo() bool {
	return js.True == bindings.HasFuncGetCpuInfo()
}

// FuncGetCpuInfo returns the function "WEBEXT.os.telemetry.getCpuInfo".
func FuncGetCpuInfo() (fn js.Func[func() js.Promise[CpuInfo]]) {
	bindings.FuncGetCpuInfo(
		js.Pointer(&fn),
	)
	return
}

// GetCpuInfo calls the function "WEBEXT.os.telemetry.getCpuInfo" directly.
func GetCpuInfo() (ret js.Promise[CpuInfo]) {
	bindings.CallGetCpuInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetCpuInfo calls the function "WEBEXT.os.telemetry.getCpuInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCpuInfo() (ret js.Promise[CpuInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCpuInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDisplayInfo returns true if the function "WEBEXT.os.telemetry.getDisplayInfo" exists.
func HasFuncGetDisplayInfo() bool {
	return js.True == bindings.HasFuncGetDisplayInfo()
}

// FuncGetDisplayInfo returns the function "WEBEXT.os.telemetry.getDisplayInfo".
func FuncGetDisplayInfo() (fn js.Func[func() js.Promise[DisplayInfo]]) {
	bindings.FuncGetDisplayInfo(
		js.Pointer(&fn),
	)
	return
}

// GetDisplayInfo calls the function "WEBEXT.os.telemetry.getDisplayInfo" directly.
func GetDisplayInfo() (ret js.Promise[DisplayInfo]) {
	bindings.CallGetDisplayInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetDisplayInfo calls the function "WEBEXT.os.telemetry.getDisplayInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDisplayInfo() (ret js.Promise[DisplayInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDisplayInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetInternetConnectivityInfo returns true if the function "WEBEXT.os.telemetry.getInternetConnectivityInfo" exists.
func HasFuncGetInternetConnectivityInfo() bool {
	return js.True == bindings.HasFuncGetInternetConnectivityInfo()
}

// FuncGetInternetConnectivityInfo returns the function "WEBEXT.os.telemetry.getInternetConnectivityInfo".
func FuncGetInternetConnectivityInfo() (fn js.Func[func() js.Promise[InternetConnectivityInfo]]) {
	bindings.FuncGetInternetConnectivityInfo(
		js.Pointer(&fn),
	)
	return
}

// GetInternetConnectivityInfo calls the function "WEBEXT.os.telemetry.getInternetConnectivityInfo" directly.
func GetInternetConnectivityInfo() (ret js.Promise[InternetConnectivityInfo]) {
	bindings.CallGetInternetConnectivityInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetInternetConnectivityInfo calls the function "WEBEXT.os.telemetry.getInternetConnectivityInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetInternetConnectivityInfo() (ret js.Promise[InternetConnectivityInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetInternetConnectivityInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetMarketingInfo returns true if the function "WEBEXT.os.telemetry.getMarketingInfo" exists.
func HasFuncGetMarketingInfo() bool {
	return js.True == bindings.HasFuncGetMarketingInfo()
}

// FuncGetMarketingInfo returns the function "WEBEXT.os.telemetry.getMarketingInfo".
func FuncGetMarketingInfo() (fn js.Func[func() js.Promise[MarketingInfo]]) {
	bindings.FuncGetMarketingInfo(
		js.Pointer(&fn),
	)
	return
}

// GetMarketingInfo calls the function "WEBEXT.os.telemetry.getMarketingInfo" directly.
func GetMarketingInfo() (ret js.Promise[MarketingInfo]) {
	bindings.CallGetMarketingInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetMarketingInfo calls the function "WEBEXT.os.telemetry.getMarketingInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetMarketingInfo() (ret js.Promise[MarketingInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetMarketingInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetMemoryInfo returns true if the function "WEBEXT.os.telemetry.getMemoryInfo" exists.
func HasFuncGetMemoryInfo() bool {
	return js.True == bindings.HasFuncGetMemoryInfo()
}

// FuncGetMemoryInfo returns the function "WEBEXT.os.telemetry.getMemoryInfo".
func FuncGetMemoryInfo() (fn js.Func[func() js.Promise[MemoryInfo]]) {
	bindings.FuncGetMemoryInfo(
		js.Pointer(&fn),
	)
	return
}

// GetMemoryInfo calls the function "WEBEXT.os.telemetry.getMemoryInfo" directly.
func GetMemoryInfo() (ret js.Promise[MemoryInfo]) {
	bindings.CallGetMemoryInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetMemoryInfo calls the function "WEBEXT.os.telemetry.getMemoryInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetMemoryInfo() (ret js.Promise[MemoryInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetMemoryInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetNonRemovableBlockDevicesInfo returns true if the function "WEBEXT.os.telemetry.getNonRemovableBlockDevicesInfo" exists.
func HasFuncGetNonRemovableBlockDevicesInfo() bool {
	return js.True == bindings.HasFuncGetNonRemovableBlockDevicesInfo()
}

// FuncGetNonRemovableBlockDevicesInfo returns the function "WEBEXT.os.telemetry.getNonRemovableBlockDevicesInfo".
func FuncGetNonRemovableBlockDevicesInfo() (fn js.Func[func() js.Promise[NonRemovableBlockDeviceInfoResponse]]) {
	bindings.FuncGetNonRemovableBlockDevicesInfo(
		js.Pointer(&fn),
	)
	return
}

// GetNonRemovableBlockDevicesInfo calls the function "WEBEXT.os.telemetry.getNonRemovableBlockDevicesInfo" directly.
func GetNonRemovableBlockDevicesInfo() (ret js.Promise[NonRemovableBlockDeviceInfoResponse]) {
	bindings.CallGetNonRemovableBlockDevicesInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetNonRemovableBlockDevicesInfo calls the function "WEBEXT.os.telemetry.getNonRemovableBlockDevicesInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetNonRemovableBlockDevicesInfo() (ret js.Promise[NonRemovableBlockDeviceInfoResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetNonRemovableBlockDevicesInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetOemData returns true if the function "WEBEXT.os.telemetry.getOemData" exists.
func HasFuncGetOemData() bool {
	return js.True == bindings.HasFuncGetOemData()
}

// FuncGetOemData returns the function "WEBEXT.os.telemetry.getOemData".
func FuncGetOemData() (fn js.Func[func() js.Promise[OemData]]) {
	bindings.FuncGetOemData(
		js.Pointer(&fn),
	)
	return
}

// GetOemData calls the function "WEBEXT.os.telemetry.getOemData" directly.
func GetOemData() (ret js.Promise[OemData]) {
	bindings.CallGetOemData(
		js.Pointer(&ret),
	)

	return
}

// TryGetOemData calls the function "WEBEXT.os.telemetry.getOemData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetOemData() (ret js.Promise[OemData], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetOemData(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetOsVersionInfo returns true if the function "WEBEXT.os.telemetry.getOsVersionInfo" exists.
func HasFuncGetOsVersionInfo() bool {
	return js.True == bindings.HasFuncGetOsVersionInfo()
}

// FuncGetOsVersionInfo returns the function "WEBEXT.os.telemetry.getOsVersionInfo".
func FuncGetOsVersionInfo() (fn js.Func[func() js.Promise[OsVersionInfo]]) {
	bindings.FuncGetOsVersionInfo(
		js.Pointer(&fn),
	)
	return
}

// GetOsVersionInfo calls the function "WEBEXT.os.telemetry.getOsVersionInfo" directly.
func GetOsVersionInfo() (ret js.Promise[OsVersionInfo]) {
	bindings.CallGetOsVersionInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetOsVersionInfo calls the function "WEBEXT.os.telemetry.getOsVersionInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetOsVersionInfo() (ret js.Promise[OsVersionInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetOsVersionInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetStatefulPartitionInfo returns true if the function "WEBEXT.os.telemetry.getStatefulPartitionInfo" exists.
func HasFuncGetStatefulPartitionInfo() bool {
	return js.True == bindings.HasFuncGetStatefulPartitionInfo()
}

// FuncGetStatefulPartitionInfo returns the function "WEBEXT.os.telemetry.getStatefulPartitionInfo".
func FuncGetStatefulPartitionInfo() (fn js.Func[func() js.Promise[StatefulPartitionInfo]]) {
	bindings.FuncGetStatefulPartitionInfo(
		js.Pointer(&fn),
	)
	return
}

// GetStatefulPartitionInfo calls the function "WEBEXT.os.telemetry.getStatefulPartitionInfo" directly.
func GetStatefulPartitionInfo() (ret js.Promise[StatefulPartitionInfo]) {
	bindings.CallGetStatefulPartitionInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetStatefulPartitionInfo calls the function "WEBEXT.os.telemetry.getStatefulPartitionInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetStatefulPartitionInfo() (ret js.Promise[StatefulPartitionInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetStatefulPartitionInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetTpmInfo returns true if the function "WEBEXT.os.telemetry.getTpmInfo" exists.
func HasFuncGetTpmInfo() bool {
	return js.True == bindings.HasFuncGetTpmInfo()
}

// FuncGetTpmInfo returns the function "WEBEXT.os.telemetry.getTpmInfo".
func FuncGetTpmInfo() (fn js.Func[func() js.Promise[TpmInfo]]) {
	bindings.FuncGetTpmInfo(
		js.Pointer(&fn),
	)
	return
}

// GetTpmInfo calls the function "WEBEXT.os.telemetry.getTpmInfo" directly.
func GetTpmInfo() (ret js.Promise[TpmInfo]) {
	bindings.CallGetTpmInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetTpmInfo calls the function "WEBEXT.os.telemetry.getTpmInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetTpmInfo() (ret js.Promise[TpmInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetTpmInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetUsbBusInfo returns true if the function "WEBEXT.os.telemetry.getUsbBusInfo" exists.
func HasFuncGetUsbBusInfo() bool {
	return js.True == bindings.HasFuncGetUsbBusInfo()
}

// FuncGetUsbBusInfo returns the function "WEBEXT.os.telemetry.getUsbBusInfo".
func FuncGetUsbBusInfo() (fn js.Func[func() js.Promise[UsbBusDevices]]) {
	bindings.FuncGetUsbBusInfo(
		js.Pointer(&fn),
	)
	return
}

// GetUsbBusInfo calls the function "WEBEXT.os.telemetry.getUsbBusInfo" directly.
func GetUsbBusInfo() (ret js.Promise[UsbBusDevices]) {
	bindings.CallGetUsbBusInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetUsbBusInfo calls the function "WEBEXT.os.telemetry.getUsbBusInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetUsbBusInfo() (ret js.Promise[UsbBusDevices], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetUsbBusInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetVpdInfo returns true if the function "WEBEXT.os.telemetry.getVpdInfo" exists.
func HasFuncGetVpdInfo() bool {
	return js.True == bindings.HasFuncGetVpdInfo()
}

// FuncGetVpdInfo returns the function "WEBEXT.os.telemetry.getVpdInfo".
func FuncGetVpdInfo() (fn js.Func[func() js.Promise[VpdInfo]]) {
	bindings.FuncGetVpdInfo(
		js.Pointer(&fn),
	)
	return
}

// GetVpdInfo calls the function "WEBEXT.os.telemetry.getVpdInfo" directly.
func GetVpdInfo() (ret js.Promise[VpdInfo]) {
	bindings.CallGetVpdInfo(
		js.Pointer(&ret),
	)

	return
}

// TryGetVpdInfo calls the function "WEBEXT.os.telemetry.getVpdInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetVpdInfo() (ret js.Promise[VpdInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetVpdInfo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

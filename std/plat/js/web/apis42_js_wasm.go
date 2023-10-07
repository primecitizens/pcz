// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

const (
	_ MLOperandType = iota

	MLOperandType_FLOAT32
	MLOperandType_FLOAT16
	MLOperandType_INT32
	MLOperandType_UINT32
	MLOperandType_INT8
	MLOperandType_UINT8
)

func (MLOperandType) FromRef(str js.Ref) MLOperandType {
	return MLOperandType(bindings.ConstOfMLOperandType(str))
}

func (x MLOperandType) String() (string, bool) {
	switch x {
	case MLOperandType_FLOAT32:
		return "float32", true
	case MLOperandType_FLOAT16:
		return "float16", true
	case MLOperandType_INT32:
		return "int32", true
	case MLOperandType_UINT32:
		return "uint32", true
	case MLOperandType_INT8:
		return "int8", true
	case MLOperandType_UINT8:
		return "uint8", true
	default:
		return "", false
	}
}

type MLOperandDescriptor struct {
	// Type is "MLOperandDescriptor.type"
	//
	// Required
	Type MLOperandType
	// Dimensions is "MLOperandDescriptor.dimensions"
	//
	// Optional
	Dimensions js.Array[uint32]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLOperandDescriptor with all fields set.
func (p MLOperandDescriptor) FromRef(ref js.Ref) MLOperandDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLOperandDescriptor in the application heap.
func (p MLOperandDescriptor) New() js.Ref {
	return bindings.MLOperandDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLOperandDescriptor) UpdateFrom(ref js.Ref) {
	bindings.MLOperandDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLOperandDescriptor) Update(ref js.Ref) {
	bindings.MLOperandDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLOperandDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Dimensions.Ref(),
	)
	p.Dimensions = p.Dimensions.FromRef(js.Undefined)
}

type MLNamedOperands = js.Record[MLOperand]

type MLHardSigmoidOptions struct {
	// Alpha is "MLHardSigmoidOptions.alpha"
	//
	// Optional, defaults to 0.2.
	//
	// NOTE: FFI_USE_Alpha MUST be set to true to make this field effective.
	Alpha float32
	// Beta is "MLHardSigmoidOptions.beta"
	//
	// Optional, defaults to 0.5.
	//
	// NOTE: FFI_USE_Beta MUST be set to true to make this field effective.
	Beta float32

	FFI_USE_Alpha bool // for Alpha.
	FFI_USE_Beta  bool // for Beta.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLHardSigmoidOptions with all fields set.
func (p MLHardSigmoidOptions) FromRef(ref js.Ref) MLHardSigmoidOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLHardSigmoidOptions in the application heap.
func (p MLHardSigmoidOptions) New() js.Ref {
	return bindings.MLHardSigmoidOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLHardSigmoidOptions) UpdateFrom(ref js.Ref) {
	bindings.MLHardSigmoidOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLHardSigmoidOptions) Update(ref js.Ref) {
	bindings.MLHardSigmoidOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLHardSigmoidOptions) FreeMembers(recursive bool) {
}

type MLGruWeightLayout uint32

const (
	_ MLGruWeightLayout = iota

	MLGruWeightLayout_ZRN
	MLGruWeightLayout_RZN
)

func (MLGruWeightLayout) FromRef(str js.Ref) MLGruWeightLayout {
	return MLGruWeightLayout(bindings.ConstOfMLGruWeightLayout(str))
}

func (x MLGruWeightLayout) String() (string, bool) {
	switch x {
	case MLGruWeightLayout_ZRN:
		return "zrn", true
	case MLGruWeightLayout_RZN:
		return "rzn", true
	default:
		return "", false
	}
}

type MLGruCellOptions struct {
	// Bias is "MLGruCellOptions.bias"
	//
	// Optional
	Bias MLOperand
	// RecurrentBias is "MLGruCellOptions.recurrentBias"
	//
	// Optional
	RecurrentBias MLOperand
	// ResetAfter is "MLGruCellOptions.resetAfter"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_ResetAfter MUST be set to true to make this field effective.
	ResetAfter bool
	// Layout is "MLGruCellOptions.layout"
	//
	// Optional, defaults to "zrn".
	Layout MLGruWeightLayout
	// Activations is "MLGruCellOptions.activations"
	//
	// Optional
	Activations js.Array[MLActivation]

	FFI_USE_ResetAfter bool // for ResetAfter.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLGruCellOptions with all fields set.
func (p MLGruCellOptions) FromRef(ref js.Ref) MLGruCellOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLGruCellOptions in the application heap.
func (p MLGruCellOptions) New() js.Ref {
	return bindings.MLGruCellOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLGruCellOptions) UpdateFrom(ref js.Ref) {
	bindings.MLGruCellOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLGruCellOptions) Update(ref js.Ref) {
	bindings.MLGruCellOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLGruCellOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Bias.Ref(),
		p.RecurrentBias.Ref(),
		p.Activations.Ref(),
	)
	p.Bias = p.Bias.FromRef(js.Undefined)
	p.RecurrentBias = p.RecurrentBias.FromRef(js.Undefined)
	p.Activations = p.Activations.FromRef(js.Undefined)
}

type MLRoundingType uint32

const (
	_ MLRoundingType = iota

	MLRoundingType_FLOOR
	MLRoundingType_CEIL
)

func (MLRoundingType) FromRef(str js.Ref) MLRoundingType {
	return MLRoundingType(bindings.ConstOfMLRoundingType(str))
}

func (x MLRoundingType) String() (string, bool) {
	switch x {
	case MLRoundingType_FLOOR:
		return "floor", true
	case MLRoundingType_CEIL:
		return "ceil", true
	default:
		return "", false
	}
}

type MLPool2dOptions struct {
	// WindowDimensions is "MLPool2dOptions.windowDimensions"
	//
	// Optional
	WindowDimensions js.Array[uint32]
	// Padding is "MLPool2dOptions.padding"
	//
	// Optional
	Padding js.Array[uint32]
	// Strides is "MLPool2dOptions.strides"
	//
	// Optional
	Strides js.Array[uint32]
	// Dilations is "MLPool2dOptions.dilations"
	//
	// Optional
	Dilations js.Array[uint32]
	// AutoPad is "MLPool2dOptions.autoPad"
	//
	// Optional, defaults to "explicit".
	AutoPad MLAutoPad
	// Layout is "MLPool2dOptions.layout"
	//
	// Optional, defaults to "nchw".
	Layout MLInputOperandLayout
	// RoundingType is "MLPool2dOptions.roundingType"
	//
	// Optional, defaults to "floor".
	RoundingType MLRoundingType
	// OutputSizes is "MLPool2dOptions.outputSizes"
	//
	// Optional
	OutputSizes js.Array[uint32]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLPool2dOptions with all fields set.
func (p MLPool2dOptions) FromRef(ref js.Ref) MLPool2dOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLPool2dOptions in the application heap.
func (p MLPool2dOptions) New() js.Ref {
	return bindings.MLPool2dOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLPool2dOptions) UpdateFrom(ref js.Ref) {
	bindings.MLPool2dOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLPool2dOptions) Update(ref js.Ref) {
	bindings.MLPool2dOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLPool2dOptions) FreeMembers(recursive bool) {
	js.Free(
		p.WindowDimensions.Ref(),
		p.Padding.Ref(),
		p.Strides.Ref(),
		p.Dilations.Ref(),
		p.OutputSizes.Ref(),
	)
	p.WindowDimensions = p.WindowDimensions.FromRef(js.Undefined)
	p.Padding = p.Padding.FromRef(js.Undefined)
	p.Strides = p.Strides.FromRef(js.Undefined)
	p.Dilations = p.Dilations.FromRef(js.Undefined)
	p.OutputSizes = p.OutputSizes.FromRef(js.Undefined)
}

type MLLinearOptions struct {
	// Alpha is "MLLinearOptions.alpha"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_Alpha MUST be set to true to make this field effective.
	Alpha float32
	// Beta is "MLLinearOptions.beta"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Beta MUST be set to true to make this field effective.
	Beta float32

	FFI_USE_Alpha bool // for Alpha.
	FFI_USE_Beta  bool // for Beta.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLLinearOptions with all fields set.
func (p MLLinearOptions) FromRef(ref js.Ref) MLLinearOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLLinearOptions in the application heap.
func (p MLLinearOptions) New() js.Ref {
	return bindings.MLLinearOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLLinearOptions) UpdateFrom(ref js.Ref) {
	bindings.MLLinearOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLLinearOptions) Update(ref js.Ref) {
	bindings.MLLinearOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLLinearOptions) FreeMembers(recursive bool) {
}

type MLLeakyReluOptions struct {
	// Alpha is "MLLeakyReluOptions.alpha"
	//
	// Optional, defaults to 0.01.
	//
	// NOTE: FFI_USE_Alpha MUST be set to true to make this field effective.
	Alpha float32

	FFI_USE_Alpha bool // for Alpha.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLLeakyReluOptions with all fields set.
func (p MLLeakyReluOptions) FromRef(ref js.Ref) MLLeakyReluOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLLeakyReluOptions in the application heap.
func (p MLLeakyReluOptions) New() js.Ref {
	return bindings.MLLeakyReluOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLLeakyReluOptions) UpdateFrom(ref js.Ref) {
	bindings.MLLeakyReluOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLLeakyReluOptions) Update(ref js.Ref) {
	bindings.MLLeakyReluOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLLeakyReluOptions) FreeMembers(recursive bool) {
}

type MLPaddingMode uint32

const (
	_ MLPaddingMode = iota

	MLPaddingMode_CONSTANT
	MLPaddingMode_EDGE
	MLPaddingMode_REFLECTION
	MLPaddingMode_SYMMETRIC
)

func (MLPaddingMode) FromRef(str js.Ref) MLPaddingMode {
	return MLPaddingMode(bindings.ConstOfMLPaddingMode(str))
}

func (x MLPaddingMode) String() (string, bool) {
	switch x {
	case MLPaddingMode_CONSTANT:
		return "constant", true
	case MLPaddingMode_EDGE:
		return "edge", true
	case MLPaddingMode_REFLECTION:
		return "reflection", true
	case MLPaddingMode_SYMMETRIC:
		return "symmetric", true
	default:
		return "", false
	}
}

type MLPadOptions struct {
	// Mode is "MLPadOptions.mode"
	//
	// Optional, defaults to "constant".
	Mode MLPaddingMode
	// Value is "MLPadOptions.value"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Value MUST be set to true to make this field effective.
	Value float32

	FFI_USE_Value bool // for Value.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLPadOptions with all fields set.
func (p MLPadOptions) FromRef(ref js.Ref) MLPadOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLPadOptions in the application heap.
func (p MLPadOptions) New() js.Ref {
	return bindings.MLPadOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLPadOptions) UpdateFrom(ref js.Ref) {
	bindings.MLPadOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLPadOptions) Update(ref js.Ref) {
	bindings.MLPadOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLPadOptions) FreeMembers(recursive bool) {
}

type MLInstanceNormalizationOptions struct {
	// Scale is "MLInstanceNormalizationOptions.scale"
	//
	// Optional
	Scale MLOperand
	// Bias is "MLInstanceNormalizationOptions.bias"
	//
	// Optional
	Bias MLOperand
	// Epsilon is "MLInstanceNormalizationOptions.epsilon"
	//
	// Optional, defaults to 1e-5.
	//
	// NOTE: FFI_USE_Epsilon MUST be set to true to make this field effective.
	Epsilon float32
	// Layout is "MLInstanceNormalizationOptions.layout"
	//
	// Optional, defaults to "nchw".
	Layout MLInputOperandLayout

	FFI_USE_Epsilon bool // for Epsilon.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLInstanceNormalizationOptions with all fields set.
func (p MLInstanceNormalizationOptions) FromRef(ref js.Ref) MLInstanceNormalizationOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLInstanceNormalizationOptions in the application heap.
func (p MLInstanceNormalizationOptions) New() js.Ref {
	return bindings.MLInstanceNormalizationOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLInstanceNormalizationOptions) UpdateFrom(ref js.Ref) {
	bindings.MLInstanceNormalizationOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLInstanceNormalizationOptions) Update(ref js.Ref) {
	bindings.MLInstanceNormalizationOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLInstanceNormalizationOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Scale.Ref(),
		p.Bias.Ref(),
	)
	p.Scale = p.Scale.FromRef(js.Undefined)
	p.Bias = p.Bias.FromRef(js.Undefined)
}

type MLSoftplusOptions struct {
	// Steepness is "MLSoftplusOptions.steepness"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_Steepness MUST be set to true to make this field effective.
	Steepness float32

	FFI_USE_Steepness bool // for Steepness.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLSoftplusOptions with all fields set.
func (p MLSoftplusOptions) FromRef(ref js.Ref) MLSoftplusOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLSoftplusOptions in the application heap.
func (p MLSoftplusOptions) New() js.Ref {
	return bindings.MLSoftplusOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLSoftplusOptions) UpdateFrom(ref js.Ref) {
	bindings.MLSoftplusOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLSoftplusOptions) Update(ref js.Ref) {
	bindings.MLSoftplusOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLSoftplusOptions) FreeMembers(recursive bool) {
}

type MLSplitOptions struct {
	// Axis is "MLSplitOptions.axis"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Axis MUST be set to true to make this field effective.
	Axis uint32

	FFI_USE_Axis bool // for Axis.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLSplitOptions with all fields set.
func (p MLSplitOptions) FromRef(ref js.Ref) MLSplitOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLSplitOptions in the application heap.
func (p MLSplitOptions) New() js.Ref {
	return bindings.MLSplitOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLSplitOptions) UpdateFrom(ref js.Ref) {
	bindings.MLSplitOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLSplitOptions) Update(ref js.Ref) {
	bindings.MLSplitOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLSplitOptions) FreeMembers(recursive bool) {
}

type MLInterpolationMode uint32

const (
	_ MLInterpolationMode = iota

	MLInterpolationMode_NEAREST_NEIGHBOR
	MLInterpolationMode_LINEAR
)

func (MLInterpolationMode) FromRef(str js.Ref) MLInterpolationMode {
	return MLInterpolationMode(bindings.ConstOfMLInterpolationMode(str))
}

func (x MLInterpolationMode) String() (string, bool) {
	switch x {
	case MLInterpolationMode_NEAREST_NEIGHBOR:
		return "nearest-neighbor", true
	case MLInterpolationMode_LINEAR:
		return "linear", true
	default:
		return "", false
	}
}

type MLResample2dOptions struct {
	// Mode is "MLResample2dOptions.mode"
	//
	// Optional, defaults to "nearest-neighbor".
	Mode MLInterpolationMode
	// Scales is "MLResample2dOptions.scales"
	//
	// Optional
	Scales js.Array[float32]
	// Sizes is "MLResample2dOptions.sizes"
	//
	// Optional
	Sizes js.Array[uint32]
	// Axes is "MLResample2dOptions.axes"
	//
	// Optional
	Axes js.Array[uint32]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLResample2dOptions with all fields set.
func (p MLResample2dOptions) FromRef(ref js.Ref) MLResample2dOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLResample2dOptions in the application heap.
func (p MLResample2dOptions) New() js.Ref {
	return bindings.MLResample2dOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLResample2dOptions) UpdateFrom(ref js.Ref) {
	bindings.MLResample2dOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLResample2dOptions) Update(ref js.Ref) {
	bindings.MLResample2dOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLResample2dOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Scales.Ref(),
		p.Sizes.Ref(),
		p.Axes.Ref(),
	)
	p.Scales = p.Scales.FromRef(js.Undefined)
	p.Sizes = p.Sizes.FromRef(js.Undefined)
	p.Axes = p.Axes.FromRef(js.Undefined)
}

type MLReduceOptions struct {
	// Axes is "MLReduceOptions.axes"
	//
	// Optional, defaults to null.
	Axes js.Array[uint32]
	// KeepDimensions is "MLReduceOptions.keepDimensions"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_KeepDimensions MUST be set to true to make this field effective.
	KeepDimensions bool

	FFI_USE_KeepDimensions bool // for KeepDimensions.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLReduceOptions with all fields set.
func (p MLReduceOptions) FromRef(ref js.Ref) MLReduceOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLReduceOptions in the application heap.
func (p MLReduceOptions) New() js.Ref {
	return bindings.MLReduceOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLReduceOptions) UpdateFrom(ref js.Ref) {
	bindings.MLReduceOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLReduceOptions) Update(ref js.Ref) {
	bindings.MLReduceOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLReduceOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Axes.Ref(),
	)
	p.Axes = p.Axes.FromRef(js.Undefined)
}

type MLRecurrentNetworkDirection uint32

const (
	_ MLRecurrentNetworkDirection = iota

	MLRecurrentNetworkDirection_FORWARD
	MLRecurrentNetworkDirection_BACKWARD
	MLRecurrentNetworkDirection_BOTH
)

func (MLRecurrentNetworkDirection) FromRef(str js.Ref) MLRecurrentNetworkDirection {
	return MLRecurrentNetworkDirection(bindings.ConstOfMLRecurrentNetworkDirection(str))
}

func (x MLRecurrentNetworkDirection) String() (string, bool) {
	switch x {
	case MLRecurrentNetworkDirection_FORWARD:
		return "forward", true
	case MLRecurrentNetworkDirection_BACKWARD:
		return "backward", true
	case MLRecurrentNetworkDirection_BOTH:
		return "both", true
	default:
		return "", false
	}
}

type MLLstmWeightLayout uint32

const (
	_ MLLstmWeightLayout = iota

	MLLstmWeightLayout_IOFG
	MLLstmWeightLayout_IFGO
)

func (MLLstmWeightLayout) FromRef(str js.Ref) MLLstmWeightLayout {
	return MLLstmWeightLayout(bindings.ConstOfMLLstmWeightLayout(str))
}

func (x MLLstmWeightLayout) String() (string, bool) {
	switch x {
	case MLLstmWeightLayout_IOFG:
		return "iofg", true
	case MLLstmWeightLayout_IFGO:
		return "ifgo", true
	default:
		return "", false
	}
}

type MLLstmOptions struct {
	// Bias is "MLLstmOptions.bias"
	//
	// Optional
	Bias MLOperand
	// RecurrentBias is "MLLstmOptions.recurrentBias"
	//
	// Optional
	RecurrentBias MLOperand
	// PeepholeWeight is "MLLstmOptions.peepholeWeight"
	//
	// Optional
	PeepholeWeight MLOperand
	// InitialHiddenState is "MLLstmOptions.initialHiddenState"
	//
	// Optional
	InitialHiddenState MLOperand
	// InitialCellState is "MLLstmOptions.initialCellState"
	//
	// Optional
	InitialCellState MLOperand
	// ReturnSequence is "MLLstmOptions.returnSequence"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ReturnSequence MUST be set to true to make this field effective.
	ReturnSequence bool
	// Direction is "MLLstmOptions.direction"
	//
	// Optional, defaults to "forward".
	Direction MLRecurrentNetworkDirection
	// Layout is "MLLstmOptions.layout"
	//
	// Optional, defaults to "iofg".
	Layout MLLstmWeightLayout
	// Activations is "MLLstmOptions.activations"
	//
	// Optional
	Activations js.Array[MLActivation]

	FFI_USE_ReturnSequence bool // for ReturnSequence.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLLstmOptions with all fields set.
func (p MLLstmOptions) FromRef(ref js.Ref) MLLstmOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLLstmOptions in the application heap.
func (p MLLstmOptions) New() js.Ref {
	return bindings.MLLstmOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLLstmOptions) UpdateFrom(ref js.Ref) {
	bindings.MLLstmOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLLstmOptions) Update(ref js.Ref) {
	bindings.MLLstmOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLLstmOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Bias.Ref(),
		p.RecurrentBias.Ref(),
		p.PeepholeWeight.Ref(),
		p.InitialHiddenState.Ref(),
		p.InitialCellState.Ref(),
		p.Activations.Ref(),
	)
	p.Bias = p.Bias.FromRef(js.Undefined)
	p.RecurrentBias = p.RecurrentBias.FromRef(js.Undefined)
	p.PeepholeWeight = p.PeepholeWeight.FromRef(js.Undefined)
	p.InitialHiddenState = p.InitialHiddenState.FromRef(js.Undefined)
	p.InitialCellState = p.InitialCellState.FromRef(js.Undefined)
	p.Activations = p.Activations.FromRef(js.Undefined)
}

type MLSqueezeOptions struct {
	// Axes is "MLSqueezeOptions.axes"
	//
	// Optional
	Axes js.Array[uint32]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLSqueezeOptions with all fields set.
func (p MLSqueezeOptions) FromRef(ref js.Ref) MLSqueezeOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLSqueezeOptions in the application heap.
func (p MLSqueezeOptions) New() js.Ref {
	return bindings.MLSqueezeOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLSqueezeOptions) UpdateFrom(ref js.Ref) {
	bindings.MLSqueezeOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLSqueezeOptions) Update(ref js.Ref) {
	bindings.MLSqueezeOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLSqueezeOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Axes.Ref(),
	)
	p.Axes = p.Axes.FromRef(js.Undefined)
}

type MLGruOptions struct {
	// Bias is "MLGruOptions.bias"
	//
	// Optional
	Bias MLOperand
	// RecurrentBias is "MLGruOptions.recurrentBias"
	//
	// Optional
	RecurrentBias MLOperand
	// InitialHiddenState is "MLGruOptions.initialHiddenState"
	//
	// Optional
	InitialHiddenState MLOperand
	// ResetAfter is "MLGruOptions.resetAfter"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_ResetAfter MUST be set to true to make this field effective.
	ResetAfter bool
	// ReturnSequence is "MLGruOptions.returnSequence"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ReturnSequence MUST be set to true to make this field effective.
	ReturnSequence bool
	// Direction is "MLGruOptions.direction"
	//
	// Optional, defaults to "forward".
	Direction MLRecurrentNetworkDirection
	// Layout is "MLGruOptions.layout"
	//
	// Optional, defaults to "zrn".
	Layout MLGruWeightLayout
	// Activations is "MLGruOptions.activations"
	//
	// Optional
	Activations js.Array[MLActivation]

	FFI_USE_ResetAfter     bool // for ResetAfter.
	FFI_USE_ReturnSequence bool // for ReturnSequence.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLGruOptions with all fields set.
func (p MLGruOptions) FromRef(ref js.Ref) MLGruOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLGruOptions in the application heap.
func (p MLGruOptions) New() js.Ref {
	return bindings.MLGruOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLGruOptions) UpdateFrom(ref js.Ref) {
	bindings.MLGruOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLGruOptions) Update(ref js.Ref) {
	bindings.MLGruOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLGruOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Bias.Ref(),
		p.RecurrentBias.Ref(),
		p.InitialHiddenState.Ref(),
		p.Activations.Ref(),
	)
	p.Bias = p.Bias.FromRef(js.Undefined)
	p.RecurrentBias = p.RecurrentBias.FromRef(js.Undefined)
	p.InitialHiddenState = p.InitialHiddenState.FromRef(js.Undefined)
	p.Activations = p.Activations.FromRef(js.Undefined)
}

type MLTransposeOptions struct {
	// Permutation is "MLTransposeOptions.permutation"
	//
	// Optional
	Permutation js.Array[uint32]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLTransposeOptions with all fields set.
func (p MLTransposeOptions) FromRef(ref js.Ref) MLTransposeOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLTransposeOptions in the application heap.
func (p MLTransposeOptions) New() js.Ref {
	return bindings.MLTransposeOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLTransposeOptions) UpdateFrom(ref js.Ref) {
	bindings.MLTransposeOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLTransposeOptions) Update(ref js.Ref) {
	bindings.MLTransposeOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLTransposeOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Permutation.Ref(),
	)
	p.Permutation = p.Permutation.FromRef(js.Undefined)
}

type MLLstmCellOptions struct {
	// Bias is "MLLstmCellOptions.bias"
	//
	// Optional
	Bias MLOperand
	// RecurrentBias is "MLLstmCellOptions.recurrentBias"
	//
	// Optional
	RecurrentBias MLOperand
	// PeepholeWeight is "MLLstmCellOptions.peepholeWeight"
	//
	// Optional
	PeepholeWeight MLOperand
	// Layout is "MLLstmCellOptions.layout"
	//
	// Optional, defaults to "iofg".
	Layout MLLstmWeightLayout
	// Activations is "MLLstmCellOptions.activations"
	//
	// Optional
	Activations js.Array[MLActivation]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLLstmCellOptions with all fields set.
func (p MLLstmCellOptions) FromRef(ref js.Ref) MLLstmCellOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MLLstmCellOptions in the application heap.
func (p MLLstmCellOptions) New() js.Ref {
	return bindings.MLLstmCellOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MLLstmCellOptions) UpdateFrom(ref js.Ref) {
	bindings.MLLstmCellOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MLLstmCellOptions) Update(ref js.Ref) {
	bindings.MLLstmCellOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MLLstmCellOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Bias.Ref(),
		p.RecurrentBias.Ref(),
		p.PeepholeWeight.Ref(),
		p.Activations.Ref(),
	)
	p.Bias = p.Bias.FromRef(js.Undefined)
	p.RecurrentBias = p.RecurrentBias.FromRef(js.Undefined)
	p.PeepholeWeight = p.PeepholeWeight.FromRef(js.Undefined)
	p.Activations = p.Activations.FromRef(js.Undefined)
}

func NewMLGraphBuilder(context MLContext) (ret MLGraphBuilder) {
	ret.ref = bindings.NewMLGraphBuilderByMLGraphBuilder(
		context.Ref())
	return
}

type MLGraphBuilder struct {
	ref js.Ref
}

func (this MLGraphBuilder) Once() MLGraphBuilder {
	this.ref.Once()
	return this
}

func (this MLGraphBuilder) Ref() js.Ref {
	return this.ref
}

func (this MLGraphBuilder) FromRef(ref js.Ref) MLGraphBuilder {
	this.ref = ref
	return this
}

func (this MLGraphBuilder) Free() {
	this.ref.Free()
}

// HasFuncInput returns true if the method "MLGraphBuilder.input" exists.
func (this MLGraphBuilder) HasFuncInput() bool {
	return js.True == bindings.HasFuncMLGraphBuilderInput(
		this.ref,
	)
}

// FuncInput returns the method "MLGraphBuilder.input".
func (this MLGraphBuilder) FuncInput() (fn js.Func[func(name js.String, descriptor MLOperandDescriptor) MLOperand]) {
	bindings.FuncMLGraphBuilderInput(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Input calls the method "MLGraphBuilder.input".
func (this MLGraphBuilder) Input(name js.String, descriptor MLOperandDescriptor) (ret MLOperand) {
	bindings.CallMLGraphBuilderInput(
		this.ref, js.Pointer(&ret),
		name.Ref(),
		js.Pointer(&descriptor),
	)

	return
}

// TryInput calls the method "MLGraphBuilder.input"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryInput(name js.String, descriptor MLOperandDescriptor) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderInput(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncConstant returns true if the method "MLGraphBuilder.constant" exists.
func (this MLGraphBuilder) HasFuncConstant() bool {
	return js.True == bindings.HasFuncMLGraphBuilderConstant(
		this.ref,
	)
}

// FuncConstant returns the method "MLGraphBuilder.constant".
func (this MLGraphBuilder) FuncConstant() (fn js.Func[func(descriptor MLOperandDescriptor, bufferView MLBufferView) MLOperand]) {
	bindings.FuncMLGraphBuilderConstant(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Constant calls the method "MLGraphBuilder.constant".
func (this MLGraphBuilder) Constant(descriptor MLOperandDescriptor, bufferView MLBufferView) (ret MLOperand) {
	bindings.CallMLGraphBuilderConstant(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
		bufferView.Ref(),
	)

	return
}

// TryConstant calls the method "MLGraphBuilder.constant"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryConstant(descriptor MLOperandDescriptor, bufferView MLBufferView) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderConstant(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
		bufferView.Ref(),
	)

	return
}

// HasFuncConstant1 returns true if the method "MLGraphBuilder.constant" exists.
func (this MLGraphBuilder) HasFuncConstant1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderConstant1(
		this.ref,
	)
}

// FuncConstant1 returns the method "MLGraphBuilder.constant".
func (this MLGraphBuilder) FuncConstant1() (fn js.Func[func(value float64, typ MLOperandType) MLOperand]) {
	bindings.FuncMLGraphBuilderConstant1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Constant1 calls the method "MLGraphBuilder.constant".
func (this MLGraphBuilder) Constant1(value float64, typ MLOperandType) (ret MLOperand) {
	bindings.CallMLGraphBuilderConstant1(
		this.ref, js.Pointer(&ret),
		float64(value),
		uint32(typ),
	)

	return
}

// TryConstant1 calls the method "MLGraphBuilder.constant"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryConstant1(value float64, typ MLOperandType) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderConstant1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
		uint32(typ),
	)

	return
}

// HasFuncConstant2 returns true if the method "MLGraphBuilder.constant" exists.
func (this MLGraphBuilder) HasFuncConstant2() bool {
	return js.True == bindings.HasFuncMLGraphBuilderConstant2(
		this.ref,
	)
}

// FuncConstant2 returns the method "MLGraphBuilder.constant".
func (this MLGraphBuilder) FuncConstant2() (fn js.Func[func(value float64) MLOperand]) {
	bindings.FuncMLGraphBuilderConstant2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Constant2 calls the method "MLGraphBuilder.constant".
func (this MLGraphBuilder) Constant2(value float64) (ret MLOperand) {
	bindings.CallMLGraphBuilderConstant2(
		this.ref, js.Pointer(&ret),
		float64(value),
	)

	return
}

// TryConstant2 calls the method "MLGraphBuilder.constant"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryConstant2(value float64) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderConstant2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)

	return
}

// HasFuncBuild returns true if the method "MLGraphBuilder.build" exists.
func (this MLGraphBuilder) HasFuncBuild() bool {
	return js.True == bindings.HasFuncMLGraphBuilderBuild(
		this.ref,
	)
}

// FuncBuild returns the method "MLGraphBuilder.build".
func (this MLGraphBuilder) FuncBuild() (fn js.Func[func(outputs MLNamedOperands) js.Promise[MLGraph]]) {
	bindings.FuncMLGraphBuilderBuild(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Build calls the method "MLGraphBuilder.build".
func (this MLGraphBuilder) Build(outputs MLNamedOperands) (ret js.Promise[MLGraph]) {
	bindings.CallMLGraphBuilderBuild(
		this.ref, js.Pointer(&ret),
		outputs.Ref(),
	)

	return
}

// TryBuild calls the method "MLGraphBuilder.build"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryBuild(outputs MLNamedOperands) (ret js.Promise[MLGraph], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderBuild(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		outputs.Ref(),
	)

	return
}

// HasFuncBuildSync returns true if the method "MLGraphBuilder.buildSync" exists.
func (this MLGraphBuilder) HasFuncBuildSync() bool {
	return js.True == bindings.HasFuncMLGraphBuilderBuildSync(
		this.ref,
	)
}

// FuncBuildSync returns the method "MLGraphBuilder.buildSync".
func (this MLGraphBuilder) FuncBuildSync() (fn js.Func[func(outputs MLNamedOperands) MLGraph]) {
	bindings.FuncMLGraphBuilderBuildSync(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BuildSync calls the method "MLGraphBuilder.buildSync".
func (this MLGraphBuilder) BuildSync(outputs MLNamedOperands) (ret MLGraph) {
	bindings.CallMLGraphBuilderBuildSync(
		this.ref, js.Pointer(&ret),
		outputs.Ref(),
	)

	return
}

// TryBuildSync calls the method "MLGraphBuilder.buildSync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryBuildSync(outputs MLNamedOperands) (ret MLGraph, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderBuildSync(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		outputs.Ref(),
	)

	return
}

// HasFuncHardSigmoid returns true if the method "MLGraphBuilder.hardSigmoid" exists.
func (this MLGraphBuilder) HasFuncHardSigmoid() bool {
	return js.True == bindings.HasFuncMLGraphBuilderHardSigmoid(
		this.ref,
	)
}

// FuncHardSigmoid returns the method "MLGraphBuilder.hardSigmoid".
func (this MLGraphBuilder) FuncHardSigmoid() (fn js.Func[func(input MLOperand, options MLHardSigmoidOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderHardSigmoid(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HardSigmoid calls the method "MLGraphBuilder.hardSigmoid".
func (this MLGraphBuilder) HardSigmoid(input MLOperand, options MLHardSigmoidOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderHardSigmoid(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryHardSigmoid calls the method "MLGraphBuilder.hardSigmoid"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryHardSigmoid(input MLOperand, options MLHardSigmoidOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderHardSigmoid(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncHardSigmoid1 returns true if the method "MLGraphBuilder.hardSigmoid" exists.
func (this MLGraphBuilder) HasFuncHardSigmoid1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderHardSigmoid1(
		this.ref,
	)
}

// FuncHardSigmoid1 returns the method "MLGraphBuilder.hardSigmoid".
func (this MLGraphBuilder) FuncHardSigmoid1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderHardSigmoid1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HardSigmoid1 calls the method "MLGraphBuilder.hardSigmoid".
func (this MLGraphBuilder) HardSigmoid1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderHardSigmoid1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryHardSigmoid1 calls the method "MLGraphBuilder.hardSigmoid"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryHardSigmoid1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderHardSigmoid1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncHardSigmoid2 returns true if the method "MLGraphBuilder.hardSigmoid" exists.
func (this MLGraphBuilder) HasFuncHardSigmoid2() bool {
	return js.True == bindings.HasFuncMLGraphBuilderHardSigmoid2(
		this.ref,
	)
}

// FuncHardSigmoid2 returns the method "MLGraphBuilder.hardSigmoid".
func (this MLGraphBuilder) FuncHardSigmoid2() (fn js.Func[func(options MLHardSigmoidOptions) MLActivation]) {
	bindings.FuncMLGraphBuilderHardSigmoid2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HardSigmoid2 calls the method "MLGraphBuilder.hardSigmoid".
func (this MLGraphBuilder) HardSigmoid2(options MLHardSigmoidOptions) (ret MLActivation) {
	bindings.CallMLGraphBuilderHardSigmoid2(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryHardSigmoid2 calls the method "MLGraphBuilder.hardSigmoid"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryHardSigmoid2(options MLHardSigmoidOptions) (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderHardSigmoid2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncHardSigmoid3 returns true if the method "MLGraphBuilder.hardSigmoid" exists.
func (this MLGraphBuilder) HasFuncHardSigmoid3() bool {
	return js.True == bindings.HasFuncMLGraphBuilderHardSigmoid3(
		this.ref,
	)
}

// FuncHardSigmoid3 returns the method "MLGraphBuilder.hardSigmoid".
func (this MLGraphBuilder) FuncHardSigmoid3() (fn js.Func[func() MLActivation]) {
	bindings.FuncMLGraphBuilderHardSigmoid3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HardSigmoid3 calls the method "MLGraphBuilder.hardSigmoid".
func (this MLGraphBuilder) HardSigmoid3() (ret MLActivation) {
	bindings.CallMLGraphBuilderHardSigmoid3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryHardSigmoid3 calls the method "MLGraphBuilder.hardSigmoid"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryHardSigmoid3() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderHardSigmoid3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGruCell returns true if the method "MLGraphBuilder.gruCell" exists.
func (this MLGraphBuilder) HasFuncGruCell() bool {
	return js.True == bindings.HasFuncMLGraphBuilderGruCell(
		this.ref,
	)
}

// FuncGruCell returns the method "MLGraphBuilder.gruCell".
func (this MLGraphBuilder) FuncGruCell() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, hiddenSize uint32, options MLGruCellOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderGruCell(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GruCell calls the method "MLGraphBuilder.gruCell".
func (this MLGraphBuilder) GruCell(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, hiddenSize uint32, options MLGruCellOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderGruCell(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		hiddenState.Ref(),
		uint32(hiddenSize),
		js.Pointer(&options),
	)

	return
}

// TryGruCell calls the method "MLGraphBuilder.gruCell"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryGruCell(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, hiddenSize uint32, options MLGruCellOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderGruCell(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		hiddenState.Ref(),
		uint32(hiddenSize),
		js.Pointer(&options),
	)

	return
}

// HasFuncGruCell1 returns true if the method "MLGraphBuilder.gruCell" exists.
func (this MLGraphBuilder) HasFuncGruCell1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderGruCell1(
		this.ref,
	)
}

// FuncGruCell1 returns the method "MLGraphBuilder.gruCell".
func (this MLGraphBuilder) FuncGruCell1() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, hiddenSize uint32) MLOperand]) {
	bindings.FuncMLGraphBuilderGruCell1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GruCell1 calls the method "MLGraphBuilder.gruCell".
func (this MLGraphBuilder) GruCell1(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, hiddenSize uint32) (ret MLOperand) {
	bindings.CallMLGraphBuilderGruCell1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		hiddenState.Ref(),
		uint32(hiddenSize),
	)

	return
}

// TryGruCell1 calls the method "MLGraphBuilder.gruCell"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryGruCell1(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, hiddenSize uint32) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderGruCell1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		hiddenState.Ref(),
		uint32(hiddenSize),
	)

	return
}

// HasFuncSlice returns true if the method "MLGraphBuilder.slice" exists.
func (this MLGraphBuilder) HasFuncSlice() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSlice(
		this.ref,
	)
}

// FuncSlice returns the method "MLGraphBuilder.slice".
func (this MLGraphBuilder) FuncSlice() (fn js.Func[func(input MLOperand, starts js.Array[uint32], sizes js.Array[uint32]) MLOperand]) {
	bindings.FuncMLGraphBuilderSlice(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Slice calls the method "MLGraphBuilder.slice".
func (this MLGraphBuilder) Slice(input MLOperand, starts js.Array[uint32], sizes js.Array[uint32]) (ret MLOperand) {
	bindings.CallMLGraphBuilderSlice(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		starts.Ref(),
		sizes.Ref(),
	)

	return
}

// TrySlice calls the method "MLGraphBuilder.slice"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySlice(input MLOperand, starts js.Array[uint32], sizes js.Array[uint32]) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSlice(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		starts.Ref(),
		sizes.Ref(),
	)

	return
}

// HasFuncAveragePool2d returns true if the method "MLGraphBuilder.averagePool2d" exists.
func (this MLGraphBuilder) HasFuncAveragePool2d() bool {
	return js.True == bindings.HasFuncMLGraphBuilderAveragePool2d(
		this.ref,
	)
}

// FuncAveragePool2d returns the method "MLGraphBuilder.averagePool2d".
func (this MLGraphBuilder) FuncAveragePool2d() (fn js.Func[func(input MLOperand, options MLPool2dOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderAveragePool2d(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AveragePool2d calls the method "MLGraphBuilder.averagePool2d".
func (this MLGraphBuilder) AveragePool2d(input MLOperand, options MLPool2dOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderAveragePool2d(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryAveragePool2d calls the method "MLGraphBuilder.averagePool2d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryAveragePool2d(input MLOperand, options MLPool2dOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderAveragePool2d(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncAveragePool2d1 returns true if the method "MLGraphBuilder.averagePool2d" exists.
func (this MLGraphBuilder) HasFuncAveragePool2d1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderAveragePool2d1(
		this.ref,
	)
}

// FuncAveragePool2d1 returns the method "MLGraphBuilder.averagePool2d".
func (this MLGraphBuilder) FuncAveragePool2d1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderAveragePool2d1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AveragePool2d1 calls the method "MLGraphBuilder.averagePool2d".
func (this MLGraphBuilder) AveragePool2d1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderAveragePool2d1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryAveragePool2d1 calls the method "MLGraphBuilder.averagePool2d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryAveragePool2d1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderAveragePool2d1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncL2Pool2d returns true if the method "MLGraphBuilder.l2Pool2d" exists.
func (this MLGraphBuilder) HasFuncL2Pool2d() bool {
	return js.True == bindings.HasFuncMLGraphBuilderL2Pool2d(
		this.ref,
	)
}

// FuncL2Pool2d returns the method "MLGraphBuilder.l2Pool2d".
func (this MLGraphBuilder) FuncL2Pool2d() (fn js.Func[func(input MLOperand, options MLPool2dOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderL2Pool2d(
		this.ref, js.Pointer(&fn),
	)
	return
}

// L2Pool2d calls the method "MLGraphBuilder.l2Pool2d".
func (this MLGraphBuilder) L2Pool2d(input MLOperand, options MLPool2dOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderL2Pool2d(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryL2Pool2d calls the method "MLGraphBuilder.l2Pool2d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryL2Pool2d(input MLOperand, options MLPool2dOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderL2Pool2d(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncL2Pool2d1 returns true if the method "MLGraphBuilder.l2Pool2d" exists.
func (this MLGraphBuilder) HasFuncL2Pool2d1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderL2Pool2d1(
		this.ref,
	)
}

// FuncL2Pool2d1 returns the method "MLGraphBuilder.l2Pool2d".
func (this MLGraphBuilder) FuncL2Pool2d1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderL2Pool2d1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// L2Pool2d1 calls the method "MLGraphBuilder.l2Pool2d".
func (this MLGraphBuilder) L2Pool2d1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderL2Pool2d1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryL2Pool2d1 calls the method "MLGraphBuilder.l2Pool2d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryL2Pool2d1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderL2Pool2d1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncMaxPool2d returns true if the method "MLGraphBuilder.maxPool2d" exists.
func (this MLGraphBuilder) HasFuncMaxPool2d() bool {
	return js.True == bindings.HasFuncMLGraphBuilderMaxPool2d(
		this.ref,
	)
}

// FuncMaxPool2d returns the method "MLGraphBuilder.maxPool2d".
func (this MLGraphBuilder) FuncMaxPool2d() (fn js.Func[func(input MLOperand, options MLPool2dOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderMaxPool2d(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MaxPool2d calls the method "MLGraphBuilder.maxPool2d".
func (this MLGraphBuilder) MaxPool2d(input MLOperand, options MLPool2dOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderMaxPool2d(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryMaxPool2d calls the method "MLGraphBuilder.maxPool2d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryMaxPool2d(input MLOperand, options MLPool2dOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderMaxPool2d(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncMaxPool2d1 returns true if the method "MLGraphBuilder.maxPool2d" exists.
func (this MLGraphBuilder) HasFuncMaxPool2d1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderMaxPool2d1(
		this.ref,
	)
}

// FuncMaxPool2d1 returns the method "MLGraphBuilder.maxPool2d".
func (this MLGraphBuilder) FuncMaxPool2d1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderMaxPool2d1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MaxPool2d1 calls the method "MLGraphBuilder.maxPool2d".
func (this MLGraphBuilder) MaxPool2d1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderMaxPool2d1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryMaxPool2d1 calls the method "MLGraphBuilder.maxPool2d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryMaxPool2d1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderMaxPool2d1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncLinear returns true if the method "MLGraphBuilder.linear" exists.
func (this MLGraphBuilder) HasFuncLinear() bool {
	return js.True == bindings.HasFuncMLGraphBuilderLinear(
		this.ref,
	)
}

// FuncLinear returns the method "MLGraphBuilder.linear".
func (this MLGraphBuilder) FuncLinear() (fn js.Func[func(input MLOperand, options MLLinearOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderLinear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Linear calls the method "MLGraphBuilder.linear".
func (this MLGraphBuilder) Linear(input MLOperand, options MLLinearOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderLinear(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryLinear calls the method "MLGraphBuilder.linear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLinear(input MLOperand, options MLLinearOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLinear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncLinear1 returns true if the method "MLGraphBuilder.linear" exists.
func (this MLGraphBuilder) HasFuncLinear1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderLinear1(
		this.ref,
	)
}

// FuncLinear1 returns the method "MLGraphBuilder.linear".
func (this MLGraphBuilder) FuncLinear1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderLinear1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Linear1 calls the method "MLGraphBuilder.linear".
func (this MLGraphBuilder) Linear1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderLinear1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryLinear1 calls the method "MLGraphBuilder.linear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLinear1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLinear1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncLinear2 returns true if the method "MLGraphBuilder.linear" exists.
func (this MLGraphBuilder) HasFuncLinear2() bool {
	return js.True == bindings.HasFuncMLGraphBuilderLinear2(
		this.ref,
	)
}

// FuncLinear2 returns the method "MLGraphBuilder.linear".
func (this MLGraphBuilder) FuncLinear2() (fn js.Func[func(options MLLinearOptions) MLActivation]) {
	bindings.FuncMLGraphBuilderLinear2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Linear2 calls the method "MLGraphBuilder.linear".
func (this MLGraphBuilder) Linear2(options MLLinearOptions) (ret MLActivation) {
	bindings.CallMLGraphBuilderLinear2(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryLinear2 calls the method "MLGraphBuilder.linear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLinear2(options MLLinearOptions) (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLinear2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncLinear3 returns true if the method "MLGraphBuilder.linear" exists.
func (this MLGraphBuilder) HasFuncLinear3() bool {
	return js.True == bindings.HasFuncMLGraphBuilderLinear3(
		this.ref,
	)
}

// FuncLinear3 returns the method "MLGraphBuilder.linear".
func (this MLGraphBuilder) FuncLinear3() (fn js.Func[func() MLActivation]) {
	bindings.FuncMLGraphBuilderLinear3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Linear3 calls the method "MLGraphBuilder.linear".
func (this MLGraphBuilder) Linear3() (ret MLActivation) {
	bindings.CallMLGraphBuilderLinear3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryLinear3 calls the method "MLGraphBuilder.linear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLinear3() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLinear3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncLeakyRelu returns true if the method "MLGraphBuilder.leakyRelu" exists.
func (this MLGraphBuilder) HasFuncLeakyRelu() bool {
	return js.True == bindings.HasFuncMLGraphBuilderLeakyRelu(
		this.ref,
	)
}

// FuncLeakyRelu returns the method "MLGraphBuilder.leakyRelu".
func (this MLGraphBuilder) FuncLeakyRelu() (fn js.Func[func(input MLOperand, options MLLeakyReluOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderLeakyRelu(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LeakyRelu calls the method "MLGraphBuilder.leakyRelu".
func (this MLGraphBuilder) LeakyRelu(input MLOperand, options MLLeakyReluOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderLeakyRelu(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryLeakyRelu calls the method "MLGraphBuilder.leakyRelu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLeakyRelu(input MLOperand, options MLLeakyReluOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLeakyRelu(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncLeakyRelu1 returns true if the method "MLGraphBuilder.leakyRelu" exists.
func (this MLGraphBuilder) HasFuncLeakyRelu1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderLeakyRelu1(
		this.ref,
	)
}

// FuncLeakyRelu1 returns the method "MLGraphBuilder.leakyRelu".
func (this MLGraphBuilder) FuncLeakyRelu1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderLeakyRelu1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LeakyRelu1 calls the method "MLGraphBuilder.leakyRelu".
func (this MLGraphBuilder) LeakyRelu1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderLeakyRelu1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryLeakyRelu1 calls the method "MLGraphBuilder.leakyRelu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLeakyRelu1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLeakyRelu1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncLeakyRelu2 returns true if the method "MLGraphBuilder.leakyRelu" exists.
func (this MLGraphBuilder) HasFuncLeakyRelu2() bool {
	return js.True == bindings.HasFuncMLGraphBuilderLeakyRelu2(
		this.ref,
	)
}

// FuncLeakyRelu2 returns the method "MLGraphBuilder.leakyRelu".
func (this MLGraphBuilder) FuncLeakyRelu2() (fn js.Func[func(options MLLeakyReluOptions) MLActivation]) {
	bindings.FuncMLGraphBuilderLeakyRelu2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LeakyRelu2 calls the method "MLGraphBuilder.leakyRelu".
func (this MLGraphBuilder) LeakyRelu2(options MLLeakyReluOptions) (ret MLActivation) {
	bindings.CallMLGraphBuilderLeakyRelu2(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryLeakyRelu2 calls the method "MLGraphBuilder.leakyRelu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLeakyRelu2(options MLLeakyReluOptions) (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLeakyRelu2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncLeakyRelu3 returns true if the method "MLGraphBuilder.leakyRelu" exists.
func (this MLGraphBuilder) HasFuncLeakyRelu3() bool {
	return js.True == bindings.HasFuncMLGraphBuilderLeakyRelu3(
		this.ref,
	)
}

// FuncLeakyRelu3 returns the method "MLGraphBuilder.leakyRelu".
func (this MLGraphBuilder) FuncLeakyRelu3() (fn js.Func[func() MLActivation]) {
	bindings.FuncMLGraphBuilderLeakyRelu3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LeakyRelu3 calls the method "MLGraphBuilder.leakyRelu".
func (this MLGraphBuilder) LeakyRelu3() (ret MLActivation) {
	bindings.CallMLGraphBuilderLeakyRelu3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryLeakyRelu3 calls the method "MLGraphBuilder.leakyRelu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLeakyRelu3() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLeakyRelu3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPad returns true if the method "MLGraphBuilder.pad" exists.
func (this MLGraphBuilder) HasFuncPad() bool {
	return js.True == bindings.HasFuncMLGraphBuilderPad(
		this.ref,
	)
}

// FuncPad returns the method "MLGraphBuilder.pad".
func (this MLGraphBuilder) FuncPad() (fn js.Func[func(input MLOperand, beginningPadding js.Array[uint32], endingPadding js.Array[uint32], options MLPadOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderPad(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Pad calls the method "MLGraphBuilder.pad".
func (this MLGraphBuilder) Pad(input MLOperand, beginningPadding js.Array[uint32], endingPadding js.Array[uint32], options MLPadOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderPad(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		beginningPadding.Ref(),
		endingPadding.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryPad calls the method "MLGraphBuilder.pad"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryPad(input MLOperand, beginningPadding js.Array[uint32], endingPadding js.Array[uint32], options MLPadOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderPad(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		beginningPadding.Ref(),
		endingPadding.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncPad1 returns true if the method "MLGraphBuilder.pad" exists.
func (this MLGraphBuilder) HasFuncPad1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderPad1(
		this.ref,
	)
}

// FuncPad1 returns the method "MLGraphBuilder.pad".
func (this MLGraphBuilder) FuncPad1() (fn js.Func[func(input MLOperand, beginningPadding js.Array[uint32], endingPadding js.Array[uint32]) MLOperand]) {
	bindings.FuncMLGraphBuilderPad1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Pad1 calls the method "MLGraphBuilder.pad".
func (this MLGraphBuilder) Pad1(input MLOperand, beginningPadding js.Array[uint32], endingPadding js.Array[uint32]) (ret MLOperand) {
	bindings.CallMLGraphBuilderPad1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		beginningPadding.Ref(),
		endingPadding.Ref(),
	)

	return
}

// TryPad1 calls the method "MLGraphBuilder.pad"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryPad1(input MLOperand, beginningPadding js.Array[uint32], endingPadding js.Array[uint32]) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderPad1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		beginningPadding.Ref(),
		endingPadding.Ref(),
	)

	return
}

// HasFuncInstanceNormalization returns true if the method "MLGraphBuilder.instanceNormalization" exists.
func (this MLGraphBuilder) HasFuncInstanceNormalization() bool {
	return js.True == bindings.HasFuncMLGraphBuilderInstanceNormalization(
		this.ref,
	)
}

// FuncInstanceNormalization returns the method "MLGraphBuilder.instanceNormalization".
func (this MLGraphBuilder) FuncInstanceNormalization() (fn js.Func[func(input MLOperand, options MLInstanceNormalizationOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderInstanceNormalization(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InstanceNormalization calls the method "MLGraphBuilder.instanceNormalization".
func (this MLGraphBuilder) InstanceNormalization(input MLOperand, options MLInstanceNormalizationOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderInstanceNormalization(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryInstanceNormalization calls the method "MLGraphBuilder.instanceNormalization"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryInstanceNormalization(input MLOperand, options MLInstanceNormalizationOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderInstanceNormalization(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncInstanceNormalization1 returns true if the method "MLGraphBuilder.instanceNormalization" exists.
func (this MLGraphBuilder) HasFuncInstanceNormalization1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderInstanceNormalization1(
		this.ref,
	)
}

// FuncInstanceNormalization1 returns the method "MLGraphBuilder.instanceNormalization".
func (this MLGraphBuilder) FuncInstanceNormalization1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderInstanceNormalization1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InstanceNormalization1 calls the method "MLGraphBuilder.instanceNormalization".
func (this MLGraphBuilder) InstanceNormalization1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderInstanceNormalization1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryInstanceNormalization1 calls the method "MLGraphBuilder.instanceNormalization"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryInstanceNormalization1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderInstanceNormalization1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncSoftplus returns true if the method "MLGraphBuilder.softplus" exists.
func (this MLGraphBuilder) HasFuncSoftplus() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSoftplus(
		this.ref,
	)
}

// FuncSoftplus returns the method "MLGraphBuilder.softplus".
func (this MLGraphBuilder) FuncSoftplus() (fn js.Func[func(input MLOperand, options MLSoftplusOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderSoftplus(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Softplus calls the method "MLGraphBuilder.softplus".
func (this MLGraphBuilder) Softplus(input MLOperand, options MLSoftplusOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderSoftplus(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TrySoftplus calls the method "MLGraphBuilder.softplus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySoftplus(input MLOperand, options MLSoftplusOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSoftplus(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncSoftplus1 returns true if the method "MLGraphBuilder.softplus" exists.
func (this MLGraphBuilder) HasFuncSoftplus1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSoftplus1(
		this.ref,
	)
}

// FuncSoftplus1 returns the method "MLGraphBuilder.softplus".
func (this MLGraphBuilder) FuncSoftplus1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderSoftplus1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Softplus1 calls the method "MLGraphBuilder.softplus".
func (this MLGraphBuilder) Softplus1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderSoftplus1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySoftplus1 calls the method "MLGraphBuilder.softplus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySoftplus1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSoftplus1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncSoftplus2 returns true if the method "MLGraphBuilder.softplus" exists.
func (this MLGraphBuilder) HasFuncSoftplus2() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSoftplus2(
		this.ref,
	)
}

// FuncSoftplus2 returns the method "MLGraphBuilder.softplus".
func (this MLGraphBuilder) FuncSoftplus2() (fn js.Func[func(options MLSoftplusOptions) MLActivation]) {
	bindings.FuncMLGraphBuilderSoftplus2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Softplus2 calls the method "MLGraphBuilder.softplus".
func (this MLGraphBuilder) Softplus2(options MLSoftplusOptions) (ret MLActivation) {
	bindings.CallMLGraphBuilderSoftplus2(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TrySoftplus2 calls the method "MLGraphBuilder.softplus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySoftplus2(options MLSoftplusOptions) (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSoftplus2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncSoftplus3 returns true if the method "MLGraphBuilder.softplus" exists.
func (this MLGraphBuilder) HasFuncSoftplus3() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSoftplus3(
		this.ref,
	)
}

// FuncSoftplus3 returns the method "MLGraphBuilder.softplus".
func (this MLGraphBuilder) FuncSoftplus3() (fn js.Func[func() MLActivation]) {
	bindings.FuncMLGraphBuilderSoftplus3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Softplus3 calls the method "MLGraphBuilder.softplus".
func (this MLGraphBuilder) Softplus3() (ret MLActivation) {
	bindings.CallMLGraphBuilderSoftplus3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySoftplus3 calls the method "MLGraphBuilder.softplus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySoftplus3() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSoftplus3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSoftsign returns true if the method "MLGraphBuilder.softsign" exists.
func (this MLGraphBuilder) HasFuncSoftsign() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSoftsign(
		this.ref,
	)
}

// FuncSoftsign returns the method "MLGraphBuilder.softsign".
func (this MLGraphBuilder) FuncSoftsign() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderSoftsign(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Softsign calls the method "MLGraphBuilder.softsign".
func (this MLGraphBuilder) Softsign(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderSoftsign(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySoftsign calls the method "MLGraphBuilder.softsign"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySoftsign(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSoftsign(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncSoftsign1 returns true if the method "MLGraphBuilder.softsign" exists.
func (this MLGraphBuilder) HasFuncSoftsign1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSoftsign1(
		this.ref,
	)
}

// FuncSoftsign1 returns the method "MLGraphBuilder.softsign".
func (this MLGraphBuilder) FuncSoftsign1() (fn js.Func[func() MLActivation]) {
	bindings.FuncMLGraphBuilderSoftsign1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Softsign1 calls the method "MLGraphBuilder.softsign".
func (this MLGraphBuilder) Softsign1() (ret MLActivation) {
	bindings.CallMLGraphBuilderSoftsign1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySoftsign1 calls the method "MLGraphBuilder.softsign"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySoftsign1() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSoftsign1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSigmoid returns true if the method "MLGraphBuilder.sigmoid" exists.
func (this MLGraphBuilder) HasFuncSigmoid() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSigmoid(
		this.ref,
	)
}

// FuncSigmoid returns the method "MLGraphBuilder.sigmoid".
func (this MLGraphBuilder) FuncSigmoid() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderSigmoid(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Sigmoid calls the method "MLGraphBuilder.sigmoid".
func (this MLGraphBuilder) Sigmoid(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderSigmoid(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySigmoid calls the method "MLGraphBuilder.sigmoid"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySigmoid(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSigmoid(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncSigmoid1 returns true if the method "MLGraphBuilder.sigmoid" exists.
func (this MLGraphBuilder) HasFuncSigmoid1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSigmoid1(
		this.ref,
	)
}

// FuncSigmoid1 returns the method "MLGraphBuilder.sigmoid".
func (this MLGraphBuilder) FuncSigmoid1() (fn js.Func[func() MLActivation]) {
	bindings.FuncMLGraphBuilderSigmoid1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Sigmoid1 calls the method "MLGraphBuilder.sigmoid".
func (this MLGraphBuilder) Sigmoid1() (ret MLActivation) {
	bindings.CallMLGraphBuilderSigmoid1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySigmoid1 calls the method "MLGraphBuilder.sigmoid"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySigmoid1() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSigmoid1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReshape returns true if the method "MLGraphBuilder.reshape" exists.
func (this MLGraphBuilder) HasFuncReshape() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReshape(
		this.ref,
	)
}

// FuncReshape returns the method "MLGraphBuilder.reshape".
func (this MLGraphBuilder) FuncReshape() (fn js.Func[func(input MLOperand, newShape js.Array[uint32]) MLOperand]) {
	bindings.FuncMLGraphBuilderReshape(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Reshape calls the method "MLGraphBuilder.reshape".
func (this MLGraphBuilder) Reshape(input MLOperand, newShape js.Array[uint32]) (ret MLOperand) {
	bindings.CallMLGraphBuilderReshape(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		newShape.Ref(),
	)

	return
}

// TryReshape calls the method "MLGraphBuilder.reshape"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReshape(input MLOperand, newShape js.Array[uint32]) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReshape(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		newShape.Ref(),
	)

	return
}

// HasFuncConv2d returns true if the method "MLGraphBuilder.conv2d" exists.
func (this MLGraphBuilder) HasFuncConv2d() bool {
	return js.True == bindings.HasFuncMLGraphBuilderConv2d(
		this.ref,
	)
}

// FuncConv2d returns the method "MLGraphBuilder.conv2d".
func (this MLGraphBuilder) FuncConv2d() (fn js.Func[func(input MLOperand, filter MLOperand, options MLConv2dOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderConv2d(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Conv2d calls the method "MLGraphBuilder.conv2d".
func (this MLGraphBuilder) Conv2d(input MLOperand, filter MLOperand, options MLConv2dOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderConv2d(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		filter.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConv2d calls the method "MLGraphBuilder.conv2d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryConv2d(input MLOperand, filter MLOperand, options MLConv2dOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderConv2d(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		filter.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncConv2d1 returns true if the method "MLGraphBuilder.conv2d" exists.
func (this MLGraphBuilder) HasFuncConv2d1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderConv2d1(
		this.ref,
	)
}

// FuncConv2d1 returns the method "MLGraphBuilder.conv2d".
func (this MLGraphBuilder) FuncConv2d1() (fn js.Func[func(input MLOperand, filter MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderConv2d1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Conv2d1 calls the method "MLGraphBuilder.conv2d".
func (this MLGraphBuilder) Conv2d1(input MLOperand, filter MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderConv2d1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		filter.Ref(),
	)

	return
}

// TryConv2d1 calls the method "MLGraphBuilder.conv2d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryConv2d1(input MLOperand, filter MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderConv2d1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		filter.Ref(),
	)

	return
}

// HasFuncSplit returns true if the method "MLGraphBuilder.split" exists.
func (this MLGraphBuilder) HasFuncSplit() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSplit(
		this.ref,
	)
}

// FuncSplit returns the method "MLGraphBuilder.split".
func (this MLGraphBuilder) FuncSplit() (fn js.Func[func(input MLOperand, splits OneOf_Uint32_ArrayUint32, options MLSplitOptions) js.Array[MLOperand]]) {
	bindings.FuncMLGraphBuilderSplit(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Split calls the method "MLGraphBuilder.split".
func (this MLGraphBuilder) Split(input MLOperand, splits OneOf_Uint32_ArrayUint32, options MLSplitOptions) (ret js.Array[MLOperand]) {
	bindings.CallMLGraphBuilderSplit(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		splits.Ref(),
		js.Pointer(&options),
	)

	return
}

// TrySplit calls the method "MLGraphBuilder.split"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySplit(input MLOperand, splits OneOf_Uint32_ArrayUint32, options MLSplitOptions) (ret js.Array[MLOperand], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSplit(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		splits.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncSplit1 returns true if the method "MLGraphBuilder.split" exists.
func (this MLGraphBuilder) HasFuncSplit1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSplit1(
		this.ref,
	)
}

// FuncSplit1 returns the method "MLGraphBuilder.split".
func (this MLGraphBuilder) FuncSplit1() (fn js.Func[func(input MLOperand, splits OneOf_Uint32_ArrayUint32) js.Array[MLOperand]]) {
	bindings.FuncMLGraphBuilderSplit1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Split1 calls the method "MLGraphBuilder.split".
func (this MLGraphBuilder) Split1(input MLOperand, splits OneOf_Uint32_ArrayUint32) (ret js.Array[MLOperand]) {
	bindings.CallMLGraphBuilderSplit1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		splits.Ref(),
	)

	return
}

// TrySplit1 calls the method "MLGraphBuilder.split"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySplit1(input MLOperand, splits OneOf_Uint32_ArrayUint32) (ret js.Array[MLOperand], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSplit1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		splits.Ref(),
	)

	return
}

// HasFuncResample2d returns true if the method "MLGraphBuilder.resample2d" exists.
func (this MLGraphBuilder) HasFuncResample2d() bool {
	return js.True == bindings.HasFuncMLGraphBuilderResample2d(
		this.ref,
	)
}

// FuncResample2d returns the method "MLGraphBuilder.resample2d".
func (this MLGraphBuilder) FuncResample2d() (fn js.Func[func(input MLOperand, options MLResample2dOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderResample2d(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Resample2d calls the method "MLGraphBuilder.resample2d".
func (this MLGraphBuilder) Resample2d(input MLOperand, options MLResample2dOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderResample2d(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryResample2d calls the method "MLGraphBuilder.resample2d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryResample2d(input MLOperand, options MLResample2dOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderResample2d(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncResample2d1 returns true if the method "MLGraphBuilder.resample2d" exists.
func (this MLGraphBuilder) HasFuncResample2d1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderResample2d1(
		this.ref,
	)
}

// FuncResample2d1 returns the method "MLGraphBuilder.resample2d".
func (this MLGraphBuilder) FuncResample2d1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderResample2d1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Resample2d1 calls the method "MLGraphBuilder.resample2d".
func (this MLGraphBuilder) Resample2d1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderResample2d1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryResample2d1 calls the method "MLGraphBuilder.resample2d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryResample2d1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderResample2d1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncHardSwish returns true if the method "MLGraphBuilder.hardSwish" exists.
func (this MLGraphBuilder) HasFuncHardSwish() bool {
	return js.True == bindings.HasFuncMLGraphBuilderHardSwish(
		this.ref,
	)
}

// FuncHardSwish returns the method "MLGraphBuilder.hardSwish".
func (this MLGraphBuilder) FuncHardSwish() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderHardSwish(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HardSwish calls the method "MLGraphBuilder.hardSwish".
func (this MLGraphBuilder) HardSwish(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderHardSwish(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryHardSwish calls the method "MLGraphBuilder.hardSwish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryHardSwish(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderHardSwish(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncHardSwish1 returns true if the method "MLGraphBuilder.hardSwish" exists.
func (this MLGraphBuilder) HasFuncHardSwish1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderHardSwish1(
		this.ref,
	)
}

// FuncHardSwish1 returns the method "MLGraphBuilder.hardSwish".
func (this MLGraphBuilder) FuncHardSwish1() (fn js.Func[func() MLActivation]) {
	bindings.FuncMLGraphBuilderHardSwish1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HardSwish1 calls the method "MLGraphBuilder.hardSwish".
func (this MLGraphBuilder) HardSwish1() (ret MLActivation) {
	bindings.CallMLGraphBuilderHardSwish1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryHardSwish1 calls the method "MLGraphBuilder.hardSwish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryHardSwish1() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderHardSwish1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSoftmax returns true if the method "MLGraphBuilder.softmax" exists.
func (this MLGraphBuilder) HasFuncSoftmax() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSoftmax(
		this.ref,
	)
}

// FuncSoftmax returns the method "MLGraphBuilder.softmax".
func (this MLGraphBuilder) FuncSoftmax() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderSoftmax(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Softmax calls the method "MLGraphBuilder.softmax".
func (this MLGraphBuilder) Softmax(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderSoftmax(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySoftmax calls the method "MLGraphBuilder.softmax"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySoftmax(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSoftmax(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncSoftmax1 returns true if the method "MLGraphBuilder.softmax" exists.
func (this MLGraphBuilder) HasFuncSoftmax1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSoftmax1(
		this.ref,
	)
}

// FuncSoftmax1 returns the method "MLGraphBuilder.softmax".
func (this MLGraphBuilder) FuncSoftmax1() (fn js.Func[func() MLActivation]) {
	bindings.FuncMLGraphBuilderSoftmax1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Softmax1 calls the method "MLGraphBuilder.softmax".
func (this MLGraphBuilder) Softmax1() (ret MLActivation) {
	bindings.CallMLGraphBuilderSoftmax1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySoftmax1 calls the method "MLGraphBuilder.softmax"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySoftmax1() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSoftmax1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncConvTranspose2d returns true if the method "MLGraphBuilder.convTranspose2d" exists.
func (this MLGraphBuilder) HasFuncConvTranspose2d() bool {
	return js.True == bindings.HasFuncMLGraphBuilderConvTranspose2d(
		this.ref,
	)
}

// FuncConvTranspose2d returns the method "MLGraphBuilder.convTranspose2d".
func (this MLGraphBuilder) FuncConvTranspose2d() (fn js.Func[func(input MLOperand, filter MLOperand, options MLConvTranspose2dOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderConvTranspose2d(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvTranspose2d calls the method "MLGraphBuilder.convTranspose2d".
func (this MLGraphBuilder) ConvTranspose2d(input MLOperand, filter MLOperand, options MLConvTranspose2dOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderConvTranspose2d(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		filter.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryConvTranspose2d calls the method "MLGraphBuilder.convTranspose2d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryConvTranspose2d(input MLOperand, filter MLOperand, options MLConvTranspose2dOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderConvTranspose2d(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		filter.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncConvTranspose2d1 returns true if the method "MLGraphBuilder.convTranspose2d" exists.
func (this MLGraphBuilder) HasFuncConvTranspose2d1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderConvTranspose2d1(
		this.ref,
	)
}

// FuncConvTranspose2d1 returns the method "MLGraphBuilder.convTranspose2d".
func (this MLGraphBuilder) FuncConvTranspose2d1() (fn js.Func[func(input MLOperand, filter MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderConvTranspose2d1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ConvTranspose2d1 calls the method "MLGraphBuilder.convTranspose2d".
func (this MLGraphBuilder) ConvTranspose2d1(input MLOperand, filter MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderConvTranspose2d1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		filter.Ref(),
	)

	return
}

// TryConvTranspose2d1 calls the method "MLGraphBuilder.convTranspose2d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryConvTranspose2d1(input MLOperand, filter MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderConvTranspose2d1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		filter.Ref(),
	)

	return
}

// HasFuncRelu returns true if the method "MLGraphBuilder.relu" exists.
func (this MLGraphBuilder) HasFuncRelu() bool {
	return js.True == bindings.HasFuncMLGraphBuilderRelu(
		this.ref,
	)
}

// FuncRelu returns the method "MLGraphBuilder.relu".
func (this MLGraphBuilder) FuncRelu() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderRelu(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Relu calls the method "MLGraphBuilder.relu".
func (this MLGraphBuilder) Relu(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderRelu(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryRelu calls the method "MLGraphBuilder.relu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryRelu(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderRelu(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncRelu1 returns true if the method "MLGraphBuilder.relu" exists.
func (this MLGraphBuilder) HasFuncRelu1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderRelu1(
		this.ref,
	)
}

// FuncRelu1 returns the method "MLGraphBuilder.relu".
func (this MLGraphBuilder) FuncRelu1() (fn js.Func[func() MLActivation]) {
	bindings.FuncMLGraphBuilderRelu1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Relu1 calls the method "MLGraphBuilder.relu".
func (this MLGraphBuilder) Relu1() (ret MLActivation) {
	bindings.CallMLGraphBuilderRelu1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRelu1 calls the method "MLGraphBuilder.relu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryRelu1() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderRelu1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAdd returns true if the method "MLGraphBuilder.add" exists.
func (this MLGraphBuilder) HasFuncAdd() bool {
	return js.True == bindings.HasFuncMLGraphBuilderAdd(
		this.ref,
	)
}

// FuncAdd returns the method "MLGraphBuilder.add".
func (this MLGraphBuilder) FuncAdd() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderAdd(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Add calls the method "MLGraphBuilder.add".
func (this MLGraphBuilder) Add(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderAdd(
		this.ref, js.Pointer(&ret),
		a.Ref(),
		b.Ref(),
	)

	return
}

// TryAdd calls the method "MLGraphBuilder.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryAdd(a MLOperand, b MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderAdd(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasFuncSub returns true if the method "MLGraphBuilder.sub" exists.
func (this MLGraphBuilder) HasFuncSub() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSub(
		this.ref,
	)
}

// FuncSub returns the method "MLGraphBuilder.sub".
func (this MLGraphBuilder) FuncSub() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderSub(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Sub calls the method "MLGraphBuilder.sub".
func (this MLGraphBuilder) Sub(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderSub(
		this.ref, js.Pointer(&ret),
		a.Ref(),
		b.Ref(),
	)

	return
}

// TrySub calls the method "MLGraphBuilder.sub"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySub(a MLOperand, b MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSub(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasFuncMul returns true if the method "MLGraphBuilder.mul" exists.
func (this MLGraphBuilder) HasFuncMul() bool {
	return js.True == bindings.HasFuncMLGraphBuilderMul(
		this.ref,
	)
}

// FuncMul returns the method "MLGraphBuilder.mul".
func (this MLGraphBuilder) FuncMul() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderMul(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Mul calls the method "MLGraphBuilder.mul".
func (this MLGraphBuilder) Mul(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderMul(
		this.ref, js.Pointer(&ret),
		a.Ref(),
		b.Ref(),
	)

	return
}

// TryMul calls the method "MLGraphBuilder.mul"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryMul(a MLOperand, b MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderMul(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasFuncDiv returns true if the method "MLGraphBuilder.div" exists.
func (this MLGraphBuilder) HasFuncDiv() bool {
	return js.True == bindings.HasFuncMLGraphBuilderDiv(
		this.ref,
	)
}

// FuncDiv returns the method "MLGraphBuilder.div".
func (this MLGraphBuilder) FuncDiv() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderDiv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Div calls the method "MLGraphBuilder.div".
func (this MLGraphBuilder) Div(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderDiv(
		this.ref, js.Pointer(&ret),
		a.Ref(),
		b.Ref(),
	)

	return
}

// TryDiv calls the method "MLGraphBuilder.div"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryDiv(a MLOperand, b MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderDiv(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasFuncMax returns true if the method "MLGraphBuilder.max" exists.
func (this MLGraphBuilder) HasFuncMax() bool {
	return js.True == bindings.HasFuncMLGraphBuilderMax(
		this.ref,
	)
}

// FuncMax returns the method "MLGraphBuilder.max".
func (this MLGraphBuilder) FuncMax() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderMax(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Max calls the method "MLGraphBuilder.max".
func (this MLGraphBuilder) Max(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderMax(
		this.ref, js.Pointer(&ret),
		a.Ref(),
		b.Ref(),
	)

	return
}

// TryMax calls the method "MLGraphBuilder.max"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryMax(a MLOperand, b MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderMax(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasFuncMin returns true if the method "MLGraphBuilder.min" exists.
func (this MLGraphBuilder) HasFuncMin() bool {
	return js.True == bindings.HasFuncMLGraphBuilderMin(
		this.ref,
	)
}

// FuncMin returns the method "MLGraphBuilder.min".
func (this MLGraphBuilder) FuncMin() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderMin(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Min calls the method "MLGraphBuilder.min".
func (this MLGraphBuilder) Min(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderMin(
		this.ref, js.Pointer(&ret),
		a.Ref(),
		b.Ref(),
	)

	return
}

// TryMin calls the method "MLGraphBuilder.min"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryMin(a MLOperand, b MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderMin(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasFuncPow returns true if the method "MLGraphBuilder.pow" exists.
func (this MLGraphBuilder) HasFuncPow() bool {
	return js.True == bindings.HasFuncMLGraphBuilderPow(
		this.ref,
	)
}

// FuncPow returns the method "MLGraphBuilder.pow".
func (this MLGraphBuilder) FuncPow() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderPow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Pow calls the method "MLGraphBuilder.pow".
func (this MLGraphBuilder) Pow(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderPow(
		this.ref, js.Pointer(&ret),
		a.Ref(),
		b.Ref(),
	)

	return
}

// TryPow calls the method "MLGraphBuilder.pow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryPow(a MLOperand, b MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderPow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasFuncReduceL1 returns true if the method "MLGraphBuilder.reduceL1" exists.
func (this MLGraphBuilder) HasFuncReduceL1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceL1(
		this.ref,
	)
}

// FuncReduceL1 returns the method "MLGraphBuilder.reduceL1".
func (this MLGraphBuilder) FuncReduceL1() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceL1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceL1 calls the method "MLGraphBuilder.reduceL1".
func (this MLGraphBuilder) ReduceL1(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceL1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryReduceL1 calls the method "MLGraphBuilder.reduceL1"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceL1(input MLOperand, options MLReduceOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceL1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncReduceL11 returns true if the method "MLGraphBuilder.reduceL1" exists.
func (this MLGraphBuilder) HasFuncReduceL11() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceL11(
		this.ref,
	)
}

// FuncReduceL11 returns the method "MLGraphBuilder.reduceL1".
func (this MLGraphBuilder) FuncReduceL11() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceL11(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceL11 calls the method "MLGraphBuilder.reduceL1".
func (this MLGraphBuilder) ReduceL11(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceL11(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceL11 calls the method "MLGraphBuilder.reduceL1"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceL11(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceL11(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncReduceL2 returns true if the method "MLGraphBuilder.reduceL2" exists.
func (this MLGraphBuilder) HasFuncReduceL2() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceL2(
		this.ref,
	)
}

// FuncReduceL2 returns the method "MLGraphBuilder.reduceL2".
func (this MLGraphBuilder) FuncReduceL2() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceL2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceL2 calls the method "MLGraphBuilder.reduceL2".
func (this MLGraphBuilder) ReduceL2(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceL2(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryReduceL2 calls the method "MLGraphBuilder.reduceL2"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceL2(input MLOperand, options MLReduceOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceL2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncReduceL21 returns true if the method "MLGraphBuilder.reduceL2" exists.
func (this MLGraphBuilder) HasFuncReduceL21() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceL21(
		this.ref,
	)
}

// FuncReduceL21 returns the method "MLGraphBuilder.reduceL2".
func (this MLGraphBuilder) FuncReduceL21() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceL21(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceL21 calls the method "MLGraphBuilder.reduceL2".
func (this MLGraphBuilder) ReduceL21(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceL21(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceL21 calls the method "MLGraphBuilder.reduceL2"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceL21(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceL21(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncReduceLogSum returns true if the method "MLGraphBuilder.reduceLogSum" exists.
func (this MLGraphBuilder) HasFuncReduceLogSum() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceLogSum(
		this.ref,
	)
}

// FuncReduceLogSum returns the method "MLGraphBuilder.reduceLogSum".
func (this MLGraphBuilder) FuncReduceLogSum() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceLogSum(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceLogSum calls the method "MLGraphBuilder.reduceLogSum".
func (this MLGraphBuilder) ReduceLogSum(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceLogSum(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryReduceLogSum calls the method "MLGraphBuilder.reduceLogSum"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceLogSum(input MLOperand, options MLReduceOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceLogSum(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncReduceLogSum1 returns true if the method "MLGraphBuilder.reduceLogSum" exists.
func (this MLGraphBuilder) HasFuncReduceLogSum1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceLogSum1(
		this.ref,
	)
}

// FuncReduceLogSum1 returns the method "MLGraphBuilder.reduceLogSum".
func (this MLGraphBuilder) FuncReduceLogSum1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceLogSum1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceLogSum1 calls the method "MLGraphBuilder.reduceLogSum".
func (this MLGraphBuilder) ReduceLogSum1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceLogSum1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceLogSum1 calls the method "MLGraphBuilder.reduceLogSum"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceLogSum1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceLogSum1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncReduceLogSumExp returns true if the method "MLGraphBuilder.reduceLogSumExp" exists.
func (this MLGraphBuilder) HasFuncReduceLogSumExp() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceLogSumExp(
		this.ref,
	)
}

// FuncReduceLogSumExp returns the method "MLGraphBuilder.reduceLogSumExp".
func (this MLGraphBuilder) FuncReduceLogSumExp() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceLogSumExp(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceLogSumExp calls the method "MLGraphBuilder.reduceLogSumExp".
func (this MLGraphBuilder) ReduceLogSumExp(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceLogSumExp(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryReduceLogSumExp calls the method "MLGraphBuilder.reduceLogSumExp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceLogSumExp(input MLOperand, options MLReduceOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceLogSumExp(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncReduceLogSumExp1 returns true if the method "MLGraphBuilder.reduceLogSumExp" exists.
func (this MLGraphBuilder) HasFuncReduceLogSumExp1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceLogSumExp1(
		this.ref,
	)
}

// FuncReduceLogSumExp1 returns the method "MLGraphBuilder.reduceLogSumExp".
func (this MLGraphBuilder) FuncReduceLogSumExp1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceLogSumExp1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceLogSumExp1 calls the method "MLGraphBuilder.reduceLogSumExp".
func (this MLGraphBuilder) ReduceLogSumExp1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceLogSumExp1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceLogSumExp1 calls the method "MLGraphBuilder.reduceLogSumExp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceLogSumExp1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceLogSumExp1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncReduceMax returns true if the method "MLGraphBuilder.reduceMax" exists.
func (this MLGraphBuilder) HasFuncReduceMax() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceMax(
		this.ref,
	)
}

// FuncReduceMax returns the method "MLGraphBuilder.reduceMax".
func (this MLGraphBuilder) FuncReduceMax() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceMax(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceMax calls the method "MLGraphBuilder.reduceMax".
func (this MLGraphBuilder) ReduceMax(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceMax(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryReduceMax calls the method "MLGraphBuilder.reduceMax"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceMax(input MLOperand, options MLReduceOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceMax(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncReduceMax1 returns true if the method "MLGraphBuilder.reduceMax" exists.
func (this MLGraphBuilder) HasFuncReduceMax1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceMax1(
		this.ref,
	)
}

// FuncReduceMax1 returns the method "MLGraphBuilder.reduceMax".
func (this MLGraphBuilder) FuncReduceMax1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceMax1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceMax1 calls the method "MLGraphBuilder.reduceMax".
func (this MLGraphBuilder) ReduceMax1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceMax1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceMax1 calls the method "MLGraphBuilder.reduceMax"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceMax1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceMax1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncReduceMean returns true if the method "MLGraphBuilder.reduceMean" exists.
func (this MLGraphBuilder) HasFuncReduceMean() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceMean(
		this.ref,
	)
}

// FuncReduceMean returns the method "MLGraphBuilder.reduceMean".
func (this MLGraphBuilder) FuncReduceMean() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceMean(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceMean calls the method "MLGraphBuilder.reduceMean".
func (this MLGraphBuilder) ReduceMean(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceMean(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryReduceMean calls the method "MLGraphBuilder.reduceMean"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceMean(input MLOperand, options MLReduceOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceMean(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncReduceMean1 returns true if the method "MLGraphBuilder.reduceMean" exists.
func (this MLGraphBuilder) HasFuncReduceMean1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceMean1(
		this.ref,
	)
}

// FuncReduceMean1 returns the method "MLGraphBuilder.reduceMean".
func (this MLGraphBuilder) FuncReduceMean1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceMean1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceMean1 calls the method "MLGraphBuilder.reduceMean".
func (this MLGraphBuilder) ReduceMean1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceMean1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceMean1 calls the method "MLGraphBuilder.reduceMean"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceMean1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceMean1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncReduceMin returns true if the method "MLGraphBuilder.reduceMin" exists.
func (this MLGraphBuilder) HasFuncReduceMin() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceMin(
		this.ref,
	)
}

// FuncReduceMin returns the method "MLGraphBuilder.reduceMin".
func (this MLGraphBuilder) FuncReduceMin() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceMin(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceMin calls the method "MLGraphBuilder.reduceMin".
func (this MLGraphBuilder) ReduceMin(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceMin(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryReduceMin calls the method "MLGraphBuilder.reduceMin"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceMin(input MLOperand, options MLReduceOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceMin(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncReduceMin1 returns true if the method "MLGraphBuilder.reduceMin" exists.
func (this MLGraphBuilder) HasFuncReduceMin1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceMin1(
		this.ref,
	)
}

// FuncReduceMin1 returns the method "MLGraphBuilder.reduceMin".
func (this MLGraphBuilder) FuncReduceMin1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceMin1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceMin1 calls the method "MLGraphBuilder.reduceMin".
func (this MLGraphBuilder) ReduceMin1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceMin1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceMin1 calls the method "MLGraphBuilder.reduceMin"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceMin1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceMin1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncReduceProduct returns true if the method "MLGraphBuilder.reduceProduct" exists.
func (this MLGraphBuilder) HasFuncReduceProduct() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceProduct(
		this.ref,
	)
}

// FuncReduceProduct returns the method "MLGraphBuilder.reduceProduct".
func (this MLGraphBuilder) FuncReduceProduct() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceProduct(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceProduct calls the method "MLGraphBuilder.reduceProduct".
func (this MLGraphBuilder) ReduceProduct(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceProduct(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryReduceProduct calls the method "MLGraphBuilder.reduceProduct"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceProduct(input MLOperand, options MLReduceOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceProduct(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncReduceProduct1 returns true if the method "MLGraphBuilder.reduceProduct" exists.
func (this MLGraphBuilder) HasFuncReduceProduct1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceProduct1(
		this.ref,
	)
}

// FuncReduceProduct1 returns the method "MLGraphBuilder.reduceProduct".
func (this MLGraphBuilder) FuncReduceProduct1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceProduct1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceProduct1 calls the method "MLGraphBuilder.reduceProduct".
func (this MLGraphBuilder) ReduceProduct1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceProduct1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceProduct1 calls the method "MLGraphBuilder.reduceProduct"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceProduct1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceProduct1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncReduceSum returns true if the method "MLGraphBuilder.reduceSum" exists.
func (this MLGraphBuilder) HasFuncReduceSum() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceSum(
		this.ref,
	)
}

// FuncReduceSum returns the method "MLGraphBuilder.reduceSum".
func (this MLGraphBuilder) FuncReduceSum() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceSum(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceSum calls the method "MLGraphBuilder.reduceSum".
func (this MLGraphBuilder) ReduceSum(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceSum(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryReduceSum calls the method "MLGraphBuilder.reduceSum"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceSum(input MLOperand, options MLReduceOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceSum(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncReduceSum1 returns true if the method "MLGraphBuilder.reduceSum" exists.
func (this MLGraphBuilder) HasFuncReduceSum1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceSum1(
		this.ref,
	)
}

// FuncReduceSum1 returns the method "MLGraphBuilder.reduceSum".
func (this MLGraphBuilder) FuncReduceSum1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceSum1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceSum1 calls the method "MLGraphBuilder.reduceSum".
func (this MLGraphBuilder) ReduceSum1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceSum1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceSum1 calls the method "MLGraphBuilder.reduceSum"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceSum1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceSum1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncReduceSumSquare returns true if the method "MLGraphBuilder.reduceSumSquare" exists.
func (this MLGraphBuilder) HasFuncReduceSumSquare() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceSumSquare(
		this.ref,
	)
}

// FuncReduceSumSquare returns the method "MLGraphBuilder.reduceSumSquare".
func (this MLGraphBuilder) FuncReduceSumSquare() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceSumSquare(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceSumSquare calls the method "MLGraphBuilder.reduceSumSquare".
func (this MLGraphBuilder) ReduceSumSquare(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceSumSquare(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryReduceSumSquare calls the method "MLGraphBuilder.reduceSumSquare"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceSumSquare(input MLOperand, options MLReduceOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceSumSquare(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncReduceSumSquare1 returns true if the method "MLGraphBuilder.reduceSumSquare" exists.
func (this MLGraphBuilder) HasFuncReduceSumSquare1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderReduceSumSquare1(
		this.ref,
	)
}

// FuncReduceSumSquare1 returns the method "MLGraphBuilder.reduceSumSquare".
func (this MLGraphBuilder) FuncReduceSumSquare1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderReduceSumSquare1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReduceSumSquare1 calls the method "MLGraphBuilder.reduceSumSquare".
func (this MLGraphBuilder) ReduceSumSquare1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceSumSquare1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceSumSquare1 calls the method "MLGraphBuilder.reduceSumSquare"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceSumSquare1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceSumSquare1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncLstm returns true if the method "MLGraphBuilder.lstm" exists.
func (this MLGraphBuilder) HasFuncLstm() bool {
	return js.True == bindings.HasFuncMLGraphBuilderLstm(
		this.ref,
	)
}

// FuncLstm returns the method "MLGraphBuilder.lstm".
func (this MLGraphBuilder) FuncLstm() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32, options MLLstmOptions) js.Array[MLOperand]]) {
	bindings.FuncMLGraphBuilderLstm(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Lstm calls the method "MLGraphBuilder.lstm".
func (this MLGraphBuilder) Lstm(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32, options MLLstmOptions) (ret js.Array[MLOperand]) {
	bindings.CallMLGraphBuilderLstm(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		uint32(steps),
		uint32(hiddenSize),
		js.Pointer(&options),
	)

	return
}

// TryLstm calls the method "MLGraphBuilder.lstm"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLstm(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32, options MLLstmOptions) (ret js.Array[MLOperand], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLstm(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		uint32(steps),
		uint32(hiddenSize),
		js.Pointer(&options),
	)

	return
}

// HasFuncLstm1 returns true if the method "MLGraphBuilder.lstm" exists.
func (this MLGraphBuilder) HasFuncLstm1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderLstm1(
		this.ref,
	)
}

// FuncLstm1 returns the method "MLGraphBuilder.lstm".
func (this MLGraphBuilder) FuncLstm1() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32) js.Array[MLOperand]]) {
	bindings.FuncMLGraphBuilderLstm1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Lstm1 calls the method "MLGraphBuilder.lstm".
func (this MLGraphBuilder) Lstm1(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32) (ret js.Array[MLOperand]) {
	bindings.CallMLGraphBuilderLstm1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		uint32(steps),
		uint32(hiddenSize),
	)

	return
}

// TryLstm1 calls the method "MLGraphBuilder.lstm"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLstm1(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32) (ret js.Array[MLOperand], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLstm1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		uint32(steps),
		uint32(hiddenSize),
	)

	return
}

// HasFuncMatmul returns true if the method "MLGraphBuilder.matmul" exists.
func (this MLGraphBuilder) HasFuncMatmul() bool {
	return js.True == bindings.HasFuncMLGraphBuilderMatmul(
		this.ref,
	)
}

// FuncMatmul returns the method "MLGraphBuilder.matmul".
func (this MLGraphBuilder) FuncMatmul() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderMatmul(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Matmul calls the method "MLGraphBuilder.matmul".
func (this MLGraphBuilder) Matmul(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderMatmul(
		this.ref, js.Pointer(&ret),
		a.Ref(),
		b.Ref(),
	)

	return
}

// TryMatmul calls the method "MLGraphBuilder.matmul"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryMatmul(a MLOperand, b MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderMatmul(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasFuncSqueeze returns true if the method "MLGraphBuilder.squeeze" exists.
func (this MLGraphBuilder) HasFuncSqueeze() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSqueeze(
		this.ref,
	)
}

// FuncSqueeze returns the method "MLGraphBuilder.squeeze".
func (this MLGraphBuilder) FuncSqueeze() (fn js.Func[func(input MLOperand, options MLSqueezeOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderSqueeze(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Squeeze calls the method "MLGraphBuilder.squeeze".
func (this MLGraphBuilder) Squeeze(input MLOperand, options MLSqueezeOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderSqueeze(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TrySqueeze calls the method "MLGraphBuilder.squeeze"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySqueeze(input MLOperand, options MLSqueezeOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSqueeze(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncSqueeze1 returns true if the method "MLGraphBuilder.squeeze" exists.
func (this MLGraphBuilder) HasFuncSqueeze1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSqueeze1(
		this.ref,
	)
}

// FuncSqueeze1 returns the method "MLGraphBuilder.squeeze".
func (this MLGraphBuilder) FuncSqueeze1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderSqueeze1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Squeeze1 calls the method "MLGraphBuilder.squeeze".
func (this MLGraphBuilder) Squeeze1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderSqueeze1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySqueeze1 calls the method "MLGraphBuilder.squeeze"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySqueeze1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSqueeze1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncTanh returns true if the method "MLGraphBuilder.tanh" exists.
func (this MLGraphBuilder) HasFuncTanh() bool {
	return js.True == bindings.HasFuncMLGraphBuilderTanh(
		this.ref,
	)
}

// FuncTanh returns the method "MLGraphBuilder.tanh".
func (this MLGraphBuilder) FuncTanh() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderTanh(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Tanh calls the method "MLGraphBuilder.tanh".
func (this MLGraphBuilder) Tanh(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderTanh(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryTanh calls the method "MLGraphBuilder.tanh"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryTanh(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderTanh(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncTanh1 returns true if the method "MLGraphBuilder.tanh" exists.
func (this MLGraphBuilder) HasFuncTanh1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderTanh1(
		this.ref,
	)
}

// FuncTanh1 returns the method "MLGraphBuilder.tanh".
func (this MLGraphBuilder) FuncTanh1() (fn js.Func[func() MLActivation]) {
	bindings.FuncMLGraphBuilderTanh1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Tanh1 calls the method "MLGraphBuilder.tanh".
func (this MLGraphBuilder) Tanh1() (ret MLActivation) {
	bindings.CallMLGraphBuilderTanh1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryTanh1 calls the method "MLGraphBuilder.tanh"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryTanh1() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderTanh1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGru returns true if the method "MLGraphBuilder.gru" exists.
func (this MLGraphBuilder) HasFuncGru() bool {
	return js.True == bindings.HasFuncMLGraphBuilderGru(
		this.ref,
	)
}

// FuncGru returns the method "MLGraphBuilder.gru".
func (this MLGraphBuilder) FuncGru() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32, options MLGruOptions) js.Array[MLOperand]]) {
	bindings.FuncMLGraphBuilderGru(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Gru calls the method "MLGraphBuilder.gru".
func (this MLGraphBuilder) Gru(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32, options MLGruOptions) (ret js.Array[MLOperand]) {
	bindings.CallMLGraphBuilderGru(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		uint32(steps),
		uint32(hiddenSize),
		js.Pointer(&options),
	)

	return
}

// TryGru calls the method "MLGraphBuilder.gru"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryGru(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32, options MLGruOptions) (ret js.Array[MLOperand], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderGru(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		uint32(steps),
		uint32(hiddenSize),
		js.Pointer(&options),
	)

	return
}

// HasFuncGru1 returns true if the method "MLGraphBuilder.gru" exists.
func (this MLGraphBuilder) HasFuncGru1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderGru1(
		this.ref,
	)
}

// FuncGru1 returns the method "MLGraphBuilder.gru".
func (this MLGraphBuilder) FuncGru1() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32) js.Array[MLOperand]]) {
	bindings.FuncMLGraphBuilderGru1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Gru1 calls the method "MLGraphBuilder.gru".
func (this MLGraphBuilder) Gru1(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32) (ret js.Array[MLOperand]) {
	bindings.CallMLGraphBuilderGru1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		uint32(steps),
		uint32(hiddenSize),
	)

	return
}

// TryGru1 calls the method "MLGraphBuilder.gru"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryGru1(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32) (ret js.Array[MLOperand], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderGru1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		uint32(steps),
		uint32(hiddenSize),
	)

	return
}

// HasFuncAbs returns true if the method "MLGraphBuilder.abs" exists.
func (this MLGraphBuilder) HasFuncAbs() bool {
	return js.True == bindings.HasFuncMLGraphBuilderAbs(
		this.ref,
	)
}

// FuncAbs returns the method "MLGraphBuilder.abs".
func (this MLGraphBuilder) FuncAbs() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderAbs(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Abs calls the method "MLGraphBuilder.abs".
func (this MLGraphBuilder) Abs(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderAbs(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryAbs calls the method "MLGraphBuilder.abs"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryAbs(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderAbs(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncCeil returns true if the method "MLGraphBuilder.ceil" exists.
func (this MLGraphBuilder) HasFuncCeil() bool {
	return js.True == bindings.HasFuncMLGraphBuilderCeil(
		this.ref,
	)
}

// FuncCeil returns the method "MLGraphBuilder.ceil".
func (this MLGraphBuilder) FuncCeil() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderCeil(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Ceil calls the method "MLGraphBuilder.ceil".
func (this MLGraphBuilder) Ceil(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderCeil(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryCeil calls the method "MLGraphBuilder.ceil"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryCeil(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderCeil(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncCos returns true if the method "MLGraphBuilder.cos" exists.
func (this MLGraphBuilder) HasFuncCos() bool {
	return js.True == bindings.HasFuncMLGraphBuilderCos(
		this.ref,
	)
}

// FuncCos returns the method "MLGraphBuilder.cos".
func (this MLGraphBuilder) FuncCos() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderCos(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Cos calls the method "MLGraphBuilder.cos".
func (this MLGraphBuilder) Cos(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderCos(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryCos calls the method "MLGraphBuilder.cos"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryCos(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderCos(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncExp returns true if the method "MLGraphBuilder.exp" exists.
func (this MLGraphBuilder) HasFuncExp() bool {
	return js.True == bindings.HasFuncMLGraphBuilderExp(
		this.ref,
	)
}

// FuncExp returns the method "MLGraphBuilder.exp".
func (this MLGraphBuilder) FuncExp() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderExp(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Exp calls the method "MLGraphBuilder.exp".
func (this MLGraphBuilder) Exp(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderExp(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryExp calls the method "MLGraphBuilder.exp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryExp(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderExp(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncFloor returns true if the method "MLGraphBuilder.floor" exists.
func (this MLGraphBuilder) HasFuncFloor() bool {
	return js.True == bindings.HasFuncMLGraphBuilderFloor(
		this.ref,
	)
}

// FuncFloor returns the method "MLGraphBuilder.floor".
func (this MLGraphBuilder) FuncFloor() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderFloor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Floor calls the method "MLGraphBuilder.floor".
func (this MLGraphBuilder) Floor(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderFloor(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryFloor calls the method "MLGraphBuilder.floor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryFloor(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderFloor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncLog returns true if the method "MLGraphBuilder.log" exists.
func (this MLGraphBuilder) HasFuncLog() bool {
	return js.True == bindings.HasFuncMLGraphBuilderLog(
		this.ref,
	)
}

// FuncLog returns the method "MLGraphBuilder.log".
func (this MLGraphBuilder) FuncLog() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderLog(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Log calls the method "MLGraphBuilder.log".
func (this MLGraphBuilder) Log(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderLog(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryLog calls the method "MLGraphBuilder.log"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLog(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLog(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncNeg returns true if the method "MLGraphBuilder.neg" exists.
func (this MLGraphBuilder) HasFuncNeg() bool {
	return js.True == bindings.HasFuncMLGraphBuilderNeg(
		this.ref,
	)
}

// FuncNeg returns the method "MLGraphBuilder.neg".
func (this MLGraphBuilder) FuncNeg() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderNeg(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Neg calls the method "MLGraphBuilder.neg".
func (this MLGraphBuilder) Neg(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderNeg(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryNeg calls the method "MLGraphBuilder.neg"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryNeg(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderNeg(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncSin returns true if the method "MLGraphBuilder.sin" exists.
func (this MLGraphBuilder) HasFuncSin() bool {
	return js.True == bindings.HasFuncMLGraphBuilderSin(
		this.ref,
	)
}

// FuncSin returns the method "MLGraphBuilder.sin".
func (this MLGraphBuilder) FuncSin() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderSin(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Sin calls the method "MLGraphBuilder.sin".
func (this MLGraphBuilder) Sin(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderSin(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySin calls the method "MLGraphBuilder.sin"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySin(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSin(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncTan returns true if the method "MLGraphBuilder.tan" exists.
func (this MLGraphBuilder) HasFuncTan() bool {
	return js.True == bindings.HasFuncMLGraphBuilderTan(
		this.ref,
	)
}

// FuncTan returns the method "MLGraphBuilder.tan".
func (this MLGraphBuilder) FuncTan() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderTan(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Tan calls the method "MLGraphBuilder.tan".
func (this MLGraphBuilder) Tan(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderTan(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryTan calls the method "MLGraphBuilder.tan"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryTan(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderTan(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncTranspose returns true if the method "MLGraphBuilder.transpose" exists.
func (this MLGraphBuilder) HasFuncTranspose() bool {
	return js.True == bindings.HasFuncMLGraphBuilderTranspose(
		this.ref,
	)
}

// FuncTranspose returns the method "MLGraphBuilder.transpose".
func (this MLGraphBuilder) FuncTranspose() (fn js.Func[func(input MLOperand, options MLTransposeOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderTranspose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Transpose calls the method "MLGraphBuilder.transpose".
func (this MLGraphBuilder) Transpose(input MLOperand, options MLTransposeOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderTranspose(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryTranspose calls the method "MLGraphBuilder.transpose"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryTranspose(input MLOperand, options MLTransposeOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderTranspose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncTranspose1 returns true if the method "MLGraphBuilder.transpose" exists.
func (this MLGraphBuilder) HasFuncTranspose1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderTranspose1(
		this.ref,
	)
}

// FuncTranspose1 returns the method "MLGraphBuilder.transpose".
func (this MLGraphBuilder) FuncTranspose1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderTranspose1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Transpose1 calls the method "MLGraphBuilder.transpose".
func (this MLGraphBuilder) Transpose1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderTranspose1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryTranspose1 calls the method "MLGraphBuilder.transpose"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryTranspose1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderTranspose1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncPrelu returns true if the method "MLGraphBuilder.prelu" exists.
func (this MLGraphBuilder) HasFuncPrelu() bool {
	return js.True == bindings.HasFuncMLGraphBuilderPrelu(
		this.ref,
	)
}

// FuncPrelu returns the method "MLGraphBuilder.prelu".
func (this MLGraphBuilder) FuncPrelu() (fn js.Func[func(input MLOperand, slope MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderPrelu(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Prelu calls the method "MLGraphBuilder.prelu".
func (this MLGraphBuilder) Prelu(input MLOperand, slope MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderPrelu(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		slope.Ref(),
	)

	return
}

// TryPrelu calls the method "MLGraphBuilder.prelu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryPrelu(input MLOperand, slope MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderPrelu(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		slope.Ref(),
	)

	return
}

// HasFuncConcat returns true if the method "MLGraphBuilder.concat" exists.
func (this MLGraphBuilder) HasFuncConcat() bool {
	return js.True == bindings.HasFuncMLGraphBuilderConcat(
		this.ref,
	)
}

// FuncConcat returns the method "MLGraphBuilder.concat".
func (this MLGraphBuilder) FuncConcat() (fn js.Func[func(inputs js.Array[MLOperand], axis uint32) MLOperand]) {
	bindings.FuncMLGraphBuilderConcat(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Concat calls the method "MLGraphBuilder.concat".
func (this MLGraphBuilder) Concat(inputs js.Array[MLOperand], axis uint32) (ret MLOperand) {
	bindings.CallMLGraphBuilderConcat(
		this.ref, js.Pointer(&ret),
		inputs.Ref(),
		uint32(axis),
	)

	return
}

// TryConcat calls the method "MLGraphBuilder.concat"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryConcat(inputs js.Array[MLOperand], axis uint32) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderConcat(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		inputs.Ref(),
		uint32(axis),
	)

	return
}

// HasFuncGemm returns true if the method "MLGraphBuilder.gemm" exists.
func (this MLGraphBuilder) HasFuncGemm() bool {
	return js.True == bindings.HasFuncMLGraphBuilderGemm(
		this.ref,
	)
}

// FuncGemm returns the method "MLGraphBuilder.gemm".
func (this MLGraphBuilder) FuncGemm() (fn js.Func[func(a MLOperand, b MLOperand, options MLGemmOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderGemm(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Gemm calls the method "MLGraphBuilder.gemm".
func (this MLGraphBuilder) Gemm(a MLOperand, b MLOperand, options MLGemmOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderGemm(
		this.ref, js.Pointer(&ret),
		a.Ref(),
		b.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryGemm calls the method "MLGraphBuilder.gemm"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryGemm(a MLOperand, b MLOperand, options MLGemmOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderGemm(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncGemm1 returns true if the method "MLGraphBuilder.gemm" exists.
func (this MLGraphBuilder) HasFuncGemm1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderGemm1(
		this.ref,
	)
}

// FuncGemm1 returns the method "MLGraphBuilder.gemm".
func (this MLGraphBuilder) FuncGemm1() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderGemm1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Gemm1 calls the method "MLGraphBuilder.gemm".
func (this MLGraphBuilder) Gemm1(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderGemm1(
		this.ref, js.Pointer(&ret),
		a.Ref(),
		b.Ref(),
	)

	return
}

// TryGemm1 calls the method "MLGraphBuilder.gemm"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryGemm1(a MLOperand, b MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderGemm1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasFuncLstmCell returns true if the method "MLGraphBuilder.lstmCell" exists.
func (this MLGraphBuilder) HasFuncLstmCell() bool {
	return js.True == bindings.HasFuncMLGraphBuilderLstmCell(
		this.ref,
	)
}

// FuncLstmCell returns the method "MLGraphBuilder.lstmCell".
func (this MLGraphBuilder) FuncLstmCell() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, cellState MLOperand, hiddenSize uint32, options MLLstmCellOptions) js.Array[MLOperand]]) {
	bindings.FuncMLGraphBuilderLstmCell(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LstmCell calls the method "MLGraphBuilder.lstmCell".
func (this MLGraphBuilder) LstmCell(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, cellState MLOperand, hiddenSize uint32, options MLLstmCellOptions) (ret js.Array[MLOperand]) {
	bindings.CallMLGraphBuilderLstmCell(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		hiddenState.Ref(),
		cellState.Ref(),
		uint32(hiddenSize),
		js.Pointer(&options),
	)

	return
}

// TryLstmCell calls the method "MLGraphBuilder.lstmCell"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLstmCell(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, cellState MLOperand, hiddenSize uint32, options MLLstmCellOptions) (ret js.Array[MLOperand], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLstmCell(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		hiddenState.Ref(),
		cellState.Ref(),
		uint32(hiddenSize),
		js.Pointer(&options),
	)

	return
}

// HasFuncLstmCell1 returns true if the method "MLGraphBuilder.lstmCell" exists.
func (this MLGraphBuilder) HasFuncLstmCell1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderLstmCell1(
		this.ref,
	)
}

// FuncLstmCell1 returns the method "MLGraphBuilder.lstmCell".
func (this MLGraphBuilder) FuncLstmCell1() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, cellState MLOperand, hiddenSize uint32) js.Array[MLOperand]]) {
	bindings.FuncMLGraphBuilderLstmCell1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LstmCell1 calls the method "MLGraphBuilder.lstmCell".
func (this MLGraphBuilder) LstmCell1(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, cellState MLOperand, hiddenSize uint32) (ret js.Array[MLOperand]) {
	bindings.CallMLGraphBuilderLstmCell1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		hiddenState.Ref(),
		cellState.Ref(),
		uint32(hiddenSize),
	)

	return
}

// TryLstmCell1 calls the method "MLGraphBuilder.lstmCell"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLstmCell1(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, cellState MLOperand, hiddenSize uint32) (ret js.Array[MLOperand], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLstmCell1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		hiddenState.Ref(),
		cellState.Ref(),
		uint32(hiddenSize),
	)

	return
}

// HasFuncBatchNormalization returns true if the method "MLGraphBuilder.batchNormalization" exists.
func (this MLGraphBuilder) HasFuncBatchNormalization() bool {
	return js.True == bindings.HasFuncMLGraphBuilderBatchNormalization(
		this.ref,
	)
}

// FuncBatchNormalization returns the method "MLGraphBuilder.batchNormalization".
func (this MLGraphBuilder) FuncBatchNormalization() (fn js.Func[func(input MLOperand, mean MLOperand, variance MLOperand, options MLBatchNormalizationOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderBatchNormalization(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BatchNormalization calls the method "MLGraphBuilder.batchNormalization".
func (this MLGraphBuilder) BatchNormalization(input MLOperand, mean MLOperand, variance MLOperand, options MLBatchNormalizationOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderBatchNormalization(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		mean.Ref(),
		variance.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryBatchNormalization calls the method "MLGraphBuilder.batchNormalization"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryBatchNormalization(input MLOperand, mean MLOperand, variance MLOperand, options MLBatchNormalizationOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderBatchNormalization(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		mean.Ref(),
		variance.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncBatchNormalization1 returns true if the method "MLGraphBuilder.batchNormalization" exists.
func (this MLGraphBuilder) HasFuncBatchNormalization1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderBatchNormalization1(
		this.ref,
	)
}

// FuncBatchNormalization1 returns the method "MLGraphBuilder.batchNormalization".
func (this MLGraphBuilder) FuncBatchNormalization1() (fn js.Func[func(input MLOperand, mean MLOperand, variance MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderBatchNormalization1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BatchNormalization1 calls the method "MLGraphBuilder.batchNormalization".
func (this MLGraphBuilder) BatchNormalization1(input MLOperand, mean MLOperand, variance MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderBatchNormalization1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		mean.Ref(),
		variance.Ref(),
	)

	return
}

// TryBatchNormalization1 calls the method "MLGraphBuilder.batchNormalization"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryBatchNormalization1(input MLOperand, mean MLOperand, variance MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderBatchNormalization1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		mean.Ref(),
		variance.Ref(),
	)

	return
}

// HasFuncElu returns true if the method "MLGraphBuilder.elu" exists.
func (this MLGraphBuilder) HasFuncElu() bool {
	return js.True == bindings.HasFuncMLGraphBuilderElu(
		this.ref,
	)
}

// FuncElu returns the method "MLGraphBuilder.elu".
func (this MLGraphBuilder) FuncElu() (fn js.Func[func(input MLOperand, options MLEluOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderElu(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Elu calls the method "MLGraphBuilder.elu".
func (this MLGraphBuilder) Elu(input MLOperand, options MLEluOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderElu(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryElu calls the method "MLGraphBuilder.elu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryElu(input MLOperand, options MLEluOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderElu(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncElu1 returns true if the method "MLGraphBuilder.elu" exists.
func (this MLGraphBuilder) HasFuncElu1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderElu1(
		this.ref,
	)
}

// FuncElu1 returns the method "MLGraphBuilder.elu".
func (this MLGraphBuilder) FuncElu1() (fn js.Func[func(input MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderElu1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Elu1 calls the method "MLGraphBuilder.elu".
func (this MLGraphBuilder) Elu1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderElu1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryElu1 calls the method "MLGraphBuilder.elu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryElu1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderElu1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncElu2 returns true if the method "MLGraphBuilder.elu" exists.
func (this MLGraphBuilder) HasFuncElu2() bool {
	return js.True == bindings.HasFuncMLGraphBuilderElu2(
		this.ref,
	)
}

// FuncElu2 returns the method "MLGraphBuilder.elu".
func (this MLGraphBuilder) FuncElu2() (fn js.Func[func(options MLEluOptions) MLActivation]) {
	bindings.FuncMLGraphBuilderElu2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Elu2 calls the method "MLGraphBuilder.elu".
func (this MLGraphBuilder) Elu2(options MLEluOptions) (ret MLActivation) {
	bindings.CallMLGraphBuilderElu2(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryElu2 calls the method "MLGraphBuilder.elu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryElu2(options MLEluOptions) (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderElu2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncElu3 returns true if the method "MLGraphBuilder.elu" exists.
func (this MLGraphBuilder) HasFuncElu3() bool {
	return js.True == bindings.HasFuncMLGraphBuilderElu3(
		this.ref,
	)
}

// FuncElu3 returns the method "MLGraphBuilder.elu".
func (this MLGraphBuilder) FuncElu3() (fn js.Func[func() MLActivation]) {
	bindings.FuncMLGraphBuilderElu3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Elu3 calls the method "MLGraphBuilder.elu".
func (this MLGraphBuilder) Elu3() (ret MLActivation) {
	bindings.CallMLGraphBuilderElu3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryElu3 calls the method "MLGraphBuilder.elu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryElu3() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderElu3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClamp returns true if the method "MLGraphBuilder.clamp" exists.
func (this MLGraphBuilder) HasFuncClamp() bool {
	return js.True == bindings.HasFuncMLGraphBuilderClamp(
		this.ref,
	)
}

// FuncClamp returns the method "MLGraphBuilder.clamp".
func (this MLGraphBuilder) FuncClamp() (fn js.Func[func(operand MLOperand, options MLClampOptions) MLOperand]) {
	bindings.FuncMLGraphBuilderClamp(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clamp calls the method "MLGraphBuilder.clamp".
func (this MLGraphBuilder) Clamp(operand MLOperand, options MLClampOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderClamp(
		this.ref, js.Pointer(&ret),
		operand.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryClamp calls the method "MLGraphBuilder.clamp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryClamp(operand MLOperand, options MLClampOptions) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderClamp(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		operand.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncClamp1 returns true if the method "MLGraphBuilder.clamp" exists.
func (this MLGraphBuilder) HasFuncClamp1() bool {
	return js.True == bindings.HasFuncMLGraphBuilderClamp1(
		this.ref,
	)
}

// FuncClamp1 returns the method "MLGraphBuilder.clamp".
func (this MLGraphBuilder) FuncClamp1() (fn js.Func[func(operand MLOperand) MLOperand]) {
	bindings.FuncMLGraphBuilderClamp1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clamp1 calls the method "MLGraphBuilder.clamp".
func (this MLGraphBuilder) Clamp1(operand MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderClamp1(
		this.ref, js.Pointer(&ret),
		operand.Ref(),
	)

	return
}

// TryClamp1 calls the method "MLGraphBuilder.clamp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryClamp1(operand MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderClamp1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		operand.Ref(),
	)

	return
}

// HasFuncClamp2 returns true if the method "MLGraphBuilder.clamp" exists.
func (this MLGraphBuilder) HasFuncClamp2() bool {
	return js.True == bindings.HasFuncMLGraphBuilderClamp2(
		this.ref,
	)
}

// FuncClamp2 returns the method "MLGraphBuilder.clamp".
func (this MLGraphBuilder) FuncClamp2() (fn js.Func[func(options MLClampOptions) MLActivation]) {
	bindings.FuncMLGraphBuilderClamp2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clamp2 calls the method "MLGraphBuilder.clamp".
func (this MLGraphBuilder) Clamp2(options MLClampOptions) (ret MLActivation) {
	bindings.CallMLGraphBuilderClamp2(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryClamp2 calls the method "MLGraphBuilder.clamp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryClamp2(options MLClampOptions) (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderClamp2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncClamp3 returns true if the method "MLGraphBuilder.clamp" exists.
func (this MLGraphBuilder) HasFuncClamp3() bool {
	return js.True == bindings.HasFuncMLGraphBuilderClamp3(
		this.ref,
	)
}

// FuncClamp3 returns the method "MLGraphBuilder.clamp".
func (this MLGraphBuilder) FuncClamp3() (fn js.Func[func() MLActivation]) {
	bindings.FuncMLGraphBuilderClamp3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clamp3 calls the method "MLGraphBuilder.clamp".
func (this MLGraphBuilder) Clamp3() (ret MLActivation) {
	bindings.CallMLGraphBuilderClamp3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClamp3 calls the method "MLGraphBuilder.clamp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryClamp3() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderClamp3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type MagnetometerLocalCoordinateSystem uint32

const (
	_ MagnetometerLocalCoordinateSystem = iota

	MagnetometerLocalCoordinateSystem_DEVICE
	MagnetometerLocalCoordinateSystem_SCREEN
)

func (MagnetometerLocalCoordinateSystem) FromRef(str js.Ref) MagnetometerLocalCoordinateSystem {
	return MagnetometerLocalCoordinateSystem(bindings.ConstOfMagnetometerLocalCoordinateSystem(str))
}

func (x MagnetometerLocalCoordinateSystem) String() (string, bool) {
	switch x {
	case MagnetometerLocalCoordinateSystem_DEVICE:
		return "device", true
	case MagnetometerLocalCoordinateSystem_SCREEN:
		return "screen", true
	default:
		return "", false
	}
}

type MagnetometerSensorOptions struct {
	// ReferenceFrame is "MagnetometerSensorOptions.referenceFrame"
	//
	// Optional, defaults to "device".
	ReferenceFrame MagnetometerLocalCoordinateSystem
	// Frequency is "MagnetometerSensorOptions.frequency"
	//
	// Optional
	//
	// NOTE: FFI_USE_Frequency MUST be set to true to make this field effective.
	Frequency float64

	FFI_USE_Frequency bool // for Frequency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MagnetometerSensorOptions with all fields set.
func (p MagnetometerSensorOptions) FromRef(ref js.Ref) MagnetometerSensorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MagnetometerSensorOptions in the application heap.
func (p MagnetometerSensorOptions) New() js.Ref {
	return bindings.MagnetometerSensorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MagnetometerSensorOptions) UpdateFrom(ref js.Ref) {
	bindings.MagnetometerSensorOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MagnetometerSensorOptions) Update(ref js.Ref) {
	bindings.MagnetometerSensorOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MagnetometerSensorOptions) FreeMembers(recursive bool) {
}

func NewMagnetometer(sensorOptions MagnetometerSensorOptions) (ret Magnetometer) {
	ret.ref = bindings.NewMagnetometerByMagnetometer(
		js.Pointer(&sensorOptions))
	return
}

func NewMagnetometerByMagnetometer1() (ret Magnetometer) {
	ret.ref = bindings.NewMagnetometerByMagnetometer1()
	return
}

type Magnetometer struct {
	Sensor
}

func (this Magnetometer) Once() Magnetometer {
	this.ref.Once()
	return this
}

func (this Magnetometer) Ref() js.Ref {
	return this.Sensor.Ref()
}

func (this Magnetometer) FromRef(ref js.Ref) Magnetometer {
	this.Sensor = this.Sensor.FromRef(ref)
	return this
}

func (this Magnetometer) Free() {
	this.ref.Free()
}

// X returns the value of property "Magnetometer.x".
//
// It returns ok=false if there is no such property.
func (this Magnetometer) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetMagnetometerX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "Magnetometer.y".
//
// It returns ok=false if there is no such property.
func (this Magnetometer) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetMagnetometerY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Z returns the value of property "Magnetometer.z".
//
// It returns ok=false if there is no such property.
func (this Magnetometer) Z() (ret float64, ok bool) {
	ok = js.True == bindings.GetMagnetometerZ(
		this.ref, js.Pointer(&ret),
	)
	return
}

type MagnetometerReadingValues struct {
	// X is "MagnetometerReadingValues.x"
	//
	// Required
	X float64
	// Y is "MagnetometerReadingValues.y"
	//
	// Required
	Y float64
	// Z is "MagnetometerReadingValues.z"
	//
	// Required
	Z float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MagnetometerReadingValues with all fields set.
func (p MagnetometerReadingValues) FromRef(ref js.Ref) MagnetometerReadingValues {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MagnetometerReadingValues in the application heap.
func (p MagnetometerReadingValues) New() js.Ref {
	return bindings.MagnetometerReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MagnetometerReadingValues) UpdateFrom(ref js.Ref) {
	bindings.MagnetometerReadingValuesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MagnetometerReadingValues) Update(ref js.Ref) {
	bindings.MagnetometerReadingValuesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MagnetometerReadingValues) FreeMembers(recursive bool) {
}

type MathMLElement struct {
	Element
}

func (this MathMLElement) Once() MathMLElement {
	this.ref.Once()
	return this
}

func (this MathMLElement) Ref() js.Ref {
	return this.Element.Ref()
}

func (this MathMLElement) FromRef(ref js.Ref) MathMLElement {
	this.Element = this.Element.FromRef(ref)
	return this
}

func (this MathMLElement) Free() {
	this.ref.Free()
}

// Style returns the value of property "MathMLElement.style".
//
// It returns ok=false if there is no such property.
func (this MathMLElement) Style() (ret CSSStyleDeclaration, ok bool) {
	ok = js.True == bindings.GetMathMLElementStyle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AttributeStyleMap returns the value of property "MathMLElement.attributeStyleMap".
//
// It returns ok=false if there is no such property.
func (this MathMLElement) AttributeStyleMap() (ret StylePropertyMap, ok bool) {
	ok = js.True == bindings.GetMathMLElementAttributeStyleMap(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Dataset returns the value of property "MathMLElement.dataset".
//
// It returns ok=false if there is no such property.
func (this MathMLElement) Dataset() (ret DOMStringMap, ok bool) {
	ok = js.True == bindings.GetMathMLElementDataset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Nonce returns the value of property "MathMLElement.nonce".
//
// It returns ok=false if there is no such property.
func (this MathMLElement) Nonce() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMathMLElementNonce(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNonce sets the value of property "MathMLElement.nonce" to val.
//
// It returns false if the property cannot be set.
func (this MathMLElement) SetNonce(val js.String) bool {
	return js.True == bindings.SetMathMLElementNonce(
		this.ref,
		val.Ref(),
	)
}

// Autofocus returns the value of property "MathMLElement.autofocus".
//
// It returns ok=false if there is no such property.
func (this MathMLElement) Autofocus() (ret bool, ok bool) {
	ok = js.True == bindings.GetMathMLElementAutofocus(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAutofocus sets the value of property "MathMLElement.autofocus" to val.
//
// It returns false if the property cannot be set.
func (this MathMLElement) SetAutofocus(val bool) bool {
	return js.True == bindings.SetMathMLElementAutofocus(
		this.ref,
		js.Bool(bool(val)),
	)
}

// TabIndex returns the value of property "MathMLElement.tabIndex".
//
// It returns ok=false if there is no such property.
func (this MathMLElement) TabIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetMathMLElementTabIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTabIndex sets the value of property "MathMLElement.tabIndex" to val.
//
// It returns false if the property cannot be set.
func (this MathMLElement) SetTabIndex(val int32) bool {
	return js.True == bindings.SetMathMLElementTabIndex(
		this.ref,
		int32(val),
	)
}

// HasFuncFocus returns true if the method "MathMLElement.focus" exists.
func (this MathMLElement) HasFuncFocus() bool {
	return js.True == bindings.HasFuncMathMLElementFocus(
		this.ref,
	)
}

// FuncFocus returns the method "MathMLElement.focus".
func (this MathMLElement) FuncFocus() (fn js.Func[func(options FocusOptions)]) {
	bindings.FuncMathMLElementFocus(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Focus calls the method "MathMLElement.focus".
func (this MathMLElement) Focus(options FocusOptions) (ret js.Void) {
	bindings.CallMathMLElementFocus(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryFocus calls the method "MathMLElement.focus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MathMLElement) TryFocus(options FocusOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMathMLElementFocus(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncFocus1 returns true if the method "MathMLElement.focus" exists.
func (this MathMLElement) HasFuncFocus1() bool {
	return js.True == bindings.HasFuncMathMLElementFocus1(
		this.ref,
	)
}

// FuncFocus1 returns the method "MathMLElement.focus".
func (this MathMLElement) FuncFocus1() (fn js.Func[func()]) {
	bindings.FuncMathMLElementFocus1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Focus1 calls the method "MathMLElement.focus".
func (this MathMLElement) Focus1() (ret js.Void) {
	bindings.CallMathMLElementFocus1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFocus1 calls the method "MathMLElement.focus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MathMLElement) TryFocus1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMathMLElementFocus1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncBlur returns true if the method "MathMLElement.blur" exists.
func (this MathMLElement) HasFuncBlur() bool {
	return js.True == bindings.HasFuncMathMLElementBlur(
		this.ref,
	)
}

// FuncBlur returns the method "MathMLElement.blur".
func (this MathMLElement) FuncBlur() (fn js.Func[func()]) {
	bindings.FuncMathMLElementBlur(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Blur calls the method "MathMLElement.blur".
func (this MathMLElement) Blur() (ret js.Void) {
	bindings.CallMathMLElementBlur(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryBlur calls the method "MathMLElement.blur"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MathMLElement) TryBlur() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMathMLElementBlur(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type MediaCapabilitiesInfo struct {
	// Supported is "MediaCapabilitiesInfo.supported"
	//
	// Required
	Supported bool
	// Smooth is "MediaCapabilitiesInfo.smooth"
	//
	// Required
	Smooth bool
	// PowerEfficient is "MediaCapabilitiesInfo.powerEfficient"
	//
	// Required
	PowerEfficient bool

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaCapabilitiesInfo with all fields set.
func (p MediaCapabilitiesInfo) FromRef(ref js.Ref) MediaCapabilitiesInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaCapabilitiesInfo in the application heap.
func (p MediaCapabilitiesInfo) New() js.Ref {
	return bindings.MediaCapabilitiesInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MediaCapabilitiesInfo) UpdateFrom(ref js.Ref) {
	bindings.MediaCapabilitiesInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaCapabilitiesInfo) Update(ref js.Ref) {
	bindings.MediaCapabilitiesInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaCapabilitiesInfo) FreeMembers(recursive bool) {
}

type MediaConfiguration struct {
	// Video is "MediaConfiguration.video"
	//
	// Optional
	//
	// NOTE: Video.FFI_USE MUST be set to true to get Video used.
	Video VideoConfiguration
	// Audio is "MediaConfiguration.audio"
	//
	// Optional
	//
	// NOTE: Audio.FFI_USE MUST be set to true to get Audio used.
	Audio AudioConfiguration

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaConfiguration with all fields set.
func (p MediaConfiguration) FromRef(ref js.Ref) MediaConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaConfiguration in the application heap.
func (p MediaConfiguration) New() js.Ref {
	return bindings.MediaConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MediaConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MediaConfigurationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaConfiguration) Update(ref js.Ref) {
	bindings.MediaConfigurationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaConfiguration) FreeMembers(recursive bool) {
	if recursive {
		p.Video.FreeMembers(true)
		p.Audio.FreeMembers(true)
	}
}

type MediaEncryptedEventInit struct {
	// InitDataType is "MediaEncryptedEventInit.initDataType"
	//
	// Optional, defaults to "".
	InitDataType js.String
	// InitData is "MediaEncryptedEventInit.initData"
	//
	// Optional, defaults to null.
	InitData js.ArrayBuffer
	// Bubbles is "MediaEncryptedEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "MediaEncryptedEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "MediaEncryptedEventInit.composed"
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

// FromRef calls UpdateFrom and returns a MediaEncryptedEventInit with all fields set.
func (p MediaEncryptedEventInit) FromRef(ref js.Ref) MediaEncryptedEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaEncryptedEventInit in the application heap.
func (p MediaEncryptedEventInit) New() js.Ref {
	return bindings.MediaEncryptedEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MediaEncryptedEventInit) UpdateFrom(ref js.Ref) {
	bindings.MediaEncryptedEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaEncryptedEventInit) Update(ref js.Ref) {
	bindings.MediaEncryptedEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaEncryptedEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.InitDataType.Ref(),
		p.InitData.Ref(),
	)
	p.InitDataType = p.InitDataType.FromRef(js.Undefined)
	p.InitData = p.InitData.FromRef(js.Undefined)
}

func NewMediaEncryptedEvent(typ js.String, eventInitDict MediaEncryptedEventInit) (ret MediaEncryptedEvent) {
	ret.ref = bindings.NewMediaEncryptedEventByMediaEncryptedEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewMediaEncryptedEventByMediaEncryptedEvent1(typ js.String) (ret MediaEncryptedEvent) {
	ret.ref = bindings.NewMediaEncryptedEventByMediaEncryptedEvent1(
		typ.Ref())
	return
}

type MediaEncryptedEvent struct {
	Event
}

func (this MediaEncryptedEvent) Once() MediaEncryptedEvent {
	this.ref.Once()
	return this
}

func (this MediaEncryptedEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this MediaEncryptedEvent) FromRef(ref js.Ref) MediaEncryptedEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this MediaEncryptedEvent) Free() {
	this.ref.Free()
}

// InitDataType returns the value of property "MediaEncryptedEvent.initDataType".
//
// It returns ok=false if there is no such property.
func (this MediaEncryptedEvent) InitDataType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaEncryptedEventInitDataType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InitData returns the value of property "MediaEncryptedEvent.initData".
//
// It returns ok=false if there is no such property.
func (this MediaEncryptedEvent) InitData() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetMediaEncryptedEventInitData(
		this.ref, js.Pointer(&ret),
	)
	return
}

type MediaKeyMessageType uint32

const (
	_ MediaKeyMessageType = iota

	MediaKeyMessageType_LICENSE_REQUEST
	MediaKeyMessageType_LICENSE_RENEWAL
	MediaKeyMessageType_LICENSE_RELEASE
	MediaKeyMessageType_INDIVIDUALIZATION_REQUEST
)

func (MediaKeyMessageType) FromRef(str js.Ref) MediaKeyMessageType {
	return MediaKeyMessageType(bindings.ConstOfMediaKeyMessageType(str))
}

func (x MediaKeyMessageType) String() (string, bool) {
	switch x {
	case MediaKeyMessageType_LICENSE_REQUEST:
		return "license-request", true
	case MediaKeyMessageType_LICENSE_RENEWAL:
		return "license-renewal", true
	case MediaKeyMessageType_LICENSE_RELEASE:
		return "license-release", true
	case MediaKeyMessageType_INDIVIDUALIZATION_REQUEST:
		return "individualization-request", true
	default:
		return "", false
	}
}

type MediaKeyMessageEventInit struct {
	// MessageType is "MediaKeyMessageEventInit.messageType"
	//
	// Required
	MessageType MediaKeyMessageType
	// Message is "MediaKeyMessageEventInit.message"
	//
	// Required
	Message js.ArrayBuffer
	// Bubbles is "MediaKeyMessageEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "MediaKeyMessageEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "MediaKeyMessageEventInit.composed"
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

// FromRef calls UpdateFrom and returns a MediaKeyMessageEventInit with all fields set.
func (p MediaKeyMessageEventInit) FromRef(ref js.Ref) MediaKeyMessageEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaKeyMessageEventInit in the application heap.
func (p MediaKeyMessageEventInit) New() js.Ref {
	return bindings.MediaKeyMessageEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MediaKeyMessageEventInit) UpdateFrom(ref js.Ref) {
	bindings.MediaKeyMessageEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaKeyMessageEventInit) Update(ref js.Ref) {
	bindings.MediaKeyMessageEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaKeyMessageEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Message.Ref(),
	)
	p.Message = p.Message.FromRef(js.Undefined)
}

func NewMediaKeyMessageEvent(typ js.String, eventInitDict MediaKeyMessageEventInit) (ret MediaKeyMessageEvent) {
	ret.ref = bindings.NewMediaKeyMessageEventByMediaKeyMessageEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type MediaKeyMessageEvent struct {
	Event
}

func (this MediaKeyMessageEvent) Once() MediaKeyMessageEvent {
	this.ref.Once()
	return this
}

func (this MediaKeyMessageEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this MediaKeyMessageEvent) FromRef(ref js.Ref) MediaKeyMessageEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this MediaKeyMessageEvent) Free() {
	this.ref.Free()
}

// MessageType returns the value of property "MediaKeyMessageEvent.messageType".
//
// It returns ok=false if there is no such property.
func (this MediaKeyMessageEvent) MessageType() (ret MediaKeyMessageType, ok bool) {
	ok = js.True == bindings.GetMediaKeyMessageEventMessageType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Message returns the value of property "MediaKeyMessageEvent.message".
//
// It returns ok=false if there is no such property.
func (this MediaKeyMessageEvent) Message() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetMediaKeyMessageEventMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

type MediaQueryListEventInit struct {
	// Media is "MediaQueryListEventInit.media"
	//
	// Optional, defaults to "".
	Media js.String
	// Matches is "MediaQueryListEventInit.matches"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Matches MUST be set to true to make this field effective.
	Matches bool
	// Bubbles is "MediaQueryListEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "MediaQueryListEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "MediaQueryListEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Matches    bool // for Matches.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaQueryListEventInit with all fields set.
func (p MediaQueryListEventInit) FromRef(ref js.Ref) MediaQueryListEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaQueryListEventInit in the application heap.
func (p MediaQueryListEventInit) New() js.Ref {
	return bindings.MediaQueryListEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MediaQueryListEventInit) UpdateFrom(ref js.Ref) {
	bindings.MediaQueryListEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaQueryListEventInit) Update(ref js.Ref) {
	bindings.MediaQueryListEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaQueryListEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Media.Ref(),
	)
	p.Media = p.Media.FromRef(js.Undefined)
}

func NewMediaQueryListEvent(typ js.String, eventInitDict MediaQueryListEventInit) (ret MediaQueryListEvent) {
	ret.ref = bindings.NewMediaQueryListEventByMediaQueryListEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewMediaQueryListEventByMediaQueryListEvent1(typ js.String) (ret MediaQueryListEvent) {
	ret.ref = bindings.NewMediaQueryListEventByMediaQueryListEvent1(
		typ.Ref())
	return
}

type MediaQueryListEvent struct {
	Event
}

func (this MediaQueryListEvent) Once() MediaQueryListEvent {
	this.ref.Once()
	return this
}

func (this MediaQueryListEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this MediaQueryListEvent) FromRef(ref js.Ref) MediaQueryListEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this MediaQueryListEvent) Free() {
	this.ref.Free()
}

// Media returns the value of property "MediaQueryListEvent.media".
//
// It returns ok=false if there is no such property.
func (this MediaQueryListEvent) Media() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaQueryListEventMedia(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Matches returns the value of property "MediaQueryListEvent.matches".
//
// It returns ok=false if there is no such property.
func (this MediaQueryListEvent) Matches() (ret bool, ok bool) {
	ok = js.True == bindings.GetMediaQueryListEventMatches(
		this.ref, js.Pointer(&ret),
	)
	return
}

type MediaRecorderOptions struct {
	// MimeType is "MediaRecorderOptions.mimeType"
	//
	// Optional, defaults to "".
	MimeType js.String
	// AudioBitsPerSecond is "MediaRecorderOptions.audioBitsPerSecond"
	//
	// Optional
	//
	// NOTE: FFI_USE_AudioBitsPerSecond MUST be set to true to make this field effective.
	AudioBitsPerSecond uint32
	// VideoBitsPerSecond is "MediaRecorderOptions.videoBitsPerSecond"
	//
	// Optional
	//
	// NOTE: FFI_USE_VideoBitsPerSecond MUST be set to true to make this field effective.
	VideoBitsPerSecond uint32
	// BitsPerSecond is "MediaRecorderOptions.bitsPerSecond"
	//
	// Optional
	//
	// NOTE: FFI_USE_BitsPerSecond MUST be set to true to make this field effective.
	BitsPerSecond uint32
	// AudioBitrateMode is "MediaRecorderOptions.audioBitrateMode"
	//
	// Optional, defaults to "variable".
	AudioBitrateMode BitrateMode
	// VideoKeyFrameIntervalDuration is "MediaRecorderOptions.videoKeyFrameIntervalDuration"
	//
	// Optional
	//
	// NOTE: FFI_USE_VideoKeyFrameIntervalDuration MUST be set to true to make this field effective.
	VideoKeyFrameIntervalDuration DOMHighResTimeStamp
	// VideoKeyFrameIntervalCount is "MediaRecorderOptions.videoKeyFrameIntervalCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_VideoKeyFrameIntervalCount MUST be set to true to make this field effective.
	VideoKeyFrameIntervalCount uint32

	FFI_USE_AudioBitsPerSecond            bool // for AudioBitsPerSecond.
	FFI_USE_VideoBitsPerSecond            bool // for VideoBitsPerSecond.
	FFI_USE_BitsPerSecond                 bool // for BitsPerSecond.
	FFI_USE_VideoKeyFrameIntervalDuration bool // for VideoKeyFrameIntervalDuration.
	FFI_USE_VideoKeyFrameIntervalCount    bool // for VideoKeyFrameIntervalCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaRecorderOptions with all fields set.
func (p MediaRecorderOptions) FromRef(ref js.Ref) MediaRecorderOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaRecorderOptions in the application heap.
func (p MediaRecorderOptions) New() js.Ref {
	return bindings.MediaRecorderOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MediaRecorderOptions) UpdateFrom(ref js.Ref) {
	bindings.MediaRecorderOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaRecorderOptions) Update(ref js.Ref) {
	bindings.MediaRecorderOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaRecorderOptions) FreeMembers(recursive bool) {
	js.Free(
		p.MimeType.Ref(),
	)
	p.MimeType = p.MimeType.FromRef(js.Undefined)
}

type RecordingState uint32

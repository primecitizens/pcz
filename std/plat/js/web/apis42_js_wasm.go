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
func (p MLOperandDescriptor) UpdateFrom(ref js.Ref) {
	bindings.MLOperandDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLOperandDescriptor) Update(ref js.Ref) {
	bindings.MLOperandDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MLHardSigmoidOptions) UpdateFrom(ref js.Ref) {
	bindings.MLHardSigmoidOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLHardSigmoidOptions) Update(ref js.Ref) {
	bindings.MLHardSigmoidOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MLGruCellOptions) UpdateFrom(ref js.Ref) {
	bindings.MLGruCellOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLGruCellOptions) Update(ref js.Ref) {
	bindings.MLGruCellOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MLPool2dOptions) UpdateFrom(ref js.Ref) {
	bindings.MLPool2dOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLPool2dOptions) Update(ref js.Ref) {
	bindings.MLPool2dOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MLLinearOptions) UpdateFrom(ref js.Ref) {
	bindings.MLLinearOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLLinearOptions) Update(ref js.Ref) {
	bindings.MLLinearOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MLLeakyReluOptions) UpdateFrom(ref js.Ref) {
	bindings.MLLeakyReluOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLLeakyReluOptions) Update(ref js.Ref) {
	bindings.MLLeakyReluOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MLPadOptions) UpdateFrom(ref js.Ref) {
	bindings.MLPadOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLPadOptions) Update(ref js.Ref) {
	bindings.MLPadOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MLInstanceNormalizationOptions) UpdateFrom(ref js.Ref) {
	bindings.MLInstanceNormalizationOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLInstanceNormalizationOptions) Update(ref js.Ref) {
	bindings.MLInstanceNormalizationOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MLSoftplusOptions) UpdateFrom(ref js.Ref) {
	bindings.MLSoftplusOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLSoftplusOptions) Update(ref js.Ref) {
	bindings.MLSoftplusOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MLSplitOptions) UpdateFrom(ref js.Ref) {
	bindings.MLSplitOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLSplitOptions) Update(ref js.Ref) {
	bindings.MLSplitOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MLResample2dOptions) UpdateFrom(ref js.Ref) {
	bindings.MLResample2dOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLResample2dOptions) Update(ref js.Ref) {
	bindings.MLResample2dOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MLReduceOptions) UpdateFrom(ref js.Ref) {
	bindings.MLReduceOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLReduceOptions) Update(ref js.Ref) {
	bindings.MLReduceOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MLLstmOptions) UpdateFrom(ref js.Ref) {
	bindings.MLLstmOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLLstmOptions) Update(ref js.Ref) {
	bindings.MLLstmOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MLSqueezeOptions) UpdateFrom(ref js.Ref) {
	bindings.MLSqueezeOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLSqueezeOptions) Update(ref js.Ref) {
	bindings.MLSqueezeOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MLGruOptions) UpdateFrom(ref js.Ref) {
	bindings.MLGruOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLGruOptions) Update(ref js.Ref) {
	bindings.MLGruOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MLTransposeOptions) UpdateFrom(ref js.Ref) {
	bindings.MLTransposeOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLTransposeOptions) Update(ref js.Ref) {
	bindings.MLTransposeOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MLLstmCellOptions) UpdateFrom(ref js.Ref) {
	bindings.MLLstmCellOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MLLstmCellOptions) Update(ref js.Ref) {
	bindings.MLLstmCellOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// HasInput returns true if the method "MLGraphBuilder.input" exists.
func (this MLGraphBuilder) HasInput() bool {
	return js.True == bindings.HasMLGraphBuilderInput(
		this.Ref(),
	)
}

// InputFunc returns the method "MLGraphBuilder.input".
func (this MLGraphBuilder) InputFunc() (fn js.Func[func(name js.String, descriptor MLOperandDescriptor) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderInputFunc(
			this.Ref(),
		),
	)
}

// Input calls the method "MLGraphBuilder.input".
func (this MLGraphBuilder) Input(name js.String, descriptor MLOperandDescriptor) (ret MLOperand) {
	bindings.CallMLGraphBuilderInput(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		js.Pointer(&descriptor),
	)

	return
}

// HasConstant returns true if the method "MLGraphBuilder.constant" exists.
func (this MLGraphBuilder) HasConstant() bool {
	return js.True == bindings.HasMLGraphBuilderConstant(
		this.Ref(),
	)
}

// ConstantFunc returns the method "MLGraphBuilder.constant".
func (this MLGraphBuilder) ConstantFunc() (fn js.Func[func(descriptor MLOperandDescriptor, bufferView MLBufferView) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderConstantFunc(
			this.Ref(),
		),
	)
}

// Constant calls the method "MLGraphBuilder.constant".
func (this MLGraphBuilder) Constant(descriptor MLOperandDescriptor, bufferView MLBufferView) (ret MLOperand) {
	bindings.CallMLGraphBuilderConstant(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
		bufferView.Ref(),
	)

	return
}

// HasConstant1 returns true if the method "MLGraphBuilder.constant" exists.
func (this MLGraphBuilder) HasConstant1() bool {
	return js.True == bindings.HasMLGraphBuilderConstant1(
		this.Ref(),
	)
}

// Constant1Func returns the method "MLGraphBuilder.constant".
func (this MLGraphBuilder) Constant1Func() (fn js.Func[func(value float64, typ MLOperandType) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderConstant1Func(
			this.Ref(),
		),
	)
}

// Constant1 calls the method "MLGraphBuilder.constant".
func (this MLGraphBuilder) Constant1(value float64, typ MLOperandType) (ret MLOperand) {
	bindings.CallMLGraphBuilderConstant1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
		uint32(typ),
	)

	return
}

// HasConstant2 returns true if the method "MLGraphBuilder.constant" exists.
func (this MLGraphBuilder) HasConstant2() bool {
	return js.True == bindings.HasMLGraphBuilderConstant2(
		this.Ref(),
	)
}

// Constant2Func returns the method "MLGraphBuilder.constant".
func (this MLGraphBuilder) Constant2Func() (fn js.Func[func(value float64) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderConstant2Func(
			this.Ref(),
		),
	)
}

// Constant2 calls the method "MLGraphBuilder.constant".
func (this MLGraphBuilder) Constant2(value float64) (ret MLOperand) {
	bindings.CallMLGraphBuilderConstant2(
		this.Ref(), js.Pointer(&ret),
		float64(value),
	)

	return
}

// TryConstant2 calls the method "MLGraphBuilder.constant"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryConstant2(value float64) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderConstant2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
	)

	return
}

// HasBuild returns true if the method "MLGraphBuilder.build" exists.
func (this MLGraphBuilder) HasBuild() bool {
	return js.True == bindings.HasMLGraphBuilderBuild(
		this.Ref(),
	)
}

// BuildFunc returns the method "MLGraphBuilder.build".
func (this MLGraphBuilder) BuildFunc() (fn js.Func[func(outputs MLNamedOperands) js.Promise[MLGraph]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderBuildFunc(
			this.Ref(),
		),
	)
}

// Build calls the method "MLGraphBuilder.build".
func (this MLGraphBuilder) Build(outputs MLNamedOperands) (ret js.Promise[MLGraph]) {
	bindings.CallMLGraphBuilderBuild(
		this.Ref(), js.Pointer(&ret),
		outputs.Ref(),
	)

	return
}

// TryBuild calls the method "MLGraphBuilder.build"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryBuild(outputs MLNamedOperands) (ret js.Promise[MLGraph], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderBuild(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		outputs.Ref(),
	)

	return
}

// HasBuildSync returns true if the method "MLGraphBuilder.buildSync" exists.
func (this MLGraphBuilder) HasBuildSync() bool {
	return js.True == bindings.HasMLGraphBuilderBuildSync(
		this.Ref(),
	)
}

// BuildSyncFunc returns the method "MLGraphBuilder.buildSync".
func (this MLGraphBuilder) BuildSyncFunc() (fn js.Func[func(outputs MLNamedOperands) MLGraph]) {
	return fn.FromRef(
		bindings.MLGraphBuilderBuildSyncFunc(
			this.Ref(),
		),
	)
}

// BuildSync calls the method "MLGraphBuilder.buildSync".
func (this MLGraphBuilder) BuildSync(outputs MLNamedOperands) (ret MLGraph) {
	bindings.CallMLGraphBuilderBuildSync(
		this.Ref(), js.Pointer(&ret),
		outputs.Ref(),
	)

	return
}

// TryBuildSync calls the method "MLGraphBuilder.buildSync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryBuildSync(outputs MLNamedOperands) (ret MLGraph, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderBuildSync(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		outputs.Ref(),
	)

	return
}

// HasHardSigmoid returns true if the method "MLGraphBuilder.hardSigmoid" exists.
func (this MLGraphBuilder) HasHardSigmoid() bool {
	return js.True == bindings.HasMLGraphBuilderHardSigmoid(
		this.Ref(),
	)
}

// HardSigmoidFunc returns the method "MLGraphBuilder.hardSigmoid".
func (this MLGraphBuilder) HardSigmoidFunc() (fn js.Func[func(input MLOperand, options MLHardSigmoidOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderHardSigmoidFunc(
			this.Ref(),
		),
	)
}

// HardSigmoid calls the method "MLGraphBuilder.hardSigmoid".
func (this MLGraphBuilder) HardSigmoid(input MLOperand, options MLHardSigmoidOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderHardSigmoid(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasHardSigmoid1 returns true if the method "MLGraphBuilder.hardSigmoid" exists.
func (this MLGraphBuilder) HasHardSigmoid1() bool {
	return js.True == bindings.HasMLGraphBuilderHardSigmoid1(
		this.Ref(),
	)
}

// HardSigmoid1Func returns the method "MLGraphBuilder.hardSigmoid".
func (this MLGraphBuilder) HardSigmoid1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderHardSigmoid1Func(
			this.Ref(),
		),
	)
}

// HardSigmoid1 calls the method "MLGraphBuilder.hardSigmoid".
func (this MLGraphBuilder) HardSigmoid1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderHardSigmoid1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryHardSigmoid1 calls the method "MLGraphBuilder.hardSigmoid"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryHardSigmoid1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderHardSigmoid1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasHardSigmoid2 returns true if the method "MLGraphBuilder.hardSigmoid" exists.
func (this MLGraphBuilder) HasHardSigmoid2() bool {
	return js.True == bindings.HasMLGraphBuilderHardSigmoid2(
		this.Ref(),
	)
}

// HardSigmoid2Func returns the method "MLGraphBuilder.hardSigmoid".
func (this MLGraphBuilder) HardSigmoid2Func() (fn js.Func[func(options MLHardSigmoidOptions) MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderHardSigmoid2Func(
			this.Ref(),
		),
	)
}

// HardSigmoid2 calls the method "MLGraphBuilder.hardSigmoid".
func (this MLGraphBuilder) HardSigmoid2(options MLHardSigmoidOptions) (ret MLActivation) {
	bindings.CallMLGraphBuilderHardSigmoid2(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryHardSigmoid2 calls the method "MLGraphBuilder.hardSigmoid"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryHardSigmoid2(options MLHardSigmoidOptions) (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderHardSigmoid2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasHardSigmoid3 returns true if the method "MLGraphBuilder.hardSigmoid" exists.
func (this MLGraphBuilder) HasHardSigmoid3() bool {
	return js.True == bindings.HasMLGraphBuilderHardSigmoid3(
		this.Ref(),
	)
}

// HardSigmoid3Func returns the method "MLGraphBuilder.hardSigmoid".
func (this MLGraphBuilder) HardSigmoid3Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderHardSigmoid3Func(
			this.Ref(),
		),
	)
}

// HardSigmoid3 calls the method "MLGraphBuilder.hardSigmoid".
func (this MLGraphBuilder) HardSigmoid3() (ret MLActivation) {
	bindings.CallMLGraphBuilderHardSigmoid3(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryHardSigmoid3 calls the method "MLGraphBuilder.hardSigmoid"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryHardSigmoid3() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderHardSigmoid3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGruCell returns true if the method "MLGraphBuilder.gruCell" exists.
func (this MLGraphBuilder) HasGruCell() bool {
	return js.True == bindings.HasMLGraphBuilderGruCell(
		this.Ref(),
	)
}

// GruCellFunc returns the method "MLGraphBuilder.gruCell".
func (this MLGraphBuilder) GruCellFunc() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, hiddenSize uint32, options MLGruCellOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderGruCellFunc(
			this.Ref(),
		),
	)
}

// GruCell calls the method "MLGraphBuilder.gruCell".
func (this MLGraphBuilder) GruCell(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, hiddenSize uint32, options MLGruCellOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderGruCell(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		hiddenState.Ref(),
		uint32(hiddenSize),
		js.Pointer(&options),
	)

	return
}

// HasGruCell1 returns true if the method "MLGraphBuilder.gruCell" exists.
func (this MLGraphBuilder) HasGruCell1() bool {
	return js.True == bindings.HasMLGraphBuilderGruCell1(
		this.Ref(),
	)
}

// GruCell1Func returns the method "MLGraphBuilder.gruCell".
func (this MLGraphBuilder) GruCell1Func() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, hiddenSize uint32) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderGruCell1Func(
			this.Ref(),
		),
	)
}

// GruCell1 calls the method "MLGraphBuilder.gruCell".
func (this MLGraphBuilder) GruCell1(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, hiddenSize uint32) (ret MLOperand) {
	bindings.CallMLGraphBuilderGruCell1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		hiddenState.Ref(),
		uint32(hiddenSize),
	)

	return
}

// HasSlice returns true if the method "MLGraphBuilder.slice" exists.
func (this MLGraphBuilder) HasSlice() bool {
	return js.True == bindings.HasMLGraphBuilderSlice(
		this.Ref(),
	)
}

// SliceFunc returns the method "MLGraphBuilder.slice".
func (this MLGraphBuilder) SliceFunc() (fn js.Func[func(input MLOperand, starts js.Array[uint32], sizes js.Array[uint32]) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSliceFunc(
			this.Ref(),
		),
	)
}

// Slice calls the method "MLGraphBuilder.slice".
func (this MLGraphBuilder) Slice(input MLOperand, starts js.Array[uint32], sizes js.Array[uint32]) (ret MLOperand) {
	bindings.CallMLGraphBuilderSlice(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		starts.Ref(),
		sizes.Ref(),
	)

	return
}

// HasAveragePool2d returns true if the method "MLGraphBuilder.averagePool2d" exists.
func (this MLGraphBuilder) HasAveragePool2d() bool {
	return js.True == bindings.HasMLGraphBuilderAveragePool2d(
		this.Ref(),
	)
}

// AveragePool2dFunc returns the method "MLGraphBuilder.averagePool2d".
func (this MLGraphBuilder) AveragePool2dFunc() (fn js.Func[func(input MLOperand, options MLPool2dOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderAveragePool2dFunc(
			this.Ref(),
		),
	)
}

// AveragePool2d calls the method "MLGraphBuilder.averagePool2d".
func (this MLGraphBuilder) AveragePool2d(input MLOperand, options MLPool2dOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderAveragePool2d(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasAveragePool2d1 returns true if the method "MLGraphBuilder.averagePool2d" exists.
func (this MLGraphBuilder) HasAveragePool2d1() bool {
	return js.True == bindings.HasMLGraphBuilderAveragePool2d1(
		this.Ref(),
	)
}

// AveragePool2d1Func returns the method "MLGraphBuilder.averagePool2d".
func (this MLGraphBuilder) AveragePool2d1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderAveragePool2d1Func(
			this.Ref(),
		),
	)
}

// AveragePool2d1 calls the method "MLGraphBuilder.averagePool2d".
func (this MLGraphBuilder) AveragePool2d1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderAveragePool2d1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryAveragePool2d1 calls the method "MLGraphBuilder.averagePool2d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryAveragePool2d1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderAveragePool2d1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasL2Pool2d returns true if the method "MLGraphBuilder.l2Pool2d" exists.
func (this MLGraphBuilder) HasL2Pool2d() bool {
	return js.True == bindings.HasMLGraphBuilderL2Pool2d(
		this.Ref(),
	)
}

// L2Pool2dFunc returns the method "MLGraphBuilder.l2Pool2d".
func (this MLGraphBuilder) L2Pool2dFunc() (fn js.Func[func(input MLOperand, options MLPool2dOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderL2Pool2dFunc(
			this.Ref(),
		),
	)
}

// L2Pool2d calls the method "MLGraphBuilder.l2Pool2d".
func (this MLGraphBuilder) L2Pool2d(input MLOperand, options MLPool2dOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderL2Pool2d(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasL2Pool2d1 returns true if the method "MLGraphBuilder.l2Pool2d" exists.
func (this MLGraphBuilder) HasL2Pool2d1() bool {
	return js.True == bindings.HasMLGraphBuilderL2Pool2d1(
		this.Ref(),
	)
}

// L2Pool2d1Func returns the method "MLGraphBuilder.l2Pool2d".
func (this MLGraphBuilder) L2Pool2d1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderL2Pool2d1Func(
			this.Ref(),
		),
	)
}

// L2Pool2d1 calls the method "MLGraphBuilder.l2Pool2d".
func (this MLGraphBuilder) L2Pool2d1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderL2Pool2d1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryL2Pool2d1 calls the method "MLGraphBuilder.l2Pool2d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryL2Pool2d1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderL2Pool2d1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasMaxPool2d returns true if the method "MLGraphBuilder.maxPool2d" exists.
func (this MLGraphBuilder) HasMaxPool2d() bool {
	return js.True == bindings.HasMLGraphBuilderMaxPool2d(
		this.Ref(),
	)
}

// MaxPool2dFunc returns the method "MLGraphBuilder.maxPool2d".
func (this MLGraphBuilder) MaxPool2dFunc() (fn js.Func[func(input MLOperand, options MLPool2dOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderMaxPool2dFunc(
			this.Ref(),
		),
	)
}

// MaxPool2d calls the method "MLGraphBuilder.maxPool2d".
func (this MLGraphBuilder) MaxPool2d(input MLOperand, options MLPool2dOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderMaxPool2d(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasMaxPool2d1 returns true if the method "MLGraphBuilder.maxPool2d" exists.
func (this MLGraphBuilder) HasMaxPool2d1() bool {
	return js.True == bindings.HasMLGraphBuilderMaxPool2d1(
		this.Ref(),
	)
}

// MaxPool2d1Func returns the method "MLGraphBuilder.maxPool2d".
func (this MLGraphBuilder) MaxPool2d1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderMaxPool2d1Func(
			this.Ref(),
		),
	)
}

// MaxPool2d1 calls the method "MLGraphBuilder.maxPool2d".
func (this MLGraphBuilder) MaxPool2d1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderMaxPool2d1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryMaxPool2d1 calls the method "MLGraphBuilder.maxPool2d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryMaxPool2d1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderMaxPool2d1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasLinear returns true if the method "MLGraphBuilder.linear" exists.
func (this MLGraphBuilder) HasLinear() bool {
	return js.True == bindings.HasMLGraphBuilderLinear(
		this.Ref(),
	)
}

// LinearFunc returns the method "MLGraphBuilder.linear".
func (this MLGraphBuilder) LinearFunc() (fn js.Func[func(input MLOperand, options MLLinearOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLinearFunc(
			this.Ref(),
		),
	)
}

// Linear calls the method "MLGraphBuilder.linear".
func (this MLGraphBuilder) Linear(input MLOperand, options MLLinearOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderLinear(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasLinear1 returns true if the method "MLGraphBuilder.linear" exists.
func (this MLGraphBuilder) HasLinear1() bool {
	return js.True == bindings.HasMLGraphBuilderLinear1(
		this.Ref(),
	)
}

// Linear1Func returns the method "MLGraphBuilder.linear".
func (this MLGraphBuilder) Linear1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLinear1Func(
			this.Ref(),
		),
	)
}

// Linear1 calls the method "MLGraphBuilder.linear".
func (this MLGraphBuilder) Linear1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderLinear1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryLinear1 calls the method "MLGraphBuilder.linear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLinear1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLinear1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasLinear2 returns true if the method "MLGraphBuilder.linear" exists.
func (this MLGraphBuilder) HasLinear2() bool {
	return js.True == bindings.HasMLGraphBuilderLinear2(
		this.Ref(),
	)
}

// Linear2Func returns the method "MLGraphBuilder.linear".
func (this MLGraphBuilder) Linear2Func() (fn js.Func[func(options MLLinearOptions) MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLinear2Func(
			this.Ref(),
		),
	)
}

// Linear2 calls the method "MLGraphBuilder.linear".
func (this MLGraphBuilder) Linear2(options MLLinearOptions) (ret MLActivation) {
	bindings.CallMLGraphBuilderLinear2(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryLinear2 calls the method "MLGraphBuilder.linear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLinear2(options MLLinearOptions) (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLinear2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasLinear3 returns true if the method "MLGraphBuilder.linear" exists.
func (this MLGraphBuilder) HasLinear3() bool {
	return js.True == bindings.HasMLGraphBuilderLinear3(
		this.Ref(),
	)
}

// Linear3Func returns the method "MLGraphBuilder.linear".
func (this MLGraphBuilder) Linear3Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLinear3Func(
			this.Ref(),
		),
	)
}

// Linear3 calls the method "MLGraphBuilder.linear".
func (this MLGraphBuilder) Linear3() (ret MLActivation) {
	bindings.CallMLGraphBuilderLinear3(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryLinear3 calls the method "MLGraphBuilder.linear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLinear3() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLinear3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasLeakyRelu returns true if the method "MLGraphBuilder.leakyRelu" exists.
func (this MLGraphBuilder) HasLeakyRelu() bool {
	return js.True == bindings.HasMLGraphBuilderLeakyRelu(
		this.Ref(),
	)
}

// LeakyReluFunc returns the method "MLGraphBuilder.leakyRelu".
func (this MLGraphBuilder) LeakyReluFunc() (fn js.Func[func(input MLOperand, options MLLeakyReluOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLeakyReluFunc(
			this.Ref(),
		),
	)
}

// LeakyRelu calls the method "MLGraphBuilder.leakyRelu".
func (this MLGraphBuilder) LeakyRelu(input MLOperand, options MLLeakyReluOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderLeakyRelu(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasLeakyRelu1 returns true if the method "MLGraphBuilder.leakyRelu" exists.
func (this MLGraphBuilder) HasLeakyRelu1() bool {
	return js.True == bindings.HasMLGraphBuilderLeakyRelu1(
		this.Ref(),
	)
}

// LeakyRelu1Func returns the method "MLGraphBuilder.leakyRelu".
func (this MLGraphBuilder) LeakyRelu1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLeakyRelu1Func(
			this.Ref(),
		),
	)
}

// LeakyRelu1 calls the method "MLGraphBuilder.leakyRelu".
func (this MLGraphBuilder) LeakyRelu1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderLeakyRelu1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryLeakyRelu1 calls the method "MLGraphBuilder.leakyRelu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLeakyRelu1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLeakyRelu1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasLeakyRelu2 returns true if the method "MLGraphBuilder.leakyRelu" exists.
func (this MLGraphBuilder) HasLeakyRelu2() bool {
	return js.True == bindings.HasMLGraphBuilderLeakyRelu2(
		this.Ref(),
	)
}

// LeakyRelu2Func returns the method "MLGraphBuilder.leakyRelu".
func (this MLGraphBuilder) LeakyRelu2Func() (fn js.Func[func(options MLLeakyReluOptions) MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLeakyRelu2Func(
			this.Ref(),
		),
	)
}

// LeakyRelu2 calls the method "MLGraphBuilder.leakyRelu".
func (this MLGraphBuilder) LeakyRelu2(options MLLeakyReluOptions) (ret MLActivation) {
	bindings.CallMLGraphBuilderLeakyRelu2(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryLeakyRelu2 calls the method "MLGraphBuilder.leakyRelu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLeakyRelu2(options MLLeakyReluOptions) (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLeakyRelu2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasLeakyRelu3 returns true if the method "MLGraphBuilder.leakyRelu" exists.
func (this MLGraphBuilder) HasLeakyRelu3() bool {
	return js.True == bindings.HasMLGraphBuilderLeakyRelu3(
		this.Ref(),
	)
}

// LeakyRelu3Func returns the method "MLGraphBuilder.leakyRelu".
func (this MLGraphBuilder) LeakyRelu3Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLeakyRelu3Func(
			this.Ref(),
		),
	)
}

// LeakyRelu3 calls the method "MLGraphBuilder.leakyRelu".
func (this MLGraphBuilder) LeakyRelu3() (ret MLActivation) {
	bindings.CallMLGraphBuilderLeakyRelu3(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryLeakyRelu3 calls the method "MLGraphBuilder.leakyRelu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLeakyRelu3() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLeakyRelu3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPad returns true if the method "MLGraphBuilder.pad" exists.
func (this MLGraphBuilder) HasPad() bool {
	return js.True == bindings.HasMLGraphBuilderPad(
		this.Ref(),
	)
}

// PadFunc returns the method "MLGraphBuilder.pad".
func (this MLGraphBuilder) PadFunc() (fn js.Func[func(input MLOperand, beginningPadding js.Array[uint32], endingPadding js.Array[uint32], options MLPadOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderPadFunc(
			this.Ref(),
		),
	)
}

// Pad calls the method "MLGraphBuilder.pad".
func (this MLGraphBuilder) Pad(input MLOperand, beginningPadding js.Array[uint32], endingPadding js.Array[uint32], options MLPadOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderPad(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		beginningPadding.Ref(),
		endingPadding.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasPad1 returns true if the method "MLGraphBuilder.pad" exists.
func (this MLGraphBuilder) HasPad1() bool {
	return js.True == bindings.HasMLGraphBuilderPad1(
		this.Ref(),
	)
}

// Pad1Func returns the method "MLGraphBuilder.pad".
func (this MLGraphBuilder) Pad1Func() (fn js.Func[func(input MLOperand, beginningPadding js.Array[uint32], endingPadding js.Array[uint32]) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderPad1Func(
			this.Ref(),
		),
	)
}

// Pad1 calls the method "MLGraphBuilder.pad".
func (this MLGraphBuilder) Pad1(input MLOperand, beginningPadding js.Array[uint32], endingPadding js.Array[uint32]) (ret MLOperand) {
	bindings.CallMLGraphBuilderPad1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		beginningPadding.Ref(),
		endingPadding.Ref(),
	)

	return
}

// HasInstanceNormalization returns true if the method "MLGraphBuilder.instanceNormalization" exists.
func (this MLGraphBuilder) HasInstanceNormalization() bool {
	return js.True == bindings.HasMLGraphBuilderInstanceNormalization(
		this.Ref(),
	)
}

// InstanceNormalizationFunc returns the method "MLGraphBuilder.instanceNormalization".
func (this MLGraphBuilder) InstanceNormalizationFunc() (fn js.Func[func(input MLOperand, options MLInstanceNormalizationOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderInstanceNormalizationFunc(
			this.Ref(),
		),
	)
}

// InstanceNormalization calls the method "MLGraphBuilder.instanceNormalization".
func (this MLGraphBuilder) InstanceNormalization(input MLOperand, options MLInstanceNormalizationOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderInstanceNormalization(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasInstanceNormalization1 returns true if the method "MLGraphBuilder.instanceNormalization" exists.
func (this MLGraphBuilder) HasInstanceNormalization1() bool {
	return js.True == bindings.HasMLGraphBuilderInstanceNormalization1(
		this.Ref(),
	)
}

// InstanceNormalization1Func returns the method "MLGraphBuilder.instanceNormalization".
func (this MLGraphBuilder) InstanceNormalization1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderInstanceNormalization1Func(
			this.Ref(),
		),
	)
}

// InstanceNormalization1 calls the method "MLGraphBuilder.instanceNormalization".
func (this MLGraphBuilder) InstanceNormalization1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderInstanceNormalization1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryInstanceNormalization1 calls the method "MLGraphBuilder.instanceNormalization"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryInstanceNormalization1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderInstanceNormalization1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasSoftplus returns true if the method "MLGraphBuilder.softplus" exists.
func (this MLGraphBuilder) HasSoftplus() bool {
	return js.True == bindings.HasMLGraphBuilderSoftplus(
		this.Ref(),
	)
}

// SoftplusFunc returns the method "MLGraphBuilder.softplus".
func (this MLGraphBuilder) SoftplusFunc() (fn js.Func[func(input MLOperand, options MLSoftplusOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSoftplusFunc(
			this.Ref(),
		),
	)
}

// Softplus calls the method "MLGraphBuilder.softplus".
func (this MLGraphBuilder) Softplus(input MLOperand, options MLSoftplusOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderSoftplus(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasSoftplus1 returns true if the method "MLGraphBuilder.softplus" exists.
func (this MLGraphBuilder) HasSoftplus1() bool {
	return js.True == bindings.HasMLGraphBuilderSoftplus1(
		this.Ref(),
	)
}

// Softplus1Func returns the method "MLGraphBuilder.softplus".
func (this MLGraphBuilder) Softplus1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSoftplus1Func(
			this.Ref(),
		),
	)
}

// Softplus1 calls the method "MLGraphBuilder.softplus".
func (this MLGraphBuilder) Softplus1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderSoftplus1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySoftplus1 calls the method "MLGraphBuilder.softplus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySoftplus1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSoftplus1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasSoftplus2 returns true if the method "MLGraphBuilder.softplus" exists.
func (this MLGraphBuilder) HasSoftplus2() bool {
	return js.True == bindings.HasMLGraphBuilderSoftplus2(
		this.Ref(),
	)
}

// Softplus2Func returns the method "MLGraphBuilder.softplus".
func (this MLGraphBuilder) Softplus2Func() (fn js.Func[func(options MLSoftplusOptions) MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSoftplus2Func(
			this.Ref(),
		),
	)
}

// Softplus2 calls the method "MLGraphBuilder.softplus".
func (this MLGraphBuilder) Softplus2(options MLSoftplusOptions) (ret MLActivation) {
	bindings.CallMLGraphBuilderSoftplus2(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TrySoftplus2 calls the method "MLGraphBuilder.softplus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySoftplus2(options MLSoftplusOptions) (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSoftplus2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasSoftplus3 returns true if the method "MLGraphBuilder.softplus" exists.
func (this MLGraphBuilder) HasSoftplus3() bool {
	return js.True == bindings.HasMLGraphBuilderSoftplus3(
		this.Ref(),
	)
}

// Softplus3Func returns the method "MLGraphBuilder.softplus".
func (this MLGraphBuilder) Softplus3Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSoftplus3Func(
			this.Ref(),
		),
	)
}

// Softplus3 calls the method "MLGraphBuilder.softplus".
func (this MLGraphBuilder) Softplus3() (ret MLActivation) {
	bindings.CallMLGraphBuilderSoftplus3(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySoftplus3 calls the method "MLGraphBuilder.softplus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySoftplus3() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSoftplus3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSoftsign returns true if the method "MLGraphBuilder.softsign" exists.
func (this MLGraphBuilder) HasSoftsign() bool {
	return js.True == bindings.HasMLGraphBuilderSoftsign(
		this.Ref(),
	)
}

// SoftsignFunc returns the method "MLGraphBuilder.softsign".
func (this MLGraphBuilder) SoftsignFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSoftsignFunc(
			this.Ref(),
		),
	)
}

// Softsign calls the method "MLGraphBuilder.softsign".
func (this MLGraphBuilder) Softsign(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderSoftsign(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySoftsign calls the method "MLGraphBuilder.softsign"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySoftsign(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSoftsign(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasSoftsign1 returns true if the method "MLGraphBuilder.softsign" exists.
func (this MLGraphBuilder) HasSoftsign1() bool {
	return js.True == bindings.HasMLGraphBuilderSoftsign1(
		this.Ref(),
	)
}

// Softsign1Func returns the method "MLGraphBuilder.softsign".
func (this MLGraphBuilder) Softsign1Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSoftsign1Func(
			this.Ref(),
		),
	)
}

// Softsign1 calls the method "MLGraphBuilder.softsign".
func (this MLGraphBuilder) Softsign1() (ret MLActivation) {
	bindings.CallMLGraphBuilderSoftsign1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySoftsign1 calls the method "MLGraphBuilder.softsign"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySoftsign1() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSoftsign1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSigmoid returns true if the method "MLGraphBuilder.sigmoid" exists.
func (this MLGraphBuilder) HasSigmoid() bool {
	return js.True == bindings.HasMLGraphBuilderSigmoid(
		this.Ref(),
	)
}

// SigmoidFunc returns the method "MLGraphBuilder.sigmoid".
func (this MLGraphBuilder) SigmoidFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSigmoidFunc(
			this.Ref(),
		),
	)
}

// Sigmoid calls the method "MLGraphBuilder.sigmoid".
func (this MLGraphBuilder) Sigmoid(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderSigmoid(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySigmoid calls the method "MLGraphBuilder.sigmoid"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySigmoid(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSigmoid(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasSigmoid1 returns true if the method "MLGraphBuilder.sigmoid" exists.
func (this MLGraphBuilder) HasSigmoid1() bool {
	return js.True == bindings.HasMLGraphBuilderSigmoid1(
		this.Ref(),
	)
}

// Sigmoid1Func returns the method "MLGraphBuilder.sigmoid".
func (this MLGraphBuilder) Sigmoid1Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSigmoid1Func(
			this.Ref(),
		),
	)
}

// Sigmoid1 calls the method "MLGraphBuilder.sigmoid".
func (this MLGraphBuilder) Sigmoid1() (ret MLActivation) {
	bindings.CallMLGraphBuilderSigmoid1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySigmoid1 calls the method "MLGraphBuilder.sigmoid"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySigmoid1() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSigmoid1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReshape returns true if the method "MLGraphBuilder.reshape" exists.
func (this MLGraphBuilder) HasReshape() bool {
	return js.True == bindings.HasMLGraphBuilderReshape(
		this.Ref(),
	)
}

// ReshapeFunc returns the method "MLGraphBuilder.reshape".
func (this MLGraphBuilder) ReshapeFunc() (fn js.Func[func(input MLOperand, newShape js.Array[uint32]) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReshapeFunc(
			this.Ref(),
		),
	)
}

// Reshape calls the method "MLGraphBuilder.reshape".
func (this MLGraphBuilder) Reshape(input MLOperand, newShape js.Array[uint32]) (ret MLOperand) {
	bindings.CallMLGraphBuilderReshape(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		newShape.Ref(),
	)

	return
}

// HasConv2d returns true if the method "MLGraphBuilder.conv2d" exists.
func (this MLGraphBuilder) HasConv2d() bool {
	return js.True == bindings.HasMLGraphBuilderConv2d(
		this.Ref(),
	)
}

// Conv2dFunc returns the method "MLGraphBuilder.conv2d".
func (this MLGraphBuilder) Conv2dFunc() (fn js.Func[func(input MLOperand, filter MLOperand, options MLConv2dOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderConv2dFunc(
			this.Ref(),
		),
	)
}

// Conv2d calls the method "MLGraphBuilder.conv2d".
func (this MLGraphBuilder) Conv2d(input MLOperand, filter MLOperand, options MLConv2dOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderConv2d(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		filter.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasConv2d1 returns true if the method "MLGraphBuilder.conv2d" exists.
func (this MLGraphBuilder) HasConv2d1() bool {
	return js.True == bindings.HasMLGraphBuilderConv2d1(
		this.Ref(),
	)
}

// Conv2d1Func returns the method "MLGraphBuilder.conv2d".
func (this MLGraphBuilder) Conv2d1Func() (fn js.Func[func(input MLOperand, filter MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderConv2d1Func(
			this.Ref(),
		),
	)
}

// Conv2d1 calls the method "MLGraphBuilder.conv2d".
func (this MLGraphBuilder) Conv2d1(input MLOperand, filter MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderConv2d1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		filter.Ref(),
	)

	return
}

// HasSplit returns true if the method "MLGraphBuilder.split" exists.
func (this MLGraphBuilder) HasSplit() bool {
	return js.True == bindings.HasMLGraphBuilderSplit(
		this.Ref(),
	)
}

// SplitFunc returns the method "MLGraphBuilder.split".
func (this MLGraphBuilder) SplitFunc() (fn js.Func[func(input MLOperand, splits OneOf_Uint32_ArrayUint32, options MLSplitOptions) js.Array[MLOperand]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSplitFunc(
			this.Ref(),
		),
	)
}

// Split calls the method "MLGraphBuilder.split".
func (this MLGraphBuilder) Split(input MLOperand, splits OneOf_Uint32_ArrayUint32, options MLSplitOptions) (ret js.Array[MLOperand]) {
	bindings.CallMLGraphBuilderSplit(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		splits.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasSplit1 returns true if the method "MLGraphBuilder.split" exists.
func (this MLGraphBuilder) HasSplit1() bool {
	return js.True == bindings.HasMLGraphBuilderSplit1(
		this.Ref(),
	)
}

// Split1Func returns the method "MLGraphBuilder.split".
func (this MLGraphBuilder) Split1Func() (fn js.Func[func(input MLOperand, splits OneOf_Uint32_ArrayUint32) js.Array[MLOperand]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSplit1Func(
			this.Ref(),
		),
	)
}

// Split1 calls the method "MLGraphBuilder.split".
func (this MLGraphBuilder) Split1(input MLOperand, splits OneOf_Uint32_ArrayUint32) (ret js.Array[MLOperand]) {
	bindings.CallMLGraphBuilderSplit1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		splits.Ref(),
	)

	return
}

// HasResample2d returns true if the method "MLGraphBuilder.resample2d" exists.
func (this MLGraphBuilder) HasResample2d() bool {
	return js.True == bindings.HasMLGraphBuilderResample2d(
		this.Ref(),
	)
}

// Resample2dFunc returns the method "MLGraphBuilder.resample2d".
func (this MLGraphBuilder) Resample2dFunc() (fn js.Func[func(input MLOperand, options MLResample2dOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderResample2dFunc(
			this.Ref(),
		),
	)
}

// Resample2d calls the method "MLGraphBuilder.resample2d".
func (this MLGraphBuilder) Resample2d(input MLOperand, options MLResample2dOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderResample2d(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasResample2d1 returns true if the method "MLGraphBuilder.resample2d" exists.
func (this MLGraphBuilder) HasResample2d1() bool {
	return js.True == bindings.HasMLGraphBuilderResample2d1(
		this.Ref(),
	)
}

// Resample2d1Func returns the method "MLGraphBuilder.resample2d".
func (this MLGraphBuilder) Resample2d1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderResample2d1Func(
			this.Ref(),
		),
	)
}

// Resample2d1 calls the method "MLGraphBuilder.resample2d".
func (this MLGraphBuilder) Resample2d1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderResample2d1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryResample2d1 calls the method "MLGraphBuilder.resample2d"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryResample2d1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderResample2d1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasHardSwish returns true if the method "MLGraphBuilder.hardSwish" exists.
func (this MLGraphBuilder) HasHardSwish() bool {
	return js.True == bindings.HasMLGraphBuilderHardSwish(
		this.Ref(),
	)
}

// HardSwishFunc returns the method "MLGraphBuilder.hardSwish".
func (this MLGraphBuilder) HardSwishFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderHardSwishFunc(
			this.Ref(),
		),
	)
}

// HardSwish calls the method "MLGraphBuilder.hardSwish".
func (this MLGraphBuilder) HardSwish(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderHardSwish(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryHardSwish calls the method "MLGraphBuilder.hardSwish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryHardSwish(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderHardSwish(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasHardSwish1 returns true if the method "MLGraphBuilder.hardSwish" exists.
func (this MLGraphBuilder) HasHardSwish1() bool {
	return js.True == bindings.HasMLGraphBuilderHardSwish1(
		this.Ref(),
	)
}

// HardSwish1Func returns the method "MLGraphBuilder.hardSwish".
func (this MLGraphBuilder) HardSwish1Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderHardSwish1Func(
			this.Ref(),
		),
	)
}

// HardSwish1 calls the method "MLGraphBuilder.hardSwish".
func (this MLGraphBuilder) HardSwish1() (ret MLActivation) {
	bindings.CallMLGraphBuilderHardSwish1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryHardSwish1 calls the method "MLGraphBuilder.hardSwish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryHardSwish1() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderHardSwish1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSoftmax returns true if the method "MLGraphBuilder.softmax" exists.
func (this MLGraphBuilder) HasSoftmax() bool {
	return js.True == bindings.HasMLGraphBuilderSoftmax(
		this.Ref(),
	)
}

// SoftmaxFunc returns the method "MLGraphBuilder.softmax".
func (this MLGraphBuilder) SoftmaxFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSoftmaxFunc(
			this.Ref(),
		),
	)
}

// Softmax calls the method "MLGraphBuilder.softmax".
func (this MLGraphBuilder) Softmax(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderSoftmax(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySoftmax calls the method "MLGraphBuilder.softmax"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySoftmax(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSoftmax(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasSoftmax1 returns true if the method "MLGraphBuilder.softmax" exists.
func (this MLGraphBuilder) HasSoftmax1() bool {
	return js.True == bindings.HasMLGraphBuilderSoftmax1(
		this.Ref(),
	)
}

// Softmax1Func returns the method "MLGraphBuilder.softmax".
func (this MLGraphBuilder) Softmax1Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSoftmax1Func(
			this.Ref(),
		),
	)
}

// Softmax1 calls the method "MLGraphBuilder.softmax".
func (this MLGraphBuilder) Softmax1() (ret MLActivation) {
	bindings.CallMLGraphBuilderSoftmax1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySoftmax1 calls the method "MLGraphBuilder.softmax"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySoftmax1() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSoftmax1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasConvTranspose2d returns true if the method "MLGraphBuilder.convTranspose2d" exists.
func (this MLGraphBuilder) HasConvTranspose2d() bool {
	return js.True == bindings.HasMLGraphBuilderConvTranspose2d(
		this.Ref(),
	)
}

// ConvTranspose2dFunc returns the method "MLGraphBuilder.convTranspose2d".
func (this MLGraphBuilder) ConvTranspose2dFunc() (fn js.Func[func(input MLOperand, filter MLOperand, options MLConvTranspose2dOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderConvTranspose2dFunc(
			this.Ref(),
		),
	)
}

// ConvTranspose2d calls the method "MLGraphBuilder.convTranspose2d".
func (this MLGraphBuilder) ConvTranspose2d(input MLOperand, filter MLOperand, options MLConvTranspose2dOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderConvTranspose2d(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		filter.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasConvTranspose2d1 returns true if the method "MLGraphBuilder.convTranspose2d" exists.
func (this MLGraphBuilder) HasConvTranspose2d1() bool {
	return js.True == bindings.HasMLGraphBuilderConvTranspose2d1(
		this.Ref(),
	)
}

// ConvTranspose2d1Func returns the method "MLGraphBuilder.convTranspose2d".
func (this MLGraphBuilder) ConvTranspose2d1Func() (fn js.Func[func(input MLOperand, filter MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderConvTranspose2d1Func(
			this.Ref(),
		),
	)
}

// ConvTranspose2d1 calls the method "MLGraphBuilder.convTranspose2d".
func (this MLGraphBuilder) ConvTranspose2d1(input MLOperand, filter MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderConvTranspose2d1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		filter.Ref(),
	)

	return
}

// HasRelu returns true if the method "MLGraphBuilder.relu" exists.
func (this MLGraphBuilder) HasRelu() bool {
	return js.True == bindings.HasMLGraphBuilderRelu(
		this.Ref(),
	)
}

// ReluFunc returns the method "MLGraphBuilder.relu".
func (this MLGraphBuilder) ReluFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReluFunc(
			this.Ref(),
		),
	)
}

// Relu calls the method "MLGraphBuilder.relu".
func (this MLGraphBuilder) Relu(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderRelu(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryRelu calls the method "MLGraphBuilder.relu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryRelu(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderRelu(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasRelu1 returns true if the method "MLGraphBuilder.relu" exists.
func (this MLGraphBuilder) HasRelu1() bool {
	return js.True == bindings.HasMLGraphBuilderRelu1(
		this.Ref(),
	)
}

// Relu1Func returns the method "MLGraphBuilder.relu".
func (this MLGraphBuilder) Relu1Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderRelu1Func(
			this.Ref(),
		),
	)
}

// Relu1 calls the method "MLGraphBuilder.relu".
func (this MLGraphBuilder) Relu1() (ret MLActivation) {
	bindings.CallMLGraphBuilderRelu1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRelu1 calls the method "MLGraphBuilder.relu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryRelu1() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderRelu1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasAdd returns true if the method "MLGraphBuilder.add" exists.
func (this MLGraphBuilder) HasAdd() bool {
	return js.True == bindings.HasMLGraphBuilderAdd(
		this.Ref(),
	)
}

// AddFunc returns the method "MLGraphBuilder.add".
func (this MLGraphBuilder) AddFunc() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderAddFunc(
			this.Ref(),
		),
	)
}

// Add calls the method "MLGraphBuilder.add".
func (this MLGraphBuilder) Add(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderAdd(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasSub returns true if the method "MLGraphBuilder.sub" exists.
func (this MLGraphBuilder) HasSub() bool {
	return js.True == bindings.HasMLGraphBuilderSub(
		this.Ref(),
	)
}

// SubFunc returns the method "MLGraphBuilder.sub".
func (this MLGraphBuilder) SubFunc() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSubFunc(
			this.Ref(),
		),
	)
}

// Sub calls the method "MLGraphBuilder.sub".
func (this MLGraphBuilder) Sub(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderSub(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasMul returns true if the method "MLGraphBuilder.mul" exists.
func (this MLGraphBuilder) HasMul() bool {
	return js.True == bindings.HasMLGraphBuilderMul(
		this.Ref(),
	)
}

// MulFunc returns the method "MLGraphBuilder.mul".
func (this MLGraphBuilder) MulFunc() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderMulFunc(
			this.Ref(),
		),
	)
}

// Mul calls the method "MLGraphBuilder.mul".
func (this MLGraphBuilder) Mul(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderMul(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasDiv returns true if the method "MLGraphBuilder.div" exists.
func (this MLGraphBuilder) HasDiv() bool {
	return js.True == bindings.HasMLGraphBuilderDiv(
		this.Ref(),
	)
}

// DivFunc returns the method "MLGraphBuilder.div".
func (this MLGraphBuilder) DivFunc() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderDivFunc(
			this.Ref(),
		),
	)
}

// Div calls the method "MLGraphBuilder.div".
func (this MLGraphBuilder) Div(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderDiv(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasMax returns true if the method "MLGraphBuilder.max" exists.
func (this MLGraphBuilder) HasMax() bool {
	return js.True == bindings.HasMLGraphBuilderMax(
		this.Ref(),
	)
}

// MaxFunc returns the method "MLGraphBuilder.max".
func (this MLGraphBuilder) MaxFunc() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderMaxFunc(
			this.Ref(),
		),
	)
}

// Max calls the method "MLGraphBuilder.max".
func (this MLGraphBuilder) Max(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderMax(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasMin returns true if the method "MLGraphBuilder.min" exists.
func (this MLGraphBuilder) HasMin() bool {
	return js.True == bindings.HasMLGraphBuilderMin(
		this.Ref(),
	)
}

// MinFunc returns the method "MLGraphBuilder.min".
func (this MLGraphBuilder) MinFunc() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderMinFunc(
			this.Ref(),
		),
	)
}

// Min calls the method "MLGraphBuilder.min".
func (this MLGraphBuilder) Min(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderMin(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasPow returns true if the method "MLGraphBuilder.pow" exists.
func (this MLGraphBuilder) HasPow() bool {
	return js.True == bindings.HasMLGraphBuilderPow(
		this.Ref(),
	)
}

// PowFunc returns the method "MLGraphBuilder.pow".
func (this MLGraphBuilder) PowFunc() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderPowFunc(
			this.Ref(),
		),
	)
}

// Pow calls the method "MLGraphBuilder.pow".
func (this MLGraphBuilder) Pow(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderPow(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasReduceL1 returns true if the method "MLGraphBuilder.reduceL1" exists.
func (this MLGraphBuilder) HasReduceL1() bool {
	return js.True == bindings.HasMLGraphBuilderReduceL1(
		this.Ref(),
	)
}

// ReduceL1Func returns the method "MLGraphBuilder.reduceL1".
func (this MLGraphBuilder) ReduceL1Func() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceL1Func(
			this.Ref(),
		),
	)
}

// ReduceL1 calls the method "MLGraphBuilder.reduceL1".
func (this MLGraphBuilder) ReduceL1(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceL1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasReduceL11 returns true if the method "MLGraphBuilder.reduceL1" exists.
func (this MLGraphBuilder) HasReduceL11() bool {
	return js.True == bindings.HasMLGraphBuilderReduceL11(
		this.Ref(),
	)
}

// ReduceL11Func returns the method "MLGraphBuilder.reduceL1".
func (this MLGraphBuilder) ReduceL11Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceL11Func(
			this.Ref(),
		),
	)
}

// ReduceL11 calls the method "MLGraphBuilder.reduceL1".
func (this MLGraphBuilder) ReduceL11(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceL11(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceL11 calls the method "MLGraphBuilder.reduceL1"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceL11(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceL11(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasReduceL2 returns true if the method "MLGraphBuilder.reduceL2" exists.
func (this MLGraphBuilder) HasReduceL2() bool {
	return js.True == bindings.HasMLGraphBuilderReduceL2(
		this.Ref(),
	)
}

// ReduceL2Func returns the method "MLGraphBuilder.reduceL2".
func (this MLGraphBuilder) ReduceL2Func() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceL2Func(
			this.Ref(),
		),
	)
}

// ReduceL2 calls the method "MLGraphBuilder.reduceL2".
func (this MLGraphBuilder) ReduceL2(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceL2(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasReduceL21 returns true if the method "MLGraphBuilder.reduceL2" exists.
func (this MLGraphBuilder) HasReduceL21() bool {
	return js.True == bindings.HasMLGraphBuilderReduceL21(
		this.Ref(),
	)
}

// ReduceL21Func returns the method "MLGraphBuilder.reduceL2".
func (this MLGraphBuilder) ReduceL21Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceL21Func(
			this.Ref(),
		),
	)
}

// ReduceL21 calls the method "MLGraphBuilder.reduceL2".
func (this MLGraphBuilder) ReduceL21(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceL21(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceL21 calls the method "MLGraphBuilder.reduceL2"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceL21(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceL21(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasReduceLogSum returns true if the method "MLGraphBuilder.reduceLogSum" exists.
func (this MLGraphBuilder) HasReduceLogSum() bool {
	return js.True == bindings.HasMLGraphBuilderReduceLogSum(
		this.Ref(),
	)
}

// ReduceLogSumFunc returns the method "MLGraphBuilder.reduceLogSum".
func (this MLGraphBuilder) ReduceLogSumFunc() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceLogSumFunc(
			this.Ref(),
		),
	)
}

// ReduceLogSum calls the method "MLGraphBuilder.reduceLogSum".
func (this MLGraphBuilder) ReduceLogSum(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceLogSum(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasReduceLogSum1 returns true if the method "MLGraphBuilder.reduceLogSum" exists.
func (this MLGraphBuilder) HasReduceLogSum1() bool {
	return js.True == bindings.HasMLGraphBuilderReduceLogSum1(
		this.Ref(),
	)
}

// ReduceLogSum1Func returns the method "MLGraphBuilder.reduceLogSum".
func (this MLGraphBuilder) ReduceLogSum1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceLogSum1Func(
			this.Ref(),
		),
	)
}

// ReduceLogSum1 calls the method "MLGraphBuilder.reduceLogSum".
func (this MLGraphBuilder) ReduceLogSum1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceLogSum1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceLogSum1 calls the method "MLGraphBuilder.reduceLogSum"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceLogSum1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceLogSum1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasReduceLogSumExp returns true if the method "MLGraphBuilder.reduceLogSumExp" exists.
func (this MLGraphBuilder) HasReduceLogSumExp() bool {
	return js.True == bindings.HasMLGraphBuilderReduceLogSumExp(
		this.Ref(),
	)
}

// ReduceLogSumExpFunc returns the method "MLGraphBuilder.reduceLogSumExp".
func (this MLGraphBuilder) ReduceLogSumExpFunc() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceLogSumExpFunc(
			this.Ref(),
		),
	)
}

// ReduceLogSumExp calls the method "MLGraphBuilder.reduceLogSumExp".
func (this MLGraphBuilder) ReduceLogSumExp(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceLogSumExp(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasReduceLogSumExp1 returns true if the method "MLGraphBuilder.reduceLogSumExp" exists.
func (this MLGraphBuilder) HasReduceLogSumExp1() bool {
	return js.True == bindings.HasMLGraphBuilderReduceLogSumExp1(
		this.Ref(),
	)
}

// ReduceLogSumExp1Func returns the method "MLGraphBuilder.reduceLogSumExp".
func (this MLGraphBuilder) ReduceLogSumExp1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceLogSumExp1Func(
			this.Ref(),
		),
	)
}

// ReduceLogSumExp1 calls the method "MLGraphBuilder.reduceLogSumExp".
func (this MLGraphBuilder) ReduceLogSumExp1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceLogSumExp1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceLogSumExp1 calls the method "MLGraphBuilder.reduceLogSumExp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceLogSumExp1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceLogSumExp1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasReduceMax returns true if the method "MLGraphBuilder.reduceMax" exists.
func (this MLGraphBuilder) HasReduceMax() bool {
	return js.True == bindings.HasMLGraphBuilderReduceMax(
		this.Ref(),
	)
}

// ReduceMaxFunc returns the method "MLGraphBuilder.reduceMax".
func (this MLGraphBuilder) ReduceMaxFunc() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceMaxFunc(
			this.Ref(),
		),
	)
}

// ReduceMax calls the method "MLGraphBuilder.reduceMax".
func (this MLGraphBuilder) ReduceMax(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceMax(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasReduceMax1 returns true if the method "MLGraphBuilder.reduceMax" exists.
func (this MLGraphBuilder) HasReduceMax1() bool {
	return js.True == bindings.HasMLGraphBuilderReduceMax1(
		this.Ref(),
	)
}

// ReduceMax1Func returns the method "MLGraphBuilder.reduceMax".
func (this MLGraphBuilder) ReduceMax1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceMax1Func(
			this.Ref(),
		),
	)
}

// ReduceMax1 calls the method "MLGraphBuilder.reduceMax".
func (this MLGraphBuilder) ReduceMax1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceMax1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceMax1 calls the method "MLGraphBuilder.reduceMax"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceMax1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceMax1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasReduceMean returns true if the method "MLGraphBuilder.reduceMean" exists.
func (this MLGraphBuilder) HasReduceMean() bool {
	return js.True == bindings.HasMLGraphBuilderReduceMean(
		this.Ref(),
	)
}

// ReduceMeanFunc returns the method "MLGraphBuilder.reduceMean".
func (this MLGraphBuilder) ReduceMeanFunc() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceMeanFunc(
			this.Ref(),
		),
	)
}

// ReduceMean calls the method "MLGraphBuilder.reduceMean".
func (this MLGraphBuilder) ReduceMean(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceMean(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasReduceMean1 returns true if the method "MLGraphBuilder.reduceMean" exists.
func (this MLGraphBuilder) HasReduceMean1() bool {
	return js.True == bindings.HasMLGraphBuilderReduceMean1(
		this.Ref(),
	)
}

// ReduceMean1Func returns the method "MLGraphBuilder.reduceMean".
func (this MLGraphBuilder) ReduceMean1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceMean1Func(
			this.Ref(),
		),
	)
}

// ReduceMean1 calls the method "MLGraphBuilder.reduceMean".
func (this MLGraphBuilder) ReduceMean1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceMean1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceMean1 calls the method "MLGraphBuilder.reduceMean"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceMean1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceMean1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasReduceMin returns true if the method "MLGraphBuilder.reduceMin" exists.
func (this MLGraphBuilder) HasReduceMin() bool {
	return js.True == bindings.HasMLGraphBuilderReduceMin(
		this.Ref(),
	)
}

// ReduceMinFunc returns the method "MLGraphBuilder.reduceMin".
func (this MLGraphBuilder) ReduceMinFunc() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceMinFunc(
			this.Ref(),
		),
	)
}

// ReduceMin calls the method "MLGraphBuilder.reduceMin".
func (this MLGraphBuilder) ReduceMin(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceMin(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasReduceMin1 returns true if the method "MLGraphBuilder.reduceMin" exists.
func (this MLGraphBuilder) HasReduceMin1() bool {
	return js.True == bindings.HasMLGraphBuilderReduceMin1(
		this.Ref(),
	)
}

// ReduceMin1Func returns the method "MLGraphBuilder.reduceMin".
func (this MLGraphBuilder) ReduceMin1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceMin1Func(
			this.Ref(),
		),
	)
}

// ReduceMin1 calls the method "MLGraphBuilder.reduceMin".
func (this MLGraphBuilder) ReduceMin1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceMin1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceMin1 calls the method "MLGraphBuilder.reduceMin"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceMin1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceMin1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasReduceProduct returns true if the method "MLGraphBuilder.reduceProduct" exists.
func (this MLGraphBuilder) HasReduceProduct() bool {
	return js.True == bindings.HasMLGraphBuilderReduceProduct(
		this.Ref(),
	)
}

// ReduceProductFunc returns the method "MLGraphBuilder.reduceProduct".
func (this MLGraphBuilder) ReduceProductFunc() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceProductFunc(
			this.Ref(),
		),
	)
}

// ReduceProduct calls the method "MLGraphBuilder.reduceProduct".
func (this MLGraphBuilder) ReduceProduct(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceProduct(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasReduceProduct1 returns true if the method "MLGraphBuilder.reduceProduct" exists.
func (this MLGraphBuilder) HasReduceProduct1() bool {
	return js.True == bindings.HasMLGraphBuilderReduceProduct1(
		this.Ref(),
	)
}

// ReduceProduct1Func returns the method "MLGraphBuilder.reduceProduct".
func (this MLGraphBuilder) ReduceProduct1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceProduct1Func(
			this.Ref(),
		),
	)
}

// ReduceProduct1 calls the method "MLGraphBuilder.reduceProduct".
func (this MLGraphBuilder) ReduceProduct1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceProduct1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceProduct1 calls the method "MLGraphBuilder.reduceProduct"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceProduct1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceProduct1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasReduceSum returns true if the method "MLGraphBuilder.reduceSum" exists.
func (this MLGraphBuilder) HasReduceSum() bool {
	return js.True == bindings.HasMLGraphBuilderReduceSum(
		this.Ref(),
	)
}

// ReduceSumFunc returns the method "MLGraphBuilder.reduceSum".
func (this MLGraphBuilder) ReduceSumFunc() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceSumFunc(
			this.Ref(),
		),
	)
}

// ReduceSum calls the method "MLGraphBuilder.reduceSum".
func (this MLGraphBuilder) ReduceSum(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceSum(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasReduceSum1 returns true if the method "MLGraphBuilder.reduceSum" exists.
func (this MLGraphBuilder) HasReduceSum1() bool {
	return js.True == bindings.HasMLGraphBuilderReduceSum1(
		this.Ref(),
	)
}

// ReduceSum1Func returns the method "MLGraphBuilder.reduceSum".
func (this MLGraphBuilder) ReduceSum1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceSum1Func(
			this.Ref(),
		),
	)
}

// ReduceSum1 calls the method "MLGraphBuilder.reduceSum".
func (this MLGraphBuilder) ReduceSum1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceSum1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceSum1 calls the method "MLGraphBuilder.reduceSum"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceSum1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceSum1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasReduceSumSquare returns true if the method "MLGraphBuilder.reduceSumSquare" exists.
func (this MLGraphBuilder) HasReduceSumSquare() bool {
	return js.True == bindings.HasMLGraphBuilderReduceSumSquare(
		this.Ref(),
	)
}

// ReduceSumSquareFunc returns the method "MLGraphBuilder.reduceSumSquare".
func (this MLGraphBuilder) ReduceSumSquareFunc() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceSumSquareFunc(
			this.Ref(),
		),
	)
}

// ReduceSumSquare calls the method "MLGraphBuilder.reduceSumSquare".
func (this MLGraphBuilder) ReduceSumSquare(input MLOperand, options MLReduceOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceSumSquare(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasReduceSumSquare1 returns true if the method "MLGraphBuilder.reduceSumSquare" exists.
func (this MLGraphBuilder) HasReduceSumSquare1() bool {
	return js.True == bindings.HasMLGraphBuilderReduceSumSquare1(
		this.Ref(),
	)
}

// ReduceSumSquare1Func returns the method "MLGraphBuilder.reduceSumSquare".
func (this MLGraphBuilder) ReduceSumSquare1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceSumSquare1Func(
			this.Ref(),
		),
	)
}

// ReduceSumSquare1 calls the method "MLGraphBuilder.reduceSumSquare".
func (this MLGraphBuilder) ReduceSumSquare1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderReduceSumSquare1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryReduceSumSquare1 calls the method "MLGraphBuilder.reduceSumSquare"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryReduceSumSquare1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderReduceSumSquare1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasLstm returns true if the method "MLGraphBuilder.lstm" exists.
func (this MLGraphBuilder) HasLstm() bool {
	return js.True == bindings.HasMLGraphBuilderLstm(
		this.Ref(),
	)
}

// LstmFunc returns the method "MLGraphBuilder.lstm".
func (this MLGraphBuilder) LstmFunc() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32, options MLLstmOptions) js.Array[MLOperand]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLstmFunc(
			this.Ref(),
		),
	)
}

// Lstm calls the method "MLGraphBuilder.lstm".
func (this MLGraphBuilder) Lstm(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32, options MLLstmOptions) (ret js.Array[MLOperand]) {
	bindings.CallMLGraphBuilderLstm(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		uint32(steps),
		uint32(hiddenSize),
		js.Pointer(&options),
	)

	return
}

// HasLstm1 returns true if the method "MLGraphBuilder.lstm" exists.
func (this MLGraphBuilder) HasLstm1() bool {
	return js.True == bindings.HasMLGraphBuilderLstm1(
		this.Ref(),
	)
}

// Lstm1Func returns the method "MLGraphBuilder.lstm".
func (this MLGraphBuilder) Lstm1Func() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32) js.Array[MLOperand]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLstm1Func(
			this.Ref(),
		),
	)
}

// Lstm1 calls the method "MLGraphBuilder.lstm".
func (this MLGraphBuilder) Lstm1(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32) (ret js.Array[MLOperand]) {
	bindings.CallMLGraphBuilderLstm1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		uint32(steps),
		uint32(hiddenSize),
	)

	return
}

// HasMatmul returns true if the method "MLGraphBuilder.matmul" exists.
func (this MLGraphBuilder) HasMatmul() bool {
	return js.True == bindings.HasMLGraphBuilderMatmul(
		this.Ref(),
	)
}

// MatmulFunc returns the method "MLGraphBuilder.matmul".
func (this MLGraphBuilder) MatmulFunc() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderMatmulFunc(
			this.Ref(),
		),
	)
}

// Matmul calls the method "MLGraphBuilder.matmul".
func (this MLGraphBuilder) Matmul(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderMatmul(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasSqueeze returns true if the method "MLGraphBuilder.squeeze" exists.
func (this MLGraphBuilder) HasSqueeze() bool {
	return js.True == bindings.HasMLGraphBuilderSqueeze(
		this.Ref(),
	)
}

// SqueezeFunc returns the method "MLGraphBuilder.squeeze".
func (this MLGraphBuilder) SqueezeFunc() (fn js.Func[func(input MLOperand, options MLSqueezeOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSqueezeFunc(
			this.Ref(),
		),
	)
}

// Squeeze calls the method "MLGraphBuilder.squeeze".
func (this MLGraphBuilder) Squeeze(input MLOperand, options MLSqueezeOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderSqueeze(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasSqueeze1 returns true if the method "MLGraphBuilder.squeeze" exists.
func (this MLGraphBuilder) HasSqueeze1() bool {
	return js.True == bindings.HasMLGraphBuilderSqueeze1(
		this.Ref(),
	)
}

// Squeeze1Func returns the method "MLGraphBuilder.squeeze".
func (this MLGraphBuilder) Squeeze1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSqueeze1Func(
			this.Ref(),
		),
	)
}

// Squeeze1 calls the method "MLGraphBuilder.squeeze".
func (this MLGraphBuilder) Squeeze1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderSqueeze1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySqueeze1 calls the method "MLGraphBuilder.squeeze"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySqueeze1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSqueeze1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasTanh returns true if the method "MLGraphBuilder.tanh" exists.
func (this MLGraphBuilder) HasTanh() bool {
	return js.True == bindings.HasMLGraphBuilderTanh(
		this.Ref(),
	)
}

// TanhFunc returns the method "MLGraphBuilder.tanh".
func (this MLGraphBuilder) TanhFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderTanhFunc(
			this.Ref(),
		),
	)
}

// Tanh calls the method "MLGraphBuilder.tanh".
func (this MLGraphBuilder) Tanh(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderTanh(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryTanh calls the method "MLGraphBuilder.tanh"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryTanh(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderTanh(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasTanh1 returns true if the method "MLGraphBuilder.tanh" exists.
func (this MLGraphBuilder) HasTanh1() bool {
	return js.True == bindings.HasMLGraphBuilderTanh1(
		this.Ref(),
	)
}

// Tanh1Func returns the method "MLGraphBuilder.tanh".
func (this MLGraphBuilder) Tanh1Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderTanh1Func(
			this.Ref(),
		),
	)
}

// Tanh1 calls the method "MLGraphBuilder.tanh".
func (this MLGraphBuilder) Tanh1() (ret MLActivation) {
	bindings.CallMLGraphBuilderTanh1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTanh1 calls the method "MLGraphBuilder.tanh"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryTanh1() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderTanh1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGru returns true if the method "MLGraphBuilder.gru" exists.
func (this MLGraphBuilder) HasGru() bool {
	return js.True == bindings.HasMLGraphBuilderGru(
		this.Ref(),
	)
}

// GruFunc returns the method "MLGraphBuilder.gru".
func (this MLGraphBuilder) GruFunc() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32, options MLGruOptions) js.Array[MLOperand]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderGruFunc(
			this.Ref(),
		),
	)
}

// Gru calls the method "MLGraphBuilder.gru".
func (this MLGraphBuilder) Gru(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32, options MLGruOptions) (ret js.Array[MLOperand]) {
	bindings.CallMLGraphBuilderGru(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		uint32(steps),
		uint32(hiddenSize),
		js.Pointer(&options),
	)

	return
}

// HasGru1 returns true if the method "MLGraphBuilder.gru" exists.
func (this MLGraphBuilder) HasGru1() bool {
	return js.True == bindings.HasMLGraphBuilderGru1(
		this.Ref(),
	)
}

// Gru1Func returns the method "MLGraphBuilder.gru".
func (this MLGraphBuilder) Gru1Func() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32) js.Array[MLOperand]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderGru1Func(
			this.Ref(),
		),
	)
}

// Gru1 calls the method "MLGraphBuilder.gru".
func (this MLGraphBuilder) Gru1(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32) (ret js.Array[MLOperand]) {
	bindings.CallMLGraphBuilderGru1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		uint32(steps),
		uint32(hiddenSize),
	)

	return
}

// HasAbs returns true if the method "MLGraphBuilder.abs" exists.
func (this MLGraphBuilder) HasAbs() bool {
	return js.True == bindings.HasMLGraphBuilderAbs(
		this.Ref(),
	)
}

// AbsFunc returns the method "MLGraphBuilder.abs".
func (this MLGraphBuilder) AbsFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderAbsFunc(
			this.Ref(),
		),
	)
}

// Abs calls the method "MLGraphBuilder.abs".
func (this MLGraphBuilder) Abs(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderAbs(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryAbs calls the method "MLGraphBuilder.abs"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryAbs(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderAbs(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasCeil returns true if the method "MLGraphBuilder.ceil" exists.
func (this MLGraphBuilder) HasCeil() bool {
	return js.True == bindings.HasMLGraphBuilderCeil(
		this.Ref(),
	)
}

// CeilFunc returns the method "MLGraphBuilder.ceil".
func (this MLGraphBuilder) CeilFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderCeilFunc(
			this.Ref(),
		),
	)
}

// Ceil calls the method "MLGraphBuilder.ceil".
func (this MLGraphBuilder) Ceil(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderCeil(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryCeil calls the method "MLGraphBuilder.ceil"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryCeil(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderCeil(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasCos returns true if the method "MLGraphBuilder.cos" exists.
func (this MLGraphBuilder) HasCos() bool {
	return js.True == bindings.HasMLGraphBuilderCos(
		this.Ref(),
	)
}

// CosFunc returns the method "MLGraphBuilder.cos".
func (this MLGraphBuilder) CosFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderCosFunc(
			this.Ref(),
		),
	)
}

// Cos calls the method "MLGraphBuilder.cos".
func (this MLGraphBuilder) Cos(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderCos(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryCos calls the method "MLGraphBuilder.cos"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryCos(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderCos(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasExp returns true if the method "MLGraphBuilder.exp" exists.
func (this MLGraphBuilder) HasExp() bool {
	return js.True == bindings.HasMLGraphBuilderExp(
		this.Ref(),
	)
}

// ExpFunc returns the method "MLGraphBuilder.exp".
func (this MLGraphBuilder) ExpFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderExpFunc(
			this.Ref(),
		),
	)
}

// Exp calls the method "MLGraphBuilder.exp".
func (this MLGraphBuilder) Exp(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderExp(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryExp calls the method "MLGraphBuilder.exp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryExp(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderExp(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFloor returns true if the method "MLGraphBuilder.floor" exists.
func (this MLGraphBuilder) HasFloor() bool {
	return js.True == bindings.HasMLGraphBuilderFloor(
		this.Ref(),
	)
}

// FloorFunc returns the method "MLGraphBuilder.floor".
func (this MLGraphBuilder) FloorFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderFloorFunc(
			this.Ref(),
		),
	)
}

// Floor calls the method "MLGraphBuilder.floor".
func (this MLGraphBuilder) Floor(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderFloor(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryFloor calls the method "MLGraphBuilder.floor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryFloor(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderFloor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasLog returns true if the method "MLGraphBuilder.log" exists.
func (this MLGraphBuilder) HasLog() bool {
	return js.True == bindings.HasMLGraphBuilderLog(
		this.Ref(),
	)
}

// LogFunc returns the method "MLGraphBuilder.log".
func (this MLGraphBuilder) LogFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLogFunc(
			this.Ref(),
		),
	)
}

// Log calls the method "MLGraphBuilder.log".
func (this MLGraphBuilder) Log(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderLog(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryLog calls the method "MLGraphBuilder.log"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryLog(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderLog(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasNeg returns true if the method "MLGraphBuilder.neg" exists.
func (this MLGraphBuilder) HasNeg() bool {
	return js.True == bindings.HasMLGraphBuilderNeg(
		this.Ref(),
	)
}

// NegFunc returns the method "MLGraphBuilder.neg".
func (this MLGraphBuilder) NegFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderNegFunc(
			this.Ref(),
		),
	)
}

// Neg calls the method "MLGraphBuilder.neg".
func (this MLGraphBuilder) Neg(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderNeg(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryNeg calls the method "MLGraphBuilder.neg"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryNeg(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderNeg(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasSin returns true if the method "MLGraphBuilder.sin" exists.
func (this MLGraphBuilder) HasSin() bool {
	return js.True == bindings.HasMLGraphBuilderSin(
		this.Ref(),
	)
}

// SinFunc returns the method "MLGraphBuilder.sin".
func (this MLGraphBuilder) SinFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSinFunc(
			this.Ref(),
		),
	)
}

// Sin calls the method "MLGraphBuilder.sin".
func (this MLGraphBuilder) Sin(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderSin(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TrySin calls the method "MLGraphBuilder.sin"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TrySin(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderSin(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasTan returns true if the method "MLGraphBuilder.tan" exists.
func (this MLGraphBuilder) HasTan() bool {
	return js.True == bindings.HasMLGraphBuilderTan(
		this.Ref(),
	)
}

// TanFunc returns the method "MLGraphBuilder.tan".
func (this MLGraphBuilder) TanFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderTanFunc(
			this.Ref(),
		),
	)
}

// Tan calls the method "MLGraphBuilder.tan".
func (this MLGraphBuilder) Tan(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderTan(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryTan calls the method "MLGraphBuilder.tan"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryTan(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderTan(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasTranspose returns true if the method "MLGraphBuilder.transpose" exists.
func (this MLGraphBuilder) HasTranspose() bool {
	return js.True == bindings.HasMLGraphBuilderTranspose(
		this.Ref(),
	)
}

// TransposeFunc returns the method "MLGraphBuilder.transpose".
func (this MLGraphBuilder) TransposeFunc() (fn js.Func[func(input MLOperand, options MLTransposeOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderTransposeFunc(
			this.Ref(),
		),
	)
}

// Transpose calls the method "MLGraphBuilder.transpose".
func (this MLGraphBuilder) Transpose(input MLOperand, options MLTransposeOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderTranspose(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasTranspose1 returns true if the method "MLGraphBuilder.transpose" exists.
func (this MLGraphBuilder) HasTranspose1() bool {
	return js.True == bindings.HasMLGraphBuilderTranspose1(
		this.Ref(),
	)
}

// Transpose1Func returns the method "MLGraphBuilder.transpose".
func (this MLGraphBuilder) Transpose1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderTranspose1Func(
			this.Ref(),
		),
	)
}

// Transpose1 calls the method "MLGraphBuilder.transpose".
func (this MLGraphBuilder) Transpose1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderTranspose1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryTranspose1 calls the method "MLGraphBuilder.transpose"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryTranspose1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderTranspose1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasPrelu returns true if the method "MLGraphBuilder.prelu" exists.
func (this MLGraphBuilder) HasPrelu() bool {
	return js.True == bindings.HasMLGraphBuilderPrelu(
		this.Ref(),
	)
}

// PreluFunc returns the method "MLGraphBuilder.prelu".
func (this MLGraphBuilder) PreluFunc() (fn js.Func[func(input MLOperand, slope MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderPreluFunc(
			this.Ref(),
		),
	)
}

// Prelu calls the method "MLGraphBuilder.prelu".
func (this MLGraphBuilder) Prelu(input MLOperand, slope MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderPrelu(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		slope.Ref(),
	)

	return
}

// HasConcat returns true if the method "MLGraphBuilder.concat" exists.
func (this MLGraphBuilder) HasConcat() bool {
	return js.True == bindings.HasMLGraphBuilderConcat(
		this.Ref(),
	)
}

// ConcatFunc returns the method "MLGraphBuilder.concat".
func (this MLGraphBuilder) ConcatFunc() (fn js.Func[func(inputs js.Array[MLOperand], axis uint32) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderConcatFunc(
			this.Ref(),
		),
	)
}

// Concat calls the method "MLGraphBuilder.concat".
func (this MLGraphBuilder) Concat(inputs js.Array[MLOperand], axis uint32) (ret MLOperand) {
	bindings.CallMLGraphBuilderConcat(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		inputs.Ref(),
		uint32(axis),
	)

	return
}

// HasGemm returns true if the method "MLGraphBuilder.gemm" exists.
func (this MLGraphBuilder) HasGemm() bool {
	return js.True == bindings.HasMLGraphBuilderGemm(
		this.Ref(),
	)
}

// GemmFunc returns the method "MLGraphBuilder.gemm".
func (this MLGraphBuilder) GemmFunc() (fn js.Func[func(a MLOperand, b MLOperand, options MLGemmOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderGemmFunc(
			this.Ref(),
		),
	)
}

// Gemm calls the method "MLGraphBuilder.gemm".
func (this MLGraphBuilder) Gemm(a MLOperand, b MLOperand, options MLGemmOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderGemm(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasGemm1 returns true if the method "MLGraphBuilder.gemm" exists.
func (this MLGraphBuilder) HasGemm1() bool {
	return js.True == bindings.HasMLGraphBuilderGemm1(
		this.Ref(),
	)
}

// Gemm1Func returns the method "MLGraphBuilder.gemm".
func (this MLGraphBuilder) Gemm1Func() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderGemm1Func(
			this.Ref(),
		),
	)
}

// Gemm1 calls the method "MLGraphBuilder.gemm".
func (this MLGraphBuilder) Gemm1(a MLOperand, b MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderGemm1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		a.Ref(),
		b.Ref(),
	)

	return
}

// HasLstmCell returns true if the method "MLGraphBuilder.lstmCell" exists.
func (this MLGraphBuilder) HasLstmCell() bool {
	return js.True == bindings.HasMLGraphBuilderLstmCell(
		this.Ref(),
	)
}

// LstmCellFunc returns the method "MLGraphBuilder.lstmCell".
func (this MLGraphBuilder) LstmCellFunc() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, cellState MLOperand, hiddenSize uint32, options MLLstmCellOptions) js.Array[MLOperand]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLstmCellFunc(
			this.Ref(),
		),
	)
}

// LstmCell calls the method "MLGraphBuilder.lstmCell".
func (this MLGraphBuilder) LstmCell(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, cellState MLOperand, hiddenSize uint32, options MLLstmCellOptions) (ret js.Array[MLOperand]) {
	bindings.CallMLGraphBuilderLstmCell(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

// HasLstmCell1 returns true if the method "MLGraphBuilder.lstmCell" exists.
func (this MLGraphBuilder) HasLstmCell1() bool {
	return js.True == bindings.HasMLGraphBuilderLstmCell1(
		this.Ref(),
	)
}

// LstmCell1Func returns the method "MLGraphBuilder.lstmCell".
func (this MLGraphBuilder) LstmCell1Func() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, cellState MLOperand, hiddenSize uint32) js.Array[MLOperand]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLstmCell1Func(
			this.Ref(),
		),
	)
}

// LstmCell1 calls the method "MLGraphBuilder.lstmCell".
func (this MLGraphBuilder) LstmCell1(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, cellState MLOperand, hiddenSize uint32) (ret js.Array[MLOperand]) {
	bindings.CallMLGraphBuilderLstmCell1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		hiddenState.Ref(),
		cellState.Ref(),
		uint32(hiddenSize),
	)

	return
}

// HasBatchNormalization returns true if the method "MLGraphBuilder.batchNormalization" exists.
func (this MLGraphBuilder) HasBatchNormalization() bool {
	return js.True == bindings.HasMLGraphBuilderBatchNormalization(
		this.Ref(),
	)
}

// BatchNormalizationFunc returns the method "MLGraphBuilder.batchNormalization".
func (this MLGraphBuilder) BatchNormalizationFunc() (fn js.Func[func(input MLOperand, mean MLOperand, variance MLOperand, options MLBatchNormalizationOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderBatchNormalizationFunc(
			this.Ref(),
		),
	)
}

// BatchNormalization calls the method "MLGraphBuilder.batchNormalization".
func (this MLGraphBuilder) BatchNormalization(input MLOperand, mean MLOperand, variance MLOperand, options MLBatchNormalizationOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderBatchNormalization(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		mean.Ref(),
		variance.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasBatchNormalization1 returns true if the method "MLGraphBuilder.batchNormalization" exists.
func (this MLGraphBuilder) HasBatchNormalization1() bool {
	return js.True == bindings.HasMLGraphBuilderBatchNormalization1(
		this.Ref(),
	)
}

// BatchNormalization1Func returns the method "MLGraphBuilder.batchNormalization".
func (this MLGraphBuilder) BatchNormalization1Func() (fn js.Func[func(input MLOperand, mean MLOperand, variance MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderBatchNormalization1Func(
			this.Ref(),
		),
	)
}

// BatchNormalization1 calls the method "MLGraphBuilder.batchNormalization".
func (this MLGraphBuilder) BatchNormalization1(input MLOperand, mean MLOperand, variance MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderBatchNormalization1(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		mean.Ref(),
		variance.Ref(),
	)

	return
}

// HasElu returns true if the method "MLGraphBuilder.elu" exists.
func (this MLGraphBuilder) HasElu() bool {
	return js.True == bindings.HasMLGraphBuilderElu(
		this.Ref(),
	)
}

// EluFunc returns the method "MLGraphBuilder.elu".
func (this MLGraphBuilder) EluFunc() (fn js.Func[func(input MLOperand, options MLEluOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderEluFunc(
			this.Ref(),
		),
	)
}

// Elu calls the method "MLGraphBuilder.elu".
func (this MLGraphBuilder) Elu(input MLOperand, options MLEluOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderElu(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasElu1 returns true if the method "MLGraphBuilder.elu" exists.
func (this MLGraphBuilder) HasElu1() bool {
	return js.True == bindings.HasMLGraphBuilderElu1(
		this.Ref(),
	)
}

// Elu1Func returns the method "MLGraphBuilder.elu".
func (this MLGraphBuilder) Elu1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderElu1Func(
			this.Ref(),
		),
	)
}

// Elu1 calls the method "MLGraphBuilder.elu".
func (this MLGraphBuilder) Elu1(input MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderElu1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryElu1 calls the method "MLGraphBuilder.elu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryElu1(input MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderElu1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasElu2 returns true if the method "MLGraphBuilder.elu" exists.
func (this MLGraphBuilder) HasElu2() bool {
	return js.True == bindings.HasMLGraphBuilderElu2(
		this.Ref(),
	)
}

// Elu2Func returns the method "MLGraphBuilder.elu".
func (this MLGraphBuilder) Elu2Func() (fn js.Func[func(options MLEluOptions) MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderElu2Func(
			this.Ref(),
		),
	)
}

// Elu2 calls the method "MLGraphBuilder.elu".
func (this MLGraphBuilder) Elu2(options MLEluOptions) (ret MLActivation) {
	bindings.CallMLGraphBuilderElu2(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryElu2 calls the method "MLGraphBuilder.elu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryElu2(options MLEluOptions) (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderElu2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasElu3 returns true if the method "MLGraphBuilder.elu" exists.
func (this MLGraphBuilder) HasElu3() bool {
	return js.True == bindings.HasMLGraphBuilderElu3(
		this.Ref(),
	)
}

// Elu3Func returns the method "MLGraphBuilder.elu".
func (this MLGraphBuilder) Elu3Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderElu3Func(
			this.Ref(),
		),
	)
}

// Elu3 calls the method "MLGraphBuilder.elu".
func (this MLGraphBuilder) Elu3() (ret MLActivation) {
	bindings.CallMLGraphBuilderElu3(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryElu3 calls the method "MLGraphBuilder.elu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryElu3() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderElu3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClamp returns true if the method "MLGraphBuilder.clamp" exists.
func (this MLGraphBuilder) HasClamp() bool {
	return js.True == bindings.HasMLGraphBuilderClamp(
		this.Ref(),
	)
}

// ClampFunc returns the method "MLGraphBuilder.clamp".
func (this MLGraphBuilder) ClampFunc() (fn js.Func[func(operand MLOperand, options MLClampOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderClampFunc(
			this.Ref(),
		),
	)
}

// Clamp calls the method "MLGraphBuilder.clamp".
func (this MLGraphBuilder) Clamp(operand MLOperand, options MLClampOptions) (ret MLOperand) {
	bindings.CallMLGraphBuilderClamp(
		this.Ref(), js.Pointer(&ret),
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
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		operand.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasClamp1 returns true if the method "MLGraphBuilder.clamp" exists.
func (this MLGraphBuilder) HasClamp1() bool {
	return js.True == bindings.HasMLGraphBuilderClamp1(
		this.Ref(),
	)
}

// Clamp1Func returns the method "MLGraphBuilder.clamp".
func (this MLGraphBuilder) Clamp1Func() (fn js.Func[func(operand MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderClamp1Func(
			this.Ref(),
		),
	)
}

// Clamp1 calls the method "MLGraphBuilder.clamp".
func (this MLGraphBuilder) Clamp1(operand MLOperand) (ret MLOperand) {
	bindings.CallMLGraphBuilderClamp1(
		this.Ref(), js.Pointer(&ret),
		operand.Ref(),
	)

	return
}

// TryClamp1 calls the method "MLGraphBuilder.clamp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryClamp1(operand MLOperand) (ret MLOperand, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderClamp1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		operand.Ref(),
	)

	return
}

// HasClamp2 returns true if the method "MLGraphBuilder.clamp" exists.
func (this MLGraphBuilder) HasClamp2() bool {
	return js.True == bindings.HasMLGraphBuilderClamp2(
		this.Ref(),
	)
}

// Clamp2Func returns the method "MLGraphBuilder.clamp".
func (this MLGraphBuilder) Clamp2Func() (fn js.Func[func(options MLClampOptions) MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderClamp2Func(
			this.Ref(),
		),
	)
}

// Clamp2 calls the method "MLGraphBuilder.clamp".
func (this MLGraphBuilder) Clamp2(options MLClampOptions) (ret MLActivation) {
	bindings.CallMLGraphBuilderClamp2(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryClamp2 calls the method "MLGraphBuilder.clamp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryClamp2(options MLClampOptions) (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderClamp2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasClamp3 returns true if the method "MLGraphBuilder.clamp" exists.
func (this MLGraphBuilder) HasClamp3() bool {
	return js.True == bindings.HasMLGraphBuilderClamp3(
		this.Ref(),
	)
}

// Clamp3Func returns the method "MLGraphBuilder.clamp".
func (this MLGraphBuilder) Clamp3Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderClamp3Func(
			this.Ref(),
		),
	)
}

// Clamp3 calls the method "MLGraphBuilder.clamp".
func (this MLGraphBuilder) Clamp3() (ret MLActivation) {
	bindings.CallMLGraphBuilderClamp3(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClamp3 calls the method "MLGraphBuilder.clamp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MLGraphBuilder) TryClamp3() (ret MLActivation, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMLGraphBuilderClamp3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
func (p MagnetometerSensorOptions) UpdateFrom(ref js.Ref) {
	bindings.MagnetometerSensorOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MagnetometerSensorOptions) Update(ref js.Ref) {
	bindings.MagnetometerSensorOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// X returns the value of property "Magnetometer.x".
//
// It returns ok=false if there is no such property.
func (this Magnetometer) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetMagnetometerX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "Magnetometer.y".
//
// It returns ok=false if there is no such property.
func (this Magnetometer) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetMagnetometerY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Z returns the value of property "Magnetometer.z".
//
// It returns ok=false if there is no such property.
func (this Magnetometer) Z() (ret float64, ok bool) {
	ok = js.True == bindings.GetMagnetometerZ(
		this.Ref(), js.Pointer(&ret),
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
func (p MagnetometerReadingValues) UpdateFrom(ref js.Ref) {
	bindings.MagnetometerReadingValuesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MagnetometerReadingValues) Update(ref js.Ref) {
	bindings.MagnetometerReadingValuesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MathMLElement struct {
	Element
}

func (this MathMLElement) Once() MathMLElement {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Style returns the value of property "MathMLElement.style".
//
// It returns ok=false if there is no such property.
func (this MathMLElement) Style() (ret CSSStyleDeclaration, ok bool) {
	ok = js.True == bindings.GetMathMLElementStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AttributeStyleMap returns the value of property "MathMLElement.attributeStyleMap".
//
// It returns ok=false if there is no such property.
func (this MathMLElement) AttributeStyleMap() (ret StylePropertyMap, ok bool) {
	ok = js.True == bindings.GetMathMLElementAttributeStyleMap(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Dataset returns the value of property "MathMLElement.dataset".
//
// It returns ok=false if there is no such property.
func (this MathMLElement) Dataset() (ret DOMStringMap, ok bool) {
	ok = js.True == bindings.GetMathMLElementDataset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Nonce returns the value of property "MathMLElement.nonce".
//
// It returns ok=false if there is no such property.
func (this MathMLElement) Nonce() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMathMLElementNonce(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetNonce sets the value of property "MathMLElement.nonce" to val.
//
// It returns false if the property cannot be set.
func (this MathMLElement) SetNonce(val js.String) bool {
	return js.True == bindings.SetMathMLElementNonce(
		this.Ref(),
		val.Ref(),
	)
}

// Autofocus returns the value of property "MathMLElement.autofocus".
//
// It returns ok=false if there is no such property.
func (this MathMLElement) Autofocus() (ret bool, ok bool) {
	ok = js.True == bindings.GetMathMLElementAutofocus(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAutofocus sets the value of property "MathMLElement.autofocus" to val.
//
// It returns false if the property cannot be set.
func (this MathMLElement) SetAutofocus(val bool) bool {
	return js.True == bindings.SetMathMLElementAutofocus(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// TabIndex returns the value of property "MathMLElement.tabIndex".
//
// It returns ok=false if there is no such property.
func (this MathMLElement) TabIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetMathMLElementTabIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTabIndex sets the value of property "MathMLElement.tabIndex" to val.
//
// It returns false if the property cannot be set.
func (this MathMLElement) SetTabIndex(val int32) bool {
	return js.True == bindings.SetMathMLElementTabIndex(
		this.Ref(),
		int32(val),
	)
}

// HasFocus returns true if the method "MathMLElement.focus" exists.
func (this MathMLElement) HasFocus() bool {
	return js.True == bindings.HasMathMLElementFocus(
		this.Ref(),
	)
}

// FocusFunc returns the method "MathMLElement.focus".
func (this MathMLElement) FocusFunc() (fn js.Func[func(options FocusOptions)]) {
	return fn.FromRef(
		bindings.MathMLElementFocusFunc(
			this.Ref(),
		),
	)
}

// Focus calls the method "MathMLElement.focus".
func (this MathMLElement) Focus(options FocusOptions) (ret js.Void) {
	bindings.CallMathMLElementFocus(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryFocus calls the method "MathMLElement.focus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MathMLElement) TryFocus(options FocusOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMathMLElementFocus(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFocus1 returns true if the method "MathMLElement.focus" exists.
func (this MathMLElement) HasFocus1() bool {
	return js.True == bindings.HasMathMLElementFocus1(
		this.Ref(),
	)
}

// Focus1Func returns the method "MathMLElement.focus".
func (this MathMLElement) Focus1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MathMLElementFocus1Func(
			this.Ref(),
		),
	)
}

// Focus1 calls the method "MathMLElement.focus".
func (this MathMLElement) Focus1() (ret js.Void) {
	bindings.CallMathMLElementFocus1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFocus1 calls the method "MathMLElement.focus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MathMLElement) TryFocus1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMathMLElementFocus1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasBlur returns true if the method "MathMLElement.blur" exists.
func (this MathMLElement) HasBlur() bool {
	return js.True == bindings.HasMathMLElementBlur(
		this.Ref(),
	)
}

// BlurFunc returns the method "MathMLElement.blur".
func (this MathMLElement) BlurFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MathMLElementBlurFunc(
			this.Ref(),
		),
	)
}

// Blur calls the method "MathMLElement.blur".
func (this MathMLElement) Blur() (ret js.Void) {
	bindings.CallMathMLElementBlur(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryBlur calls the method "MathMLElement.blur"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MathMLElement) TryBlur() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMathMLElementBlur(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
func (p MediaCapabilitiesInfo) UpdateFrom(ref js.Ref) {
	bindings.MediaCapabilitiesInfoJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaCapabilitiesInfo) Update(ref js.Ref) {
	bindings.MediaCapabilitiesInfoJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MediaConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MediaConfigurationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaConfiguration) Update(ref js.Ref) {
	bindings.MediaConfigurationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p MediaEncryptedEventInit) UpdateFrom(ref js.Ref) {
	bindings.MediaEncryptedEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaEncryptedEventInit) Update(ref js.Ref) {
	bindings.MediaEncryptedEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// InitDataType returns the value of property "MediaEncryptedEvent.initDataType".
//
// It returns ok=false if there is no such property.
func (this MediaEncryptedEvent) InitDataType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaEncryptedEventInitDataType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// InitData returns the value of property "MediaEncryptedEvent.initData".
//
// It returns ok=false if there is no such property.
func (this MediaEncryptedEvent) InitData() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetMediaEncryptedEventInitData(
		this.Ref(), js.Pointer(&ret),
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
func (p MediaKeyMessageEventInit) UpdateFrom(ref js.Ref) {
	bindings.MediaKeyMessageEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaKeyMessageEventInit) Update(ref js.Ref) {
	bindings.MediaKeyMessageEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// MessageType returns the value of property "MediaKeyMessageEvent.messageType".
//
// It returns ok=false if there is no such property.
func (this MediaKeyMessageEvent) MessageType() (ret MediaKeyMessageType, ok bool) {
	ok = js.True == bindings.GetMediaKeyMessageEventMessageType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Message returns the value of property "MediaKeyMessageEvent.message".
//
// It returns ok=false if there is no such property.
func (this MediaKeyMessageEvent) Message() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetMediaKeyMessageEventMessage(
		this.Ref(), js.Pointer(&ret),
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
func (p MediaQueryListEventInit) UpdateFrom(ref js.Ref) {
	bindings.MediaQueryListEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaQueryListEventInit) Update(ref js.Ref) {
	bindings.MediaQueryListEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Media returns the value of property "MediaQueryListEvent.media".
//
// It returns ok=false if there is no such property.
func (this MediaQueryListEvent) Media() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaQueryListEventMedia(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Matches returns the value of property "MediaQueryListEvent.matches".
//
// It returns ok=false if there is no such property.
func (this MediaQueryListEvent) Matches() (ret bool, ok bool) {
	ok = js.True == bindings.GetMediaQueryListEventMatches(
		this.Ref(), js.Pointer(&ret),
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
func (p MediaRecorderOptions) UpdateFrom(ref js.Ref) {
	bindings.MediaRecorderOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaRecorderOptions) Update(ref js.Ref) {
	bindings.MediaRecorderOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RecordingState uint32

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

// New creates a new {0x140004cc0e0 MLOperandDescriptor MLOperandDescriptor [// MLOperandDescriptor] [0x140008d0b40 0x140008d0be0] 0x140008be4b0 {0 0}} in the application heap.
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
	Alpha float32
	// Beta is "MLHardSigmoidOptions.beta"
	//
	// Optional, defaults to 0.5.
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

// New creates a new {0x140004cc0e0 MLHardSigmoidOptions MLHardSigmoidOptions [// MLHardSigmoidOptions] [0x140008d0c80 0x140008d0dc0 0x140008d0d20 0x140008d0e60] 0x140008be540 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 MLGruCellOptions MLGruCellOptions [// MLGruCellOptions] [0x140008d0f00 0x140008d0fa0 0x140008d1040 0x140008d1180 0x140008d1220 0x140008d10e0] 0x140008be5b8 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 MLPool2dOptions MLPool2dOptions [// MLPool2dOptions] [0x140008d12c0 0x140008d1360 0x140008d1400 0x140008d14a0 0x140008d1540 0x140008d15e0 0x140008d1680 0x140008d1720] 0x140008be630 {0 0}} in the application heap.
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
	Alpha float32
	// Beta is "MLLinearOptions.beta"
	//
	// Optional, defaults to 0.
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

// New creates a new {0x140004cc0e0 MLLinearOptions MLLinearOptions [// MLLinearOptions] [0x140008d17c0 0x140008d1900 0x140008d1860 0x140008d19a0] 0x140008be6a8 {0 0}} in the application heap.
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
	Alpha float32

	FFI_USE_Alpha bool // for Alpha.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLLeakyReluOptions with all fields set.
func (p MLLeakyReluOptions) FromRef(ref js.Ref) MLLeakyReluOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MLLeakyReluOptions MLLeakyReluOptions [// MLLeakyReluOptions] [0x140008d1a40 0x140008d1ae0] 0x140008be720 {0 0}} in the application heap.
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
	Value float32

	FFI_USE_Value bool // for Value.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLPadOptions with all fields set.
func (p MLPadOptions) FromRef(ref js.Ref) MLPadOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MLPadOptions MLPadOptions [// MLPadOptions] [0x140008d1b80 0x140008d1c20 0x140008d1cc0] 0x140008be780 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 MLInstanceNormalizationOptions MLInstanceNormalizationOptions [// MLInstanceNormalizationOptions] [0x140008d1d60 0x140008d1e00 0x140008d1ea0 0x140008dc000 0x140008d1f40] 0x140008be7c8 {0 0}} in the application heap.
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
	Steepness float32

	FFI_USE_Steepness bool // for Steepness.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLSoftplusOptions with all fields set.
func (p MLSoftplusOptions) FromRef(ref js.Ref) MLSoftplusOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MLSoftplusOptions MLSoftplusOptions [// MLSoftplusOptions] [0x140008dc0a0 0x140008dc140] 0x140008be840 {0 0}} in the application heap.
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
	Axis uint32

	FFI_USE_Axis bool // for Axis.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLSplitOptions with all fields set.
func (p MLSplitOptions) FromRef(ref js.Ref) MLSplitOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MLSplitOptions MLSplitOptions [// MLSplitOptions] [0x140008dc1e0 0x140008dc280] 0x140008be8d0 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 MLResample2dOptions MLResample2dOptions [// MLResample2dOptions] [0x140008dc320 0x140008dc3c0 0x140008dc460 0x140008dc500] 0x140008be930 {0 0}} in the application heap.
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
	KeepDimensions bool

	FFI_USE_KeepDimensions bool // for KeepDimensions.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MLReduceOptions with all fields set.
func (p MLReduceOptions) FromRef(ref js.Ref) MLReduceOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MLReduceOptions MLReduceOptions [// MLReduceOptions] [0x140008dc5a0 0x140008dc640 0x140008dc6e0] 0x140008beaf8 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 MLLstmOptions MLLstmOptions [// MLLstmOptions] [0x140008dc780 0x140008dc820 0x140008dc8c0 0x140008dc960 0x140008dca00 0x140008dcaa0 0x140008dcbe0 0x140008dcc80 0x140008dcd20 0x140008dcb40] 0x140008bec48 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 MLSqueezeOptions MLSqueezeOptions [// MLSqueezeOptions] [0x140008dcdc0] 0x140008becc0 {0 0}} in the application heap.
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
	ResetAfter bool
	// ReturnSequence is "MLGruOptions.returnSequence"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 MLGruOptions MLGruOptions [// MLGruOptions] [0x140008dce60 0x140008dcf00 0x140008dcfa0 0x140008dd040 0x140008dd180 0x140008dd2c0 0x140008dd360 0x140008dd400 0x140008dd0e0 0x140008dd220] 0x140008bed50 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 MLTransposeOptions MLTransposeOptions [// MLTransposeOptions] [0x140008dd4a0] 0x140008beea0 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 MLLstmCellOptions MLLstmCellOptions [// MLLstmCellOptions] [0x140008dd540 0x140008dd5e0 0x140008dd680 0x140008dd720 0x140008dd7c0] 0x140008bef18 {0 0}} in the application heap.
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

func NewMLGraphBuilder(context MLContext) MLGraphBuilder {
	return MLGraphBuilder{}.FromRef(
		bindings.NewMLGraphBuilderByMLGraphBuilder(
			context.Ref()),
	)
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

// Input calls the method "MLGraphBuilder.input".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Input(name js.String, descriptor MLOperandDescriptor) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderInput(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		js.Pointer(&descriptor),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// InputFunc returns the method "MLGraphBuilder.input".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) InputFunc() (fn js.Func[func(name js.String, descriptor MLOperandDescriptor) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderInputFunc(
			this.Ref(),
		),
	)
}

// Constant calls the method "MLGraphBuilder.constant".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Constant(descriptor MLOperandDescriptor, bufferView MLBufferView) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderConstant(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
		bufferView.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ConstantFunc returns the method "MLGraphBuilder.constant".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ConstantFunc() (fn js.Func[func(descriptor MLOperandDescriptor, bufferView MLBufferView) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderConstantFunc(
			this.Ref(),
		),
	)
}

// Constant1 calls the method "MLGraphBuilder.constant".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Constant1(value float64, typ MLOperandType) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderConstant1(
		this.Ref(), js.Pointer(&_ok),
		float64(value),
		uint32(typ),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// Constant1Func returns the method "MLGraphBuilder.constant".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Constant1Func() (fn js.Func[func(value float64, typ MLOperandType) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderConstant1Func(
			this.Ref(),
		),
	)
}

// Constant2 calls the method "MLGraphBuilder.constant".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Constant2(value float64) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderConstant2(
		this.Ref(), js.Pointer(&_ok),
		float64(value),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// Constant2Func returns the method "MLGraphBuilder.constant".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Constant2Func() (fn js.Func[func(value float64) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderConstant2Func(
			this.Ref(),
		),
	)
}

// Build calls the method "MLGraphBuilder.build".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Build(outputs MLNamedOperands) (js.Promise[MLGraph], bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderBuild(
		this.Ref(), js.Pointer(&_ok),
		outputs.Ref(),
	)

	return js.Promise[MLGraph]{}.FromRef(_ret), _ok
}

// BuildFunc returns the method "MLGraphBuilder.build".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) BuildFunc() (fn js.Func[func(outputs MLNamedOperands) js.Promise[MLGraph]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderBuildFunc(
			this.Ref(),
		),
	)
}

// BuildSync calls the method "MLGraphBuilder.buildSync".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) BuildSync(outputs MLNamedOperands) (MLGraph, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderBuildSync(
		this.Ref(), js.Pointer(&_ok),
		outputs.Ref(),
	)

	return MLGraph{}.FromRef(_ret), _ok
}

// BuildSyncFunc returns the method "MLGraphBuilder.buildSync".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) BuildSyncFunc() (fn js.Func[func(outputs MLNamedOperands) MLGraph]) {
	return fn.FromRef(
		bindings.MLGraphBuilderBuildSyncFunc(
			this.Ref(),
		),
	)
}

// HardSigmoid calls the method "MLGraphBuilder.hardSigmoid".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) HardSigmoid(input MLOperand, options MLHardSigmoidOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderHardSigmoid(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// HardSigmoidFunc returns the method "MLGraphBuilder.hardSigmoid".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) HardSigmoidFunc() (fn js.Func[func(input MLOperand, options MLHardSigmoidOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderHardSigmoidFunc(
			this.Ref(),
		),
	)
}

// HardSigmoid1 calls the method "MLGraphBuilder.hardSigmoid".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) HardSigmoid1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderHardSigmoid1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// HardSigmoid1Func returns the method "MLGraphBuilder.hardSigmoid".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) HardSigmoid1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderHardSigmoid1Func(
			this.Ref(),
		),
	)
}

// HardSigmoid2 calls the method "MLGraphBuilder.hardSigmoid".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) HardSigmoid2(options MLHardSigmoidOptions) (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderHardSigmoid2(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// HardSigmoid2Func returns the method "MLGraphBuilder.hardSigmoid".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) HardSigmoid2Func() (fn js.Func[func(options MLHardSigmoidOptions) MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderHardSigmoid2Func(
			this.Ref(),
		),
	)
}

// HardSigmoid3 calls the method "MLGraphBuilder.hardSigmoid".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) HardSigmoid3() (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderHardSigmoid3(
		this.Ref(), js.Pointer(&_ok),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// HardSigmoid3Func returns the method "MLGraphBuilder.hardSigmoid".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) HardSigmoid3Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderHardSigmoid3Func(
			this.Ref(),
		),
	)
}

// GruCell calls the method "MLGraphBuilder.gruCell".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) GruCell(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, hiddenSize uint32, options MLGruCellOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderGruCell(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		hiddenState.Ref(),
		uint32(hiddenSize),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// GruCellFunc returns the method "MLGraphBuilder.gruCell".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) GruCellFunc() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, hiddenSize uint32, options MLGruCellOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderGruCellFunc(
			this.Ref(),
		),
	)
}

// GruCell1 calls the method "MLGraphBuilder.gruCell".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) GruCell1(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, hiddenSize uint32) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderGruCell1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		hiddenState.Ref(),
		uint32(hiddenSize),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// GruCell1Func returns the method "MLGraphBuilder.gruCell".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) GruCell1Func() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, hiddenSize uint32) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderGruCell1Func(
			this.Ref(),
		),
	)
}

// Slice calls the method "MLGraphBuilder.slice".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Slice(input MLOperand, starts js.Array[uint32], sizes js.Array[uint32]) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSlice(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		starts.Ref(),
		sizes.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// SliceFunc returns the method "MLGraphBuilder.slice".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) SliceFunc() (fn js.Func[func(input MLOperand, starts js.Array[uint32], sizes js.Array[uint32]) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSliceFunc(
			this.Ref(),
		),
	)
}

// AveragePool2d calls the method "MLGraphBuilder.averagePool2d".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) AveragePool2d(input MLOperand, options MLPool2dOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderAveragePool2d(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// AveragePool2dFunc returns the method "MLGraphBuilder.averagePool2d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) AveragePool2dFunc() (fn js.Func[func(input MLOperand, options MLPool2dOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderAveragePool2dFunc(
			this.Ref(),
		),
	)
}

// AveragePool2d1 calls the method "MLGraphBuilder.averagePool2d".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) AveragePool2d1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderAveragePool2d1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// AveragePool2d1Func returns the method "MLGraphBuilder.averagePool2d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) AveragePool2d1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderAveragePool2d1Func(
			this.Ref(),
		),
	)
}

// L2Pool2d calls the method "MLGraphBuilder.l2Pool2d".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) L2Pool2d(input MLOperand, options MLPool2dOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderL2Pool2d(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// L2Pool2dFunc returns the method "MLGraphBuilder.l2Pool2d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) L2Pool2dFunc() (fn js.Func[func(input MLOperand, options MLPool2dOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderL2Pool2dFunc(
			this.Ref(),
		),
	)
}

// L2Pool2d1 calls the method "MLGraphBuilder.l2Pool2d".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) L2Pool2d1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderL2Pool2d1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// L2Pool2d1Func returns the method "MLGraphBuilder.l2Pool2d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) L2Pool2d1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderL2Pool2d1Func(
			this.Ref(),
		),
	)
}

// MaxPool2d calls the method "MLGraphBuilder.maxPool2d".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) MaxPool2d(input MLOperand, options MLPool2dOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderMaxPool2d(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// MaxPool2dFunc returns the method "MLGraphBuilder.maxPool2d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) MaxPool2dFunc() (fn js.Func[func(input MLOperand, options MLPool2dOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderMaxPool2dFunc(
			this.Ref(),
		),
	)
}

// MaxPool2d1 calls the method "MLGraphBuilder.maxPool2d".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) MaxPool2d1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderMaxPool2d1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// MaxPool2d1Func returns the method "MLGraphBuilder.maxPool2d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) MaxPool2d1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderMaxPool2d1Func(
			this.Ref(),
		),
	)
}

// Linear calls the method "MLGraphBuilder.linear".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Linear(input MLOperand, options MLLinearOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderLinear(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// LinearFunc returns the method "MLGraphBuilder.linear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) LinearFunc() (fn js.Func[func(input MLOperand, options MLLinearOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLinearFunc(
			this.Ref(),
		),
	)
}

// Linear1 calls the method "MLGraphBuilder.linear".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Linear1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderLinear1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// Linear1Func returns the method "MLGraphBuilder.linear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Linear1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLinear1Func(
			this.Ref(),
		),
	)
}

// Linear2 calls the method "MLGraphBuilder.linear".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Linear2(options MLLinearOptions) (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderLinear2(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// Linear2Func returns the method "MLGraphBuilder.linear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Linear2Func() (fn js.Func[func(options MLLinearOptions) MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLinear2Func(
			this.Ref(),
		),
	)
}

// Linear3 calls the method "MLGraphBuilder.linear".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Linear3() (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderLinear3(
		this.Ref(), js.Pointer(&_ok),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// Linear3Func returns the method "MLGraphBuilder.linear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Linear3Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLinear3Func(
			this.Ref(),
		),
	)
}

// LeakyRelu calls the method "MLGraphBuilder.leakyRelu".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) LeakyRelu(input MLOperand, options MLLeakyReluOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderLeakyRelu(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// LeakyReluFunc returns the method "MLGraphBuilder.leakyRelu".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) LeakyReluFunc() (fn js.Func[func(input MLOperand, options MLLeakyReluOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLeakyReluFunc(
			this.Ref(),
		),
	)
}

// LeakyRelu1 calls the method "MLGraphBuilder.leakyRelu".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) LeakyRelu1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderLeakyRelu1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// LeakyRelu1Func returns the method "MLGraphBuilder.leakyRelu".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) LeakyRelu1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLeakyRelu1Func(
			this.Ref(),
		),
	)
}

// LeakyRelu2 calls the method "MLGraphBuilder.leakyRelu".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) LeakyRelu2(options MLLeakyReluOptions) (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderLeakyRelu2(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// LeakyRelu2Func returns the method "MLGraphBuilder.leakyRelu".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) LeakyRelu2Func() (fn js.Func[func(options MLLeakyReluOptions) MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLeakyRelu2Func(
			this.Ref(),
		),
	)
}

// LeakyRelu3 calls the method "MLGraphBuilder.leakyRelu".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) LeakyRelu3() (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderLeakyRelu3(
		this.Ref(), js.Pointer(&_ok),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// LeakyRelu3Func returns the method "MLGraphBuilder.leakyRelu".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) LeakyRelu3Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLeakyRelu3Func(
			this.Ref(),
		),
	)
}

// Pad calls the method "MLGraphBuilder.pad".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Pad(input MLOperand, beginningPadding js.Array[uint32], endingPadding js.Array[uint32], options MLPadOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderPad(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		beginningPadding.Ref(),
		endingPadding.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// PadFunc returns the method "MLGraphBuilder.pad".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) PadFunc() (fn js.Func[func(input MLOperand, beginningPadding js.Array[uint32], endingPadding js.Array[uint32], options MLPadOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderPadFunc(
			this.Ref(),
		),
	)
}

// Pad1 calls the method "MLGraphBuilder.pad".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Pad1(input MLOperand, beginningPadding js.Array[uint32], endingPadding js.Array[uint32]) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderPad1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		beginningPadding.Ref(),
		endingPadding.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// Pad1Func returns the method "MLGraphBuilder.pad".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Pad1Func() (fn js.Func[func(input MLOperand, beginningPadding js.Array[uint32], endingPadding js.Array[uint32]) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderPad1Func(
			this.Ref(),
		),
	)
}

// InstanceNormalization calls the method "MLGraphBuilder.instanceNormalization".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) InstanceNormalization(input MLOperand, options MLInstanceNormalizationOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderInstanceNormalization(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// InstanceNormalizationFunc returns the method "MLGraphBuilder.instanceNormalization".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) InstanceNormalizationFunc() (fn js.Func[func(input MLOperand, options MLInstanceNormalizationOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderInstanceNormalizationFunc(
			this.Ref(),
		),
	)
}

// InstanceNormalization1 calls the method "MLGraphBuilder.instanceNormalization".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) InstanceNormalization1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderInstanceNormalization1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// InstanceNormalization1Func returns the method "MLGraphBuilder.instanceNormalization".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) InstanceNormalization1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderInstanceNormalization1Func(
			this.Ref(),
		),
	)
}

// Softplus calls the method "MLGraphBuilder.softplus".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Softplus(input MLOperand, options MLSoftplusOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSoftplus(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// SoftplusFunc returns the method "MLGraphBuilder.softplus".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) SoftplusFunc() (fn js.Func[func(input MLOperand, options MLSoftplusOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSoftplusFunc(
			this.Ref(),
		),
	)
}

// Softplus1 calls the method "MLGraphBuilder.softplus".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Softplus1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSoftplus1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// Softplus1Func returns the method "MLGraphBuilder.softplus".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Softplus1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSoftplus1Func(
			this.Ref(),
		),
	)
}

// Softplus2 calls the method "MLGraphBuilder.softplus".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Softplus2(options MLSoftplusOptions) (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSoftplus2(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// Softplus2Func returns the method "MLGraphBuilder.softplus".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Softplus2Func() (fn js.Func[func(options MLSoftplusOptions) MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSoftplus2Func(
			this.Ref(),
		),
	)
}

// Softplus3 calls the method "MLGraphBuilder.softplus".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Softplus3() (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSoftplus3(
		this.Ref(), js.Pointer(&_ok),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// Softplus3Func returns the method "MLGraphBuilder.softplus".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Softplus3Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSoftplus3Func(
			this.Ref(),
		),
	)
}

// Softsign calls the method "MLGraphBuilder.softsign".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Softsign(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSoftsign(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// SoftsignFunc returns the method "MLGraphBuilder.softsign".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) SoftsignFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSoftsignFunc(
			this.Ref(),
		),
	)
}

// Softsign1 calls the method "MLGraphBuilder.softsign".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Softsign1() (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSoftsign1(
		this.Ref(), js.Pointer(&_ok),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// Softsign1Func returns the method "MLGraphBuilder.softsign".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Softsign1Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSoftsign1Func(
			this.Ref(),
		),
	)
}

// Sigmoid calls the method "MLGraphBuilder.sigmoid".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Sigmoid(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSigmoid(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// SigmoidFunc returns the method "MLGraphBuilder.sigmoid".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) SigmoidFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSigmoidFunc(
			this.Ref(),
		),
	)
}

// Sigmoid1 calls the method "MLGraphBuilder.sigmoid".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Sigmoid1() (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSigmoid1(
		this.Ref(), js.Pointer(&_ok),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// Sigmoid1Func returns the method "MLGraphBuilder.sigmoid".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Sigmoid1Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSigmoid1Func(
			this.Ref(),
		),
	)
}

// Reshape calls the method "MLGraphBuilder.reshape".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Reshape(input MLOperand, newShape js.Array[uint32]) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReshape(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		newShape.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReshapeFunc returns the method "MLGraphBuilder.reshape".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReshapeFunc() (fn js.Func[func(input MLOperand, newShape js.Array[uint32]) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReshapeFunc(
			this.Ref(),
		),
	)
}

// Conv2d calls the method "MLGraphBuilder.conv2d".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Conv2d(input MLOperand, filter MLOperand, options MLConv2dOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderConv2d(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		filter.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// Conv2dFunc returns the method "MLGraphBuilder.conv2d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Conv2dFunc() (fn js.Func[func(input MLOperand, filter MLOperand, options MLConv2dOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderConv2dFunc(
			this.Ref(),
		),
	)
}

// Conv2d1 calls the method "MLGraphBuilder.conv2d".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Conv2d1(input MLOperand, filter MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderConv2d1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		filter.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// Conv2d1Func returns the method "MLGraphBuilder.conv2d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Conv2d1Func() (fn js.Func[func(input MLOperand, filter MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderConv2d1Func(
			this.Ref(),
		),
	)
}

// Split calls the method "MLGraphBuilder.split".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Split(input MLOperand, splits OneOf_Uint32_ArrayUint32, options MLSplitOptions) (js.Array[MLOperand], bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSplit(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		splits.Ref(),
		js.Pointer(&options),
	)

	return js.Array[MLOperand]{}.FromRef(_ret), _ok
}

// SplitFunc returns the method "MLGraphBuilder.split".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) SplitFunc() (fn js.Func[func(input MLOperand, splits OneOf_Uint32_ArrayUint32, options MLSplitOptions) js.Array[MLOperand]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSplitFunc(
			this.Ref(),
		),
	)
}

// Split1 calls the method "MLGraphBuilder.split".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Split1(input MLOperand, splits OneOf_Uint32_ArrayUint32) (js.Array[MLOperand], bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSplit1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		splits.Ref(),
	)

	return js.Array[MLOperand]{}.FromRef(_ret), _ok
}

// Split1Func returns the method "MLGraphBuilder.split".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Split1Func() (fn js.Func[func(input MLOperand, splits OneOf_Uint32_ArrayUint32) js.Array[MLOperand]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSplit1Func(
			this.Ref(),
		),
	)
}

// Resample2d calls the method "MLGraphBuilder.resample2d".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Resample2d(input MLOperand, options MLResample2dOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderResample2d(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// Resample2dFunc returns the method "MLGraphBuilder.resample2d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Resample2dFunc() (fn js.Func[func(input MLOperand, options MLResample2dOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderResample2dFunc(
			this.Ref(),
		),
	)
}

// Resample2d1 calls the method "MLGraphBuilder.resample2d".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Resample2d1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderResample2d1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// Resample2d1Func returns the method "MLGraphBuilder.resample2d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Resample2d1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderResample2d1Func(
			this.Ref(),
		),
	)
}

// HardSwish calls the method "MLGraphBuilder.hardSwish".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) HardSwish(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderHardSwish(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// HardSwishFunc returns the method "MLGraphBuilder.hardSwish".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) HardSwishFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderHardSwishFunc(
			this.Ref(),
		),
	)
}

// HardSwish1 calls the method "MLGraphBuilder.hardSwish".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) HardSwish1() (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderHardSwish1(
		this.Ref(), js.Pointer(&_ok),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// HardSwish1Func returns the method "MLGraphBuilder.hardSwish".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) HardSwish1Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderHardSwish1Func(
			this.Ref(),
		),
	)
}

// Softmax calls the method "MLGraphBuilder.softmax".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Softmax(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSoftmax(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// SoftmaxFunc returns the method "MLGraphBuilder.softmax".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) SoftmaxFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSoftmaxFunc(
			this.Ref(),
		),
	)
}

// Softmax1 calls the method "MLGraphBuilder.softmax".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Softmax1() (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSoftmax1(
		this.Ref(), js.Pointer(&_ok),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// Softmax1Func returns the method "MLGraphBuilder.softmax".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Softmax1Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSoftmax1Func(
			this.Ref(),
		),
	)
}

// ConvTranspose2d calls the method "MLGraphBuilder.convTranspose2d".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ConvTranspose2d(input MLOperand, filter MLOperand, options MLConvTranspose2dOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderConvTranspose2d(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		filter.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ConvTranspose2dFunc returns the method "MLGraphBuilder.convTranspose2d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ConvTranspose2dFunc() (fn js.Func[func(input MLOperand, filter MLOperand, options MLConvTranspose2dOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderConvTranspose2dFunc(
			this.Ref(),
		),
	)
}

// ConvTranspose2d1 calls the method "MLGraphBuilder.convTranspose2d".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ConvTranspose2d1(input MLOperand, filter MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderConvTranspose2d1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		filter.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ConvTranspose2d1Func returns the method "MLGraphBuilder.convTranspose2d".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ConvTranspose2d1Func() (fn js.Func[func(input MLOperand, filter MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderConvTranspose2d1Func(
			this.Ref(),
		),
	)
}

// Relu calls the method "MLGraphBuilder.relu".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Relu(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderRelu(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReluFunc returns the method "MLGraphBuilder.relu".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReluFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReluFunc(
			this.Ref(),
		),
	)
}

// Relu1 calls the method "MLGraphBuilder.relu".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Relu1() (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderRelu1(
		this.Ref(), js.Pointer(&_ok),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// Relu1Func returns the method "MLGraphBuilder.relu".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Relu1Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderRelu1Func(
			this.Ref(),
		),
	)
}

// Add calls the method "MLGraphBuilder.add".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Add(a MLOperand, b MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderAdd(
		this.Ref(), js.Pointer(&_ok),
		a.Ref(),
		b.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// AddFunc returns the method "MLGraphBuilder.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) AddFunc() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderAddFunc(
			this.Ref(),
		),
	)
}

// Sub calls the method "MLGraphBuilder.sub".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Sub(a MLOperand, b MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSub(
		this.Ref(), js.Pointer(&_ok),
		a.Ref(),
		b.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// SubFunc returns the method "MLGraphBuilder.sub".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) SubFunc() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSubFunc(
			this.Ref(),
		),
	)
}

// Mul calls the method "MLGraphBuilder.mul".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Mul(a MLOperand, b MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderMul(
		this.Ref(), js.Pointer(&_ok),
		a.Ref(),
		b.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// MulFunc returns the method "MLGraphBuilder.mul".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) MulFunc() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderMulFunc(
			this.Ref(),
		),
	)
}

// Div calls the method "MLGraphBuilder.div".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Div(a MLOperand, b MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderDiv(
		this.Ref(), js.Pointer(&_ok),
		a.Ref(),
		b.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// DivFunc returns the method "MLGraphBuilder.div".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) DivFunc() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderDivFunc(
			this.Ref(),
		),
	)
}

// Max calls the method "MLGraphBuilder.max".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Max(a MLOperand, b MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderMax(
		this.Ref(), js.Pointer(&_ok),
		a.Ref(),
		b.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// MaxFunc returns the method "MLGraphBuilder.max".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) MaxFunc() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderMaxFunc(
			this.Ref(),
		),
	)
}

// Min calls the method "MLGraphBuilder.min".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Min(a MLOperand, b MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderMin(
		this.Ref(), js.Pointer(&_ok),
		a.Ref(),
		b.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// MinFunc returns the method "MLGraphBuilder.min".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) MinFunc() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderMinFunc(
			this.Ref(),
		),
	)
}

// Pow calls the method "MLGraphBuilder.pow".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Pow(a MLOperand, b MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderPow(
		this.Ref(), js.Pointer(&_ok),
		a.Ref(),
		b.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// PowFunc returns the method "MLGraphBuilder.pow".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) PowFunc() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderPowFunc(
			this.Ref(),
		),
	)
}

// ReduceL1 calls the method "MLGraphBuilder.reduceL1".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceL1(input MLOperand, options MLReduceOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceL1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceL1Func returns the method "MLGraphBuilder.reduceL1".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceL1Func() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceL1Func(
			this.Ref(),
		),
	)
}

// ReduceL11 calls the method "MLGraphBuilder.reduceL1".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceL11(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceL11(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceL11Func returns the method "MLGraphBuilder.reduceL1".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceL11Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceL11Func(
			this.Ref(),
		),
	)
}

// ReduceL2 calls the method "MLGraphBuilder.reduceL2".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceL2(input MLOperand, options MLReduceOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceL2(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceL2Func returns the method "MLGraphBuilder.reduceL2".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceL2Func() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceL2Func(
			this.Ref(),
		),
	)
}

// ReduceL21 calls the method "MLGraphBuilder.reduceL2".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceL21(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceL21(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceL21Func returns the method "MLGraphBuilder.reduceL2".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceL21Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceL21Func(
			this.Ref(),
		),
	)
}

// ReduceLogSum calls the method "MLGraphBuilder.reduceLogSum".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceLogSum(input MLOperand, options MLReduceOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceLogSum(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceLogSumFunc returns the method "MLGraphBuilder.reduceLogSum".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceLogSumFunc() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceLogSumFunc(
			this.Ref(),
		),
	)
}

// ReduceLogSum1 calls the method "MLGraphBuilder.reduceLogSum".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceLogSum1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceLogSum1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceLogSum1Func returns the method "MLGraphBuilder.reduceLogSum".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceLogSum1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceLogSum1Func(
			this.Ref(),
		),
	)
}

// ReduceLogSumExp calls the method "MLGraphBuilder.reduceLogSumExp".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceLogSumExp(input MLOperand, options MLReduceOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceLogSumExp(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceLogSumExpFunc returns the method "MLGraphBuilder.reduceLogSumExp".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceLogSumExpFunc() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceLogSumExpFunc(
			this.Ref(),
		),
	)
}

// ReduceLogSumExp1 calls the method "MLGraphBuilder.reduceLogSumExp".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceLogSumExp1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceLogSumExp1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceLogSumExp1Func returns the method "MLGraphBuilder.reduceLogSumExp".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceLogSumExp1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceLogSumExp1Func(
			this.Ref(),
		),
	)
}

// ReduceMax calls the method "MLGraphBuilder.reduceMax".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceMax(input MLOperand, options MLReduceOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceMax(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceMaxFunc returns the method "MLGraphBuilder.reduceMax".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceMaxFunc() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceMaxFunc(
			this.Ref(),
		),
	)
}

// ReduceMax1 calls the method "MLGraphBuilder.reduceMax".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceMax1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceMax1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceMax1Func returns the method "MLGraphBuilder.reduceMax".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceMax1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceMax1Func(
			this.Ref(),
		),
	)
}

// ReduceMean calls the method "MLGraphBuilder.reduceMean".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceMean(input MLOperand, options MLReduceOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceMean(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceMeanFunc returns the method "MLGraphBuilder.reduceMean".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceMeanFunc() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceMeanFunc(
			this.Ref(),
		),
	)
}

// ReduceMean1 calls the method "MLGraphBuilder.reduceMean".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceMean1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceMean1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceMean1Func returns the method "MLGraphBuilder.reduceMean".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceMean1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceMean1Func(
			this.Ref(),
		),
	)
}

// ReduceMin calls the method "MLGraphBuilder.reduceMin".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceMin(input MLOperand, options MLReduceOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceMin(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceMinFunc returns the method "MLGraphBuilder.reduceMin".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceMinFunc() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceMinFunc(
			this.Ref(),
		),
	)
}

// ReduceMin1 calls the method "MLGraphBuilder.reduceMin".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceMin1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceMin1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceMin1Func returns the method "MLGraphBuilder.reduceMin".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceMin1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceMin1Func(
			this.Ref(),
		),
	)
}

// ReduceProduct calls the method "MLGraphBuilder.reduceProduct".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceProduct(input MLOperand, options MLReduceOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceProduct(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceProductFunc returns the method "MLGraphBuilder.reduceProduct".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceProductFunc() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceProductFunc(
			this.Ref(),
		),
	)
}

// ReduceProduct1 calls the method "MLGraphBuilder.reduceProduct".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceProduct1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceProduct1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceProduct1Func returns the method "MLGraphBuilder.reduceProduct".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceProduct1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceProduct1Func(
			this.Ref(),
		),
	)
}

// ReduceSum calls the method "MLGraphBuilder.reduceSum".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceSum(input MLOperand, options MLReduceOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceSum(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceSumFunc returns the method "MLGraphBuilder.reduceSum".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceSumFunc() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceSumFunc(
			this.Ref(),
		),
	)
}

// ReduceSum1 calls the method "MLGraphBuilder.reduceSum".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceSum1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceSum1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceSum1Func returns the method "MLGraphBuilder.reduceSum".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceSum1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceSum1Func(
			this.Ref(),
		),
	)
}

// ReduceSumSquare calls the method "MLGraphBuilder.reduceSumSquare".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceSumSquare(input MLOperand, options MLReduceOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceSumSquare(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceSumSquareFunc returns the method "MLGraphBuilder.reduceSumSquare".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceSumSquareFunc() (fn js.Func[func(input MLOperand, options MLReduceOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceSumSquareFunc(
			this.Ref(),
		),
	)
}

// ReduceSumSquare1 calls the method "MLGraphBuilder.reduceSumSquare".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) ReduceSumSquare1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderReduceSumSquare1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ReduceSumSquare1Func returns the method "MLGraphBuilder.reduceSumSquare".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ReduceSumSquare1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderReduceSumSquare1Func(
			this.Ref(),
		),
	)
}

// Lstm calls the method "MLGraphBuilder.lstm".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Lstm(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32, options MLLstmOptions) (js.Array[MLOperand], bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderLstm(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		uint32(steps),
		uint32(hiddenSize),
		js.Pointer(&options),
	)

	return js.Array[MLOperand]{}.FromRef(_ret), _ok
}

// LstmFunc returns the method "MLGraphBuilder.lstm".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) LstmFunc() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32, options MLLstmOptions) js.Array[MLOperand]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLstmFunc(
			this.Ref(),
		),
	)
}

// Lstm1 calls the method "MLGraphBuilder.lstm".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Lstm1(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32) (js.Array[MLOperand], bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderLstm1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		uint32(steps),
		uint32(hiddenSize),
	)

	return js.Array[MLOperand]{}.FromRef(_ret), _ok
}

// Lstm1Func returns the method "MLGraphBuilder.lstm".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Lstm1Func() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32) js.Array[MLOperand]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLstm1Func(
			this.Ref(),
		),
	)
}

// Matmul calls the method "MLGraphBuilder.matmul".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Matmul(a MLOperand, b MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderMatmul(
		this.Ref(), js.Pointer(&_ok),
		a.Ref(),
		b.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// MatmulFunc returns the method "MLGraphBuilder.matmul".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) MatmulFunc() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderMatmulFunc(
			this.Ref(),
		),
	)
}

// Squeeze calls the method "MLGraphBuilder.squeeze".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Squeeze(input MLOperand, options MLSqueezeOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSqueeze(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// SqueezeFunc returns the method "MLGraphBuilder.squeeze".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) SqueezeFunc() (fn js.Func[func(input MLOperand, options MLSqueezeOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSqueezeFunc(
			this.Ref(),
		),
	)
}

// Squeeze1 calls the method "MLGraphBuilder.squeeze".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Squeeze1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSqueeze1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// Squeeze1Func returns the method "MLGraphBuilder.squeeze".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Squeeze1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSqueeze1Func(
			this.Ref(),
		),
	)
}

// Tanh calls the method "MLGraphBuilder.tanh".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Tanh(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderTanh(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// TanhFunc returns the method "MLGraphBuilder.tanh".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) TanhFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderTanhFunc(
			this.Ref(),
		),
	)
}

// Tanh1 calls the method "MLGraphBuilder.tanh".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Tanh1() (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderTanh1(
		this.Ref(), js.Pointer(&_ok),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// Tanh1Func returns the method "MLGraphBuilder.tanh".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Tanh1Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderTanh1Func(
			this.Ref(),
		),
	)
}

// Gru calls the method "MLGraphBuilder.gru".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Gru(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32, options MLGruOptions) (js.Array[MLOperand], bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderGru(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		uint32(steps),
		uint32(hiddenSize),
		js.Pointer(&options),
	)

	return js.Array[MLOperand]{}.FromRef(_ret), _ok
}

// GruFunc returns the method "MLGraphBuilder.gru".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) GruFunc() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32, options MLGruOptions) js.Array[MLOperand]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderGruFunc(
			this.Ref(),
		),
	)
}

// Gru1 calls the method "MLGraphBuilder.gru".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Gru1(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32) (js.Array[MLOperand], bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderGru1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		uint32(steps),
		uint32(hiddenSize),
	)

	return js.Array[MLOperand]{}.FromRef(_ret), _ok
}

// Gru1Func returns the method "MLGraphBuilder.gru".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Gru1Func() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, steps uint32, hiddenSize uint32) js.Array[MLOperand]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderGru1Func(
			this.Ref(),
		),
	)
}

// Abs calls the method "MLGraphBuilder.abs".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Abs(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderAbs(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// AbsFunc returns the method "MLGraphBuilder.abs".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) AbsFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderAbsFunc(
			this.Ref(),
		),
	)
}

// Ceil calls the method "MLGraphBuilder.ceil".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Ceil(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderCeil(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// CeilFunc returns the method "MLGraphBuilder.ceil".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) CeilFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderCeilFunc(
			this.Ref(),
		),
	)
}

// Cos calls the method "MLGraphBuilder.cos".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Cos(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderCos(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// CosFunc returns the method "MLGraphBuilder.cos".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) CosFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderCosFunc(
			this.Ref(),
		),
	)
}

// Exp calls the method "MLGraphBuilder.exp".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Exp(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderExp(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ExpFunc returns the method "MLGraphBuilder.exp".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ExpFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderExpFunc(
			this.Ref(),
		),
	)
}

// Floor calls the method "MLGraphBuilder.floor".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Floor(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderFloor(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// FloorFunc returns the method "MLGraphBuilder.floor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) FloorFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderFloorFunc(
			this.Ref(),
		),
	)
}

// Log calls the method "MLGraphBuilder.log".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Log(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderLog(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// LogFunc returns the method "MLGraphBuilder.log".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) LogFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLogFunc(
			this.Ref(),
		),
	)
}

// Neg calls the method "MLGraphBuilder.neg".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Neg(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderNeg(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// NegFunc returns the method "MLGraphBuilder.neg".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) NegFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderNegFunc(
			this.Ref(),
		),
	)
}

// Sin calls the method "MLGraphBuilder.sin".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Sin(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderSin(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// SinFunc returns the method "MLGraphBuilder.sin".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) SinFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderSinFunc(
			this.Ref(),
		),
	)
}

// Tan calls the method "MLGraphBuilder.tan".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Tan(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderTan(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// TanFunc returns the method "MLGraphBuilder.tan".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) TanFunc() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderTanFunc(
			this.Ref(),
		),
	)
}

// Transpose calls the method "MLGraphBuilder.transpose".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Transpose(input MLOperand, options MLTransposeOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderTranspose(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// TransposeFunc returns the method "MLGraphBuilder.transpose".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) TransposeFunc() (fn js.Func[func(input MLOperand, options MLTransposeOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderTransposeFunc(
			this.Ref(),
		),
	)
}

// Transpose1 calls the method "MLGraphBuilder.transpose".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Transpose1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderTranspose1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// Transpose1Func returns the method "MLGraphBuilder.transpose".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Transpose1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderTranspose1Func(
			this.Ref(),
		),
	)
}

// Prelu calls the method "MLGraphBuilder.prelu".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Prelu(input MLOperand, slope MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderPrelu(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		slope.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// PreluFunc returns the method "MLGraphBuilder.prelu".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) PreluFunc() (fn js.Func[func(input MLOperand, slope MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderPreluFunc(
			this.Ref(),
		),
	)
}

// Concat calls the method "MLGraphBuilder.concat".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Concat(inputs js.Array[MLOperand], axis uint32) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderConcat(
		this.Ref(), js.Pointer(&_ok),
		inputs.Ref(),
		uint32(axis),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ConcatFunc returns the method "MLGraphBuilder.concat".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ConcatFunc() (fn js.Func[func(inputs js.Array[MLOperand], axis uint32) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderConcatFunc(
			this.Ref(),
		),
	)
}

// Gemm calls the method "MLGraphBuilder.gemm".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Gemm(a MLOperand, b MLOperand, options MLGemmOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderGemm(
		this.Ref(), js.Pointer(&_ok),
		a.Ref(),
		b.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// GemmFunc returns the method "MLGraphBuilder.gemm".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) GemmFunc() (fn js.Func[func(a MLOperand, b MLOperand, options MLGemmOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderGemmFunc(
			this.Ref(),
		),
	)
}

// Gemm1 calls the method "MLGraphBuilder.gemm".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Gemm1(a MLOperand, b MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderGemm1(
		this.Ref(), js.Pointer(&_ok),
		a.Ref(),
		b.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// Gemm1Func returns the method "MLGraphBuilder.gemm".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Gemm1Func() (fn js.Func[func(a MLOperand, b MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderGemm1Func(
			this.Ref(),
		),
	)
}

// LstmCell calls the method "MLGraphBuilder.lstmCell".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) LstmCell(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, cellState MLOperand, hiddenSize uint32, options MLLstmCellOptions) (js.Array[MLOperand], bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderLstmCell(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		hiddenState.Ref(),
		cellState.Ref(),
		uint32(hiddenSize),
		js.Pointer(&options),
	)

	return js.Array[MLOperand]{}.FromRef(_ret), _ok
}

// LstmCellFunc returns the method "MLGraphBuilder.lstmCell".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) LstmCellFunc() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, cellState MLOperand, hiddenSize uint32, options MLLstmCellOptions) js.Array[MLOperand]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLstmCellFunc(
			this.Ref(),
		),
	)
}

// LstmCell1 calls the method "MLGraphBuilder.lstmCell".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) LstmCell1(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, cellState MLOperand, hiddenSize uint32) (js.Array[MLOperand], bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderLstmCell1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		weight.Ref(),
		recurrentWeight.Ref(),
		hiddenState.Ref(),
		cellState.Ref(),
		uint32(hiddenSize),
	)

	return js.Array[MLOperand]{}.FromRef(_ret), _ok
}

// LstmCell1Func returns the method "MLGraphBuilder.lstmCell".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) LstmCell1Func() (fn js.Func[func(input MLOperand, weight MLOperand, recurrentWeight MLOperand, hiddenState MLOperand, cellState MLOperand, hiddenSize uint32) js.Array[MLOperand]]) {
	return fn.FromRef(
		bindings.MLGraphBuilderLstmCell1Func(
			this.Ref(),
		),
	)
}

// BatchNormalization calls the method "MLGraphBuilder.batchNormalization".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) BatchNormalization(input MLOperand, mean MLOperand, variance MLOperand, options MLBatchNormalizationOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderBatchNormalization(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		mean.Ref(),
		variance.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// BatchNormalizationFunc returns the method "MLGraphBuilder.batchNormalization".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) BatchNormalizationFunc() (fn js.Func[func(input MLOperand, mean MLOperand, variance MLOperand, options MLBatchNormalizationOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderBatchNormalizationFunc(
			this.Ref(),
		),
	)
}

// BatchNormalization1 calls the method "MLGraphBuilder.batchNormalization".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) BatchNormalization1(input MLOperand, mean MLOperand, variance MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderBatchNormalization1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		mean.Ref(),
		variance.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// BatchNormalization1Func returns the method "MLGraphBuilder.batchNormalization".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) BatchNormalization1Func() (fn js.Func[func(input MLOperand, mean MLOperand, variance MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderBatchNormalization1Func(
			this.Ref(),
		),
	)
}

// Elu calls the method "MLGraphBuilder.elu".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Elu(input MLOperand, options MLEluOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderElu(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// EluFunc returns the method "MLGraphBuilder.elu".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) EluFunc() (fn js.Func[func(input MLOperand, options MLEluOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderEluFunc(
			this.Ref(),
		),
	)
}

// Elu1 calls the method "MLGraphBuilder.elu".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Elu1(input MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderElu1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// Elu1Func returns the method "MLGraphBuilder.elu".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Elu1Func() (fn js.Func[func(input MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderElu1Func(
			this.Ref(),
		),
	)
}

// Elu2 calls the method "MLGraphBuilder.elu".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Elu2(options MLEluOptions) (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderElu2(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// Elu2Func returns the method "MLGraphBuilder.elu".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Elu2Func() (fn js.Func[func(options MLEluOptions) MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderElu2Func(
			this.Ref(),
		),
	)
}

// Elu3 calls the method "MLGraphBuilder.elu".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Elu3() (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderElu3(
		this.Ref(), js.Pointer(&_ok),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// Elu3Func returns the method "MLGraphBuilder.elu".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Elu3Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderElu3Func(
			this.Ref(),
		),
	)
}

// Clamp calls the method "MLGraphBuilder.clamp".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Clamp(operand MLOperand, options MLClampOptions) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderClamp(
		this.Ref(), js.Pointer(&_ok),
		operand.Ref(),
		js.Pointer(&options),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// ClampFunc returns the method "MLGraphBuilder.clamp".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) ClampFunc() (fn js.Func[func(operand MLOperand, options MLClampOptions) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderClampFunc(
			this.Ref(),
		),
	)
}

// Clamp1 calls the method "MLGraphBuilder.clamp".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Clamp1(operand MLOperand) (MLOperand, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderClamp1(
		this.Ref(), js.Pointer(&_ok),
		operand.Ref(),
	)

	return MLOperand{}.FromRef(_ret), _ok
}

// Clamp1Func returns the method "MLGraphBuilder.clamp".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Clamp1Func() (fn js.Func[func(operand MLOperand) MLOperand]) {
	return fn.FromRef(
		bindings.MLGraphBuilderClamp1Func(
			this.Ref(),
		),
	)
}

// Clamp2 calls the method "MLGraphBuilder.clamp".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Clamp2(options MLClampOptions) (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderClamp2(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// Clamp2Func returns the method "MLGraphBuilder.clamp".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Clamp2Func() (fn js.Func[func(options MLClampOptions) MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderClamp2Func(
			this.Ref(),
		),
	)
}

// Clamp3 calls the method "MLGraphBuilder.clamp".
//
// The returned bool will be false if there is no such method.
func (this MLGraphBuilder) Clamp3() (MLActivation, bool) {
	var _ok bool
	_ret := bindings.CallMLGraphBuilderClamp3(
		this.Ref(), js.Pointer(&_ok),
	)

	return MLActivation{}.FromRef(_ret), _ok
}

// Clamp3Func returns the method "MLGraphBuilder.clamp".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MLGraphBuilder) Clamp3Func() (fn js.Func[func() MLActivation]) {
	return fn.FromRef(
		bindings.MLGraphBuilderClamp3Func(
			this.Ref(),
		),
	)
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
	Frequency float64

	FFI_USE_Frequency bool // for Frequency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MagnetometerSensorOptions with all fields set.
func (p MagnetometerSensorOptions) FromRef(ref js.Ref) MagnetometerSensorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MagnetometerSensorOptions MagnetometerSensorOptions [// MagnetometerSensorOptions] [0x140008dd860 0x140008dd900 0x140008dd9a0] 0x140008bf020 {0 0}} in the application heap.
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

func NewMagnetometer(sensorOptions MagnetometerSensorOptions) Magnetometer {
	return Magnetometer{}.FromRef(
		bindings.NewMagnetometerByMagnetometer(
			js.Pointer(&sensorOptions)),
	)
}

func NewMagnetometerByMagnetometer1() Magnetometer {
	return Magnetometer{}.FromRef(
		bindings.NewMagnetometerByMagnetometer1(),
	)
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
// The returned bool will be false if there is no such property.
func (this Magnetometer) X() (float64, bool) {
	var _ok bool
	_ret := bindings.GetMagnetometerX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Y returns the value of property "Magnetometer.y".
//
// The returned bool will be false if there is no such property.
func (this Magnetometer) Y() (float64, bool) {
	var _ok bool
	_ret := bindings.GetMagnetometerY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Z returns the value of property "Magnetometer.z".
//
// The returned bool will be false if there is no such property.
func (this Magnetometer) Z() (float64, bool) {
	var _ok bool
	_ret := bindings.GetMagnetometerZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
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

// New creates a new {0x140004cc0e0 MagnetometerReadingValues MagnetometerReadingValues [// MagnetometerReadingValues] [0x140008ddae0 0x140008ddb80 0x140008ddc20] 0x140008bf050 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this MathMLElement) Style() (CSSStyleDeclaration, bool) {
	var _ok bool
	_ret := bindings.GetMathMLElementStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSStyleDeclaration{}.FromRef(_ret), _ok
}

// AttributeStyleMap returns the value of property "MathMLElement.attributeStyleMap".
//
// The returned bool will be false if there is no such property.
func (this MathMLElement) AttributeStyleMap() (StylePropertyMap, bool) {
	var _ok bool
	_ret := bindings.GetMathMLElementAttributeStyleMap(
		this.Ref(), js.Pointer(&_ok),
	)
	return StylePropertyMap{}.FromRef(_ret), _ok
}

// Dataset returns the value of property "MathMLElement.dataset".
//
// The returned bool will be false if there is no such property.
func (this MathMLElement) Dataset() (DOMStringMap, bool) {
	var _ok bool
	_ret := bindings.GetMathMLElementDataset(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMStringMap{}.FromRef(_ret), _ok
}

// Nonce returns the value of property "MathMLElement.nonce".
//
// The returned bool will be false if there is no such property.
func (this MathMLElement) Nonce() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMathMLElementNonce(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Nonce sets the value of property "MathMLElement.nonce" to val.
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
// The returned bool will be false if there is no such property.
func (this MathMLElement) Autofocus() (bool, bool) {
	var _ok bool
	_ret := bindings.GetMathMLElementAutofocus(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Autofocus sets the value of property "MathMLElement.autofocus" to val.
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
// The returned bool will be false if there is no such property.
func (this MathMLElement) TabIndex() (int32, bool) {
	var _ok bool
	_ret := bindings.GetMathMLElementTabIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// TabIndex sets the value of property "MathMLElement.tabIndex" to val.
//
// It returns false if the property cannot be set.
func (this MathMLElement) SetTabIndex(val int32) bool {
	return js.True == bindings.SetMathMLElementTabIndex(
		this.Ref(),
		int32(val),
	)
}

// Focus calls the method "MathMLElement.focus".
//
// The returned bool will be false if there is no such method.
func (this MathMLElement) Focus(options FocusOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMathMLElementFocus(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FocusFunc returns the method "MathMLElement.focus".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MathMLElement) FocusFunc() (fn js.Func[func(options FocusOptions)]) {
	return fn.FromRef(
		bindings.MathMLElementFocusFunc(
			this.Ref(),
		),
	)
}

// Focus1 calls the method "MathMLElement.focus".
//
// The returned bool will be false if there is no such method.
func (this MathMLElement) Focus1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMathMLElementFocus1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Focus1Func returns the method "MathMLElement.focus".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MathMLElement) Focus1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MathMLElementFocus1Func(
			this.Ref(),
		),
	)
}

// Blur calls the method "MathMLElement.blur".
//
// The returned bool will be false if there is no such method.
func (this MathMLElement) Blur() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMathMLElementBlur(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlurFunc returns the method "MathMLElement.blur".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MathMLElement) BlurFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MathMLElementBlurFunc(
			this.Ref(),
		),
	)
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

// New creates a new {0x140004cc0e0 MediaCapabilitiesInfo MediaCapabilitiesInfo [// MediaCapabilitiesInfo] [0x140008ddd60 0x140008dde00 0x140008ddea0] 0x140008bf0f8 {0 0}} in the application heap.
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
	Video VideoConfiguration
	// Audio is "MediaConfiguration.audio"
	//
	// Optional
	Audio AudioConfiguration

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaConfiguration with all fields set.
func (p MediaConfiguration) FromRef(ref js.Ref) MediaConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MediaConfiguration MediaConfiguration [// MediaConfiguration] [0x140008ddf40 0x140008ec000] 0x140008bf128 {0 0}} in the application heap.
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
	Bubbles bool
	// Cancelable is "MediaEncryptedEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "MediaEncryptedEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 MediaEncryptedEventInit MediaEncryptedEventInit [// MediaEncryptedEventInit] [0x140008ec0a0 0x140008ec140 0x140008ec1e0 0x140008ec320 0x140008ec460 0x140008ec280 0x140008ec3c0 0x140008ec500] 0x140008bf158 {0 0}} in the application heap.
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

func NewMediaEncryptedEvent(typ js.String, eventInitDict MediaEncryptedEventInit) MediaEncryptedEvent {
	return MediaEncryptedEvent{}.FromRef(
		bindings.NewMediaEncryptedEventByMediaEncryptedEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewMediaEncryptedEventByMediaEncryptedEvent1(typ js.String) MediaEncryptedEvent {
	return MediaEncryptedEvent{}.FromRef(
		bindings.NewMediaEncryptedEventByMediaEncryptedEvent1(
			typ.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this MediaEncryptedEvent) InitDataType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaEncryptedEventInitDataType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// InitData returns the value of property "MediaEncryptedEvent.initData".
//
// The returned bool will be false if there is no such property.
func (this MediaEncryptedEvent) InitData() (js.ArrayBuffer, bool) {
	var _ok bool
	_ret := bindings.GetMediaEncryptedEventInitData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.ArrayBuffer{}.FromRef(_ret), _ok
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
	Bubbles bool
	// Cancelable is "MediaKeyMessageEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "MediaKeyMessageEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 MediaKeyMessageEventInit MediaKeyMessageEventInit [// MediaKeyMessageEventInit] [0x140008ec640 0x140008ec6e0 0x140008ec780 0x140008ec8c0 0x140008eca00 0x140008ec820 0x140008ec960 0x140008ecaa0] 0x140008bf1a0 {0 0}} in the application heap.
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

func NewMediaKeyMessageEvent(typ js.String, eventInitDict MediaKeyMessageEventInit) MediaKeyMessageEvent {
	return MediaKeyMessageEvent{}.FromRef(
		bindings.NewMediaKeyMessageEventByMediaKeyMessageEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
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
// The returned bool will be false if there is no such property.
func (this MediaKeyMessageEvent) MessageType() (MediaKeyMessageType, bool) {
	var _ok bool
	_ret := bindings.GetMediaKeyMessageEventMessageType(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaKeyMessageType(_ret), _ok
}

// Message returns the value of property "MediaKeyMessageEvent.message".
//
// The returned bool will be false if there is no such property.
func (this MediaKeyMessageEvent) Message() (js.ArrayBuffer, bool) {
	var _ok bool
	_ret := bindings.GetMediaKeyMessageEventMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.ArrayBuffer{}.FromRef(_ret), _ok
}

type MediaQueryListEventInit struct {
	// Media is "MediaQueryListEventInit.media"
	//
	// Optional, defaults to "".
	Media js.String
	// Matches is "MediaQueryListEventInit.matches"
	//
	// Optional, defaults to false.
	Matches bool
	// Bubbles is "MediaQueryListEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "MediaQueryListEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "MediaQueryListEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 MediaQueryListEventInit MediaQueryListEventInit [// MediaQueryListEventInit] [0x140008ecbe0 0x140008ecc80 0x140008ecdc0 0x140008ecf00 0x140008ed040 0x140008ecd20 0x140008ece60 0x140008ecfa0 0x140008ed0e0] 0x140008bf320 {0 0}} in the application heap.
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

func NewMediaQueryListEvent(typ js.String, eventInitDict MediaQueryListEventInit) MediaQueryListEvent {
	return MediaQueryListEvent{}.FromRef(
		bindings.NewMediaQueryListEventByMediaQueryListEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewMediaQueryListEventByMediaQueryListEvent1(typ js.String) MediaQueryListEvent {
	return MediaQueryListEvent{}.FromRef(
		bindings.NewMediaQueryListEventByMediaQueryListEvent1(
			typ.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this MediaQueryListEvent) Media() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaQueryListEventMedia(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Matches returns the value of property "MediaQueryListEvent.matches".
//
// The returned bool will be false if there is no such property.
func (this MediaQueryListEvent) Matches() (bool, bool) {
	var _ok bool
	_ret := bindings.GetMediaQueryListEventMatches(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

type MediaRecorderOptions struct {
	// MimeType is "MediaRecorderOptions.mimeType"
	//
	// Optional, defaults to "".
	MimeType js.String
	// AudioBitsPerSecond is "MediaRecorderOptions.audioBitsPerSecond"
	//
	// Optional
	AudioBitsPerSecond uint32
	// VideoBitsPerSecond is "MediaRecorderOptions.videoBitsPerSecond"
	//
	// Optional
	VideoBitsPerSecond uint32
	// BitsPerSecond is "MediaRecorderOptions.bitsPerSecond"
	//
	// Optional
	BitsPerSecond uint32
	// AudioBitrateMode is "MediaRecorderOptions.audioBitrateMode"
	//
	// Optional, defaults to "variable".
	AudioBitrateMode BitrateMode
	// VideoKeyFrameIntervalDuration is "MediaRecorderOptions.videoKeyFrameIntervalDuration"
	//
	// Optional
	VideoKeyFrameIntervalDuration DOMHighResTimeStamp
	// VideoKeyFrameIntervalCount is "MediaRecorderOptions.videoKeyFrameIntervalCount"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 MediaRecorderOptions MediaRecorderOptions [// MediaRecorderOptions] [0x140008ed220 0x140008ed2c0 0x140008ed400 0x140008ed540 0x140008ed680 0x140008ed720 0x140008ed860 0x140008ed360 0x140008ed4a0 0x140008ed5e0 0x140008ed7c0 0x140008ed900] 0x140008bf398 {0 0}} in the application heap.
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

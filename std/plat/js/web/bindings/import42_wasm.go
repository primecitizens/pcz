// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

type (
	_ unsafe.Pointer
	_ js.Ref
)

//go:wasmimport plat/js/web constof_MLOperandType
//go:noescape
func ConstOfMLOperandType(str js.Ref) uint32

//go:wasmimport plat/js/web store_MLOperandDescriptor
//go:noescape
func MLOperandDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLOperandDescriptor
//go:noescape
func MLOperandDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MLHardSigmoidOptions
//go:noescape
func MLHardSigmoidOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLHardSigmoidOptions
//go:noescape
func MLHardSigmoidOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MLGruWeightLayout
//go:noescape
func ConstOfMLGruWeightLayout(str js.Ref) uint32

//go:wasmimport plat/js/web store_MLGruCellOptions
//go:noescape
func MLGruCellOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLGruCellOptions
//go:noescape
func MLGruCellOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MLRoundingType
//go:noescape
func ConstOfMLRoundingType(str js.Ref) uint32

//go:wasmimport plat/js/web store_MLPool2dOptions
//go:noescape
func MLPool2dOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLPool2dOptions
//go:noescape
func MLPool2dOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MLLinearOptions
//go:noescape
func MLLinearOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLLinearOptions
//go:noescape
func MLLinearOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MLLeakyReluOptions
//go:noescape
func MLLeakyReluOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLLeakyReluOptions
//go:noescape
func MLLeakyReluOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MLPaddingMode
//go:noescape
func ConstOfMLPaddingMode(str js.Ref) uint32

//go:wasmimport plat/js/web store_MLPadOptions
//go:noescape
func MLPadOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLPadOptions
//go:noescape
func MLPadOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MLInstanceNormalizationOptions
//go:noescape
func MLInstanceNormalizationOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLInstanceNormalizationOptions
//go:noescape
func MLInstanceNormalizationOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MLSoftplusOptions
//go:noescape
func MLSoftplusOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLSoftplusOptions
//go:noescape
func MLSoftplusOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MLSplitOptions
//go:noescape
func MLSplitOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLSplitOptions
//go:noescape
func MLSplitOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MLInterpolationMode
//go:noescape
func ConstOfMLInterpolationMode(str js.Ref) uint32

//go:wasmimport plat/js/web store_MLResample2dOptions
//go:noescape
func MLResample2dOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLResample2dOptions
//go:noescape
func MLResample2dOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MLReduceOptions
//go:noescape
func MLReduceOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLReduceOptions
//go:noescape
func MLReduceOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MLRecurrentNetworkDirection
//go:noescape
func ConstOfMLRecurrentNetworkDirection(str js.Ref) uint32

//go:wasmimport plat/js/web constof_MLLstmWeightLayout
//go:noescape
func ConstOfMLLstmWeightLayout(str js.Ref) uint32

//go:wasmimport plat/js/web store_MLLstmOptions
//go:noescape
func MLLstmOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLLstmOptions
//go:noescape
func MLLstmOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MLSqueezeOptions
//go:noescape
func MLSqueezeOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLSqueezeOptions
//go:noescape
func MLSqueezeOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MLGruOptions
//go:noescape
func MLGruOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLGruOptions
//go:noescape
func MLGruOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MLTransposeOptions
//go:noescape
func MLTransposeOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLTransposeOptions
//go:noescape
func MLTransposeOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MLLstmCellOptions
//go:noescape
func MLLstmCellOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLLstmCellOptions
//go:noescape
func MLLstmCellOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_MLGraphBuilder_MLGraphBuilder
//go:noescape
func NewMLGraphBuilderByMLGraphBuilder(
	context js.Ref) js.Ref

//go:wasmimport plat/js/web has_MLGraphBuilder_Input
//go:noescape
func HasFuncMLGraphBuilderInput(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Input
//go:noescape
func FuncMLGraphBuilderInput(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Input
//go:noescape
func CallMLGraphBuilderInput(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Input
//go:noescape
func TryMLGraphBuilderInput(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Constant
//go:noescape
func HasFuncMLGraphBuilderConstant(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Constant
//go:noescape
func FuncMLGraphBuilderConstant(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Constant
//go:noescape
func CallMLGraphBuilderConstant(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer,
	bufferView js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Constant
//go:noescape
func TryMLGraphBuilderConstant(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer,
	bufferView js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Constant1
//go:noescape
func HasFuncMLGraphBuilderConstant1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Constant1
//go:noescape
func FuncMLGraphBuilderConstant1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Constant1
//go:noescape
func CallMLGraphBuilderConstant1(
	this js.Ref, retPtr unsafe.Pointer,
	value float64,
	typ uint32)

//go:wasmimport plat/js/web try_MLGraphBuilder_Constant1
//go:noescape
func TryMLGraphBuilderConstant1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64,
	typ uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Constant2
//go:noescape
func HasFuncMLGraphBuilderConstant2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Constant2
//go:noescape
func FuncMLGraphBuilderConstant2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Constant2
//go:noescape
func CallMLGraphBuilderConstant2(
	this js.Ref, retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_MLGraphBuilder_Constant2
//go:noescape
func TryMLGraphBuilderConstant2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Build
//go:noescape
func HasFuncMLGraphBuilderBuild(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Build
//go:noescape
func FuncMLGraphBuilderBuild(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Build
//go:noescape
func CallMLGraphBuilderBuild(
	this js.Ref, retPtr unsafe.Pointer,
	outputs js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Build
//go:noescape
func TryMLGraphBuilderBuild(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	outputs js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_BuildSync
//go:noescape
func HasFuncMLGraphBuilderBuildSync(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_BuildSync
//go:noescape
func FuncMLGraphBuilderBuildSync(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_BuildSync
//go:noescape
func CallMLGraphBuilderBuildSync(
	this js.Ref, retPtr unsafe.Pointer,
	outputs js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_BuildSync
//go:noescape
func TryMLGraphBuilderBuildSync(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	outputs js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_HardSigmoid
//go:noescape
func HasFuncMLGraphBuilderHardSigmoid(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_HardSigmoid
//go:noescape
func FuncMLGraphBuilderHardSigmoid(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_HardSigmoid
//go:noescape
func CallMLGraphBuilderHardSigmoid(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_HardSigmoid
//go:noescape
func TryMLGraphBuilderHardSigmoid(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_HardSigmoid1
//go:noescape
func HasFuncMLGraphBuilderHardSigmoid1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_HardSigmoid1
//go:noescape
func FuncMLGraphBuilderHardSigmoid1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_HardSigmoid1
//go:noescape
func CallMLGraphBuilderHardSigmoid1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_HardSigmoid1
//go:noescape
func TryMLGraphBuilderHardSigmoid1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_HardSigmoid2
//go:noescape
func HasFuncMLGraphBuilderHardSigmoid2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_HardSigmoid2
//go:noescape
func FuncMLGraphBuilderHardSigmoid2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_HardSigmoid2
//go:noescape
func CallMLGraphBuilderHardSigmoid2(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_HardSigmoid2
//go:noescape
func TryMLGraphBuilderHardSigmoid2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_HardSigmoid3
//go:noescape
func HasFuncMLGraphBuilderHardSigmoid3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_HardSigmoid3
//go:noescape
func FuncMLGraphBuilderHardSigmoid3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_HardSigmoid3
//go:noescape
func CallMLGraphBuilderHardSigmoid3(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_HardSigmoid3
//go:noescape
func TryMLGraphBuilderHardSigmoid3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_GruCell
//go:noescape
func HasFuncMLGraphBuilderGruCell(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_GruCell
//go:noescape
func FuncMLGraphBuilderGruCell(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_GruCell
//go:noescape
func CallMLGraphBuilderGruCell(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	weight js.Ref,
	recurrentWeight js.Ref,
	hiddenState js.Ref,
	hiddenSize uint32,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_GruCell
//go:noescape
func TryMLGraphBuilderGruCell(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	weight js.Ref,
	recurrentWeight js.Ref,
	hiddenState js.Ref,
	hiddenSize uint32,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_GruCell1
//go:noescape
func HasFuncMLGraphBuilderGruCell1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_GruCell1
//go:noescape
func FuncMLGraphBuilderGruCell1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_GruCell1
//go:noescape
func CallMLGraphBuilderGruCell1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	weight js.Ref,
	recurrentWeight js.Ref,
	hiddenState js.Ref,
	hiddenSize uint32)

//go:wasmimport plat/js/web try_MLGraphBuilder_GruCell1
//go:noescape
func TryMLGraphBuilderGruCell1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	weight js.Ref,
	recurrentWeight js.Ref,
	hiddenState js.Ref,
	hiddenSize uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Slice
//go:noescape
func HasFuncMLGraphBuilderSlice(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Slice
//go:noescape
func FuncMLGraphBuilderSlice(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Slice
//go:noescape
func CallMLGraphBuilderSlice(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	starts js.Ref,
	sizes js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Slice
//go:noescape
func TryMLGraphBuilderSlice(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	starts js.Ref,
	sizes js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_AveragePool2d
//go:noescape
func HasFuncMLGraphBuilderAveragePool2d(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_AveragePool2d
//go:noescape
func FuncMLGraphBuilderAveragePool2d(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_AveragePool2d
//go:noescape
func CallMLGraphBuilderAveragePool2d(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_AveragePool2d
//go:noescape
func TryMLGraphBuilderAveragePool2d(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_AveragePool2d1
//go:noescape
func HasFuncMLGraphBuilderAveragePool2d1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_AveragePool2d1
//go:noescape
func FuncMLGraphBuilderAveragePool2d1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_AveragePool2d1
//go:noescape
func CallMLGraphBuilderAveragePool2d1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_AveragePool2d1
//go:noescape
func TryMLGraphBuilderAveragePool2d1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_L2Pool2d
//go:noescape
func HasFuncMLGraphBuilderL2Pool2d(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_L2Pool2d
//go:noescape
func FuncMLGraphBuilderL2Pool2d(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_L2Pool2d
//go:noescape
func CallMLGraphBuilderL2Pool2d(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_L2Pool2d
//go:noescape
func TryMLGraphBuilderL2Pool2d(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_L2Pool2d1
//go:noescape
func HasFuncMLGraphBuilderL2Pool2d1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_L2Pool2d1
//go:noescape
func FuncMLGraphBuilderL2Pool2d1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_L2Pool2d1
//go:noescape
func CallMLGraphBuilderL2Pool2d1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_L2Pool2d1
//go:noescape
func TryMLGraphBuilderL2Pool2d1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_MaxPool2d
//go:noescape
func HasFuncMLGraphBuilderMaxPool2d(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_MaxPool2d
//go:noescape
func FuncMLGraphBuilderMaxPool2d(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_MaxPool2d
//go:noescape
func CallMLGraphBuilderMaxPool2d(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_MaxPool2d
//go:noescape
func TryMLGraphBuilderMaxPool2d(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_MaxPool2d1
//go:noescape
func HasFuncMLGraphBuilderMaxPool2d1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_MaxPool2d1
//go:noescape
func FuncMLGraphBuilderMaxPool2d1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_MaxPool2d1
//go:noescape
func CallMLGraphBuilderMaxPool2d1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_MaxPool2d1
//go:noescape
func TryMLGraphBuilderMaxPool2d1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Linear
//go:noescape
func HasFuncMLGraphBuilderLinear(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Linear
//go:noescape
func FuncMLGraphBuilderLinear(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Linear
//go:noescape
func CallMLGraphBuilderLinear(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Linear
//go:noescape
func TryMLGraphBuilderLinear(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Linear1
//go:noescape
func HasFuncMLGraphBuilderLinear1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Linear1
//go:noescape
func FuncMLGraphBuilderLinear1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Linear1
//go:noescape
func CallMLGraphBuilderLinear1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Linear1
//go:noescape
func TryMLGraphBuilderLinear1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Linear2
//go:noescape
func HasFuncMLGraphBuilderLinear2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Linear2
//go:noescape
func FuncMLGraphBuilderLinear2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Linear2
//go:noescape
func CallMLGraphBuilderLinear2(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Linear2
//go:noescape
func TryMLGraphBuilderLinear2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Linear3
//go:noescape
func HasFuncMLGraphBuilderLinear3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Linear3
//go:noescape
func FuncMLGraphBuilderLinear3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Linear3
//go:noescape
func CallMLGraphBuilderLinear3(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Linear3
//go:noescape
func TryMLGraphBuilderLinear3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_LeakyRelu
//go:noescape
func HasFuncMLGraphBuilderLeakyRelu(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_LeakyRelu
//go:noescape
func FuncMLGraphBuilderLeakyRelu(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_LeakyRelu
//go:noescape
func CallMLGraphBuilderLeakyRelu(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_LeakyRelu
//go:noescape
func TryMLGraphBuilderLeakyRelu(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_LeakyRelu1
//go:noescape
func HasFuncMLGraphBuilderLeakyRelu1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_LeakyRelu1
//go:noescape
func FuncMLGraphBuilderLeakyRelu1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_LeakyRelu1
//go:noescape
func CallMLGraphBuilderLeakyRelu1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_LeakyRelu1
//go:noescape
func TryMLGraphBuilderLeakyRelu1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_LeakyRelu2
//go:noescape
func HasFuncMLGraphBuilderLeakyRelu2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_LeakyRelu2
//go:noescape
func FuncMLGraphBuilderLeakyRelu2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_LeakyRelu2
//go:noescape
func CallMLGraphBuilderLeakyRelu2(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_LeakyRelu2
//go:noescape
func TryMLGraphBuilderLeakyRelu2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_LeakyRelu3
//go:noescape
func HasFuncMLGraphBuilderLeakyRelu3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_LeakyRelu3
//go:noescape
func FuncMLGraphBuilderLeakyRelu3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_LeakyRelu3
//go:noescape
func CallMLGraphBuilderLeakyRelu3(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_LeakyRelu3
//go:noescape
func TryMLGraphBuilderLeakyRelu3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Pad
//go:noescape
func HasFuncMLGraphBuilderPad(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Pad
//go:noescape
func FuncMLGraphBuilderPad(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Pad
//go:noescape
func CallMLGraphBuilderPad(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	beginningPadding js.Ref,
	endingPadding js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Pad
//go:noescape
func TryMLGraphBuilderPad(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	beginningPadding js.Ref,
	endingPadding js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Pad1
//go:noescape
func HasFuncMLGraphBuilderPad1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Pad1
//go:noescape
func FuncMLGraphBuilderPad1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Pad1
//go:noescape
func CallMLGraphBuilderPad1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	beginningPadding js.Ref,
	endingPadding js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Pad1
//go:noescape
func TryMLGraphBuilderPad1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	beginningPadding js.Ref,
	endingPadding js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_InstanceNormalization
//go:noescape
func HasFuncMLGraphBuilderInstanceNormalization(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_InstanceNormalization
//go:noescape
func FuncMLGraphBuilderInstanceNormalization(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_InstanceNormalization
//go:noescape
func CallMLGraphBuilderInstanceNormalization(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_InstanceNormalization
//go:noescape
func TryMLGraphBuilderInstanceNormalization(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_InstanceNormalization1
//go:noescape
func HasFuncMLGraphBuilderInstanceNormalization1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_InstanceNormalization1
//go:noescape
func FuncMLGraphBuilderInstanceNormalization1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_InstanceNormalization1
//go:noescape
func CallMLGraphBuilderInstanceNormalization1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_InstanceNormalization1
//go:noescape
func TryMLGraphBuilderInstanceNormalization1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Softplus
//go:noescape
func HasFuncMLGraphBuilderSoftplus(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Softplus
//go:noescape
func FuncMLGraphBuilderSoftplus(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Softplus
//go:noescape
func CallMLGraphBuilderSoftplus(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Softplus
//go:noescape
func TryMLGraphBuilderSoftplus(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Softplus1
//go:noescape
func HasFuncMLGraphBuilderSoftplus1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Softplus1
//go:noescape
func FuncMLGraphBuilderSoftplus1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Softplus1
//go:noescape
func CallMLGraphBuilderSoftplus1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Softplus1
//go:noescape
func TryMLGraphBuilderSoftplus1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Softplus2
//go:noescape
func HasFuncMLGraphBuilderSoftplus2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Softplus2
//go:noescape
func FuncMLGraphBuilderSoftplus2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Softplus2
//go:noescape
func CallMLGraphBuilderSoftplus2(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Softplus2
//go:noescape
func TryMLGraphBuilderSoftplus2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Softplus3
//go:noescape
func HasFuncMLGraphBuilderSoftplus3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Softplus3
//go:noescape
func FuncMLGraphBuilderSoftplus3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Softplus3
//go:noescape
func CallMLGraphBuilderSoftplus3(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Softplus3
//go:noescape
func TryMLGraphBuilderSoftplus3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Softsign
//go:noescape
func HasFuncMLGraphBuilderSoftsign(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Softsign
//go:noescape
func FuncMLGraphBuilderSoftsign(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Softsign
//go:noescape
func CallMLGraphBuilderSoftsign(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Softsign
//go:noescape
func TryMLGraphBuilderSoftsign(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Softsign1
//go:noescape
func HasFuncMLGraphBuilderSoftsign1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Softsign1
//go:noescape
func FuncMLGraphBuilderSoftsign1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Softsign1
//go:noescape
func CallMLGraphBuilderSoftsign1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Softsign1
//go:noescape
func TryMLGraphBuilderSoftsign1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Sigmoid
//go:noescape
func HasFuncMLGraphBuilderSigmoid(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Sigmoid
//go:noescape
func FuncMLGraphBuilderSigmoid(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Sigmoid
//go:noescape
func CallMLGraphBuilderSigmoid(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Sigmoid
//go:noescape
func TryMLGraphBuilderSigmoid(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Sigmoid1
//go:noescape
func HasFuncMLGraphBuilderSigmoid1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Sigmoid1
//go:noescape
func FuncMLGraphBuilderSigmoid1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Sigmoid1
//go:noescape
func CallMLGraphBuilderSigmoid1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Sigmoid1
//go:noescape
func TryMLGraphBuilderSigmoid1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Reshape
//go:noescape
func HasFuncMLGraphBuilderReshape(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Reshape
//go:noescape
func FuncMLGraphBuilderReshape(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Reshape
//go:noescape
func CallMLGraphBuilderReshape(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	newShape js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Reshape
//go:noescape
func TryMLGraphBuilderReshape(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	newShape js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Conv2d
//go:noescape
func HasFuncMLGraphBuilderConv2d(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Conv2d
//go:noescape
func FuncMLGraphBuilderConv2d(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Conv2d
//go:noescape
func CallMLGraphBuilderConv2d(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	filter js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Conv2d
//go:noescape
func TryMLGraphBuilderConv2d(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	filter js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Conv2d1
//go:noescape
func HasFuncMLGraphBuilderConv2d1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Conv2d1
//go:noescape
func FuncMLGraphBuilderConv2d1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Conv2d1
//go:noescape
func CallMLGraphBuilderConv2d1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	filter js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Conv2d1
//go:noescape
func TryMLGraphBuilderConv2d1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	filter js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Split
//go:noescape
func HasFuncMLGraphBuilderSplit(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Split
//go:noescape
func FuncMLGraphBuilderSplit(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Split
//go:noescape
func CallMLGraphBuilderSplit(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	splits js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Split
//go:noescape
func TryMLGraphBuilderSplit(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	splits js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Split1
//go:noescape
func HasFuncMLGraphBuilderSplit1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Split1
//go:noescape
func FuncMLGraphBuilderSplit1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Split1
//go:noescape
func CallMLGraphBuilderSplit1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	splits js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Split1
//go:noescape
func TryMLGraphBuilderSplit1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	splits js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Resample2d
//go:noescape
func HasFuncMLGraphBuilderResample2d(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Resample2d
//go:noescape
func FuncMLGraphBuilderResample2d(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Resample2d
//go:noescape
func CallMLGraphBuilderResample2d(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Resample2d
//go:noescape
func TryMLGraphBuilderResample2d(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Resample2d1
//go:noescape
func HasFuncMLGraphBuilderResample2d1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Resample2d1
//go:noescape
func FuncMLGraphBuilderResample2d1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Resample2d1
//go:noescape
func CallMLGraphBuilderResample2d1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Resample2d1
//go:noescape
func TryMLGraphBuilderResample2d1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_HardSwish
//go:noescape
func HasFuncMLGraphBuilderHardSwish(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_HardSwish
//go:noescape
func FuncMLGraphBuilderHardSwish(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_HardSwish
//go:noescape
func CallMLGraphBuilderHardSwish(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_HardSwish
//go:noescape
func TryMLGraphBuilderHardSwish(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_HardSwish1
//go:noescape
func HasFuncMLGraphBuilderHardSwish1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_HardSwish1
//go:noescape
func FuncMLGraphBuilderHardSwish1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_HardSwish1
//go:noescape
func CallMLGraphBuilderHardSwish1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_HardSwish1
//go:noescape
func TryMLGraphBuilderHardSwish1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Softmax
//go:noescape
func HasFuncMLGraphBuilderSoftmax(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Softmax
//go:noescape
func FuncMLGraphBuilderSoftmax(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Softmax
//go:noescape
func CallMLGraphBuilderSoftmax(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Softmax
//go:noescape
func TryMLGraphBuilderSoftmax(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Softmax1
//go:noescape
func HasFuncMLGraphBuilderSoftmax1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Softmax1
//go:noescape
func FuncMLGraphBuilderSoftmax1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Softmax1
//go:noescape
func CallMLGraphBuilderSoftmax1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Softmax1
//go:noescape
func TryMLGraphBuilderSoftmax1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ConvTranspose2d
//go:noescape
func HasFuncMLGraphBuilderConvTranspose2d(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ConvTranspose2d
//go:noescape
func FuncMLGraphBuilderConvTranspose2d(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ConvTranspose2d
//go:noescape
func CallMLGraphBuilderConvTranspose2d(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	filter js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_ConvTranspose2d
//go:noescape
func TryMLGraphBuilderConvTranspose2d(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	filter js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ConvTranspose2d1
//go:noescape
func HasFuncMLGraphBuilderConvTranspose2d1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ConvTranspose2d1
//go:noescape
func FuncMLGraphBuilderConvTranspose2d1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ConvTranspose2d1
//go:noescape
func CallMLGraphBuilderConvTranspose2d1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	filter js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_ConvTranspose2d1
//go:noescape
func TryMLGraphBuilderConvTranspose2d1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	filter js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Relu
//go:noescape
func HasFuncMLGraphBuilderRelu(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Relu
//go:noescape
func FuncMLGraphBuilderRelu(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Relu
//go:noescape
func CallMLGraphBuilderRelu(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Relu
//go:noescape
func TryMLGraphBuilderRelu(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Relu1
//go:noescape
func HasFuncMLGraphBuilderRelu1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Relu1
//go:noescape
func FuncMLGraphBuilderRelu1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Relu1
//go:noescape
func CallMLGraphBuilderRelu1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Relu1
//go:noescape
func TryMLGraphBuilderRelu1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Add
//go:noescape
func HasFuncMLGraphBuilderAdd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Add
//go:noescape
func FuncMLGraphBuilderAdd(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Add
//go:noescape
func CallMLGraphBuilderAdd(
	this js.Ref, retPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Add
//go:noescape
func TryMLGraphBuilderAdd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Sub
//go:noescape
func HasFuncMLGraphBuilderSub(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Sub
//go:noescape
func FuncMLGraphBuilderSub(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Sub
//go:noescape
func CallMLGraphBuilderSub(
	this js.Ref, retPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Sub
//go:noescape
func TryMLGraphBuilderSub(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Mul
//go:noescape
func HasFuncMLGraphBuilderMul(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Mul
//go:noescape
func FuncMLGraphBuilderMul(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Mul
//go:noescape
func CallMLGraphBuilderMul(
	this js.Ref, retPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Mul
//go:noescape
func TryMLGraphBuilderMul(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Div
//go:noescape
func HasFuncMLGraphBuilderDiv(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Div
//go:noescape
func FuncMLGraphBuilderDiv(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Div
//go:noescape
func CallMLGraphBuilderDiv(
	this js.Ref, retPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Div
//go:noescape
func TryMLGraphBuilderDiv(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Max
//go:noescape
func HasFuncMLGraphBuilderMax(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Max
//go:noescape
func FuncMLGraphBuilderMax(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Max
//go:noescape
func CallMLGraphBuilderMax(
	this js.Ref, retPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Max
//go:noescape
func TryMLGraphBuilderMax(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Min
//go:noescape
func HasFuncMLGraphBuilderMin(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Min
//go:noescape
func FuncMLGraphBuilderMin(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Min
//go:noescape
func CallMLGraphBuilderMin(
	this js.Ref, retPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Min
//go:noescape
func TryMLGraphBuilderMin(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Pow
//go:noescape
func HasFuncMLGraphBuilderPow(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Pow
//go:noescape
func FuncMLGraphBuilderPow(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Pow
//go:noescape
func CallMLGraphBuilderPow(
	this js.Ref, retPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Pow
//go:noescape
func TryMLGraphBuilderPow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceL1
//go:noescape
func HasFuncMLGraphBuilderReduceL1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceL1
//go:noescape
func FuncMLGraphBuilderReduceL1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceL1
//go:noescape
func CallMLGraphBuilderReduceL1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceL1
//go:noescape
func TryMLGraphBuilderReduceL1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceL11
//go:noescape
func HasFuncMLGraphBuilderReduceL11(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceL11
//go:noescape
func FuncMLGraphBuilderReduceL11(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceL11
//go:noescape
func CallMLGraphBuilderReduceL11(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceL11
//go:noescape
func TryMLGraphBuilderReduceL11(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceL2
//go:noescape
func HasFuncMLGraphBuilderReduceL2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceL2
//go:noescape
func FuncMLGraphBuilderReduceL2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceL2
//go:noescape
func CallMLGraphBuilderReduceL2(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceL2
//go:noescape
func TryMLGraphBuilderReduceL2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceL21
//go:noescape
func HasFuncMLGraphBuilderReduceL21(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceL21
//go:noescape
func FuncMLGraphBuilderReduceL21(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceL21
//go:noescape
func CallMLGraphBuilderReduceL21(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceL21
//go:noescape
func TryMLGraphBuilderReduceL21(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceLogSum
//go:noescape
func HasFuncMLGraphBuilderReduceLogSum(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceLogSum
//go:noescape
func FuncMLGraphBuilderReduceLogSum(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceLogSum
//go:noescape
func CallMLGraphBuilderReduceLogSum(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceLogSum
//go:noescape
func TryMLGraphBuilderReduceLogSum(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceLogSum1
//go:noescape
func HasFuncMLGraphBuilderReduceLogSum1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceLogSum1
//go:noescape
func FuncMLGraphBuilderReduceLogSum1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceLogSum1
//go:noescape
func CallMLGraphBuilderReduceLogSum1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceLogSum1
//go:noescape
func TryMLGraphBuilderReduceLogSum1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceLogSumExp
//go:noescape
func HasFuncMLGraphBuilderReduceLogSumExp(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceLogSumExp
//go:noescape
func FuncMLGraphBuilderReduceLogSumExp(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceLogSumExp
//go:noescape
func CallMLGraphBuilderReduceLogSumExp(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceLogSumExp
//go:noescape
func TryMLGraphBuilderReduceLogSumExp(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceLogSumExp1
//go:noescape
func HasFuncMLGraphBuilderReduceLogSumExp1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceLogSumExp1
//go:noescape
func FuncMLGraphBuilderReduceLogSumExp1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceLogSumExp1
//go:noescape
func CallMLGraphBuilderReduceLogSumExp1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceLogSumExp1
//go:noescape
func TryMLGraphBuilderReduceLogSumExp1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceMax
//go:noescape
func HasFuncMLGraphBuilderReduceMax(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceMax
//go:noescape
func FuncMLGraphBuilderReduceMax(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceMax
//go:noescape
func CallMLGraphBuilderReduceMax(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceMax
//go:noescape
func TryMLGraphBuilderReduceMax(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceMax1
//go:noescape
func HasFuncMLGraphBuilderReduceMax1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceMax1
//go:noescape
func FuncMLGraphBuilderReduceMax1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceMax1
//go:noescape
func CallMLGraphBuilderReduceMax1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceMax1
//go:noescape
func TryMLGraphBuilderReduceMax1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceMean
//go:noescape
func HasFuncMLGraphBuilderReduceMean(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceMean
//go:noescape
func FuncMLGraphBuilderReduceMean(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceMean
//go:noescape
func CallMLGraphBuilderReduceMean(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceMean
//go:noescape
func TryMLGraphBuilderReduceMean(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceMean1
//go:noescape
func HasFuncMLGraphBuilderReduceMean1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceMean1
//go:noescape
func FuncMLGraphBuilderReduceMean1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceMean1
//go:noescape
func CallMLGraphBuilderReduceMean1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceMean1
//go:noescape
func TryMLGraphBuilderReduceMean1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceMin
//go:noescape
func HasFuncMLGraphBuilderReduceMin(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceMin
//go:noescape
func FuncMLGraphBuilderReduceMin(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceMin
//go:noescape
func CallMLGraphBuilderReduceMin(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceMin
//go:noescape
func TryMLGraphBuilderReduceMin(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceMin1
//go:noescape
func HasFuncMLGraphBuilderReduceMin1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceMin1
//go:noescape
func FuncMLGraphBuilderReduceMin1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceMin1
//go:noescape
func CallMLGraphBuilderReduceMin1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceMin1
//go:noescape
func TryMLGraphBuilderReduceMin1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceProduct
//go:noescape
func HasFuncMLGraphBuilderReduceProduct(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceProduct
//go:noescape
func FuncMLGraphBuilderReduceProduct(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceProduct
//go:noescape
func CallMLGraphBuilderReduceProduct(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceProduct
//go:noescape
func TryMLGraphBuilderReduceProduct(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceProduct1
//go:noescape
func HasFuncMLGraphBuilderReduceProduct1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceProduct1
//go:noescape
func FuncMLGraphBuilderReduceProduct1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceProduct1
//go:noescape
func CallMLGraphBuilderReduceProduct1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceProduct1
//go:noescape
func TryMLGraphBuilderReduceProduct1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceSum
//go:noescape
func HasFuncMLGraphBuilderReduceSum(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceSum
//go:noescape
func FuncMLGraphBuilderReduceSum(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceSum
//go:noescape
func CallMLGraphBuilderReduceSum(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceSum
//go:noescape
func TryMLGraphBuilderReduceSum(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceSum1
//go:noescape
func HasFuncMLGraphBuilderReduceSum1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceSum1
//go:noescape
func FuncMLGraphBuilderReduceSum1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceSum1
//go:noescape
func CallMLGraphBuilderReduceSum1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceSum1
//go:noescape
func TryMLGraphBuilderReduceSum1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceSumSquare
//go:noescape
func HasFuncMLGraphBuilderReduceSumSquare(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceSumSquare
//go:noescape
func FuncMLGraphBuilderReduceSumSquare(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceSumSquare
//go:noescape
func CallMLGraphBuilderReduceSumSquare(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceSumSquare
//go:noescape
func TryMLGraphBuilderReduceSumSquare(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_ReduceSumSquare1
//go:noescape
func HasFuncMLGraphBuilderReduceSumSquare1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_ReduceSumSquare1
//go:noescape
func FuncMLGraphBuilderReduceSumSquare1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_ReduceSumSquare1
//go:noescape
func CallMLGraphBuilderReduceSumSquare1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_ReduceSumSquare1
//go:noescape
func TryMLGraphBuilderReduceSumSquare1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Lstm
//go:noescape
func HasFuncMLGraphBuilderLstm(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Lstm
//go:noescape
func FuncMLGraphBuilderLstm(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Lstm
//go:noescape
func CallMLGraphBuilderLstm(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	weight js.Ref,
	recurrentWeight js.Ref,
	steps uint32,
	hiddenSize uint32,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Lstm
//go:noescape
func TryMLGraphBuilderLstm(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	weight js.Ref,
	recurrentWeight js.Ref,
	steps uint32,
	hiddenSize uint32,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Lstm1
//go:noescape
func HasFuncMLGraphBuilderLstm1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Lstm1
//go:noescape
func FuncMLGraphBuilderLstm1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Lstm1
//go:noescape
func CallMLGraphBuilderLstm1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	weight js.Ref,
	recurrentWeight js.Ref,
	steps uint32,
	hiddenSize uint32)

//go:wasmimport plat/js/web try_MLGraphBuilder_Lstm1
//go:noescape
func TryMLGraphBuilderLstm1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	weight js.Ref,
	recurrentWeight js.Ref,
	steps uint32,
	hiddenSize uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Matmul
//go:noescape
func HasFuncMLGraphBuilderMatmul(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Matmul
//go:noescape
func FuncMLGraphBuilderMatmul(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Matmul
//go:noescape
func CallMLGraphBuilderMatmul(
	this js.Ref, retPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Matmul
//go:noescape
func TryMLGraphBuilderMatmul(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Squeeze
//go:noescape
func HasFuncMLGraphBuilderSqueeze(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Squeeze
//go:noescape
func FuncMLGraphBuilderSqueeze(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Squeeze
//go:noescape
func CallMLGraphBuilderSqueeze(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Squeeze
//go:noescape
func TryMLGraphBuilderSqueeze(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Squeeze1
//go:noescape
func HasFuncMLGraphBuilderSqueeze1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Squeeze1
//go:noescape
func FuncMLGraphBuilderSqueeze1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Squeeze1
//go:noescape
func CallMLGraphBuilderSqueeze1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Squeeze1
//go:noescape
func TryMLGraphBuilderSqueeze1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Tanh
//go:noescape
func HasFuncMLGraphBuilderTanh(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Tanh
//go:noescape
func FuncMLGraphBuilderTanh(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Tanh
//go:noescape
func CallMLGraphBuilderTanh(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Tanh
//go:noescape
func TryMLGraphBuilderTanh(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Tanh1
//go:noescape
func HasFuncMLGraphBuilderTanh1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Tanh1
//go:noescape
func FuncMLGraphBuilderTanh1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Tanh1
//go:noescape
func CallMLGraphBuilderTanh1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Tanh1
//go:noescape
func TryMLGraphBuilderTanh1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Gru
//go:noescape
func HasFuncMLGraphBuilderGru(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Gru
//go:noescape
func FuncMLGraphBuilderGru(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Gru
//go:noescape
func CallMLGraphBuilderGru(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	weight js.Ref,
	recurrentWeight js.Ref,
	steps uint32,
	hiddenSize uint32,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Gru
//go:noescape
func TryMLGraphBuilderGru(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	weight js.Ref,
	recurrentWeight js.Ref,
	steps uint32,
	hiddenSize uint32,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Gru1
//go:noescape
func HasFuncMLGraphBuilderGru1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Gru1
//go:noescape
func FuncMLGraphBuilderGru1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Gru1
//go:noescape
func CallMLGraphBuilderGru1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	weight js.Ref,
	recurrentWeight js.Ref,
	steps uint32,
	hiddenSize uint32)

//go:wasmimport plat/js/web try_MLGraphBuilder_Gru1
//go:noescape
func TryMLGraphBuilderGru1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	weight js.Ref,
	recurrentWeight js.Ref,
	steps uint32,
	hiddenSize uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Abs
//go:noescape
func HasFuncMLGraphBuilderAbs(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Abs
//go:noescape
func FuncMLGraphBuilderAbs(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Abs
//go:noescape
func CallMLGraphBuilderAbs(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Abs
//go:noescape
func TryMLGraphBuilderAbs(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Ceil
//go:noescape
func HasFuncMLGraphBuilderCeil(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Ceil
//go:noescape
func FuncMLGraphBuilderCeil(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Ceil
//go:noescape
func CallMLGraphBuilderCeil(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Ceil
//go:noescape
func TryMLGraphBuilderCeil(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Cos
//go:noescape
func HasFuncMLGraphBuilderCos(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Cos
//go:noescape
func FuncMLGraphBuilderCos(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Cos
//go:noescape
func CallMLGraphBuilderCos(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Cos
//go:noescape
func TryMLGraphBuilderCos(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Exp
//go:noescape
func HasFuncMLGraphBuilderExp(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Exp
//go:noescape
func FuncMLGraphBuilderExp(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Exp
//go:noescape
func CallMLGraphBuilderExp(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Exp
//go:noescape
func TryMLGraphBuilderExp(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Floor
//go:noescape
func HasFuncMLGraphBuilderFloor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Floor
//go:noescape
func FuncMLGraphBuilderFloor(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Floor
//go:noescape
func CallMLGraphBuilderFloor(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Floor
//go:noescape
func TryMLGraphBuilderFloor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Log
//go:noescape
func HasFuncMLGraphBuilderLog(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Log
//go:noescape
func FuncMLGraphBuilderLog(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Log
//go:noescape
func CallMLGraphBuilderLog(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Log
//go:noescape
func TryMLGraphBuilderLog(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Neg
//go:noescape
func HasFuncMLGraphBuilderNeg(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Neg
//go:noescape
func FuncMLGraphBuilderNeg(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Neg
//go:noescape
func CallMLGraphBuilderNeg(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Neg
//go:noescape
func TryMLGraphBuilderNeg(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Sin
//go:noescape
func HasFuncMLGraphBuilderSin(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Sin
//go:noescape
func FuncMLGraphBuilderSin(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Sin
//go:noescape
func CallMLGraphBuilderSin(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Sin
//go:noescape
func TryMLGraphBuilderSin(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Tan
//go:noescape
func HasFuncMLGraphBuilderTan(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Tan
//go:noescape
func FuncMLGraphBuilderTan(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Tan
//go:noescape
func CallMLGraphBuilderTan(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Tan
//go:noescape
func TryMLGraphBuilderTan(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Transpose
//go:noescape
func HasFuncMLGraphBuilderTranspose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Transpose
//go:noescape
func FuncMLGraphBuilderTranspose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Transpose
//go:noescape
func CallMLGraphBuilderTranspose(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Transpose
//go:noescape
func TryMLGraphBuilderTranspose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Transpose1
//go:noescape
func HasFuncMLGraphBuilderTranspose1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Transpose1
//go:noescape
func FuncMLGraphBuilderTranspose1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Transpose1
//go:noescape
func CallMLGraphBuilderTranspose1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Transpose1
//go:noescape
func TryMLGraphBuilderTranspose1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Prelu
//go:noescape
func HasFuncMLGraphBuilderPrelu(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Prelu
//go:noescape
func FuncMLGraphBuilderPrelu(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Prelu
//go:noescape
func CallMLGraphBuilderPrelu(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	slope js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Prelu
//go:noescape
func TryMLGraphBuilderPrelu(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	slope js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Concat
//go:noescape
func HasFuncMLGraphBuilderConcat(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Concat
//go:noescape
func FuncMLGraphBuilderConcat(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Concat
//go:noescape
func CallMLGraphBuilderConcat(
	this js.Ref, retPtr unsafe.Pointer,
	inputs js.Ref,
	axis uint32)

//go:wasmimport plat/js/web try_MLGraphBuilder_Concat
//go:noescape
func TryMLGraphBuilderConcat(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	inputs js.Ref,
	axis uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Gemm
//go:noescape
func HasFuncMLGraphBuilderGemm(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Gemm
//go:noescape
func FuncMLGraphBuilderGemm(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Gemm
//go:noescape
func CallMLGraphBuilderGemm(
	this js.Ref, retPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Gemm
//go:noescape
func TryMLGraphBuilderGemm(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Gemm1
//go:noescape
func HasFuncMLGraphBuilderGemm1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Gemm1
//go:noescape
func FuncMLGraphBuilderGemm1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Gemm1
//go:noescape
func CallMLGraphBuilderGemm1(
	this js.Ref, retPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Gemm1
//go:noescape
func TryMLGraphBuilderGemm1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	a js.Ref,
	b js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_LstmCell
//go:noescape
func HasFuncMLGraphBuilderLstmCell(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_LstmCell
//go:noescape
func FuncMLGraphBuilderLstmCell(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_LstmCell
//go:noescape
func CallMLGraphBuilderLstmCell(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	weight js.Ref,
	recurrentWeight js.Ref,
	hiddenState js.Ref,
	cellState js.Ref,
	hiddenSize uint32,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_LstmCell
//go:noescape
func TryMLGraphBuilderLstmCell(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	weight js.Ref,
	recurrentWeight js.Ref,
	hiddenState js.Ref,
	cellState js.Ref,
	hiddenSize uint32,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_LstmCell1
//go:noescape
func HasFuncMLGraphBuilderLstmCell1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_LstmCell1
//go:noescape
func FuncMLGraphBuilderLstmCell1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_LstmCell1
//go:noescape
func CallMLGraphBuilderLstmCell1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	weight js.Ref,
	recurrentWeight js.Ref,
	hiddenState js.Ref,
	cellState js.Ref,
	hiddenSize uint32)

//go:wasmimport plat/js/web try_MLGraphBuilder_LstmCell1
//go:noescape
func TryMLGraphBuilderLstmCell1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	weight js.Ref,
	recurrentWeight js.Ref,
	hiddenState js.Ref,
	cellState js.Ref,
	hiddenSize uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_BatchNormalization
//go:noescape
func HasFuncMLGraphBuilderBatchNormalization(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_BatchNormalization
//go:noescape
func FuncMLGraphBuilderBatchNormalization(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_BatchNormalization
//go:noescape
func CallMLGraphBuilderBatchNormalization(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	mean js.Ref,
	variance js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_BatchNormalization
//go:noescape
func TryMLGraphBuilderBatchNormalization(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	mean js.Ref,
	variance js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_BatchNormalization1
//go:noescape
func HasFuncMLGraphBuilderBatchNormalization1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_BatchNormalization1
//go:noescape
func FuncMLGraphBuilderBatchNormalization1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_BatchNormalization1
//go:noescape
func CallMLGraphBuilderBatchNormalization1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	mean js.Ref,
	variance js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_BatchNormalization1
//go:noescape
func TryMLGraphBuilderBatchNormalization1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	mean js.Ref,
	variance js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Elu
//go:noescape
func HasFuncMLGraphBuilderElu(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Elu
//go:noescape
func FuncMLGraphBuilderElu(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Elu
//go:noescape
func CallMLGraphBuilderElu(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Elu
//go:noescape
func TryMLGraphBuilderElu(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Elu1
//go:noescape
func HasFuncMLGraphBuilderElu1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Elu1
//go:noescape
func FuncMLGraphBuilderElu1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Elu1
//go:noescape
func CallMLGraphBuilderElu1(
	this js.Ref, retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Elu1
//go:noescape
func TryMLGraphBuilderElu1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Elu2
//go:noescape
func HasFuncMLGraphBuilderElu2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Elu2
//go:noescape
func FuncMLGraphBuilderElu2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Elu2
//go:noescape
func CallMLGraphBuilderElu2(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Elu2
//go:noescape
func TryMLGraphBuilderElu2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Elu3
//go:noescape
func HasFuncMLGraphBuilderElu3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Elu3
//go:noescape
func FuncMLGraphBuilderElu3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Elu3
//go:noescape
func CallMLGraphBuilderElu3(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Elu3
//go:noescape
func TryMLGraphBuilderElu3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Clamp
//go:noescape
func HasFuncMLGraphBuilderClamp(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Clamp
//go:noescape
func FuncMLGraphBuilderClamp(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Clamp
//go:noescape
func CallMLGraphBuilderClamp(
	this js.Ref, retPtr unsafe.Pointer,
	operand js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Clamp
//go:noescape
func TryMLGraphBuilderClamp(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	operand js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Clamp1
//go:noescape
func HasFuncMLGraphBuilderClamp1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Clamp1
//go:noescape
func FuncMLGraphBuilderClamp1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Clamp1
//go:noescape
func CallMLGraphBuilderClamp1(
	this js.Ref, retPtr unsafe.Pointer,
	operand js.Ref)

//go:wasmimport plat/js/web try_MLGraphBuilder_Clamp1
//go:noescape
func TryMLGraphBuilderClamp1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	operand js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Clamp2
//go:noescape
func HasFuncMLGraphBuilderClamp2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Clamp2
//go:noescape
func FuncMLGraphBuilderClamp2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Clamp2
//go:noescape
func CallMLGraphBuilderClamp2(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Clamp2
//go:noescape
func TryMLGraphBuilderClamp2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLGraphBuilder_Clamp3
//go:noescape
func HasFuncMLGraphBuilderClamp3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLGraphBuilder_Clamp3
//go:noescape
func FuncMLGraphBuilderClamp3(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MLGraphBuilder_Clamp3
//go:noescape
func CallMLGraphBuilderClamp3(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MLGraphBuilder_Clamp3
//go:noescape
func TryMLGraphBuilderClamp3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_MagnetometerLocalCoordinateSystem
//go:noescape
func ConstOfMagnetometerLocalCoordinateSystem(str js.Ref) uint32

//go:wasmimport plat/js/web store_MagnetometerSensorOptions
//go:noescape
func MagnetometerSensorOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MagnetometerSensorOptions
//go:noescape
func MagnetometerSensorOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_Magnetometer_Magnetometer
//go:noescape
func NewMagnetometerByMagnetometer(
	sensorOptions unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_Magnetometer_Magnetometer1
//go:noescape
func NewMagnetometerByMagnetometer1() js.Ref

//go:wasmimport plat/js/web get_Magnetometer_X
//go:noescape
func GetMagnetometerX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Magnetometer_Y
//go:noescape
func GetMagnetometerY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Magnetometer_Z
//go:noescape
func GetMagnetometerZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_MagnetometerReadingValues
//go:noescape
func MagnetometerReadingValuesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MagnetometerReadingValues
//go:noescape
func MagnetometerReadingValuesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_MathMLElement_Style
//go:noescape
func GetMathMLElementStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MathMLElement_AttributeStyleMap
//go:noescape
func GetMathMLElementAttributeStyleMap(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MathMLElement_Dataset
//go:noescape
func GetMathMLElementDataset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MathMLElement_Nonce
//go:noescape
func GetMathMLElementNonce(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_MathMLElement_Nonce
//go:noescape
func SetMathMLElementNonce(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_MathMLElement_Autofocus
//go:noescape
func GetMathMLElementAutofocus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_MathMLElement_Autofocus
//go:noescape
func SetMathMLElementAutofocus(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_MathMLElement_TabIndex
//go:noescape
func GetMathMLElementTabIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_MathMLElement_TabIndex
//go:noescape
func SetMathMLElementTabIndex(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web has_MathMLElement_Focus
//go:noescape
func HasFuncMathMLElementFocus(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MathMLElement_Focus
//go:noescape
func FuncMathMLElementFocus(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MathMLElement_Focus
//go:noescape
func CallMathMLElementFocus(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_MathMLElement_Focus
//go:noescape
func TryMathMLElementFocus(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MathMLElement_Focus1
//go:noescape
func HasFuncMathMLElementFocus1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MathMLElement_Focus1
//go:noescape
func FuncMathMLElementFocus1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MathMLElement_Focus1
//go:noescape
func CallMathMLElementFocus1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MathMLElement_Focus1
//go:noescape
func TryMathMLElementFocus1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MathMLElement_Blur
//go:noescape
func HasFuncMathMLElementBlur(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MathMLElement_Blur
//go:noescape
func FuncMathMLElementBlur(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_MathMLElement_Blur
//go:noescape
func CallMathMLElementBlur(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MathMLElement_Blur
//go:noescape
func TryMathMLElementBlur(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_MediaCapabilitiesInfo
//go:noescape
func MediaCapabilitiesInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaCapabilitiesInfo
//go:noescape
func MediaCapabilitiesInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaConfiguration
//go:noescape
func MediaConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaConfiguration
//go:noescape
func MediaConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaEncryptedEventInit
//go:noescape
func MediaEncryptedEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaEncryptedEventInit
//go:noescape
func MediaEncryptedEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_MediaEncryptedEvent_MediaEncryptedEvent
//go:noescape
func NewMediaEncryptedEventByMediaEncryptedEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_MediaEncryptedEvent_MediaEncryptedEvent1
//go:noescape
func NewMediaEncryptedEventByMediaEncryptedEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_MediaEncryptedEvent_InitDataType
//go:noescape
func GetMediaEncryptedEventInitDataType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaEncryptedEvent_InitData
//go:noescape
func GetMediaEncryptedEventInitData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_MediaKeyMessageType
//go:noescape
func ConstOfMediaKeyMessageType(str js.Ref) uint32

//go:wasmimport plat/js/web store_MediaKeyMessageEventInit
//go:noescape
func MediaKeyMessageEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaKeyMessageEventInit
//go:noescape
func MediaKeyMessageEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_MediaKeyMessageEvent_MediaKeyMessageEvent
//go:noescape
func NewMediaKeyMessageEventByMediaKeyMessageEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_MediaKeyMessageEvent_MessageType
//go:noescape
func GetMediaKeyMessageEventMessageType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaKeyMessageEvent_Message
//go:noescape
func GetMediaKeyMessageEventMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_MediaQueryListEventInit
//go:noescape
func MediaQueryListEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaQueryListEventInit
//go:noescape
func MediaQueryListEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_MediaQueryListEvent_MediaQueryListEvent
//go:noescape
func NewMediaQueryListEventByMediaQueryListEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_MediaQueryListEvent_MediaQueryListEvent1
//go:noescape
func NewMediaQueryListEventByMediaQueryListEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_MediaQueryListEvent_Media
//go:noescape
func GetMediaQueryListEventMedia(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_MediaQueryListEvent_Matches
//go:noescape
func GetMediaQueryListEventMatches(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_MediaRecorderOptions
//go:noescape
func MediaRecorderOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaRecorderOptions
//go:noescape
func MediaRecorderOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
